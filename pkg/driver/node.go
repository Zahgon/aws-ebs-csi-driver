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
	"time"

	csi "github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/driver/internal"
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/mounter"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	// default file system type to be used when it is not provided.
	defaultFsType = FSTypeExt4

	// VolumeOperationAlreadyExists is message fmt returned to CO when there is another in-flight call on the given volumeID.
	VolumeOperationAlreadyExists = "An operation with the given volume=%q is already in progress"
)

var (
	ValidFSTypes = map[string]struct{}{
		FSTypeExt3: {},
		FSTypeExt4: {},
		FSTypeXfs:  {},
		FSTypeNtfs: {},
	}
)

var (
	// nodeCaps represents the capability of node service.
	nodeCaps = []csi.NodeServiceCapability_RPC_Type{
		csi.NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME,
		csi.NodeServiceCapability_RPC_EXPAND_VOLUME,
		csi.NodeServiceCapability_RPC_GET_VOLUME_STATS,
	}
)

const (
	// taintWatcherDuration is the maximum duration for the not-ready taint watcher to run.
	taintWatcherDuration = 10 * time.Minute
)

// NodeService represents the node service of CSI driver.
type NodeService struct {
	metadata metadata.MetadataService
	mounter  mounter.Mounter
	inFlight *internal.InFlight
	options  *Options
	csi.UnimplementedNodeServer
}

// NewNodeService creates a new node service.
func NewNodeService(o *Options, md metadata.MetadataService, m mounter.Mounter, k kubernetes.Interface) *NodeService {
	_ = "STUB: not implemented"

	// Watch for the agent‑not‑ready taint for up to one minute and remove it
	// as soon as allocatable is available.
	return nil
}

func (d *NodeService) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the access type is block, do nothing for stage

// When exists is true it means target path was created but device isn't mounted.
// We don't want to do anything in that case and let the operation proceed.
// Otherwise we need to create the target directory.

// If target path does not exist we need to create the directory where volume will be staged

// Check if a device is mounted in target directory

// This operation (NodeStageVolume) MUST be idempotent.
// If the volume corresponding to the volume_id is already staged to the staging_target_path,
// and is identical to the specified volume_capability the Plugin MUST reply 0 OK.

// FormatAndMount will format only if needed

func (d *NodeService) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if target directory is a mount point. GetDeviceNameFromMount
// given a mnt point, finds the device from /proc/mounts
// returns the device name, reference count, and error code

// From the spec: If the volume corresponding to the volume_id
// is not staged to the staging_target_path, the Plugin MUST
// reply 0 OK.

func (d *NodeService) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VolumeCapability is optional, if specified, use that as source of truth

// Noop for Block NodeExpandVolume

// TODO use util.GenericResizeFS
// VolumeCapability is nil, check if volumePath point to a block device

// Skip resizing for Block NodeExpandVolume

func (d *NodeService) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *NodeService) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *NodeService) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *NodeService) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *NodeService) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// to my surprise ARN's string representation is not empty for empty ARN

func (d *NodeService) nodePublishVolumeForBlock(req *csi.NodePublishVolumeRequest, mountOptions []string) error {
	_ = "STUB: not implemented"
	return nil
}

// create the global mount path if it is missing
// Path in the form of /var/lib/kubelet/plugins/kubernetes.io/csi/volumeDevices/publish/{volumeName}

// Create the mount point as a file since bind mount device node requires it to be a file
// This implementation detail is relied upon by the NVMECollector,
// which discovers block devices by parsing /proc/self/mountinfo. The bind mount
// created here ensures block devices appear in mountinfo even without a filesystem.

// Checking if the target file is already mounted with a device.

// isMounted checks if target is mounted. It does NOT return an error if target
// doesn't exist.
func (d *NodeService) isMounted(_ string, target string) (bool, error) {
	_ = "STUB: not implemented"
	/*
		Checking if it's a mount point using IsLikelyNotMountPoint. There are three different return values,
		1. true, err when the directory does not exist or corrupted.
		2. false, nil when the path is already mounted with a device.
		3. true, nil when the path is not mounted with any device.
	*/return false, nil
}

// Checking if the path exists and error is related to Corrupted Mount, in that case, the system could unmount and mount.

// After successful unmount, the device is ready to be mounted.

// Do not return os.IsNotExist error. Other errors were handled above.  The
// Existence of the target should be checked by the caller explicitly and
// independently because sometimes prior to mount it is expected not to exist
// (in Windows, the target must NOT exist before a symlink is created at it)
// and in others it is an error (in Linux, the target mount directory must
// exist before mount is called on it)

func (d *NodeService) nodePublishVolumeForFileSystem(req *csi.NodePublishVolumeRequest, mountOptions []string, mode *csi.VolumeCapability_Mount) error {
	_ = "STUB: not implemented"
	return nil
}

// Checking if the target directory is already mounted with a device.

// getVolumesLimit returns the limit of volumes that the node supports.
func (d *NodeService) getVolumesLimit() int64 { _ = "STUB: not implemented"; return 0 }

// Calculate reserved volume attachments (additional EBS volumes)

// Auto-detect number of reserved volume attachments - plus 1 to account for the root volume

// For shared attachment types, subtract ENIs

// Safety measure: Never return a limit of below 1, as Kubernetes will treat it as infinite

// hasMountOption returns a boolean indicating whether the given
// slice already contains a mount option. This is used to prevent
// passing duplicate option to the mount command.
func hasMountOption(options []string, opt string) bool { _ = "STUB: not implemented"; return false }

// collectMountOptions returns array of mount options from
// VolumeCapability_MountVolume and special mount options for
// given filesystem.
func collectMountOptions(fsType string, mntFlags []string) []string {
	_ = "STUB: not implemented"
	return nil
}

// By default, xfs does not allow mounting of two volumes with the same filesystem uuid.
// Force ignore this uuid to be able to mount volume + its clone / restored snapshot on the same node.

// startNotReadyTaintWatcher launches a short‑lived Node informer that removes the
// ebs.csi.aws.com/agent‑not‑ready taint. The informer is stopped after maxWatchDuration.
func startNotReadyTaintWatcher(clientset kubernetes.Interface, maxWatchDuration time.Duration) {
	_ = "STUB: not implemented"
	return
}

// Resync every 5 seconds in case of networking or other rare issue

// Add additional 90 seconds to the context for the last-try taint removal

// Node has no taint, do nothing

// Another removal thread is already running, do nothing

// Our node is probably stale, get a new copy

// Continue retrying with old node

// Check if taint was already removed by another attempt

// Taint is gone, we're done

// Update to fresh node for next retry

// Continue retrying

// We either removed the taint, or there was no taint to remove

// Informer doesn't have permission - cancel context
// to avoid spamming logs with informer errors

// Context likely cancelled because of permissions error - log at higher
// verbosity in this case to avoid spamming logs of users that have
// modified their permissions to opt out

// Immediate scan in case the taint is already present and no event fires

// Informer is operational - wait for maxWatchDuration for it to handle Node updates

// Try to remove the taint one last time in case we got extremely unlucky with the informer
// We still try this even if the informer failed, as we may only be missing the watch permission

func hasNotReadyTaint(n *corev1.Node) bool { _ = "STUB: not implemented"; return false }

// JSONPatch struct for JSON patch operations.
type JSONPatch struct {
	OP    string `json:"op,omitempty"`
	Path  string `json:"path,omitempty"`
	Value any    `json:"value"`
}

// removeNotReadyTaint removes the taint ebs.csi.aws.com/agent-not-ready from the local node
// This taint can be optionally applied by users to prevent startup race conditions such as
// https://github.com/kubernetes/kubernetes/issues/95911
func removeNotReadyTaint(ctx context.Context, clientset kubernetes.Interface, node *corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func checkAllocatable(ctx context.Context, clientset kubernetes.Interface, nodeName string) error {
	_ = "STUB: not implemented"
	return nil
}

func recheckFormattingOptionParameter(context map[string]string, key string, fsConfigs map[string]fileSystemConfig, fsType string) (value string, err error) {
	_ = "STUB: not implemented"
	return "", nil

	// This check is already performed on the controller side
	// However, because it is potentially security-sensitive, we redo it here to be safe
}

// In the case that the default fstype does not support custom sizes we could
// be using an invalid fstype, so recheck that here
