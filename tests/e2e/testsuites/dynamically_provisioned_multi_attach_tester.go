/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
   http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package testsuites

import (
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/tests/e2e/driver"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
)

type DynamicallyProvisionedMultiAttachTest struct {
	CSIDriver  driver.DynamicPVTestDriver
	Pods       []PodDetails
	VolumeMode VolumeMode
	AccessMode v1.PersistentVolumeAccessMode
	VolumeType string
	RunningPod bool
	PendingPVC bool
}

func (t *DynamicallyProvisionedMultiAttachTest) Run(client clientset.Interface, namespace *v1.Namespace) {
	_ = "STUB: not implemented"
	// Setup StorageClass and PVC
	return
}
