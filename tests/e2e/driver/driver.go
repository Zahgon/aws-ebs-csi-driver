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

package driver

import (
	volumesnapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1"
	v1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
)

const (
	VolumeSnapshotClassKind = "VolumeSnapshotClass"
	SnapshotAPIVersion      = "snapshot.storage.k8s.io/v1"
)

type PVTestDriver interface {
	DynamicPVTestDriver
	PreProvisionedVolumeTestDriver
	VolumeSnapshotTestDriver
}

// DynamicPVTestDriver represents an interface for a CSI driver that supports DynamicPV.
type DynamicPVTestDriver interface {
	// GetDynamicProvisionStorageClass returns a StorageClass dynamic provision Persistent Volume
	GetDynamicProvisionStorageClass(parameters map[string]string, mountOptions []string, reclaimPolicy *v1.PersistentVolumeReclaimPolicy, volumeExpansion *bool, bindingMode *storagev1.VolumeBindingMode, allowedTopologyValues []string, namespace string) *storagev1.StorageClass
}

// PreProvisionedVolumeTestDriver represents an interface for a CSI driver that supports pre-provisioned volume.
type PreProvisionedVolumeTestDriver interface {
	// GetPersistentVolume returns a PersistentVolume with pre-provisioned volumeHandle
	GetPersistentVolume(volumeID string, fsType string, size string, reclaimPolicy *v1.PersistentVolumeReclaimPolicy, namespace string, accessMode v1.PersistentVolumeAccessMode, volumeMode v1.PersistentVolumeMode) *v1.PersistentVolume
}

type VolumeSnapshotTestDriver interface {
	GetVolumeSnapshotClass(namespace string, parameters map[string]string) *volumesnapshotv1.VolumeSnapshotClass
}

func getStorageClass(
	generateName string,
	provisioner string,
	parameters map[string]string,
	mountOptions []string,
	reclaimPolicy *v1.PersistentVolumeReclaimPolicy,
	volumeExpansion *bool,
	bindingMode *storagev1.VolumeBindingMode,
	allowedTopologies []v1.TopologySelectorTerm,
) *storagev1.StorageClass {
	_ = "STUB: not implemented"
	return nil
}

func getVolumeSnapshotClass(generateName string, provisioner string, parameters map[string]string) *volumesnapshotv1.VolumeSnapshotClass {
	_ = "STUB: not implemented"
	return nil
}
