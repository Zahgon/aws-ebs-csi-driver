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

// ModifyVolumeTest will provision pod with attached volume, and test that modifying its pvc will modify the associated pv.
type ModifyVolumeTest struct {
	CreateVolumeParameters                map[string]string
	ModifyVolumeParameters                map[string]string
	ShouldResizeVolume                    bool
	ShouldTestInvalidModificationRecovery bool
	ExternalResizerOnly                   bool
}

var (
	invalidAnnotations = map[string]string{
		Iops: "1",
	}
	volumeSize = "10Gi" // Different from driver.MinimumSizeForVolumeType to simplify iops, throughput, volumeType modification
)

type ModifyTestType int64

const (
	VolumeModifierForK8s ModifyTestType = iota
	ExternalResizer
)

func (modifyVolumeTest *ModifyVolumeTest) Run(c clientset.Interface, ns *v1.Namespace, ebsDriver driver.PVTestDriver, testType ModifyTestType) {
	_ = "STUB: not implemented"
	return
}

func attemptInvalidModification(c clientset.Interface, ns *v1.Namespace, testVolume *TestPersistentVolumeClaim) {
	_ = "STUB: not implemented"
	return
}
