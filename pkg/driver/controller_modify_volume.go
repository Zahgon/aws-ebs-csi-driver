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

package driver

import (
	"context"

	"github.com/awslabs/volume-modifier-for-k8s/pkg/rpc"
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/coalescer"
)

const (
	ModificationKeyVolumeType = "type"
	// DeprecatedModificationKeyVolumeType is retained for backwards compatibility, but not recommended.
	DeprecatedModificationKeyVolumeType = "volumeType"

	ModificationKeyIOPS = "iops"

	ModificationKeyThroughput = "throughput"

	ModificationAddTag = "tagSpecification"

	ModificationDeleteTag = "tagDeletion"
)

type modifyVolumeRequest struct {
	newSize           int64
	modifyDiskOptions cloud.ModifyDiskOptions
	modifyTagsOptions cloud.ModifyTagsOptions
}

func (d *ControllerService) GetCSIDriverModificationCapability(
	_ context.Context,
	_ *rpc.GetCSIDriverModificationCapabilityRequest,
) (*rpc.GetCSIDriverModificationCapabilityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *ControllerService) ModifyVolumeProperties(
	ctx context.Context,
	req *rpc.ModifyVolumePropertiesRequest,
) (*rpc.ModifyVolumePropertiesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newModifyVolumeCoalescer(c cloud.Cloud, o *Options) coalescer.Coalescer[modifyVolumeRequest, int32] {
	_ = "STUB: not implemented"
	return nil
}

func mergeModifyVolumeRequest(input modifyVolumeRequest, existing modifyVolumeRequest) (modifyVolumeRequest, error) {
	_ = "STUB: not implemented"
	return *new(modifyVolumeRequest), nil
}

func executeModifyTagsRequest(volumeID string, options modifyVolumeRequest, c cloud.Cloud, ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func executeModifyVolumeRequest(c cloud.Cloud) func(string, modifyVolumeRequest) (int32, error) {
	_ = "STUB: not implemented"
	return nil
}

// Returning Internal error instead of InvaliArgument because at this point any tag modifications have succeeded.
// It would not be correct to return an error that is considered infeasible by the resizer if the volume was already modified in any way.

// No change to the volume was requested, so return an empty result with no error

func parseModifyVolumeParameters(params map[string]string) (*modifyVolumeRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
