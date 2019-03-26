// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package purge_test

import (
	"errors"

	"github.com/bborbe/kafka-k8s-topic-controller/mocks"
	"github.com/bborbe/kafka-k8s-topic-controller/purge"
	"github.com/bborbe/kafka-k8s-topic-controller/topic"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Purge", func() {
	var purger purge.Purger
	var k8sConnector *mocks.K8sConnector
	var kafkaConnector *mocks.KafkaConnector
	BeforeEach(func() {
		k8sConnector = &mocks.K8sConnector{}
		kafkaConnector = &mocks.KafkaConnector{}
		purger = purge.NewPurger(k8sConnector, kafkaConnector)
	})
	It("runs without error", func() {
		err := purger.Purge()
		Expect(err).To(BeNil())
	})
	It("returns an error if list k8s topics failed", func() {
		k8sConnector.TopicsReturns(nil, errors.New("banana"))
		err := purger.Purge()
		Expect(err).NotTo(BeNil())
	})
	It("returns an error if list kafka topics failed", func() {
		kafkaConnector.TopicsReturns(nil, errors.New("banana"))
		err := purger.Purge()
		Expect(err).NotTo(BeNil())
	})
	It("returns an error if delete topic fails", func() {
		kafkaConnector.TopicsReturns([]topic.Topic{
			{
				Name: "b",
			},
		}, nil)
		kafkaConnector.DeleteTopicReturns(errors.New("banana"))
		err := purger.Purge()
		Expect(err).NotTo(BeNil())
	})
	It("delete nothing if topics are equal", func() {
		k8sConnector.TopicsReturns([]topic.Topic{
			{
				Name: "a",
			},
		}, nil)
		kafkaConnector.TopicsReturns([]topic.Topic{
			{
				Name: "a",
			},
		}, nil)
		err := purger.Purge()
		Expect(err).To(BeNil())
		Expect(kafkaConnector.DeleteTopicCallCount()).To(Equal(0))
	})
	It("delete topic if it not exists", func() {
		k8sConnector.TopicsReturns([]topic.Topic{
			{
				Name: "a",
			},
		}, nil)
		kafkaConnector.TopicsReturns([]topic.Topic{
			{
				Name: "b",
			},
		}, nil)
		err := purger.Purge()
		Expect(err).To(BeNil())
		Expect(kafkaConnector.DeleteTopicCallCount()).To(Equal(1))
		deleteTopic := kafkaConnector.DeleteTopicArgsForCall(0)
		Expect(deleteTopic.Name).To(Equal("b"))
	})
	It("delete nothing if more topics exists in k8s", func() {
		k8sConnector.TopicsReturns([]topic.Topic{
			{
				Name: "a",
			},
			{
				Name: "b",
			},
		}, nil)
		kafkaConnector.TopicsReturns([]topic.Topic{
			{
				Name: "a",
			},
		}, nil)
		err := purger.Purge()
		Expect(err).To(BeNil())
		Expect(kafkaConnector.DeleteTopicCallCount()).To(Equal(0))
	})
	It("delete no topic start with unterscore", func() {
		k8sConnector.TopicsReturns([]topic.Topic{}, nil)
		kafkaConnector.TopicsReturns([]topic.Topic{
			{
				Name: "_system_topic",
			},
		}, nil)
		err := purger.Purge()
		Expect(err).To(BeNil())
		Expect(kafkaConnector.DeleteTopicCallCount()).To(Equal(0))
	})
})
