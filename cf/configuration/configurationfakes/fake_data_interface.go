// This file was generated by counterfeiter
package configurationfakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/configuration"
)

type FakeDataInterface struct {
	JsonMarshalV3Stub        func() ([]byte, error)
	jsonMarshalV3Mutex       sync.RWMutex
	jsonMarshalV3ArgsForCall []struct{}
	jsonMarshalV3Returns     struct {
		result1 []byte
		result2 error
	}
	JsonUnmarshalV3Stub        func([]byte) error
	jsonUnmarshalV3Mutex       sync.RWMutex
	jsonUnmarshalV3ArgsForCall []struct {
		arg1 []byte
	}
	jsonUnmarshalV3Returns struct {
		result1 error
	}
}

func (fake *FakeDataInterface) JsonMarshalV3() ([]byte, error) {
	fake.jsonMarshalV3Mutex.Lock()
	fake.jsonMarshalV3ArgsForCall = append(fake.jsonMarshalV3ArgsForCall, struct{}{})
	fake.jsonMarshalV3Mutex.Unlock()
	if fake.JsonMarshalV3Stub != nil {
		return fake.JsonMarshalV3Stub()
	} else {
		return fake.jsonMarshalV3Returns.result1, fake.jsonMarshalV3Returns.result2
	}
}

func (fake *FakeDataInterface) JsonMarshalV3CallCount() int {
	fake.jsonMarshalV3Mutex.RLock()
	defer fake.jsonMarshalV3Mutex.RUnlock()
	return len(fake.jsonMarshalV3ArgsForCall)
}

func (fake *FakeDataInterface) JsonMarshalV3Returns(result1 []byte, result2 error) {
	fake.JsonMarshalV3Stub = nil
	fake.jsonMarshalV3Returns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeDataInterface) JsonUnmarshalV3(arg1 []byte) error {
	fake.jsonUnmarshalV3Mutex.Lock()
	fake.jsonUnmarshalV3ArgsForCall = append(fake.jsonUnmarshalV3ArgsForCall, struct {
		arg1 []byte
	}{arg1})
	fake.jsonUnmarshalV3Mutex.Unlock()
	if fake.JsonUnmarshalV3Stub != nil {
		return fake.JsonUnmarshalV3Stub(arg1)
	} else {
		return fake.jsonUnmarshalV3Returns.result1
	}
}

func (fake *FakeDataInterface) JsonUnmarshalV3CallCount() int {
	fake.jsonUnmarshalV3Mutex.RLock()
	defer fake.jsonUnmarshalV3Mutex.RUnlock()
	return len(fake.jsonUnmarshalV3ArgsForCall)
}

func (fake *FakeDataInterface) JsonUnmarshalV3ArgsForCall(i int) []byte {
	fake.jsonUnmarshalV3Mutex.RLock()
	defer fake.jsonUnmarshalV3Mutex.RUnlock()
	return fake.jsonUnmarshalV3ArgsForCall[i].arg1
}

func (fake *FakeDataInterface) JsonUnmarshalV3Returns(result1 error) {
	fake.JsonUnmarshalV3Stub = nil
	fake.jsonUnmarshalV3Returns = struct {
		result1 error
	}{result1}
}

var _ configuration.DataInterface = new(FakeDataInterface)
