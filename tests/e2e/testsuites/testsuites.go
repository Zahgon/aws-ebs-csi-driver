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
	"fmt"
	"time"

	volumesnapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1"
	awscloud "github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/cloud"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	apps "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	clientset "k8s.io/client-go/kubernetes"
	restclientset "k8s.io/client-go/rest"
)

const (
	execTimeout = 10 * time.Second
	// Some pods can take much longer to get ready due to volume attach/detach latency.
	slowPodStartTimeout = 15 * time.Minute
	// Description that will printed during tests.
	failedConditionDescription = "Error status code"

	volumeSnapshotNameStatic         = "volume-snapshot-tester"
	volumeSnapshotContenetNameStatic = "volume-snapshot-content-tester"
)

type TestStorageClass struct {
	client       clientset.Interface
	storageClass *storagev1.StorageClass
	namespace    *v1.Namespace
}

func NewTestStorageClass(c clientset.Interface, ns *v1.Namespace, sc *storagev1.StorageClass) *TestStorageClass {
	_ = "STUB: not implemented"
	return nil
}

func (t *TestStorageClass) Create() storagev1.StorageClass {
	_ = "STUB: not implemented"
	return *new(storagev1.StorageClass)
}

func (t *TestStorageClass) Cleanup() { _ = "STUB: not implemented"; return }

type TestVolumeSnapshotClass struct {
	client              restclientset.Interface
	volumeSnapshotClass *volumesnapshotv1.VolumeSnapshotClass
	namespace           *v1.Namespace
}

func NewTestVolumeSnapshotClass(c restclientset.Interface, ns *v1.Namespace, vsc *volumesnapshotv1.VolumeSnapshotClass) *TestVolumeSnapshotClass {
	_ = "STUB: not implemented"
	return nil
}

func (t *TestVolumeSnapshotClass) Create() { _ = "STUB: not implemented"; return }

func (t *TestVolumeSnapshotClass) CreateSnapshot(pvc *v1.PersistentVolumeClaim) *volumesnapshotv1.VolumeSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (t *TestVolumeSnapshotClass) CreateStaticVolumeSnapshot(vsc *volumesnapshotv1.VolumeSnapshotContent) *volumesnapshotv1.VolumeSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (t *TestVolumeSnapshotClass) CreateStaticVolumeSnapshotContent(snapshotID string) *volumesnapshotv1.VolumeSnapshotContent {
	_ = "STUB: not implemented"
	return nil
}

func (t *TestVolumeSnapshotClass) UpdateStaticVolumeSnapshotContent(volumeSnapshot *volumesnapshotv1.VolumeSnapshot, volumeSnapshotContent *volumesnapshotv1.VolumeSnapshotContent) {
	_ = "STUB: not implemented"
	return
}

func (t *TestVolumeSnapshotClass) ReadyToUse(snapshot *volumesnapshotv1.VolumeSnapshot) {
	_ = "STUB: not implemented"
	return
}

func (t *TestVolumeSnapshotClass) unlockSnapshot(vs *volumesnapshotv1.VolumeSnapshot) {
	_ = "STUB: not implemented"
	return
}

// Snapshot not found or error, skip unlock

func (t *TestVolumeSnapshotClass) DeleteSnapshot(vs *volumesnapshotv1.VolumeSnapshot) {
	_ = "STUB: not implemented"
	return
}

func (t *TestVolumeSnapshotClass) DeleteVolumeSnapshotContent(vsc *volumesnapshotv1.VolumeSnapshotContent) {
	_ = "STUB: not implemented"
	return
}

//nolint

func (t *TestVolumeSnapshotClass) Cleanup() { _ = "STUB: not implemented"; return }

func (t *TestVolumeSnapshotClass) waitForSnapshotDeleted(ns string, snapshotName string, poll, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *TestVolumeSnapshotClass) waitForVolumeSnapshotContentDeleted(vscName string, poll, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

type TestPreProvisionedPersistentVolume struct {
	client                    clientset.Interface
	persistentVolume          *v1.PersistentVolume
	requestedPersistentVolume *v1.PersistentVolume
}

func NewTestPreProvisionedPersistentVolume(c clientset.Interface, pv *v1.PersistentVolume) *TestPreProvisionedPersistentVolume {
	_ = "STUB: not implemented"
	return nil
}

func (pv *TestPreProvisionedPersistentVolume) Create() v1.PersistentVolume {
	_ = "STUB: not implemented"
	return *new(v1.PersistentVolume)
}

type TestPersistentVolumeClaim struct {
	client                         clientset.Interface
	claimSize                      string
	volumeMode                     v1.PersistentVolumeMode
	accessMode                     v1.PersistentVolumeAccessMode
	storageClass                   *storagev1.StorageClass
	namespace                      *v1.Namespace
	persistentVolume               *v1.PersistentVolume
	persistentVolumeClaim          *v1.PersistentVolumeClaim
	requestedPersistentVolumeClaim *v1.PersistentVolumeClaim
	dataSource                     *v1.TypedLocalObjectReference
}

func NewTestPersistentVolumeClaim(c clientset.Interface, ns *v1.Namespace, claimSize string, volumeMode VolumeMode, sc *storagev1.StorageClass, accessMode v1.PersistentVolumeAccessMode) *TestPersistentVolumeClaim {
	_ = "STUB: not implemented"
	return nil
}

func NewTestPersistentVolumeClaimWithDataSource(c clientset.Interface, ns *v1.Namespace, claimSize string, volumeMode VolumeMode, sc *storagev1.StorageClass, dataSource *v1.TypedLocalObjectReference, accessMode v1.PersistentVolumeAccessMode) *TestPersistentVolumeClaim {
	_ = "STUB: not implemented"
	return nil
}

func (t *TestPersistentVolumeClaim) Create() { _ = "STUB: not implemented"; return }

func (t *TestPersistentVolumeClaim) ValidateProvisionedPersistentVolume() {
	_ = "STUB: not implemented"

	// Get the bound PersistentVolume
	return
}

// Check sizes

// Check PV properties

// If storageClass is nil, PV was pre-provisioned with these values already set

// Since we're chaging our topology key, assume we have the values below to compare:
// NodeSelectorTerms: [{[{topology.ebs.csi.aws.com/zone In [us-west-2a]} {topology.kubernetes.io/zone In [us-west-2a]}] []}]
// AllowedTopologies: [{[{topology.ebs.csi.aws.com/zone [us-west-2a us-west-2b us-west-2c]}]}]
// As you can see tests might fail depending on the ordering of the NodeSelectorTerms. That's why we're doing this "hack".
// This is a quick fix to unblock the PRs we have. We really need to improve this. TODO

// additional sanity check so we can catch an unintended test case that'd hide failures

func (t *TestPersistentVolumeClaim) WaitForBound() v1.PersistentVolumeClaim {
	_ = "STUB: not implemented"
	return *new(v1.PersistentVolumeClaim)
}

// Get new copy of the claim

func generatePVC(namespace, storageClassName, claimSize string, volumeMode v1.PersistentVolumeMode, dataSource *v1.TypedLocalObjectReference, accessMode v1.PersistentVolumeAccessMode) *v1.PersistentVolumeClaim {
	_ = "STUB: not implemented"
	return nil
}

func (t *TestPersistentVolumeClaim) Cleanup() { _ = "STUB: not implemented"; return }

// Wait for the PV to get deleted if reclaim policy is Delete. (If it's
// Retain, there's no use waiting because the PV won't be auto-deleted and
// it's expected for the caller to do it.) Technically, the first few delete
// attempts may fail, as the volume is still attached to a node because
// kubelet is slowly cleaning up the previous pod, however it should succeed
// in a couple of minutes.

// Wait for the PVC to be deleted

func (t *TestPersistentVolumeClaim) ReclaimPolicy() v1.PersistentVolumeReclaimPolicy {
	_ = "STUB: not implemented"
	return *new(v1.PersistentVolumeReclaimPolicy)
}

func (t *TestPersistentVolumeClaim) WaitForPersistentVolumePhase(phase v1.PersistentVolumePhase) {
	_ = "STUB: not implemented"
	return
}

func (t *TestPersistentVolumeClaim) DeleteBoundPersistentVolume() {
	_ = "STUB: not implemented"
	return
}

func (t *TestPersistentVolumeClaim) DeleteBackingVolume(cloud awscloud.Cloud) {
	_ = "STUB: not implemented"
	return
}

type TestDeployment struct {
	client     clientset.Interface
	deployment *apps.Deployment
	namespace  *v1.Namespace
	podName    string
}

func NewTestDeployment(c clientset.Interface, ns *v1.Namespace, command string, pvc *v1.PersistentVolumeClaim, volumeName, mountPath string, readOnly bool) *TestDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (t *TestDeployment) Create() { _ = "STUB: not implemented"; return }

// always get first pod as there should only be one

func (t *TestDeployment) WaitForPodReady() { _ = "STUB: not implemented"; return }

// always get first pod as there should only be one

func (t *TestDeployment) Exec(command []string, expectedString string) {
	_ = "STUB: not implemented"
	return
}

func (t *TestDeployment) DeletePodAndWait() { _ = "STUB: not implemented"; return }

func (t *TestDeployment) Cleanup() { _ = "STUB: not implemented"; return }

func (t *TestDeployment) Logs() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// waitForPersistentVolumeClaimDeleted waits for a PersistentVolumeClaim to be removed from the system until timeout occurs, whichever comes first.
func waitForPersistentVolumeClaimDeleted(c clientset.Interface, ns string, pvcName string, poll, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

type TestPod struct {
	client    clientset.Interface
	pod       *v1.Pod
	namespace *v1.Namespace
}

func NewTestPod(c clientset.Interface, ns *v1.Namespace, command string) *TestPod {
	_ = "STUB: not implemented"
	return nil
}

func (t *TestPod) Create() { _ = "STUB: not implemented"; return }

func (t *TestPod) GetName() string { _ = "STUB: not implemented"; return "" }

func (t *TestPod) WaitForSuccess() { _ = "STUB: not implemented"; return }

func (t *TestPod) WaitForRunning() { _ = "STUB: not implemented"; return }

// Ideally this would be in "k8s.io/kubernetes/test/e2e/framework"
// Similar to framework.WaitForPodSuccessInNamespace.
var podFailedCondition = func(pod *v1.Pod) (bool, error) {
	switch pod.Status.Phase {
	case v1.PodFailed:
		By("Saw pod failure")
		return true, nil
	case v1.PodSucceeded:
		return true, fmt.Errorf("pod %q successed with reason: %q, message: %q", pod.Name, pod.Status.Reason, pod.Status.Message)
	case v1.PodPending, v1.PodRunning, v1.PodUnknown:
		return false, nil
	default:
		return false, nil
	}
}

func (t *TestPod) WaitForFailure() { _ = "STUB: not implemented"; return }

func (t *TestPod) SetupVolume(pvc *v1.PersistentVolumeClaim, name, mountPath string, readOnly bool) {
	_ = "STUB: not implemented"
	return
}

func (t *TestPod) SetupRawBlockVolume(pvc *v1.PersistentVolumeClaim, name, devicePath string) {
	_ = "STUB: not implemented"
	return
}

func (t *TestPod) SetNodeSelector(nodeSelector map[string]string) {
	_ = "STUB: not implemented"
	return
}

func (t *TestPod) Cleanup() { _ = "STUB: not implemented"; return }

func (t *TestPod) Logs() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func cleanupPodOrFail(client clientset.Interface, name, namespace string) {
	_ = "STUB: not implemented"
	return
}

func podLogs(client clientset.Interface, name, namespace string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
