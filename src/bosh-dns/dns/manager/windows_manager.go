// +build windows

package manager

import (
	"net"
	"os"
	"path/filepath"
	"syscall"
	"unicode/utf16"
	"unsafe"

	"golang.org/x/sys/windows"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
)

const prependDNSServer = `
param ($DNSAddress = $(throw "DNSAddress parameter is required."))

$ErrorActionPreference = "Stop"

function DnsServers($interface) {
  return (Get-DnsClientServerAddress -InterfaceAlias $interface -AddressFamily ipv4 -ErrorAction Stop).ServerAddresses
}

try {
  # identify our interface
  [array]$routeable_interfaces = Get-WmiObject Win32_NetworkAdapterConfiguration | Where { $_.IpAddress -AND ($_.IpAddress | Where { $addr = [Net.IPAddress] $_; $addr.AddressFamily -eq "InterNetwork" -AND ($addr.address -BAND ([Net.IPAddress] "255.255.0.0").address) -ne ([Net.IPAddress] "169.254.0.0").address }) }
  $interface = (Get-WmiObject Win32_NetworkAdapter | Where { $_.DeviceID -eq $routeable_interfaces[0].Index }).netconnectionid

  # avoid prepending if we happen to already be at the top to try and avoid races
  [array]$servers = DnsServers($interface)
  if($servers[0] -eq $DNSAddress) {
    Exit 0
  }

  Set-DnsClientServerAddress -InterfaceAlias $interface -ServerAddresses (,$DNSAddress + $servers)

  # read back the servers in case set silently failed
  [array]$servers = DnsServers($interface)
  if($servers[0] -ne $DNSAddress) {
      Write-Error "Failed to set '${DNSAddress}' as the first dns client server address"
  }
} catch {
  $Host.UI.WriteErrorLine($_.Exception.Message)
  Exit 1
}

Exit 0
`

type windowsManager struct {
	runner boshsys.CmdRunner
	fs     boshsys.FileSystem
}

func NewWindowsManager(runner boshsys.CmdRunner, fs boshsys.FileSystem) *windowsManager {
	return &windowsManager{runner: runner, fs: fs}
}

func (manager *windowsManager) SetPrimary(address string) error {
	servers, err := manager.Read()
	if err != nil {
		return err
	}

	if len(servers) > 0 && servers[0] == address {
		return nil
	}

	scriptName, err := manager.writeScript("prepend-dns-server", prependDNSServer)
	if err != nil {
		return bosherr.WrapError(err, "Creating prepend-dns-server.ps1")
	}
	defer manager.fs.RemoveAll(filepath.Dir(scriptName))

	_, _, _, err = manager.runner.RunCommand("powershell.exe", scriptName, address)
	if err != nil {
		return bosherr.WrapError(err, "Executing prepend-dns-server.ps1")
	}

	return nil
}

func (manager *windowsManager) Read() ([]string, error) {
	servers, err := dnsResolvers()
	if err != nil {
		return nil, bosherr.WrapError(err, "Getting list of current DNS Servers")
	}

	return servers, nil
}

func (manager *windowsManager) writeScript(name, contents string) (string, error) {
	dir, err := manager.fs.TempDir(name)
	if err != nil {
		return "", err
	}

	scriptName := filepath.Join(dir, name+".ps1")
	err = manager.fs.WriteFileString(scriptName, contents)
	if err != nil {
		return "", err
	}

	err = manager.fs.Chmod(scriptName, 0700)
	if err != nil {
		return "", err
	}

	return scriptName, nil
}

func dnsResolvers() ([]string, error) {
	addresses, err := getAllPhysicalInterface()
	if err != nil {
		return nil, err
	}

	var resolvers []string

	for _, addr := range addresses {
		for aa := addr.FirstDnsServerAddress; aa != nil; aa = aa.Next {
			resolvers = append(resolvers, sockaddrToIP(aa.Address))
		}
	}

	return resolvers, nil
}

func sockaddrToIP(sockaddr windows.SocketAddress) (string, err) {
	sa, err := aa.Address.Sockaddr.Sockaddr()
	if err != nil {
		return "", os.NewSyscallError("sockaddr", err)
	}

	switch sa := sa.(type) {
	case *syscall.SockaddrInet4:
		ifa := &net.IPAddr{IP: net.IPv4(sa.Addr[0], sa.Addr[1], sa.Addr[2], sa.Addr[3])}
		return ifa.String(), nil
	case *syscall.SockaddrInet6:
		ifa := &net.IPAddr{IP: make(net.IP, net.IPv6len)}
		copy(ifa.IP, sa.Addr[:])
		return ifa.String(), nil
	}

	return "", nil
}

const (
	IfOperStatusUp            = 1
	IF_TYPE_SOFTWARE_LOOPBACK = 24
	IF_TYPE_TUNNEL            = 131
)

const hexDigit = "0123456789abcdef"

func adapterAddresses() ([]*windows.IpAdapterAddresses, error) {
	var b []byte
	l := uint32(15000) // recommended initial size
	for {
		b = make([]byte, l)
		err := windows.GetAdaptersAddresses(syscall.AF_UNSPEC, windows.GAA_FLAG_INCLUDE_PREFIX, 0, (*windows.IpAdapterAddresses)(unsafe.Pointer(&b[0])), &l)
		if err == nil {
			if l == 0 {
				return nil, nil
			}
			break
		}
		if err.(syscall.Errno) != syscall.ERROR_BUFFER_OVERFLOW {
			return nil, os.NewSyscallError("getadaptersaddresses", err)
		}
		if l <= uint32(len(b)) {
			return nil, os.NewSyscallError("getadaptersaddresses", err)
		}
	}
	var aas []*windows.IpAdapterAddresses
	for aa := (*windows.IpAdapterAddresses)(unsafe.Pointer(&b[0])); aa != nil; aa = aa.Next {
		aas = append(aas, aa)
	}
	return aas, nil
}

func bytePtrToString(p *uint8) string {
	a := (*[10000]uint8)(unsafe.Pointer(p))
	i := 0
	for a[i] != 0 {
		i++
	}
	return string(a[:i])
}

func physicalAddrToString(physAddr [8]byte) string {
	if len(physAddr) == 0 {
		return ""
	}
	buf := make([]byte, 0, len(physAddr)*3-1)
	for i, b := range physAddr {
		if i > 0 {
			buf = append(buf, ':')
		}
		buf = append(buf, hexDigit[b>>4])
		buf = append(buf, hexDigit[b&0xF])
	}
	return string(buf)
}

func cStringToString(cs *uint16) (s string) {
	if cs != nil {
		us := make([]uint16, 0, 256)
		for p := uintptr(unsafe.Pointer(cs)); ; p += 2 {
			u := *(*uint16)(unsafe.Pointer(p))
			if u == 0 {
				return string(utf16.Decode(us))
			}
			us = append(us, u)
		}
	}
	return ""
}

// Gets all physical interfaces based on filter results, ignoring all VM, Loopback and Tunnel interfaces.
func getAllPhysicalInterface() []*windows.IpAdapterAddresses {
	aa, _ := adapterAddresses()

	var outInterfaces []*windows.IpAdapterAddresses

	for _, pa := range aa {
		mac := physicalAddrToString(pa.PhysicalAddress)
		name := "\\Device\\NPF_" + bytePtrToString(pa.AdapterName)

		if pa.IfType != uint32(IF_TYPE_SOFTWARE_LOOPBACK) && pa.IfType != uint32(IF_TYPE_TUNNEL) &&
			pa.OperStatus == uint32(IfOperStatusUp) && isPhysicalInterface(mac) {
			outInterfaces = append(outInterfaces, pa)
		}
	}

	return outInterfaces
}
