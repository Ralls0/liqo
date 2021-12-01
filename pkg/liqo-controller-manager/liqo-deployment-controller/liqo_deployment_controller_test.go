// Copyright 2019-2021 The Liqo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package liqodeploymentctrl_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/kubectl/pkg/scheme"
	ctrl "sigs.k8s.io/controller-runtime"

	liqodeploymentctrl "github.com/liqotech/liqo/pkg/liqo-controller-manager/liqo-deployment-controller"
)

var _ = Describe("Reconcile", func() {
	var (
		ctx context.Context
		res ctrl.Result
		err error

		req = ctrl.Request{}
	)

	BeforeEach(func() {
		ctx = context.TODO()
	})

	JustBeforeEach(func() {
		r := &liqodeploymentctrl.Reconciler{
			Client: k8sClient,
			Scheme: scheme.Scheme,
		}

		res, err = r.Reconcile(ctx, req)
	})

	It("should not error", func() {
		Expect(res).To(BeNil())
		Expect(err).NotTo(HaveOccurred())
	})
})
