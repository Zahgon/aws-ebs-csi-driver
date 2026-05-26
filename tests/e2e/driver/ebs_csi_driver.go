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
	True = "true"
)

// Implement DynamicPVTestDriver interface.
type ebsCSIDriver struct {
	driverName string
}

// InitEbsCSIDriver returns ebsCSIDriver that implements DynamicPVTestDriver interface.
func InitEbsCSIDriver() PVTestDriver { _ = "STUB: not implemented"; return *new(PVTestDriver) }

func (d *ebsCSIDriver) GetDynamicProvisionStorageClass(parameters map[string]string, mountOptions []string, reclaimPolicy *v1.PersistentVolumeReclaimPolicy, volumeExpansion *bool, bindingMode *storagev1.VolumeBindingMode, allowedTopologyValues []string, namespace string) *storagev1.StorageClass {
	_ = "STUB: not implemented"
	return nil
}

func (d *ebsCSIDriver) GetVolumeSnapshotClass(namespace string, parameters map[string]string) *volumesnapshotv1.VolumeSnapshotClass {
	_ = "STUB: not implemented"
	return nil
}

func (d *ebsCSIDriver) GetPersistentVolume(volumeID string, fsType string, size string, reclaimPolicy *v1.PersistentVolumeReclaimPolicy, namespace string, accessMode v1.PersistentVolumeAccessMode, volumeMode v1.PersistentVolumeMode) *v1.PersistentVolume {
	_ = "STUB: not implemented"
	return nil
}

// Default to Retain ReclaimPolicy for pre-provisioned volumes

// TODO remove if https://github.com/kubernetes-csi/external-provisioner/issues/202 is fixed

// MinimumSizeForVolumeType returns the minimum disk size for each volumeType.
func MinimumSizeForVolumeType(volumeType string) string { _ = "STUB: not implemented"; return "" }
