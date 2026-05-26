/*
Copyright 2018 The Kubernetes Authors.

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

// FormatOptionTest will provision required StorageClass(es), PVC(s) and Pod(s) in order to test that volumes with
// a specified custom format options will mount and are able to be later resized.
type FormatOptionTest struct {
	CreateVolumeParameters map[string]string
}

func (t *FormatOptionTest) Run(client clientset.Interface, namespace *v1.Namespace, ebsDriver driver.PVTestDriver) {
	_ = "STUB: not implemented"
	return
}

func createPodWithVolume(client clientset.Interface, namespace *v1.Namespace, cmd string, testPvc *TestPersistentVolumeClaim, volumeDetails *VolumeDetails) *TestPod {
	_ = "STUB: not implemented"
	return nil
}
