// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/bborbe/kafka-k8s-topic-controller/purge"
	"github.com/bborbe/run"

	"github.com/Shopify/sarama"
	"github.com/bborbe/argument"
	flag "github.com/bborbe/flagenv"
	"github.com/bborbe/kafka-k8s-topic-controller/k8s"
	"github.com/bborbe/kafka-k8s-topic-controller/kafka"
	"github.com/golang/glog"
	"github.com/pkg/errors"
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
	Kubeconfig    string        `required:"true" arg:"kubeconfig" env:"KUBECONFIG" usage:"Path to k8s config"`
	KafkaBrokers  string        `required:"true" arg:"kafka-brokers" env:"KAFKA_BROKERS" usage:"Comma seperated list of Kafka brokers"`
	Purge         bool          `required:"true" arg:"purge" env:"PURGE" usage:"Purge all not existing Topics"`
	PurgeInterval time.Duration `required:"true" arg:"purge-interval" env:"PURGE_INTERVAL" usage:"Time between purges" default:"5m"`
}

func (a *application) run(ctx context.Context) error {
	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0

	clusterAdmin, err := sarama.NewClusterAdmin(strings.Split(a.KafkaBrokers, ","), config)
	if err != nil {
		return errors.Wrap(err, "create cluster admin failed")
	}
	defer clusterAdmin.Close()

	k8sConnector := k8s.NewConnector(a.Kubeconfig)
	if err := k8sConnector.SetupCustomResourceDefinition(); err != nil {
		return err
	}
	kafkaConnector := kafka.NewConnector(clusterAdmin)
	purger := purge.NewPurger(k8sConnector, kafkaConnector)
	return run.CancelOnFirstFinish(
		ctx,
		func(ctx context.Context) error {
			return k8sConnector.Listen(
				ctx,
				k8s.NewResourceEventHandler(
					kafkaConnector,
				),
			)
		},
		func(ctx context.Context) error {
			for {
				select {
				case <-ctx.Done():
					return nil
				case <-time.NewTicker(a.PurgeInterval).C:
					err := purger.Purge()
					if err != nil {
						return err
					}
				}
			}
		},
	)
}
