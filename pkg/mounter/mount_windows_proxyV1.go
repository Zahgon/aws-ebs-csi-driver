//go:build windows

/*
Copyright 2024 The Kubernetes Authors.

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
	"os"

	diskclient "github.com/kubernetes-csi/csi-proxy/client/groups/disk/v1"
	fsclient "github.com/kubernetes-csi/csi-proxy/client/groups/filesystem/v1"
	volumeclient "github.com/kubernetes-csi/csi-proxy/client/groups/volume/v1"

	mountutils "k8s.io/mount-utils"
)

type CSIProxyMounter struct {
	FsClient     *fsclient.Client
	DiskClient   *diskclient.Client
	VolumeClient *volumeclient.Client
}

// NewSafeMounter returns a new instance of SafeFormatAndMount.
func NewSafeMounter() (*mountutils.SafeFormatAndMount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Mount just creates a soft link at target pointing to source.
func (mounter *CSIProxyMounter) Mount(source string, target string, fstype string, options []string) error {
	_ = "STUB: not implemented"
	// Mount is called after the format is done.
	// TODO: Confirm that fstype is empty.
	return nil
}

func (mounter *CSIProxyMounter) Unmount(target string) error {
	_ = "STUB: not implemented"
	// Find the volume id
	return nil
}

// Call UnmountVolume CSI proxy function which flushes data cache to disk and removes the global staging path

// Cleanup stage path

// Get disk number

// Offline the disk

// Rmdir - delete the given directory
func (mounter *CSIProxyMounter) Rmdir(path string) error { _ = "STUB: not implemented"; return nil }

func (mounter *CSIProxyMounter) WriteVolumeCache(target string) { _ = "STUB: not implemented"; return }

func (mounter *CSIProxyMounter) List() ([]mountutils.MountPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mounter *CSIProxyMounter) IsMountPointMatch(mp mountutils.MountPoint, dir string) bool {
	_ = "STUB: not implemented"
	return false

	// IsMountPoint: determines if a directory is a mountpoint.
}

func (mounter *CSIProxyMounter) IsMountPoint(file string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsLikelyMountPoint - If the directory does not exists, the function will return os.ErrNotExist error.
//
//	If the path exists, call to CSI proxy will check if its a link, if its a link then existence of target
//	path is checked.
func (mounter *CSIProxyMounter) IsLikelyNotMountPoint(path string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (mounter *CSIProxyMounter) PathIsDevice(pathname string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (mounter *CSIProxyMounter) DeviceOpened(pathname string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GetDeviceNameFromMount returns the disk number for a mount path.
func (mounter *CSIProxyMounter) GetDeviceNameFromMount(mountPath, _ string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Get disk number

func (mounter *CSIProxyMounter) MakeRShared(path string) error {
	_ = "STUB: not implemented"
	return nil
}

func (mounter *CSIProxyMounter) MakeFile(pathname string) error {
	_ = "STUB: not implemented"
	return nil
}

// MakeDir - Creates a directory. The CSI proxy takes in context information.
// Currently the make dir is only used from the staging code path, hence we call it
// with Plugin context..
func (mounter *CSIProxyMounter) MakeDir(pathname string) error {
	_ = "STUB: not implemented"
	return nil
}

// ExistsPath - Checks if a path exists. Unlike util ExistsPath, this call does not perform follow link.
func (mounter *CSIProxyMounter) ExistsPath(path string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (mounter *CSIProxyMounter) EvalHostSymlinks(pathname string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (mounter *CSIProxyMounter) GetMountRefs(pathname string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mounter *CSIProxyMounter) GetFSGroup(pathname string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (mounter *CSIProxyMounter) GetSELinuxSupport(pathname string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (mounter *CSIProxyMounter) GetMode(pathname string) (os.FileMode, error) {
	_ = "STUB: not implemented"
	return *new(os.FileMode), nil
}

func (mounter *CSIProxyMounter) MountSensitive(source string, target string, fstype string, options []string, sensitiveOptions []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (mounter *CSIProxyMounter) MountSensitiveWithoutSystemd(source string, target string, fstype string, options []string, sensitiveOptions []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (mounter *CSIProxyMounter) MountSensitiveWithoutSystemdWithMountFlags(source string, target string, fstype string, options []string, sensitiveOptions []string, mountFlags []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Rescan would trigger an update storage cache via the CSI proxy.
func (mounter *CSIProxyMounter) Rescan() error {
	_ = "STUB: not implemented"
	// Call Rescan from disk APIs of CSI Proxy.
	return nil
}

// FindDiskByLun - given a lun number, find out the corresponding disk
func (mounter *CSIProxyMounter) FindDiskByLun(lun string) (diskNum string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List all disk locations and match the lun id being requested for.
// If match is found then return back the disk number.

// FormatAndMount - accepts the source disk number, target path to mount, the fstype to format with and options to be used.
func (mounter *CSIProxyMounter) FormatAndMountSensitiveWithFormatOptions(source string, target string, fstype string, options []string, sensitiveOptions []string, formatOptions []string) error {
	_ = "STUB: not implemented"
	// sensitiveOptions is not supported on Windows because we have no reasonable way to control what the csi-proxy does
	return nil
}

// formatOptions is not supported on Windows because the csi-proxy does not allow supplying format arguments
// This limitation will be addressed in the future with privileged Windows containers

// Call PartitionDisk CSI proxy call to partition the disk and return the volume id

// Ensure the disk is online before mounting.

// List the volumes on the given disk.

// TODO: consider partitions and choose the right partition.
// For now just choose the first volume.

// Check if the volume is formatted.

// If the volume is not formatted, then format it, else proceed to mount.

// TODO: Accept the filesystem and other options

// Mount the volume by calling the CSI proxy call.

// ResizeVolume resizes the volume at given mount path
func (mounter *CSIProxyMounter) ResizeVolume(deviceMountPath string) (bool, error) {
	_ = "STUB: not implemented"
	// Find the volume id
	return false, nil
}

// Resize volume

// GetVolumeSizeInBytes returns the size of the volume in bytes
func (mounter *CSIProxyMounter) GetVolumeSizeInBytes(deviceMountPath string) (int64, error) {
	_ = "STUB: not implemented"
	// Find the volume id
	return 0, nil
}

// Get size of the volume

// GetDeviceSize returns the size of the disk in bytes
func (mounter *CSIProxyMounter) GetDeviceSize(devicePath string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//Get size of the disk

func (mounter *CSIProxyMounter) CanSafelySkipMountPointCheck() bool {
	_ = "STUB: not implemented"
	return false
}

func (mounter *CSIProxyMounter) FindDevicePath(devicePath, volumeID, _, _ string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
