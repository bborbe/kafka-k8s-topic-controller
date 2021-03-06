// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"github.com/bborbe/kafka-k8s-topic-controller/k8s"
	"github.com/bborbe/kafka-k8s-topic-controller/topic"
	"k8s.io/client-go/tools/cache"
)

type K8sConnector struct {
	ListenStub        func(context.Context, cache.ResourceEventHandler) error
	listenMutex       sync.RWMutex
	listenArgsForCall []struct {
		arg1 context.Context
		arg2 cache.ResourceEventHandler
	}
	listenReturns struct {
		result1 error
	}
	listenReturnsOnCall map[int]struct {
		result1 error
	}
	SetupCustomResourceDefinitionStub        func() error
	setupCustomResourceDefinitionMutex       sync.RWMutex
	setupCustomResourceDefinitionArgsForCall []struct {
	}
	setupCustomResourceDefinitionReturns struct {
		result1 error
	}
	setupCustomResourceDefinitionReturnsOnCall map[int]struct {
		result1 error
	}
	TopicsStub        func() ([]topic.Topic, error)
	topicsMutex       sync.RWMutex
	topicsArgsForCall []struct {
	}
	topicsReturns struct {
		result1 []topic.Topic
		result2 error
	}
	topicsReturnsOnCall map[int]struct {
		result1 []topic.Topic
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *K8sConnector) Listen(arg1 context.Context, arg2 cache.ResourceEventHandler) error {
	fake.listenMutex.Lock()
	ret, specificReturn := fake.listenReturnsOnCall[len(fake.listenArgsForCall)]
	fake.listenArgsForCall = append(fake.listenArgsForCall, struct {
		arg1 context.Context
		arg2 cache.ResourceEventHandler
	}{arg1, arg2})
	fake.recordInvocation("Listen", []interface{}{arg1, arg2})
	fake.listenMutex.Unlock()
	if fake.ListenStub != nil {
		return fake.ListenStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.listenReturns
	return fakeReturns.result1
}

func (fake *K8sConnector) ListenCallCount() int {
	fake.listenMutex.RLock()
	defer fake.listenMutex.RUnlock()
	return len(fake.listenArgsForCall)
}

func (fake *K8sConnector) ListenCalls(stub func(context.Context, cache.ResourceEventHandler) error) {
	fake.listenMutex.Lock()
	defer fake.listenMutex.Unlock()
	fake.ListenStub = stub
}

func (fake *K8sConnector) ListenArgsForCall(i int) (context.Context, cache.ResourceEventHandler) {
	fake.listenMutex.RLock()
	defer fake.listenMutex.RUnlock()
	argsForCall := fake.listenArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *K8sConnector) ListenReturns(result1 error) {
	fake.listenMutex.Lock()
	defer fake.listenMutex.Unlock()
	fake.ListenStub = nil
	fake.listenReturns = struct {
		result1 error
	}{result1}
}

func (fake *K8sConnector) ListenReturnsOnCall(i int, result1 error) {
	fake.listenMutex.Lock()
	defer fake.listenMutex.Unlock()
	fake.ListenStub = nil
	if fake.listenReturnsOnCall == nil {
		fake.listenReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.listenReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *K8sConnector) SetupCustomResourceDefinition() error {
	fake.setupCustomResourceDefinitionMutex.Lock()
	ret, specificReturn := fake.setupCustomResourceDefinitionReturnsOnCall[len(fake.setupCustomResourceDefinitionArgsForCall)]
	fake.setupCustomResourceDefinitionArgsForCall = append(fake.setupCustomResourceDefinitionArgsForCall, struct {
	}{})
	fake.recordInvocation("SetupCustomResourceDefinition", []interface{}{})
	fake.setupCustomResourceDefinitionMutex.Unlock()
	if fake.SetupCustomResourceDefinitionStub != nil {
		return fake.SetupCustomResourceDefinitionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setupCustomResourceDefinitionReturns
	return fakeReturns.result1
}

func (fake *K8sConnector) SetupCustomResourceDefinitionCallCount() int {
	fake.setupCustomResourceDefinitionMutex.RLock()
	defer fake.setupCustomResourceDefinitionMutex.RUnlock()
	return len(fake.setupCustomResourceDefinitionArgsForCall)
}

func (fake *K8sConnector) SetupCustomResourceDefinitionCalls(stub func() error) {
	fake.setupCustomResourceDefinitionMutex.Lock()
	defer fake.setupCustomResourceDefinitionMutex.Unlock()
	fake.SetupCustomResourceDefinitionStub = stub
}

func (fake *K8sConnector) SetupCustomResourceDefinitionReturns(result1 error) {
	fake.setupCustomResourceDefinitionMutex.Lock()
	defer fake.setupCustomResourceDefinitionMutex.Unlock()
	fake.SetupCustomResourceDefinitionStub = nil
	fake.setupCustomResourceDefinitionReturns = struct {
		result1 error
	}{result1}
}

func (fake *K8sConnector) SetupCustomResourceDefinitionReturnsOnCall(i int, result1 error) {
	fake.setupCustomResourceDefinitionMutex.Lock()
	defer fake.setupCustomResourceDefinitionMutex.Unlock()
	fake.SetupCustomResourceDefinitionStub = nil
	if fake.setupCustomResourceDefinitionReturnsOnCall == nil {
		fake.setupCustomResourceDefinitionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setupCustomResourceDefinitionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *K8sConnector) Topics() ([]topic.Topic, error) {
	fake.topicsMutex.Lock()
	ret, specificReturn := fake.topicsReturnsOnCall[len(fake.topicsArgsForCall)]
	fake.topicsArgsForCall = append(fake.topicsArgsForCall, struct {
	}{})
	fake.recordInvocation("Topics", []interface{}{})
	fake.topicsMutex.Unlock()
	if fake.TopicsStub != nil {
		return fake.TopicsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.topicsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *K8sConnector) TopicsCallCount() int {
	fake.topicsMutex.RLock()
	defer fake.topicsMutex.RUnlock()
	return len(fake.topicsArgsForCall)
}

func (fake *K8sConnector) TopicsCalls(stub func() ([]topic.Topic, error)) {
	fake.topicsMutex.Lock()
	defer fake.topicsMutex.Unlock()
	fake.TopicsStub = stub
}

func (fake *K8sConnector) TopicsReturns(result1 []topic.Topic, result2 error) {
	fake.topicsMutex.Lock()
	defer fake.topicsMutex.Unlock()
	fake.TopicsStub = nil
	fake.topicsReturns = struct {
		result1 []topic.Topic
		result2 error
	}{result1, result2}
}

func (fake *K8sConnector) TopicsReturnsOnCall(i int, result1 []topic.Topic, result2 error) {
	fake.topicsMutex.Lock()
	defer fake.topicsMutex.Unlock()
	fake.TopicsStub = nil
	if fake.topicsReturnsOnCall == nil {
		fake.topicsReturnsOnCall = make(map[int]struct {
			result1 []topic.Topic
			result2 error
		})
	}
	fake.topicsReturnsOnCall[i] = struct {
		result1 []topic.Topic
		result2 error
	}{result1, result2}
}

func (fake *K8sConnector) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listenMutex.RLock()
	defer fake.listenMutex.RUnlock()
	fake.setupCustomResourceDefinitionMutex.RLock()
	defer fake.setupCustomResourceDefinitionMutex.RUnlock()
	fake.topicsMutex.RLock()
	defer fake.topicsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *K8sConnector) recordInvocation(key string, args []interface{}) {
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

var _ k8s.Connector = new(K8sConnector)
