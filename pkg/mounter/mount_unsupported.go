//go:build darwin

/*
Copyright 2019 The Kubernetes Authors.

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

package mounter

import (
	mountutils "k8s.io/mount-utils"
)

const (
	stubMessage = "nodeService is unsupported for this platform"
)

/*
NOTE: This is stub implementation of nodeService so that maintainers without access to a Linux/Windows workstation can
run driver e2e tests.
*/

func NewSafeMounter() (*mountutils.SafeFormatAndMount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSafeMounterV2() (*mountutils.SafeFormatAndMount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *NodeMounter) FindDevicePath(devicePath, volumeID, partition, region string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *NodeMounter) PreparePublishTarget(target string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *NodeMounter) IsBlockDevice(fullPath string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *NodeMounter) GetBlockSizeBytes(devicePath string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m NodeMounter) GetDeviceNameFromMount(mountPath string) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func (m NodeMounter) IsCorruptedMnt(err error) bool { _ = "STUB: not implemented"; return false }

func (m *NodeMounter) MakeFile(path string) error { _ = "STUB: not implemented"; return nil }

func (m *NodeMounter) MakeDir(path string) error { _ = "STUB: not implemented"; return nil }

func (m *NodeMounter) PathExists(path string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *NodeMounter) Resize(devicePath, deviceMountPath string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *NodeMounter) NeedResize(devicePath string, deviceMountPath string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *NodeMounter) Unpublish(path string) error { _ = "STUB: not implemented"; return nil }

func (m *NodeMounter) Unstage(path string) error { _ = "STUB: not implemented"; return nil }

func (m *NodeMounter) GetVolumeStats(volumePath string) (VolumeStats, error) {
	_ = "STUB: not implemented"
	return *new(VolumeStats), nil
}
