//go:build windows

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
	"errors"
)

var (
	ErrUnsupportedMounter = errors.New("unsupported mounter type")
)

func (m *NodeMounter) FindDevicePath(devicePath, volumeID, _, _ string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *NodeMounter) PreparePublishTarget(target string) error {
	_ = "STUB: not implemented"
	// On Windows, Mount will create the parent of target and mklink (create a symbolic link) at target later, so don't create a
	// directory at target now. Otherwise mklink will error: "Cannot create a file when that file already exists".
	// Instead, delete the target if it already exists (like if it was created by kubelet <1.20)
	// https://github.com/kubernetes/kubernetes/pull/88759
	return nil
}

// If the target does not exist, no action is necessary

// Handle different mounter implementations

// IsBlockDevice checks if the given path is a block device
func (m *NodeMounter) IsBlockDevice(fullPath string) (bool, error) {
	_ = "STUB: not implemented"

	// getBlockSizeBytes gets the size of the disk in bytes
	return false, nil
}

func (m *NodeMounter) GetBlockSizeBytes(devicePath string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *NodeMounter) FormatAndMountSensitiveWithFormatOptions(source string, target string, fstype string, options []string, sensitiveOptions []string, formatOptions []string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetDeviceNameFromMount returns the volume ID for a mount path.
// The ref count returned is always 1 or 0 because csi-proxy doesn't provide a
// way to determine the actual ref count (as opposed to Linux where the mount
// table gets read). In practice this shouldn't matter, as in the NodeStage
// case the ref count is ignored and in the NodeUnstage case, the ref count
// being >1 is just a warning.
// Command to determine ref count would be something like:
// Get-Volume -UniqueId "\\?\Volume{7c3da0c1-0000-0000-0000-010000000000}\" | Get-Partition | Select AccessPaths
func (m *NodeMounter) GetDeviceNameFromMount(mountPath string) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

// HACK change csi-proxy behavior instead of relying on fragile internal
// implementation details!
// if err contains '"(Get-Item...).Target, output: , error: <nil>' then the
// internal Get-Item cmdlet didn't fail but no item/device was found at the
// path so we should return empty string and nil error just like the Linux
// implementation would.

// handleGetDeviceNameFromMountError processes common error patterns for GetDeviceNameFromMount
func handleGetDeviceNameFromMountError(err error) (string, int, error) {
	_ = "STUB: not implemented"
	// Handling common error pattern as seen in previous implementations
	return "", 0, nil
}

// No device found, but not an error condition

// IsCorruptedMnt return true if err is about corrupted mount point
func (m *NodeMounter) IsCorruptedMnt(err error) bool { _ = "STUB: not implemented"; return false }

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

// Refresh the host storage cache before resizing to ensure Windows has an updated view of disk geometry.

// NeedResize called at NodeStage to ensure file system is the correct size
func (m *NodeMounter) NeedResize(devicePath, deviceMountPath string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Tolerate one block difference (4096 bytes)

// Unmount volume from target path
func (m *NodeMounter) Unpublish(target string) error { _ = "STUB: not implemented"; return nil }

// Unmount volume from staging path
// usually this staging path is a global directory on the node
func (m *NodeMounter) Unstage(target string) error { _ = "STUB: not implemented"; return nil }

// GetVolumeStats acquires byte statistics of filesystem at volumePath.
func (m *NodeMounter) GetVolumeStats(volumePath string) (VolumeStats, error) {
	_ = "STUB: not implemented"
	return *new(VolumeStats), nil
}

// Need to ensure directory path to prevent error code: The directory name is invalid. (#99173)
// See https://docs.microsoft.com/en-us/windows/win32/debug/system-error-codes--0-499-
