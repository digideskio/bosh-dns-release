// This file was generated by counterfeiter
package internalfakes

import (
	"net"
	"sync"

	"github.com/miekg/dns"
)

type FakeResponseWriter struct {
	LocalAddrStub        func() net.Addr
	localAddrMutex       sync.RWMutex
	localAddrArgsForCall []struct{}
	localAddrReturns     struct {
		result1 net.Addr
	}
	localAddrReturnsOnCall map[int]struct {
		result1 net.Addr
	}
	RemoteAddrStub        func() net.Addr
	remoteAddrMutex       sync.RWMutex
	remoteAddrArgsForCall []struct{}
	remoteAddrReturns     struct {
		result1 net.Addr
	}
	remoteAddrReturnsOnCall map[int]struct {
		result1 net.Addr
	}
	WriteMsgStub        func(*dns.Msg) error
	writeMsgMutex       sync.RWMutex
	writeMsgArgsForCall []struct {
		arg1 *dns.Msg
	}
	writeMsgReturns struct {
		result1 error
	}
	writeMsgReturnsOnCall map[int]struct {
		result1 error
	}
	WriteStub        func([]byte) (int, error)
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		arg1 []byte
	}
	writeReturns struct {
		result1 int
		result2 error
	}
	writeReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	TsigStatusStub        func() error
	tsigStatusMutex       sync.RWMutex
	tsigStatusArgsForCall []struct{}
	tsigStatusReturns     struct {
		result1 error
	}
	tsigStatusReturnsOnCall map[int]struct {
		result1 error
	}
	TsigTimersOnlyStub        func(bool)
	tsigTimersOnlyMutex       sync.RWMutex
	tsigTimersOnlyArgsForCall []struct {
		arg1 bool
	}
	HijackStub        func()
	hijackMutex       sync.RWMutex
	hijackArgsForCall []struct{}
	invocations       map[string][][]interface{}
	invocationsMutex  sync.RWMutex
}

func (fake *FakeResponseWriter) LocalAddr() net.Addr {
	fake.localAddrMutex.Lock()
	ret, specificReturn := fake.localAddrReturnsOnCall[len(fake.localAddrArgsForCall)]
	fake.localAddrArgsForCall = append(fake.localAddrArgsForCall, struct{}{})
	fake.recordInvocation("LocalAddr", []interface{}{})
	fake.localAddrMutex.Unlock()
	if fake.LocalAddrStub != nil {
		return fake.LocalAddrStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.localAddrReturns.result1
}

func (fake *FakeResponseWriter) LocalAddrCallCount() int {
	fake.localAddrMutex.RLock()
	defer fake.localAddrMutex.RUnlock()
	return len(fake.localAddrArgsForCall)
}

func (fake *FakeResponseWriter) LocalAddrReturns(result1 net.Addr) {
	fake.LocalAddrStub = nil
	fake.localAddrReturns = struct {
		result1 net.Addr
	}{result1}
}

func (fake *FakeResponseWriter) LocalAddrReturnsOnCall(i int, result1 net.Addr) {
	fake.LocalAddrStub = nil
	if fake.localAddrReturnsOnCall == nil {
		fake.localAddrReturnsOnCall = make(map[int]struct {
			result1 net.Addr
		})
	}
	fake.localAddrReturnsOnCall[i] = struct {
		result1 net.Addr
	}{result1}
}

func (fake *FakeResponseWriter) RemoteAddr() net.Addr {
	fake.remoteAddrMutex.Lock()
	ret, specificReturn := fake.remoteAddrReturnsOnCall[len(fake.remoteAddrArgsForCall)]
	fake.remoteAddrArgsForCall = append(fake.remoteAddrArgsForCall, struct{}{})
	fake.recordInvocation("RemoteAddr", []interface{}{})
	fake.remoteAddrMutex.Unlock()
	if fake.RemoteAddrStub != nil {
		return fake.RemoteAddrStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.remoteAddrReturns.result1
}

func (fake *FakeResponseWriter) RemoteAddrCallCount() int {
	fake.remoteAddrMutex.RLock()
	defer fake.remoteAddrMutex.RUnlock()
	return len(fake.remoteAddrArgsForCall)
}

func (fake *FakeResponseWriter) RemoteAddrReturns(result1 net.Addr) {
	fake.RemoteAddrStub = nil
	fake.remoteAddrReturns = struct {
		result1 net.Addr
	}{result1}
}

func (fake *FakeResponseWriter) RemoteAddrReturnsOnCall(i int, result1 net.Addr) {
	fake.RemoteAddrStub = nil
	if fake.remoteAddrReturnsOnCall == nil {
		fake.remoteAddrReturnsOnCall = make(map[int]struct {
			result1 net.Addr
		})
	}
	fake.remoteAddrReturnsOnCall[i] = struct {
		result1 net.Addr
	}{result1}
}

func (fake *FakeResponseWriter) WriteMsg(arg1 *dns.Msg) error {
	fake.writeMsgMutex.Lock()
	ret, specificReturn := fake.writeMsgReturnsOnCall[len(fake.writeMsgArgsForCall)]
	fake.writeMsgArgsForCall = append(fake.writeMsgArgsForCall, struct {
		arg1 *dns.Msg
	}{arg1})
	fake.recordInvocation("WriteMsg", []interface{}{arg1})
	fake.writeMsgMutex.Unlock()
	if fake.WriteMsgStub != nil {
		return fake.WriteMsgStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.writeMsgReturns.result1
}

func (fake *FakeResponseWriter) WriteMsgCallCount() int {
	fake.writeMsgMutex.RLock()
	defer fake.writeMsgMutex.RUnlock()
	return len(fake.writeMsgArgsForCall)
}

func (fake *FakeResponseWriter) WriteMsgArgsForCall(i int) *dns.Msg {
	fake.writeMsgMutex.RLock()
	defer fake.writeMsgMutex.RUnlock()
	return fake.writeMsgArgsForCall[i].arg1
}

func (fake *FakeResponseWriter) WriteMsgReturns(result1 error) {
	fake.WriteMsgStub = nil
	fake.writeMsgReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResponseWriter) WriteMsgReturnsOnCall(i int, result1 error) {
	fake.WriteMsgStub = nil
	if fake.writeMsgReturnsOnCall == nil {
		fake.writeMsgReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeMsgReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResponseWriter) Write(arg1 []byte) (int, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.writeMutex.Lock()
	ret, specificReturn := fake.writeReturnsOnCall[len(fake.writeArgsForCall)]
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("Write", []interface{}{arg1Copy})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		return fake.WriteStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.writeReturns.result1, fake.writeReturns.result2
}

func (fake *FakeResponseWriter) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *FakeResponseWriter) WriteArgsForCall(i int) []byte {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return fake.writeArgsForCall[i].arg1
}

func (fake *FakeResponseWriter) WriteReturns(result1 int, result2 error) {
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeResponseWriter) WriteReturnsOnCall(i int, result1 int, result2 error) {
	fake.WriteStub = nil
	if fake.writeReturnsOnCall == nil {
		fake.writeReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.writeReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeResponseWriter) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.closeReturns.result1
}

func (fake *FakeResponseWriter) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeResponseWriter) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResponseWriter) CloseReturnsOnCall(i int, result1 error) {
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResponseWriter) TsigStatus() error {
	fake.tsigStatusMutex.Lock()
	ret, specificReturn := fake.tsigStatusReturnsOnCall[len(fake.tsigStatusArgsForCall)]
	fake.tsigStatusArgsForCall = append(fake.tsigStatusArgsForCall, struct{}{})
	fake.recordInvocation("TsigStatus", []interface{}{})
	fake.tsigStatusMutex.Unlock()
	if fake.TsigStatusStub != nil {
		return fake.TsigStatusStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.tsigStatusReturns.result1
}

func (fake *FakeResponseWriter) TsigStatusCallCount() int {
	fake.tsigStatusMutex.RLock()
	defer fake.tsigStatusMutex.RUnlock()
	return len(fake.tsigStatusArgsForCall)
}

func (fake *FakeResponseWriter) TsigStatusReturns(result1 error) {
	fake.TsigStatusStub = nil
	fake.tsigStatusReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResponseWriter) TsigStatusReturnsOnCall(i int, result1 error) {
	fake.TsigStatusStub = nil
	if fake.tsigStatusReturnsOnCall == nil {
		fake.tsigStatusReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.tsigStatusReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResponseWriter) TsigTimersOnly(arg1 bool) {
	fake.tsigTimersOnlyMutex.Lock()
	fake.tsigTimersOnlyArgsForCall = append(fake.tsigTimersOnlyArgsForCall, struct {
		arg1 bool
	}{arg1})
	fake.recordInvocation("TsigTimersOnly", []interface{}{arg1})
	fake.tsigTimersOnlyMutex.Unlock()
	if fake.TsigTimersOnlyStub != nil {
		fake.TsigTimersOnlyStub(arg1)
	}
}

func (fake *FakeResponseWriter) TsigTimersOnlyCallCount() int {
	fake.tsigTimersOnlyMutex.RLock()
	defer fake.tsigTimersOnlyMutex.RUnlock()
	return len(fake.tsigTimersOnlyArgsForCall)
}

func (fake *FakeResponseWriter) TsigTimersOnlyArgsForCall(i int) bool {
	fake.tsigTimersOnlyMutex.RLock()
	defer fake.tsigTimersOnlyMutex.RUnlock()
	return fake.tsigTimersOnlyArgsForCall[i].arg1
}

func (fake *FakeResponseWriter) Hijack() {
	fake.hijackMutex.Lock()
	fake.hijackArgsForCall = append(fake.hijackArgsForCall, struct{}{})
	fake.recordInvocation("Hijack", []interface{}{})
	fake.hijackMutex.Unlock()
	if fake.HijackStub != nil {
		fake.HijackStub()
	}
}

func (fake *FakeResponseWriter) HijackCallCount() int {
	fake.hijackMutex.RLock()
	defer fake.hijackMutex.RUnlock()
	return len(fake.hijackArgsForCall)
}

func (fake *FakeResponseWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.localAddrMutex.RLock()
	defer fake.localAddrMutex.RUnlock()
	fake.remoteAddrMutex.RLock()
	defer fake.remoteAddrMutex.RUnlock()
	fake.writeMsgMutex.RLock()
	defer fake.writeMsgMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.tsigStatusMutex.RLock()
	defer fake.tsigStatusMutex.RUnlock()
	fake.tsigTimersOnlyMutex.RLock()
	defer fake.tsigTimersOnlyMutex.RUnlock()
	fake.hijackMutex.RLock()
	defer fake.hijackMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeResponseWriter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ dns.ResponseWriter = new(FakeResponseWriter)
