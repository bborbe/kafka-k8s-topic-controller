// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"github.com/bborbe/kafka-k8s-topic-controller/k8s"
	v1 "github.com/bborbe/kafka-k8s-topic-controller/k8s/apis/kafka.benjamin-borbe.de/v1"
	"github.com/bborbe/kafka-k8s-topic-controller/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/client-go/tools/cache"
)

var _ = Describe("ResourceEventHandler", func() {
	var resourceEventHandler cache.ResourceEventHandler
	var eventHandler *mocks.K8sEventHandler
	BeforeEach(func() {
		eventHandler = &mocks.K8sEventHandler{}
		resourceEventHandler = k8s.NewResourceEventHandler(eventHandler)
	})

	It("create topic OnAdd", func() {
		resourceEventHandler.OnAdd(&v1.Topic{
			Spec: v1.TopicSpec{
				Name:              "my-topic",
				NumPartitions:     2,
				ReplicationFactor: 3,
			},
		})
		Expect(eventHandler.CreateTopicCallCount()).To(Equal(1))
		topic := eventHandler.CreateTopicArgsForCall(0)
		Expect(topic.Name).To(Equal("my-topic"))
		Expect(topic.NumPartitions).To(Equal(int32(2)))
		Expect(topic.ReplicationFactor).To(Equal(int16(3)))
	})
	It("delete topic OnDelete", func() {
		resourceEventHandler.OnDelete(&v1.Topic{
			Spec: v1.TopicSpec{
				Name:              "my-topic",
				NumPartitions:     2,
				ReplicationFactor: 3,
			},
		})
		Expect(eventHandler.DeleteTopicCallCount()).To(Equal(1))
		topic := eventHandler.DeleteTopicArgsForCall(0)
		Expect(topic.Name).To(Equal("my-topic"))
		Expect(topic.NumPartitions).To(Equal(int32(2)))
		Expect(topic.ReplicationFactor).To(Equal(int16(3)))
	})
	It("update topic OnUpdate", func() {
		resourceEventHandler.OnUpdate(&v1.Topic{
			Spec: v1.TopicSpec{
				Name:              "old-topic",
				NumPartitions:     1,
				ReplicationFactor: 2,
			},
		}, &v1.Topic{
			Spec: v1.TopicSpec{
				Name:              "new-topic",
				NumPartitions:     3,
				ReplicationFactor: 4,
			},
		})
		Expect(eventHandler.UpdateTopicCallCount()).To(Equal(1))
		topicOld, topicNew := eventHandler.UpdateTopicArgsForCall(0)
		Expect(topicOld.Name).To(Equal("old-topic"))
		Expect(topicOld.NumPartitions).To(Equal(int32(1)))
		Expect(topicOld.ReplicationFactor).To(Equal(int16(2)))
		Expect(topicNew.Name).To(Equal("new-topic"))
		Expect(topicNew.NumPartitions).To(Equal(int32(3)))
		Expect(topicNew.ReplicationFactor).To(Equal(int16(4)))
	})
})
