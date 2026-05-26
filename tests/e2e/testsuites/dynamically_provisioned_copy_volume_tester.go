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
	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
)

// DynamicallyProvisionedCopyVolumeTest will provision required StorageClass(es), PVC(s) and Pod(s)
// Waiting for the PV provisioner to create a new PV
// Testing if the Pod(s) can write and read to mounted volumes
// Create a copy of the volume using PVC as data source, validate the data is copied, and then write and read to it again
// This test only supports a single volume.
type DynamicallyProvisionedCopyVolumeTest struct {
	CSIDriver    driver.PVTestDriver
	Pod          PodDetails
	ClonedPod    PodDetails
	ValidateFunc func()
}

func (t *DynamicallyProvisionedCopyVolumeTest) Run(client clientset.Interface, namespace *v1.Namespace) {
	_ = "STUB: not implemented"
	return
}
