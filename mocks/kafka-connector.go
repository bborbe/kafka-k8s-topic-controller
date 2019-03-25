// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/bborbe/kafka-k8s-topic-controller/kafka"
	"github.com/bborbe/kafka-k8s-topic-controller/topic"
)

type KafkaConnector struct {
	CreateTopicStub        func(topic.Topic) error
	createTopicMutex       sync.RWMutex
	createTopicArgsForCall []struct {
		arg1 topic.Topic
	}
	createTopicReturns struct {
		result1 error
	}
	createTopicReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteTopicStub        func(topic.Topic) error
	deleteTopicMutex       sync.RWMutex
	deleteTopicArgsForCall []struct {
		arg1 topic.Topic
	}
	deleteTopicReturns struct {
		result1 error
	}
	deleteTopicReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateTopicStub        func(topic.Topic, topic.Topic) error
	updateTopicMutex       sync.RWMutex
	updateTopicArgsForCall []struct {
		arg1 topic.Topic
		arg2 topic.Topic
	}
	updateTopicReturns struct {
		result1 error
	}
	updateTopicReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *KafkaConnector) CreateTopic(arg1 topic.Topic) error {
	fake.createTopicMutex.Lock()
	ret, specificReturn := fake.createTopicReturnsOnCall[len(fake.createTopicArgsForCall)]
	fake.createTopicArgsForCall = append(fake.createTopicArgsForCall, struct {
		arg1 topic.Topic
	}{arg1})
	fake.recordInvocation("CreateTopic", []interface{}{arg1})
	fake.createTopicMutex.Unlock()
	if fake.CreateTopicStub != nil {
		return fake.CreateTopicStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createTopicReturns
	return fakeReturns.result1
}

func (fake *KafkaConnector) CreateTopicCallCount() int {
	fake.createTopicMutex.RLock()
	defer fake.createTopicMutex.RUnlock()
	return len(fake.createTopicArgsForCall)
}

func (fake *KafkaConnector) CreateTopicCalls(stub func(topic.Topic) error) {
	fake.createTopicMutex.Lock()
	defer fake.createTopicMutex.Unlock()
	fake.CreateTopicStub = stub
}

func (fake *KafkaConnector) CreateTopicArgsForCall(i int) topic.Topic {
	fake.createTopicMutex.RLock()
	defer fake.createTopicMutex.RUnlock()
	argsForCall := fake.createTopicArgsForCall[i]
	return argsForCall.arg1
}

func (fake *KafkaConnector) CreateTopicReturns(result1 error) {
	fake.createTopicMutex.Lock()
	defer fake.createTopicMutex.Unlock()
	fake.CreateTopicStub = nil
	fake.createTopicReturns = struct {
		result1 error
	}{result1}
}

func (fake *KafkaConnector) CreateTopicReturnsOnCall(i int, result1 error) {
	fake.createTopicMutex.Lock()
	defer fake.createTopicMutex.Unlock()
	fake.CreateTopicStub = nil
	if fake.createTopicReturnsOnCall == nil {
		fake.createTopicReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createTopicReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *KafkaConnector) DeleteTopic(arg1 topic.Topic) error {
	fake.deleteTopicMutex.Lock()
	ret, specificReturn := fake.deleteTopicReturnsOnCall[len(fake.deleteTopicArgsForCall)]
	fake.deleteTopicArgsForCall = append(fake.deleteTopicArgsForCall, struct {
		arg1 topic.Topic
	}{arg1})
	fake.recordInvocation("DeleteTopic", []interface{}{arg1})
	fake.deleteTopicMutex.Unlock()
	if fake.DeleteTopicStub != nil {
		return fake.DeleteTopicStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteTopicReturns
	return fakeReturns.result1
}

func (fake *KafkaConnector) DeleteTopicCallCount() int {
	fake.deleteTopicMutex.RLock()
	defer fake.deleteTopicMutex.RUnlock()
	return len(fake.deleteTopicArgsForCall)
}

func (fake *KafkaConnector) DeleteTopicCalls(stub func(topic.Topic) error) {
	fake.deleteTopicMutex.Lock()
	defer fake.deleteTopicMutex.Unlock()
	fake.DeleteTopicStub = stub
}

func (fake *KafkaConnector) DeleteTopicArgsForCall(i int) topic.Topic {
	fake.deleteTopicMutex.RLock()
	defer fake.deleteTopicMutex.RUnlock()
	argsForCall := fake.deleteTopicArgsForCall[i]
	return argsForCall.arg1
}

func (fake *KafkaConnector) DeleteTopicReturns(result1 error) {
	fake.deleteTopicMutex.Lock()
	defer fake.deleteTopicMutex.Unlock()
	fake.DeleteTopicStub = nil
	fake.deleteTopicReturns = struct {
		result1 error
	}{result1}
}

func (fake *KafkaConnector) DeleteTopicReturnsOnCall(i int, result1 error) {
	fake.deleteTopicMutex.Lock()
	defer fake.deleteTopicMutex.Unlock()
	fake.DeleteTopicStub = nil
	if fake.deleteTopicReturnsOnCall == nil {
		fake.deleteTopicReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteTopicReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *KafkaConnector) UpdateTopic(arg1 topic.Topic, arg2 topic.Topic) error {
	fake.updateTopicMutex.Lock()
	ret, specificReturn := fake.updateTopicReturnsOnCall[len(fake.updateTopicArgsForCall)]
	fake.updateTopicArgsForCall = append(fake.updateTopicArgsForCall, struct {
		arg1 topic.Topic
		arg2 topic.Topic
	}{arg1, arg2})
	fake.recordInvocation("UpdateTopic", []interface{}{arg1, arg2})
	fake.updateTopicMutex.Unlock()
	if fake.UpdateTopicStub != nil {
		return fake.UpdateTopicStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateTopicReturns
	return fakeReturns.result1
}

func (fake *KafkaConnector) UpdateTopicCallCount() int {
	fake.updateTopicMutex.RLock()
	defer fake.updateTopicMutex.RUnlock()
	return len(fake.updateTopicArgsForCall)
}

func (fake *KafkaConnector) UpdateTopicCalls(stub func(topic.Topic, topic.Topic) error) {
	fake.updateTopicMutex.Lock()
	defer fake.updateTopicMutex.Unlock()
	fake.UpdateTopicStub = stub
}

func (fake *KafkaConnector) UpdateTopicArgsForCall(i int) (topic.Topic, topic.Topic) {
	fake.updateTopicMutex.RLock()
	defer fake.updateTopicMutex.RUnlock()
	argsForCall := fake.updateTopicArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *KafkaConnector) UpdateTopicReturns(result1 error) {
	fake.updateTopicMutex.Lock()
	defer fake.updateTopicMutex.Unlock()
	fake.UpdateTopicStub = nil
	fake.updateTopicReturns = struct {
		result1 error
	}{result1}
}

func (fake *KafkaConnector) UpdateTopicReturnsOnCall(i int, result1 error) {
	fake.updateTopicMutex.Lock()
	defer fake.updateTopicMutex.Unlock()
	fake.UpdateTopicStub = nil
	if fake.updateTopicReturnsOnCall == nil {
		fake.updateTopicReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateTopicReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *KafkaConnector) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createTopicMutex.RLock()
	defer fake.createTopicMutex.RUnlock()
	fake.deleteTopicMutex.RLock()
	defer fake.deleteTopicMutex.RUnlock()
	fake.updateTopicMutex.RLock()
	defer fake.updateTopicMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *KafkaConnector) recordInvocation(key string, args []interface{}) {
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

var _ kafka.Connector = new(KafkaConnector)
