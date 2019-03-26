// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package purge

import (
	"strings"

	"github.com/bborbe/kafka-k8s-topic-controller/k8s"
	"github.com/bborbe/kafka-k8s-topic-controller/kafka"
	"github.com/golang/glog"
	"github.com/pkg/errors"
)

//go:generate counterfeiter -o ../mocks/purger.go --fake-name Purger . Purger
type Purger interface {
	Purge() error
}

func NewPurger(
	k8sConnector k8s.Connector,
	kafkaConnector kafka.Connector,
) Purger {
	return &purger{
		k8sConnector:   k8sConnector,
		kafkaConnector: kafkaConnector,
	}
}

type purger struct {
	k8sConnector   k8s.Connector
	kafkaConnector kafka.Connector
}

func (p *purger) Purge() error {
	k8sTopics, err := p.k8sConnector.Topics()
	if err != nil {
		return errors.Wrap(err, "get topic from k8s failed")
	}
	kafkaTopics, err := p.kafkaConnector.Topics()
	if err != nil {
		return errors.Wrap(err, "list topics from kafka failed")
	}
	for _, kafkaTopic := range kafkaTopics {
		if strings.HasPrefix(kafkaTopic.Name, "_") {
			glog.V(3).Infof("ignore topic %s", kafkaTopic.Name)
			continue
		}
		missing := true
		for _, k8sTopic := range k8sTopics {
			if k8sTopic.Name == kafkaTopic.Name {
				missing = false
			}
		}
		if missing {
			err := p.kafkaConnector.DeleteTopic(kafkaTopic)
			if err != nil {
				return errors.Wrap(err, "delete topic failed")
			}
		}
	}
	return nil
}
