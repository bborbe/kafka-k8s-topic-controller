// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package topic_test

import (
	"github.com/bborbe/kafka-k8s-topic-controller/topic"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Topic", func() {
	var topicA topic.Topic
	var topicB topic.Topic
	BeforeEach(func() {
		topicA = topic.Topic{
			Name:              "a",
			NumPartitions:     1,
			ReplicationFactor: 2,
		}
		topicB = topic.Topic{
			Name:              "a",
			NumPartitions:     1,
			ReplicationFactor: 2,
		}
	})

	It("is equal", func() {
		Expect(topicA.Equals(topicB)).To(BeTrue())
	})
	It("is not equal if name is different", func() {
		topicB.Name = "b"
		Expect(topicA.Equals(topicB)).To(BeFalse())
	})
	It("is not equal if partitions is different", func() {
		topicB.NumPartitions = 3
		Expect(topicA.Equals(topicB)).To(BeFalse())
	})
	It("is not equal if replicas is different", func() {
		topicB.ReplicationFactor = 3
		Expect(topicA.Equals(topicB)).To(BeFalse())
	})
})
