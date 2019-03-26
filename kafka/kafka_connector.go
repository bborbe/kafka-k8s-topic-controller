// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/bborbe/kafka-k8s-topic-controller/topic"
	"github.com/golang/glog"
	"github.com/pkg/errors"
)

//go:generate counterfeiter -o ../mocks/kafka-clusteradmin.go --fake-name KafkaClusterAdmin . ClusterAdmin
type ClusterAdmin interface {
	ListTopics() (map[string]sarama.TopicDetail, error)
	DeleteTopic(topic string) error
	CreateTopic(topic string, detail *sarama.TopicDetail, validateOnly bool) error
}

//go:generate counterfeiter -o ../mocks/kafka-connector.go --fake-name KafkaConnector . Connector
type Connector interface {
	CreateTopic(topic topic.Topic) error
	DeleteTopic(topic topic.Topic) error
	UpdateTopic(oldTopic topic.Topic, newTopic topic.Topic) error
	Topics() ([]topic.Topic, error)
}

func NewConnector(clusterAdmin ClusterAdmin) Connector {
	return &connector{
		clusterAdmin: clusterAdmin,
	}
}

type connector struct {
	clusterAdmin ClusterAdmin
}

func (c *connector) Topics() ([]topic.Topic, error) {
	topicDetails, err := c.clusterAdmin.ListTopics()
	if err != nil {
		return nil, errors.Wrap(err, "list kafka topics failed")
	}
	var result []topic.Topic
	for name, data := range topicDetails {
		result = append(result, topic.Topic{
			Name:              name,
			NumPartitions:     data.NumPartitions,
			ReplicationFactor: data.ReplicationFactor,
		})
	}
	return result, nil
}

func (c *connector) CreateTopic(topic topic.Topic) error {
	glog.V(2).Infof("create topic %s", topic.Name)
	err := c.clusterAdmin.CreateTopic(
		topic.Name,
		&sarama.TopicDetail{
			NumPartitions:     topic.NumPartitions,
			ReplicationFactor: topic.ReplicationFactor,
			ReplicaAssignment: make(map[int32][]int32),
			ConfigEntries:     make(map[string]*string),
		},
		false,
	)
	if err != nil {
		return errors.Wrapf(err, "create topic %s failed", topic.Name)
	}
	glog.V(1).Infof("topic %s created", topic.Name)
	return nil
}

func (c *connector) UpdateTopic(oldTopic topic.Topic, newTopic topic.Topic) error {
	if oldTopic.Equals(newTopic) {
		glog.V(2).Infof("topic unchanged => skip")
		return nil
	}
	glog.V(2).Infof("update topic %s", newTopic.Name)
	if err := c.DeleteTopic(oldTopic); err != nil {
		glog.Warningf("delete topic failed: %v", err)
	}
	if err := c.CreateTopic(newTopic); err != nil {
		glog.Warningf("create topic failed: %v", err)
	}
	return nil
}

func (c *connector) DeleteTopic(topic topic.Topic) error {
	glog.V(2).Infof("delete topic %s", topic.Name)
	if err := c.clusterAdmin.DeleteTopic(topic.Name); err != nil {
		return errors.Wrapf(err, "delete topic %s failed", topic.Name)
	}
	glog.V(1).Infof("topic %s deleted", topic.Name)
	return nil
}
