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

package driver

import (
	"context"

	"github.com/awslabs/volume-modifier-for-k8s/pkg/rpc"
	csi "github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/coalescer"
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/driver/internal"
)

// Supported access modes.
const (
	SingleNodeWriter     = csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER
	MultiNodeMultiWriter = csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER
)

var (
	// controllerCaps represents the capability of controller service.
	controllerCaps = []csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_CLONE_VOLUME,
		csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_SNAPSHOT,
		csi.ControllerServiceCapability_RPC_LIST_SNAPSHOTS,
		csi.ControllerServiceCapability_RPC_EXPAND_VOLUME,
		csi.ControllerServiceCapability_RPC_MODIFY_VOLUME,
	}
)

const trueStr = "true"
const isManagedByDriver = trueStr

// ControllerService represents the controller service of CSI driver.
type ControllerService struct {
	cloud                 cloud.Cloud
	inFlight              *internal.InFlight
	options               *Options
	modifyVolumeCoalescer coalescer.Coalescer[modifyVolumeRequest, int32]
	rpc.UnimplementedModifyServer
	csi.UnimplementedControllerServer
}

// NewControllerService creates a new controller service.
func NewControllerService(c cloud.Cloud, o *Options) *ControllerService {
	_ = "STUB: not implemented"
	return nil
}

func (d *ControllerService) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if a request is already in-flight

// "Values specified in mutable_parameters MUST take precedence over the values from parameters."
// https://github.com/container-storage-interface/spec/blob/master/spec.md#createvolume

// fill volume tags - set cluster tags first so user tags can override them

// create or clone a new volume

func validateCreateVolumeRequest(req *csi.CreateVolumeRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *ControllerService) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if a request is already in-flight

func validateDeleteVolumeRequest(req *csi.DeleteVolumeRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *ControllerService) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateControllerPublishVolumeRequest(req *csi.ControllerPublishVolumeRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *ControllerService) controllerPublishVolumeNodeLocal(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *ControllerService) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateControllerUnpublishVolumeRequest(req *csi.ControllerUnpublishVolumeRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *ControllerService) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *ControllerService) GetCapacity(ctx context.Context, req *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *ControllerService) ListVolumes(ctx context.Context, req *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *ControllerService) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Node-local volumes don't need GetDiskByID validation

// For node-local volumes, allow RWX

func (d *ControllerService) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if this is a raw block device, no expansion should be necessary on the node

func (d *ControllerService) ControllerModifyVolume(ctx context.Context, req *csi.ControllerModifyVolumeRequest) (*csi.ControllerModifyVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *ControllerService) ControllerGetVolume(ctx context.Context, req *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isValidVolumeCapabilities(v []*csi.VolumeCapability) bool {
	_ = "STUB: not implemented"
	return false
}

func isValidCapability(c *csi.VolumeCapability) bool { _ = "STUB: not implemented"; return false }

//nolint:exhaustive

func isNodeLocalVolume(volumeID string) bool { _ = "STUB: not implemented"; return false }

func isValidCapabilityForNodeLocal(c *csi.VolumeCapability) bool {
	_ = "STUB: not implemented"
	return false
}

func isBlock(capability *csi.VolumeCapability) bool { _ = "STUB: not implemented"; return false }

func isValidVolumeContext(volContext map[string]string) bool {
	_ = "STUB: not implemented"
	// There could be multiple volume attributes in the volumeContext map
	// Validate here case by case
	return false
}

func (d *ControllerService) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if a request is already in-flight

// Check if the availability zone is supported for fast snapshot restore

func validateCreateSnapshotRequest(req *csi.CreateSnapshotRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *ControllerService) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if a request is already in-flight

func validateDeleteSnapshotRequest(req *csi.DeleteSnapshotRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *ControllerService) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// pickAvailabilityZone selects 1 zone given topology requirement.
// if not found, empty string is returned.
func pickAvailabilityZone(requirement *csi.TopologyRequirement) string {
	_ = "STUB: not implemented"
	return ""
}

func pickAvailabilityZoneID(requirement *csi.TopologyRequirement) string {
	_ = "STUB: not implemented"
	return ""
}

func getOutpostArn(requirement *csi.TopologyRequirement) string {
	_ = "STUB: not implemented"
	return ""
}

// Check if source volumes topology matches with clones requisite topology requirements.
func checkSourceTopology(requirement *csi.TopologyRequirement, sourceVolumeZone string, sourceVolumeOutpostArn string, sourceVolumeZoneID string) error {
	_ = "STUB: not implemented"
	return nil
}

func newCreateVolumeResponse(disk *cloud.Disk, ctx map[string]string) *csi.CreateVolumeResponse {
	_ = "STUB: not implemented"
	return nil
}

func newCreateSnapshotResponse(snapshot *cloud.Snapshot) *csi.CreateSnapshotResponse {
	_ = "STUB: not implemented"
	return nil
}

func newListSnapshotsResponse(cloudResponse *cloud.ListSnapshotsResponse) *csi.ListSnapshotsResponse {
	_ = "STUB: not implemented"
	return nil
}

func newListSnapshotsResponseEntry(snapshot *cloud.Snapshot) *csi.ListSnapshotsResponse_Entry {
	_ = "STUB: not implemented"
	return nil
}

func getVolSizeBytes(req *csi.CreateVolumeRequest) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// BuildOutpostArn returns the string representation of the outpost ARN from the given csi.TopologyRequirement.segments.
func BuildOutpostArn(segments map[string]string) string { _ = "STUB: not implemented"; return "" }

func validateFormattingOption(volumeCapabilities []*csi.VolumeCapability, paramName string, fsConfigs map[string]fileSystemConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func isTrue(value string) bool { _ = "STUB: not implemented"; return false }

func (d *ControllerService) cleanupSnapshotOnError(ctx context.Context, snapshotID, snapshotName string, originalErr error, errorMsg string) error {
	_ = "STUB: not implemented"
	return nil
}
