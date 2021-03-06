package manager

import (
	"path/filepath"

	"strings"

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
	runner         boshsys.CmdRunner
	fs             boshsys.FileSystem
	adapterFetcher AdapterFetcher
}

//go:generate counterfeiter . AdapterFetcher

type AdapterFetcher interface {
	Adapters() ([]Adapter, error)
}

func NewWindowsManager(runner boshsys.CmdRunner, fs boshsys.FileSystem, adapterFetcher AdapterFetcher) *windowsManager {
	return &windowsManager{runner: runner, fs: fs, adapterFetcher: adapterFetcher}
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
	var servers []string

	adapters, err := manager.getPhysicalAdapters()
	if err != nil {
		return nil, bosherr.WrapError(err, "Getting list of current DNS Servers")
	}

	for _, adapter := range adapters {
		servers = append(servers, adapter.DNSServerAddresses...)
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

const (
	IfOperStatusUp         uint32 = 1
	IfTypeSoftwareLoopback uint32 = 24
	IfTypeTunnel           uint32 = 131
)

func (manager *windowsManager) getPhysicalAdapters() ([]Adapter, error) {
	var result []Adapter

	adapters, err := manager.adapterFetcher.Adapters()
	if err != nil {
		return nil, err
	}
	for _, adapter := range adapters {
		if adapter.IfType != IfTypeSoftwareLoopback && adapter.IfType != IfTypeTunnel &&
			adapter.OperStatus == IfOperStatusUp && adapter.isPhysicalInterface() {
			result = append(result, adapter)
		}
	}

	return result, nil
}

type Adapter struct {
	IfType             uint32
	OperStatus         uint32
	PhysicalAddress    string
	DNSServerAddresses []string
}

// Mac Address parts to look for, and identify non physical devices. There may be more, update me!
var macAddrPartsToFilter []string = []string{
	"00:03:FF",       // Microsoft Hyper-V, Virtual Server, Virtual PC
	"0A:00:27",       // VirtualBox
	"00:00:00:00:00", // Teredo Tunneling Pseudo-Interface
	"00:50:56",       // VMware ESX 3, Server, Workstation, Player
	"00:1C:14",       // VMware ESX 3, Server, Workstation, Player
	"00:0C:29",       // VMware ESX 3, Server, Workstation, Player
	"00:05:69",       // VMware ESX 3, Server, Workstation, Player
	"00:1C:42",       // Microsoft Hyper-V, Virtual Server, Virtual PC
	"00:0F:4B",       // Virtual Iron 4
	"00:16:3E",       // Red Hat Xen, Oracle VM, XenSource, Novell Xen
	"08:00:27",       // Sun xVM VirtualBox
}

// Filters the possible physical interface address by comparing it to known popular VM Software adresses
// and Teredo Tunneling Pseudo-Interface.
func (a Adapter) isPhysicalInterface() bool {
	for _, macPart := range macAddrPartsToFilter {
		if strings.HasPrefix(strings.ToLower(a.PhysicalAddress), strings.ToLower(macPart)) {
			return false
		}
	}

	return true
}
