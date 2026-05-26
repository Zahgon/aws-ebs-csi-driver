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
	storagev1 "k8s.io/api/storage/v1"
	clientset "k8s.io/client-go/kubernetes"
	restclientset "k8s.io/client-go/rest"
)

type PodDetails struct {
	Cmd     string
	Volumes []VolumeDetails
}

type VolumeDetails struct {
	MountOptions               []string
	ClaimSize                  string
	ReclaimPolicy              *v1.PersistentVolumeReclaimPolicy
	AllowVolumeExpansion       *bool
	VolumeBindingMode          *storagev1.VolumeBindingMode
	AccessMode                 v1.PersistentVolumeAccessMode
	AllowedTopologyValues      []string
	VolumeMode                 VolumeMode
	VolumeMount                VolumeMountDetails
	VolumeDevice               VolumeDeviceDetails
	CreateVolumeParameters     map[string]string // Optional, used when dynamically-provisioned volumes
	VolumeID                   string            // Optional, used with pre-provisioned volumes
	PreProvisionedVolumeFsType string            // Optional, used with pre-provisioned volumes
	DataSource                 *DataSource       // Optional, used with PVCs created from snapshots
}

type VolumeMode int

const (
	FileSystem VolumeMode = iota
	Block
)

const (
	VolumeSnapshotKind        = "VolumeSnapshot"
	PersistentVolumeClaimKind = "PersistentVolumeClaim"
	VolumeSnapshotContentKind = "VolumeSnapshotContent"
	SnapshotAPIVersion        = "snapshot.storage.k8s.io/v1"
	APIVersionv1              = "v1"
)

var (
	SnapshotAPIGroup = "snapshot.storage.k8s.io"
)

type VolumeMountDetails struct {
	NameGenerate      string
	MountPathGenerate string
	ReadOnly          bool
}

type VolumeDeviceDetails struct {
	NameGenerate string
	DevicePath   string
}

type DataSource struct {
	Name string
	Kind string
}

func (pod *PodDetails) SetupWithDynamicVolumes(client clientset.Interface, namespace *v1.Namespace, csiDriver driver.DynamicPVTestDriver) (*TestPod, []func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pod *PodDetails) SetupWithPreProvisionedVolumes(client clientset.Interface, namespace *v1.Namespace, csiDriver driver.PreProvisionedVolumeTestDriver) (*TestPod, []func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pod *PodDetails) SetupDeployment(client clientset.Interface, namespace *v1.Namespace, csiDriver driver.DynamicPVTestDriver) (*TestDeployment, []func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (volume *VolumeDetails) SetupDynamicPersistentVolumeClaim(client clientset.Interface, namespace *v1.Namespace, csiDriver driver.DynamicPVTestDriver) (*TestPersistentVolumeClaim, []func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PV will not be ready until PVC is used in a pod when volumeBindingMode: WaitForFirstConsumer

func (volume *VolumeDetails) SetupPreProvisionedPersistentVolumeClaim(client clientset.Interface, namespace *v1.Namespace, csiDriver driver.PreProvisionedVolumeTestDriver) (*TestPersistentVolumeClaim, []func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CreateVolumeSnapshotClass(client restclientset.Interface, namespace *v1.Namespace, csiDriver driver.VolumeSnapshotTestDriver, vscParameters map[string]string) (*TestVolumeSnapshotClass, func()) {
	_ = "STUB: not implemented"
	return nil, nil
}
