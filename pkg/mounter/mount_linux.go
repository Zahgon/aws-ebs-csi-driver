//go:build linux

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
	nvmeDiskPartitionSuffix = "p"
	diskPartitionSuffix     = ""
)

func NewSafeMounter() (*mountutils.SafeFormatAndMount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSafeMounterV2() (*mountutils.SafeFormatAndMount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindDevicePath finds path of device and verifies its existence
// if the device is not nvme, return the path directly
// if the device is nvme, finds and returns the nvme device path eg. /dev/nvme1n1.
func (m *NodeMounter) FindDevicePath(devicePath, volumeID, partition, region string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// If the given path exists, the device MAY be nvme. Further, it MAY be a
// symlink to the nvme device path like:
// | $ stat /dev/xvdba
// | File: ‘/dev/xvdba’ -> ‘nvme1n1’
// Since these are maybes, not guarantees, the search for the nvme device
// path below must happen and must rely on volume ID

// AWS recommends identifying devices by volume ID
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nvme-ebs-volumes.html),
// so find the nvme device path using volume ID. This is the magic name on
// which AWS presents NVME devices under /dev/disk/by-id/. For example,
// vol-0fab1d5e3f72a5e23 creates a symlink at
// /dev/disk/by-id/nvme-Amazon_Elastic_Block_Store_vol0fab1d5e3f72a5e23

// findNvmeVolume looks for the nvme volume with the specified name
// It follows the symlink (if it exists) and returns the absolute path to the device.
func findNvmeVolume(findName string) (device string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Find the target, resolving to an absolute path
// For example, /dev/disk/by-id/nvme-Amazon_Elastic_Block_Store_vol0fab1d5e3f72a5e23 -> ../../nvme2n1

// execRunner is a helper to inject exec.Comamnd().CombinedOutput() for verifyVolumeSerialMatch
// Tests use a mocked version that does not actually execute any binaries.
func execRunner(name string, arg ...string) ([]byte, error) {
	_ = "STUB: not implemented"
	// TODO: Pass context down from driver.go
	return nil, nil
}

// verifyVolumeSerialMatch checks the volume serial of the device against the expected volume.
func verifyVolumeSerialMatch(canonicalDevicePath string, strippedVolumeName string, execRunner func(string, ...string) ([]byte, error)) error {
	_ = "STUB: not implemented"
	// As a security precaution, check the device name looks like a real device name before passing anything to exec
	return nil
}

// In some rare cases, a race condition can lead to the /dev/disk/by-id/ symlink becoming out of date
// See https://github.com/kubernetes-sigs/aws-ebs-csi-driver/issues/1224 for more info
// Attempt to use lsblk to double check that the nvme device selected was the correct volume

// Look for an EBS volume ID in the output, compare all matches against what we expect
// (in some rare cases there may be multiple matches due to lsblk printing partitions)
// If no volume ID is in the output (non-Nitro instances, SBE devices, etc) silently proceed

// If the command fails (for example, because lsblk is not available), silently ignore the error and proceed

// PreparePublishTarget creates the target directory for the volume to be mounted.
func (m *NodeMounter) PreparePublishTarget(target string) error {
	_ = "STUB: not implemented"
	return nil
}

// IsBlockDevice checks if the given path is a block device.
func (m *NodeMounter) IsBlockDevice(fullPath string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GetBlockSizeBytes gets the size of the disk in bytes.
func (m *NodeMounter) GetBlockSizeBytes(devicePath string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// appendPartition appends the partition to the device path.
func (m *NodeMounter) appendPartition(devicePath, partition string) string {
	_ = "STUB: not implemented"
	return ""
}

// GetDeviceNameFromMount returns the volume ID for a mount path.
func (m *NodeMounter) GetDeviceNameFromMount(mountPath string) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

// IsCorruptedMnt return true if err is about corrupted mount point.
func (m *NodeMounter) IsCorruptedMnt(err error) bool { _ = "STUB: not implemented"; return false }

// MakeFile function is mirrored in ./sanity_test.go to make sure sanity test covered this block of code
// Please mirror the change to func MakeFile in ./sanity_test.go.
func (m *NodeMounter) MakeFile(path string) error { _ = "STUB: not implemented"; return nil }

// MakeDir function is mirrored in ./sanity_test.go to make sure sanity test covered this block of code
// Please mirror the change to func MakeFile in ./sanity_test.go.
func (m *NodeMounter) MakeDir(path string) error { _ = "STUB: not implemented"; return nil }

// PathExists function is mirrored in ./sanity_test.go to make sure sanity test covered this block of code
// Please mirror the change to func MakeFile in ./sanity_test.go.
func (m *NodeMounter) PathExists(path string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Resize resizes the filesystem of the given devicePath.
func (m *NodeMounter) Resize(devicePath, deviceMountPath string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// NeedResize checks if the filesystem of the given devicePath needs to be resized.
func (m *NodeMounter) NeedResize(devicePath string, deviceMountPath string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Unpublish unmounts the given path.
func (m *NodeMounter) Unpublish(path string) error {
	_ = "STUB: not implemented"
	// On linux, unpublish and unstage both perform an unmount
	return nil
}

// Unstage unmounts the given path.
func (m *NodeMounter) Unstage(path string) error { _ = "STUB: not implemented"; return nil }

// Ignore the error when it contains "not mounted", because that indicates the
// world is already in the desired state
//
// mount-utils attempts to detect this on its own but fails when running on
// a read-only root filesystem, which our manifests use by default

// GetVolumeStats acquires byte and inode statistics of filesystem at volumePath.
func (m *NodeMounter) GetVolumeStats(volumePath string) (VolumeStats, error) {
	_ = "STUB: not implemented"
	return *new(VolumeStats), nil
}

// Get byte stats (safely)

// Get inode stats (safely)
