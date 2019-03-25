// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Kafka K8s Topic Controller", func() {
	It("Compiles", func() {
		var err error
		_, err = gexec.Build("github.com/bborbe/kafka-k8s-topic-controller")
		Expect(err).NotTo(HaveOccurred())
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kafka K8s Topic Controller Suite")
}
