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
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/cloud/metadata"
)

type fakeCloud struct {
	fakeMetadata     *metadata.Metadata
	mountPath        string
	disks            map[string]*cloud.Disk
	snapshots        map[string]*cloud.Snapshot
	snapshotNameToID map[string]string
}

func newFakeCloud(fmd *metadata.Metadata, mp string) *fakeCloud {
	_ = "STUB: not implemented"
	return nil
}

func (d *fakeCloud) CreateDisk(ctx context.Context, volumeID string, diskOptions *cloud.DiskOptions) (*cloud.Disk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *fakeCloud) DeleteDisk(ctx context.Context, volumeID string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (d *fakeCloud) GetDiskByID(ctx context.Context, volumeID string) (*cloud.Disk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *fakeCloud) CreateSnapshot(ctx context.Context, volumeID string, opts *cloud.SnapshotOptions) (*cloud.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *fakeCloud) DeleteSnapshot(ctx context.Context, snapshotID string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (d *fakeCloud) GetSnapshotByID(ctx context.Context, snapshotID string) (*cloud.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *fakeCloud) GetSnapshotByName(ctx context.Context, name string) (*cloud.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *fakeCloud) GetInstancesPatching(ctx context.Context, nodeIDs []string) ([]*types.Instance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *fakeCloud) ListSnapshots(ctx context.Context, sourceVolumeID string, maxResults int32, nextToken string) (*cloud.ListSnapshotsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *fakeCloud) AttachDisk(ctx context.Context, volumeID string, instanceID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *fakeCloud) DetachDisk(ctx context.Context, volumeID string, instanceID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *fakeCloud) ResizeOrModifyDisk(ctx context.Context, volumeID string, newSizeBytes int64, modifyOptions *cloud.ModifyDiskOptions) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *fakeCloud) AvailabilityZones(ctx context.Context) (map[string]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *fakeCloud) EnableFastSnapshotRestores(ctx context.Context, availabilityZones []string, snapshotID string) (*ec2.EnableFastSnapshotRestoresOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *fakeCloud) LockSnapshot(ctx context.Context, lockOptions *cloud.SnapshotLockOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *fakeCloud) GetDiskByName(ctx context.Context, name string, capacityBytes int64) (*cloud.Disk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *fakeCloud) ModifyTags(ctx context.Context, volumeID string, tagOptions cloud.ModifyTagsOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *fakeCloud) WaitForAttachmentState(ctx context.Context, expectedState types.VolumeAttachmentState, volumeID string, expectedInstance string, expectedDevice string, alreadyAssigned bool, expectedCardIndex *int32) (*types.VolumeAttachment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *fakeCloud) IsVolumeInitialized(ctx context.Context, volumeID string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (d *fakeCloud) DryRun(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (d *fakeCloud) GetVolumeIDByNodeAndDevice(ctx context.Context, nodeID, deviceName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
