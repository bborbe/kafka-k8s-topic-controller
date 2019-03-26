// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	v1 "github.com/bborbe/kafka-k8s-topic-controller/k8s/apis/kafka.benjamin-borbe.de/v1"
	"github.com/bborbe/kafka-k8s-topic-controller/topic"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"k8s.io/client-go/tools/cache"
)

//go:generate counterfeiter -o ../mocks/k8s-eventhandler.go --fake-name K8sEventHandler . EventHandler
type EventHandler interface {
	CreateTopic(topic topic.Topic) error
	DeleteTopic(topic topic.Topic) error
	UpdateTopic(oldTopic topic.Topic, newTopic topic.Topic) error
}

func NewResourceEventHandler(handler EventHandler) cache.ResourceEventHandler {
	return &resourceEventHandler{
		handler: handler,
	}
}

type resourceEventHandler struct {
	handler EventHandler
}

func (r *resourceEventHandler) OnAdd(obj interface{}) {
	topic, err := objToTopic(obj)
	if err != nil {
		glog.Warningf("convert topic failed: %v", err)
		return
	}
	if err := r.handler.CreateTopic(*topic); err != nil {
		glog.Warningf("create topic failed: %v", err)
	}
}

func (r *resourceEventHandler) OnUpdate(oldObj, newObj interface{}) {
	oldTopic, err := objToTopic(oldObj)
	if err != nil {
		glog.Warningf("convert topic failed: %v", err)
		return
	}
	newTopic, err := objToTopic(newObj)
	if err != nil {
		glog.Warningf("convert topic failed: %v", err)
		return
	}
	if err := r.handler.UpdateTopic(*oldTopic, *newTopic); err != nil {
		glog.Warningf("update topic failed: %v", err)
		return
	}
}

func (r *resourceEventHandler) OnDelete(obj interface{}) {
	topic, err := objToTopic(obj)
	if err != nil {
		glog.Warningf("convert topic failed: %v", err)
		return
	}
	if err := r.handler.DeleteTopic(*topic); err != nil {
		glog.Warningf("delete topic failed: %v", err)
	}
}

func objToTopic(obj interface{}) (*topic.Topic, error) {
	switch t := obj.(type) {
	case *v1.Topic:
		result := manifestToTopic(*t)
		return &result, nil
	default:
		return nil, errors.Errorf("not supporeted type %T", obj)
	}
}

func manifestToTopic(t v1.Topic) topic.Topic {
	return topic.Topic{
		Name:              t.Spec.Name,
		NumPartitions:     t.Spec.NumPartitions,
		ReplicationFactor: t.Spec.ReplicationFactor,
	}
}
