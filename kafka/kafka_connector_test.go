// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kafka_test

import (
	"errors"

	"github.com/Shopify/sarama"
	"github.com/bborbe/kafka-k8s-topic-controller/kafka"
	"github.com/bborbe/kafka-k8s-topic-controller/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Kafak Connector", func() {
	var kafkaConnector kafka.Connector
	var clusterAdmin *mocks.KafkaClusterAdmin
	BeforeEach(func() {
		clusterAdmin = &mocks.KafkaClusterAdmin{}
		kafkaConnector = kafka.NewConnector(clusterAdmin)
	})
	Context("Topcis", func() {
		It("returns no error", func() {
			clusterAdmin.ListTopicsReturns(map[string]sarama.TopicDetail{
				"my-topic": {
					NumPartitions:     2,
					ReplicationFactor: 3,
				},
			}, nil)
			topics, err := kafkaConnector.Topics()
			Expect(err).To(BeNil())
			Expect(topics).NotTo(BeNil())
			Expect(topics).To(HaveLen(1))
			Expect(topics[0].Name).To(Equal("my-topic"))
			Expect(topics[0].NumPartitions).To(Equal(int32(2)))
			Expect(topics[0].ReplicationFactor).To(Equal(int16(3)))
		})
		It("returns error if clusteradmin fails", func() {
			clusterAdmin.ListTopicsReturns(nil, errors.New("banana"))
			topics, err := kafkaConnector.Topics()
			Expect(err).NotTo(BeNil())
			Expect(topics).To(BeNil())
		})
	})
})
