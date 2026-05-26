// Copyright 2024 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the 'License');
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an 'AS IS' BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sanity

import (
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/mounter"
	"k8s.io/mount-utils"
)

type fakeMounter struct {
	mounts map[string]string
}

func newFakeMounter() *fakeMounter { _ = "STUB: not implemented"; return nil }

func (m *fakeMounter) FindDevicePath(devicePath, volumeID, partition, region string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *fakeMounter) PreparePublishTarget(target string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *fakeMounter) IsBlockDevice(fullPath string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *fakeMounter) GetBlockSizeBytes(devicePath string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *fakeMounter) GetDeviceNameFromMount(mountPath string) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func (m *fakeMounter) IsCorruptedMnt(err error) bool { _ = "STUB: not implemented"; return false }

func (m *fakeMounter) MakeFile(path string) error { _ = "STUB: not implemented"; return nil }

func (m *fakeMounter) MakeDir(path string) error { _ = "STUB: not implemented"; return nil }

func (m *fakeMounter) PathExists(path string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *fakeMounter) Resize(devicePath, deviceMountPath string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *fakeMounter) NeedResize(devicePath string, deviceMountPath string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *fakeMounter) Unpublish(path string) error { _ = "STUB: not implemented"; return nil }

func (m *fakeMounter) Unstage(path string) error { _ = "STUB: not implemented"; return nil }

func (m *fakeMounter) Mount(source string, target string, fstype string, options []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *fakeMounter) CanSafelySkipMountPointCheck() bool { _ = "STUB: not implemented"; return false }

func (m *fakeMounter) FormatAndMountSensitiveWithFormatOptions(source, target, fstype string, options, sensitiveOptions, formatOptions []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *fakeMounter) GetMountRefs(pathname string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *fakeMounter) IsLikelyNotMountPoint(file string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *fakeMounter) IsMountPoint(file string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *fakeMounter) List() ([]mount.MountPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *fakeMounter) MountSensitive(source, target, fstype string, options, sensitiveOptions []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *fakeMounter) MountSensitiveWithoutSystemd(source, target, fstype string, options, sensitiveOptions []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *fakeMounter) MountSensitiveWithoutSystemdWithMountFlags(source, target, fstype string, options, sensitiveOptions, mountFlags []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *fakeMounter) Unmount(target string) error { _ = "STUB: not implemented"; return nil }

func (m *fakeMounter) GetVolumeStats(volumePath string) (mounter.VolumeStats, error) {
	_ = "STUB: not implemented"
	return *new(mounter.VolumeStats), nil
}
