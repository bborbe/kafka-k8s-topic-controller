// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	"context"
	"time"

	"github.com/bborbe/kafka-k8s-topic-controller/topic"

	"github.com/bborbe/kafka-k8s-topic-controller/k8s/client/clientset/versioned"
	"github.com/bborbe/kafka-k8s-topic-controller/k8s/client/informers/externalversions"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensionsClient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

const defaultResync = 5 * time.Minute

//go:generate counterfeiter -o ../mocks/k8s-connector.go --fake-name K8sConnector . Connector
type Connector interface {
	SetupCustomResourceDefinition() error
	Listen(
		ctx context.Context,
		resourceEventHandler cache.ResourceEventHandler,
	) error
	Topics() ([]topic.Topic, error)
}

func NewConnector(kubeconfig string) Connector {
	return &connector{
		kubeconfig: kubeconfig,
	}
}

type connector struct {
	kubeconfig string
}

func (c *connector) Topics() ([]topic.Topic, error) {
	config, err := c.createKubernetesConfig()
	if err != nil {
		return nil, errors.Wrap(err, "build k8s config failed")
	}
	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "build clientset failed")
	}
	topicList, err := clientset.KafkaV1().Topics("").List(metav1.ListOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "list topic failed")
	}
	var result []topic.Topic
	for _, item := range topicList.Items {
		result = append(result, manifestToTopic(item))
	}
	return result, nil
}

func (c *connector) Listen(
	ctx context.Context,
	resourceEventHandler cache.ResourceEventHandler,
) error {
	config, err := c.createKubernetesConfig()
	if err != nil {
		return errors.Wrap(err, "build k8s config failed")
	}
	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		return errors.Wrap(err, "build clientset failed")
	}
	informerFactory := externalversions.NewSharedInformerFactory(clientset, defaultResync)
	informerFactory.
		Kafka().
		V1().
		Topics().
		Informer().
		AddEventHandler(resourceEventHandler)

	stopCh := make(chan struct{})
	glog.V(2).Infof("listen for events")
	informerFactory.Start(stopCh)
	select {
	case <-ctx.Done():
		glog.V(0).Infof("listen canceled")
	case <-stopCh:
		glog.V(0).Infof("listen stopped")
	}
	return nil
}

func (c *connector) SetupCustomResourceDefinition() error {
	config, err := c.createKubernetesConfig()
	if err != nil {
		return errors.Wrap(err, "build k8s config failed")
	}
	clientset, err := apiextensionsClient.NewForConfig(config)
	if err != nil {
		return errors.Wrap(err, "build clientset failed")
	}
	name := "topics.kafka.benjamin-borbe.de"
	spec := v1beta1.CustomResourceDefinitionSpec{
		Group: "kafka.benjamin-borbe.de",
		Names: v1beta1.CustomResourceDefinitionNames{
			Kind:     "Topic",
			ListKind: "TopicList",
			Plural:   "topics",
			Singular: "topic",
		},
		Scope: "Namespaced",
		Versions: []v1beta1.CustomResourceDefinitionVersion{
			{
				Name:    "v1",
				Served:  true,
				Storage: true,
			},
		},
	}
	customResourceDefinition, err := clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Get(name, metav1.GetOptions{})
	if err != nil {
		glog.V(2).Infof("get CustomResourceDefinition %s failed => create", name)
		_, err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Create(&v1beta1.CustomResourceDefinition{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "apiextensions.k8s.io/v1beta1",
				Kind:       "CustomResourceDefinition",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
			Spec: spec,
		})
		if err != nil {
			return errors.Wrap(err, "create CustomResourceDefinition failed")
		}
		glog.V(2).Infof("CustomResourceDefinitions %s created", name)
		return nil
	}
	customResourceDefinition.Spec = spec
	_, err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Update(customResourceDefinition)
	if err != nil {
		return errors.Wrap(err, "update CustomResourceDefinition failed")
	}
	glog.V(2).Infof("CustomResourceDefinitions %s updated", name)
	return nil
}

func (c *connector) createKubernetesConfig() (*rest.Config, error) {
	if len(c.kubeconfig) > 0 {
		glog.V(3).Infof("create kube config from flags")
		return clientcmd.BuildConfigFromFlags("", c.kubeconfig)
	}
	glog.V(3).Infof("create in cluster kube config")
	return rest.InClusterConfig()
}
