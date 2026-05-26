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
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	clientset "k8s.io/client-go/kubernetes"
)

const (
	DefaultVolumeName = "test-volume-1"
	DefaultMountPath  = "/mnt/default-mount"

	DefaultIopsIoVolumes = "100"

	DefaultSizeIncreaseGi = int32(1)

	DefaultModificationTimeout   = 3 * time.Minute
	DefaultResizeTimout          = 1 * time.Minute
	DefaultK8sAPIPollingInterval = 5 * time.Second

	Iops       = "iops"
	Throughput = "throughput"
	VolumeType = "type"
	TagSpec    = "tagSpecification"
	TagDel     = "tagDeletion"
	Encrypted  = "encrypted"
)

var DefaultGeneratedVolumeMount = VolumeMountDetails{
	NameGenerate:      "test-volume-",
	MountPathGenerate: "/mnt/test-",
}

// PodCmdWriteToVolume returns pod command that would write to mounted volume.
func PodCmdWriteToVolume(volumeMountPath string) string { _ = "STUB: not implemented"; return "" }

// PodCmdContinuousWrite returns pod command that would continuously write to mounted volume.
func PodCmdContinuousWrite(volumeMountPath string) string { _ = "STUB: not implemented"; return "" }

// PodCmdGrepVolumeData returns pod command that would check that a volume was written to by PodCmdWriteToVolume.
func PodCmdGrepVolumeData(volumeMountPath string) string { _ = "STUB: not implemented"; return "" }

// IncreasePvcObjectStorage increases `storage` of a K8s PVC object by specified Gigabytes.
func IncreasePvcObjectStorage(pvc *v1.PersistentVolumeClaim, sizeIncreaseGi int32) resource.Quantity {
	_ = "STUB: not implemented"
	return *new(resource.Quantity)
}

// WaitForPvToResize waiting for pvc size to be resized to desired size.
func WaitForPvToResize(c clientset.Interface, ns *v1.Namespace, pvName string, desiredSize resource.Quantity, timeout time.Duration, interval time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// ResizeTestPvc increases size of given `TestPersistentVolumeClaim` by specified Gigabytes.
func ResizeTestPvc(client clientset.Interface, namespace *v1.Namespace, testPvc *TestPersistentVolumeClaim, sizeIncreaseGi int32) (updatedSize resource.Quantity) {
	_ = "STUB: not implemented"
	return *new(resource.Quantity)
}

// AnnotatePvc annotates supplied k8s pvc object with supplied annotations.
func AnnotatePvc(pvc *v1.PersistentVolumeClaim, annotations map[string]string) {
	_ = "STUB: not implemented"
	return
}

// CheckPvAnnotations checks whether supplied k8s pv object contains supplied annotations.
func CheckPvAnnotations(pv *v1.PersistentVolume, annotations map[string]string) bool {
	_ = "STUB: not implemented"
	return false
}

// WaitForPvToModify waiting for PV to be modified.
func WaitForPvToModify(c clientset.Interface, ns *v1.Namespace, pvName string, expectedAnnotations map[string]string, timeout time.Duration, interval time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForVacToApplyToPv waits for a PV's VAC to match the PVC's VAC.
func WaitForVacToApplyToPv(c clientset.Interface, ns *v1.Namespace, pvName string, expectedVac string, timeout time.Duration, interval time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func CreateVolumeDetails(createVolumeParameters map[string]string, volumeSize string) *VolumeDetails {
	_ = "STUB: not implemented"
	return nil
}

func PrefixAnnotations(prefix string, parameters map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

type ExpectedParameters struct {
	Size       *int32
	IOPS       *int32
	Throughput *int32
	VolumeType *string
	Encrypted  *bool
}

func BuildExpectedParameters(params map[string]string, claimSize string) ExpectedParameters {
	_ = "STUB: not implemented"
	return *new(ExpectedParameters)
}

func VerifyVolumeProperties(volumeID string, verification ExpectedParameters) {
	_ = "STUB: not implemented"
	return
}
