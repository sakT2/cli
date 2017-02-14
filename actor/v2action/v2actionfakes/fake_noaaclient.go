// This file was generated by counterfeiter
package v2actionfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"github.com/cloudfoundry/sonde-go/events"
)

type FakeNOAAClient struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	TailingLogsStub        func(appGuid, authToken string) (<-chan *events.LogMessage, <-chan error)
	tailingLogsMutex       sync.RWMutex
	tailingLogsArgsForCall []struct {
		appGuid   string
		authToken string
	}
	tailingLogsReturns struct {
		result1 <-chan *events.LogMessage
		result2 <-chan error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNOAAClient) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	} else {
		return fake.closeReturns.result1
	}
}

func (fake *FakeNOAAClient) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeNOAAClient) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNOAAClient) TailingLogs(appGuid string, authToken string) (<-chan *events.LogMessage, <-chan error) {
	fake.tailingLogsMutex.Lock()
	fake.tailingLogsArgsForCall = append(fake.tailingLogsArgsForCall, struct {
		appGuid   string
		authToken string
	}{appGuid, authToken})
	fake.recordInvocation("TailingLogs", []interface{}{appGuid, authToken})
	fake.tailingLogsMutex.Unlock()
	if fake.TailingLogsStub != nil {
		return fake.TailingLogsStub(appGuid, authToken)
	} else {
		return fake.tailingLogsReturns.result1, fake.tailingLogsReturns.result2
	}
}

func (fake *FakeNOAAClient) TailingLogsCallCount() int {
	fake.tailingLogsMutex.RLock()
	defer fake.tailingLogsMutex.RUnlock()
	return len(fake.tailingLogsArgsForCall)
}

func (fake *FakeNOAAClient) TailingLogsArgsForCall(i int) (string, string) {
	fake.tailingLogsMutex.RLock()
	defer fake.tailingLogsMutex.RUnlock()
	return fake.tailingLogsArgsForCall[i].appGuid, fake.tailingLogsArgsForCall[i].authToken
}

func (fake *FakeNOAAClient) TailingLogsReturns(result1 <-chan *events.LogMessage, result2 <-chan error) {
	fake.TailingLogsStub = nil
	fake.tailingLogsReturns = struct {
		result1 <-chan *events.LogMessage
		result2 <-chan error
	}{result1, result2}
}

func (fake *FakeNOAAClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.tailingLogsMutex.RLock()
	defer fake.tailingLogsMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeNOAAClient) recordInvocation(key string, args []interface{}) {
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

var _ v2action.NOAAClient = new(FakeNOAAClient)