// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/bborbe/argument"
	flag "github.com/bborbe/flagenv"
	"github.com/bborbe/kafka-k8s-topic-controller/client/clientset/versioned"
	"github.com/bborbe/kafka-k8s-topic-controller/client/informers/externalversions"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	runtime.GOMAXPROCS(runtime.NumCPU())
	_ = flag.Set("logtostderr", "true")

	app := &application{}
	if err := argument.Parse(app); err != nil {
		glog.Exitf("parse args failed: %v", err)
	}

	glog.V(1).Infof("application started")
	if err := app.run(contextWithSig(context.Background())); err != nil {
		glog.Exitf("application failed: %+v", err)
	}
	glog.V(1).Infof("application finished")
	os.Exit(0)
}

func contextWithSig(ctx context.Context) context.Context {
	ctxWithCancel, cancel := context.WithCancel(ctx)
	go func() {
		defer cancel()

		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-signalCh:
		case <-ctx.Done():
		}
	}()
	return ctxWithCancel
}

type application struct {
	Kubeconfig string `required:"true" arg:"kubeconfig" env:"KUBECONFIG"`
}

func (a *application) run(ctx context.Context) error {
	clientset, err := a.createKubernetesClientset()
	if err != nil {
		return err
	}
	informerFactory := externalversions.NewSharedInformerFactory(clientset, 30*time.Second)
	informerFactory.Example().V1().Databases().Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{})
	select {
	case <-ctx.Done():
		return nil
	}
}

func (a *application) createKubernetesClientset() (versioned.Interface, error) {
	config, err := a.createKubernetesConfig()
	if err != nil {
		return nil, errors.Wrap(err, "build k8s config failed")
	}
	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "build clientset failed")
	}
	return clientset, nil
}

func (a *application) createKubernetesConfig() (*rest.Config, error) {
	if len(a.Kubeconfig) > 0 {
		glog.V(2).Infof("create kube config from flags")
		return clientcmd.BuildConfigFromFlags("", a.Kubeconfig)
	}
	glog.V(2).Infof("create in cluster kube config")
	return rest.InClusterConfig()
}
