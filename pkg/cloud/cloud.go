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

package cloud

import (
	"context"
	"errors"
	"regexp"
	"sync"
	"sync/atomic"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/batcher"
	dm "github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/cloud/devicemanager"
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/expiringcache"
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/util"
	"k8s.io/apimachinery/pkg/util/wait"
)

// AWS volume types.
const (
	// VolumeTypeIO1 represents a provisioned IOPS SSD type of volume.
	VolumeTypeIO1 = "io1"
	// VolumeTypeIO2 represents a provisioned IOPS SSD type of volume.
	VolumeTypeIO2 = "io2"
	// VolumeTypeGP2 represents a general purpose SSD type of volume.
	VolumeTypeGP2 = "gp2"
	// VolumeTypeGP3 represents a general purpose SSD type of volume.
	VolumeTypeGP3 = "gp3"
	// VolumeTypeSC1 represents a cold HDD (sc1) type of volume.
	VolumeTypeSC1 = "sc1"
	// VolumeTypeST1 represents a throughput-optimized HDD type of volume.
	VolumeTypeST1 = "st1"
	// VolumeTypeStandard represents a previous type of  volume.
	VolumeTypeStandard = "standard"
)

// AWS provisioning limits.
// Source: http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
const (
	io1MinTotalIOPS    = 100
	io1FallbackMaxIOPS = 64000
	io1MaxIOPSPerGB    = 50
	io2MinTotalIOPS    = 100
	io2FallbackMaxIOPS = 256000
	io2MaxIOPSPerGB    = 1000
	gp3FallbackMaxIOPS = 16000
	gp3MinTotalIOPS    = 3000
	gp3MaxIOPSPerGB    = 500
)

var (
	ValidVolumeTypes = []string{
		VolumeTypeIO1,
		VolumeTypeIO2,
		VolumeTypeGP2,
		VolumeTypeGP3,
		VolumeTypeSC1,
		VolumeTypeST1,
		VolumeTypeStandard,
	}
)

const (
	cacheForgetDelay          = 1 * time.Hour
	volInitCacheForgetDelay   = 6 * time.Hour
	iopsLimitCacheForgetDelay = 12 * time.Hour

	dryRunInterval = 3 * time.Hour

	getCallerIdentityRetryDelay = 30 * time.Second

	// stuckAttachingTimeout is the duration after which an attachment stuck in "attaching" state
	// will be detached to allow a retry.
	stuckAttachingTimeout = 90 * time.Second
)

// VolumeStatusInitializingState is const reported by EC2 DescribeVolumeStatus which AWS SDK does not have type for.
const (
	VolumeStatusInitializingState = "initializing"
	VolumeStatusInitializedState  = "completed"
)

// Defaults.
const (
	// DefaultVolumeSize represents the default volume size.
	DefaultVolumeSize int64 = 100 * util.GiB
)

// Tags.
const (
	// VolumeNameTagKey is the key value that refers to the volume's name.
	VolumeNameTagKey = "CSIVolumeName"
	// SnapshotNameTagKey is the key value that refers to the snapshot's name.
	SnapshotNameTagKey = "CSIVolumeSnapshotName"
	// KubernetesTagKeyPrefix is the prefix of the key value that is reserved for Kubernetes.
	KubernetesTagKeyPrefix = "kubernetes.io"
)

// Tags that depend on driver name (initialized in NewCloud).
var (
	// AwsEbsDriverTagKey is the tag to identify if a volume/snapshot is managed by ebs csi driver.
	AwsEbsDriverTagKey string
	// AllowAutoIOPSIncreaseOnModifyKey is the tag key for allowing IOPS increase on resizing if IOPSPerGB is set to ensure desired ratio is maintained.
	AllowAutoIOPSIncreaseOnModifyKey string
	// IOPSPerGBKey represents the tag key for IOPS per GB.
	IOPSPerGBKey string
)

// Batcher.
const (
	volumeIDBatcher volumeBatcherType = iota
	volumeTagBatcher

	snapshotIDBatcher snapshotBatcherType = iota
	snapshotTagBatcher
)

const (
	batchDescribeTimeout = 30 * time.Second

	// Minimizes RPC latency and EC2 API calls. Tuned via scalability tests.
	batchMaxDelay = 500 * time.Millisecond

	// Tuned for EC2 DescribeVolumeStatus -- as of July 2025 it takes up to 5 min for initialization info to be updated.
	slowVolumeStatusBatchMaxDelay = 2 * time.Minute
	fastVolumeStatusBatchMaxDelay = 500 * time.Millisecond
)

const (
	// maxInstancesDescribed is the maximum number of instances described in each EC2 Describe Instances call.
	maxInstancesDescribed = 1000
)

var (
	// ErrMultiDisks is an error that is returned when multiple
	// disks are found with the same volume name.
	ErrMultiDisks = errors.New("multiple disks with same name")

	// ErrDiskExistsDiffSize is an error that is returned if a disk with a given
	// name, but different size, is found.
	ErrDiskExistsDiffSize = errors.New("there is already a disk with same name and different size")

	// ErrSourceNotFound is returned when a volume's source is not found when provisioning using a source (snapshot or volume).
	ErrSourceNotFound = errors.New("source was not found")

	// ErrNotFound is returned when a resource is not found.
	ErrNotFound = errors.New("resource was not found")

	// ErrIdempotentParameterMismatch is returned when another request with same idempotent token is in-flight.
	ErrIdempotentParameterMismatch = errors.New("parameters on this idempotent request are inconsistent with parameters used in previous request(s)")

	// ErrAlreadyExists is returned when a resource is already existent.
	ErrAlreadyExists = errors.New("resource already exists")

	// ErrMultiSnapshots is returned when multiple snapshots are found
	// with the same ID.
	ErrMultiSnapshots = errors.New("multiple snapshots with the same name found")

	// ErrInvalidMaxResults is returned when a MaxResults pagination parameter is between 1 and 4.
	ErrInvalidMaxResults = errors.New("maxResults parameter must be 0 or greater than or equal to 5")

	// ErrVolumeNotBeingModified is returned if volume being described is not being modified.
	ErrVolumeNotBeingModified = errors.New("volume is not being modified")

	// ErrInvalidArgument is returned if parameters were rejected by cloud provider.
	ErrInvalidArgument = errors.New("invalid argument")

	// ErrInvalidRequest is returned if parameters were rejected by driver.
	ErrInvalidRequest = errors.New("invalid request")

	// ErrLimitExceeded is returned if a user exceeds a quota.
	ErrLimitExceeded = errors.New("limit exceeded")
)

// Set during build time via -ldflags.
var driverVersion string

// AWS error codes.
const (
	ValidationException = "ValidationException"
)

// Regex Patterns.
var (
	// For getting IOPS limit from gp3/io1 error.
	// Error example it is used for: "An error occurred (InvalidParameterValue) when calling the CreateVolume operation: Volume iops of 200000 is too high; maximum is 80000".
	nonIo2ErrRegex = regexp.MustCompile(`(?i)volume iops.*is too high.*maximum is (\d+)`)

	// For getting IOPS limit from io2 error.
	// Error example it is used for: "An error occurred (InvalidParameterCombination) when calling the CreateVolume operation: io2 volumes configured with greater than 64 TiB or 256K IOPS or 1000:1 IOPS:GB ratio are not supported".
	io2ErrRegex = regexp.MustCompile(`(?i)(\d+)K IOPS`)

	volumeIDRegex   = regexp.MustCompile(util.VolumeIDRegex)
	snapshotIDRegex = regexp.MustCompile(util.SnapshotIDRegex)
	instanceIDRegex = regexp.MustCompile(util.InstanceIDRegex)
)

var invalidParameterErrorCodes = map[string]struct{}{
	"InvalidParameter":            {},
	"InvalidParameterCombination": {},
	"InvalidParameterDependency":  {},
	"InvalidParameterValue":       {},
	"UnknownParameter":            {},
	"UnknownVolumeType":           {},
	"UnsupportedOperation":        {},
	"ValidationError":             {},
}

// Disk represents a EBS volume.
type Disk struct {
	VolumeID           string
	CapacityGiB        int32
	AvailabilityZone   string
	AvailabilityZoneID string
	SourceVolumeID     string
	SnapshotID         string
	OutpostArn         string
	KmsKeyID           string
	Attachments        []string
}

// DiskOptions represents parameters to create an EBS volume.
type DiskOptions struct {
	CapacityBytes          int64
	Tags                   map[string]string
	VolumeType             string
	IOPSPerGB              int32
	AllowIOPSPerGBIncrease bool
	IOPS                   int32
	Throughput             int32
	AvailabilityZone       string
	AvailabilityZoneID     string
	OutpostArn             string
	Encrypted              bool
	MultiAttachEnabled     bool
	// KmsKeyID represents a fully qualified resource name to the key to use for encryption.
	// example: arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef
	KmsKeyID                 string
	SnapshotID               string
	SourceVolumeID           string
	VolumeInitializationRate int32
}

// ModifyDiskOptions represents parameters to modify an EBS volume.
type ModifyDiskOptions struct {
	VolumeType                string
	IOPS                      int32
	Throughput                int32
	IOPSPerGB                 int32
	AllowIopsIncreaseOnResize bool
}

// iopsLimits represents the IOPS limits set by EBS of a volume dependent on the volume type.
type iopsLimits struct {
	maxIops      int32
	minIops      int32
	maxIopsPerGb int32
}

// getVolumeLimitsParams represents the AZ parameters that getVolumeLimits will use to make the DryRun CreateVolume call.
type getVolumeLimitsParams struct {
	availabilityZone   string
	availabilityZoneId string
	outpostArn         string
}

// ModifyTagsOptions represents parameter to modify the tags of an existing EBS volume.
type ModifyTagsOptions struct {
	TagsToAdd    map[string]string
	TagsToDelete []string
}

// Snapshot represents an EBS volume snapshot.
type Snapshot struct {
	SnapshotID     string
	SourceVolumeID string
	Size           int32
	CreationTime   time.Time
	ReadyToUse     bool
}

// ListSnapshotsResponse is the container for our snapshots along with a pagination token to pass back to the caller.
type ListSnapshotsResponse struct {
	Snapshots []*Snapshot
	NextToken string
}

// SnapshotOptions represents parameters to create an EBS snapshot.
type SnapshotOptions struct {
	Tags       map[string]string
	OutpostArn string
}

// SnapshotLockOptions represents parameters to lock an EBS snapshot.
type SnapshotLockOptions struct {
	SnapshotId     *string
	LockMode       types.LockMode
	CoolOffPeriod  *int32
	ExpirationDate *time.Time
	LockDuration   *int32
}

// ec2ListSnapshotsResponse is a helper struct returned from the AWS API calling function to the main ListSnapshots function.
type ec2ListSnapshotsResponse struct {
	Snapshots []types.Snapshot
	NextToken *string
}

// volumeWaitParameters dictates how to poll for volume events.
// E.g. how often to check if volume is created after an EC2 CreateVolume call.
type volumeWaitParameters struct {
	creationInitialDelay time.Duration
	creationBackoff      wait.Backoff
	modificationBackoff  wait.Backoff
	attachmentBackoff    wait.Backoff
}

var (
	vwp = volumeWaitParameters{
		// Based on our testing in us-west-2 and ap-south-1, the median/p99 time until volume creation is ~1.5/~4 seconds.
		// We have found that the following parameters are optimal for minimizing provisioning time and DescribeVolumes calls
		// we queue DescribeVolume calls after [1.25, 0.5, 0.75, 1.125, 1.7, 2.5, 3] seconds.
		// In total, we wait for ~60 seconds.
		creationInitialDelay: 1250 * time.Millisecond,
		creationBackoff: wait.Backoff{
			Duration: 500 * time.Millisecond,
			Factor:   1.5,
			Steps:    11,
		},

		// Most attach/detach operations on AWS finish within 1-4 seconds.
		// By using 1 second starting interval with a backoff of 1.8,
		// we get [1, 1.8, 3.24, 5.832000000000001, 10.4976].
		// In total, we wait for 2601 seconds.
		attachmentBackoff: wait.Backoff{
			Duration: 1 * time.Second,
			Factor:   1.8,
			Steps:    13,
		},

		modificationBackoff: wait.Backoff{
			Duration: 1 * time.Second,
			Factor:   1.7,
			Steps:    10,
		},
	}
)

// volumeBatcherType is an enum representing the types of volume batchers available.
type volumeBatcherType int

// snapshotBatcherType is an enum representing the types of snapshot batchers available.
type snapshotBatcherType int

// batcherManager maintains a collection of batchers for different types of tasks.
type batcherManager struct {
	volumeIDBatcher             *batcher.Batcher[string, *types.Volume]
	volumeTagBatcher            *batcher.Batcher[string, *types.Volume]
	instanceIDBatcher           *batcher.Batcher[string, *types.Instance]
	snapshotIDBatcher           *batcher.Batcher[string, *types.Snapshot]
	snapshotTagBatcher          *batcher.Batcher[string, *types.Snapshot]
	volumeModificationIDBatcher *batcher.Batcher[string, *types.VolumeModification]
	volumeStatusIDBatcherSlow   *batcher.Batcher[string, *types.VolumeStatusItem]
	volumeStatusIDBatcherFast   *batcher.Batcher[string, *types.VolumeStatusItem]
}

type cloud struct {
	awsConfig             aws.Config
	region                string
	ec2                   util.EC2API
	sm                    util.SageMakerAPI
	dm                    dm.DeviceManager
	bm                    *batcherManager
	rm                    *retryManager
	vwp                   volumeWaitParameters
	likelyBadDeviceNames  expiringcache.ExpiringCache[string, sync.Map]
	latestClientTokens    expiringcache.ExpiringCache[string, int]
	volumeInitializations expiringcache.ExpiringCache[string, volumeInitialization]
	latestIOPSLimits      expiringcache.ExpiringCache[string, iopsLimits]
	cardCountCache        expiringcache.ExpiringCache[string, int]
	accountID             string
	accountIDOnce         sync.Once
	attemptDryRun         atomic.Bool
}

var _ Cloud = &cloud{}

// initVariables initializes variables that depend on driver name.
// Separated into a separate function from NewCloud so it can be called in tests.
func initVariables() { _ = "STUB: not implemented"; return }

// NewCloud returns a new instance of AWS cloud
// It panics if session is invalid.
func NewCloud(region string, awsSdkDebugLog bool, userAgentExtra string, batchingEnabled bool, deprecatedMetrics bool) Cloud {
	_ = "STUB: not implemented"
	return *new(Cloud)
}

// Set the env var so that the session appends custom user agent string

// This middlware should always be last so it sees an unmangled error

// Allow custom SageMaker endpoint for testing

// Default clients if plugin is not in use or does not implement client override.

// Ensure an EC2 Dry-run API call is made on startup and every dryRunInterval

// newBatcherManager initializes a new instance of batcherManager.
// Each batcher's `entries` set to maximum results returned by relevant EC2 API call without pagination.
// Each batcher's `delay` minimizes RPC latency and EC2 API calls. Tuned via scalability tests.
func newBatcherManager(svc util.EC2API) *batcherManager { _ = "STUB: not implemented"; return nil }

func removeLikelyBadIds(cache expiringcache.ExpiringCache[string, struct{}], input []string) (goodIds []string, likelyBadIds []string) {
	_ = "STUB: not implemented"
	// Iterate backwards to safely remove values without affecting indices of remaining items
	return nil, nil
}

// execBatchDescribeVolumes executes a batched DescribeVolumes API call depending on the type of batcher.
func execBatchDescribeVolumes(svc util.EC2API, input []string, batcher volumeBatcherType, cache expiringcache.ExpiringCache[string, struct{}]) (map[string]*types.Volume, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// batchDescribeVolumes processes a DescribeVolumes request. Depending on the request,
// it determines the appropriate batcher to use, queues the task, and waits for the result.
func (c *cloud) batchDescribeVolumes(request *ec2.DescribeVolumesInput) (*types.Volume, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// extractVolumeKey retrieves the key associated with a given volume based on the batcher type.
// For the volumeIDBatcher type, it returns the volume's ID.
// For other types, it searches for the VolumeNameTagKey within the volume's tags.
func extractVolumeKey(v *types.Volume, batcher volumeBatcherType) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *cloud) CreateDisk(ctx context.Context, volumeName string, diskOptions *DiskOptions) (*Disk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If no volume type is specified, GP3 is used as default for newly created volumes.

// The first client token used for any volume is the volume name as provided via CSI
// However, if a volume fails to create asyncronously (that is, the CreateVolume call
// succeeds but the volume ultimately fails to create), the client token is burned until
// EC2 forgets about its use (measured as 12 hours under normal conditions)
//
// To prevent becoming stuck for 12 hours when this occurs, we sequentially append "-2",
// "-3", "-4", etc to the volume name before hashing on the subsequent attempt after a
// volume fails to create because of an IdempotentParameterMismatch AWS error
// The most recent appended value is stored in an expiring cache to prevent memory leaks

// We use a sha256 hash to guarantee the token that is less than or equal to 64 characters

// EC2 API does NOT handle idempotency correctly when a theoretical volume
// would put the caller over a limit for their account
//
// To avoid leaking volumes, make a DescribeVolumes call here

// Call DescribeVolumes directly as there is a high chance this volume
// will return a NotFound error and would poison a batch call

// Volume with requested name exists, continue with it

// This should in theory be impossible, but if the API
// changes or breaks it would cause a panic, so handle it

func (c *cloud) createCloneHelper(ctx context.Context, input *ec2.CopyVolumesInput, iops int32, throughput int32) (int32, string, string, error) {
	_ = "STUB: not implemented"
	return 0, "", "", nil
}

func (c *cloud) createVolumeHelper(ctx context.Context, diskOptions *DiskOptions, input *ec2.CreateVolumeInput, iops int32, throughput int32, zone string, zoneID string) (int32, string, string, error) {
	_ = "STUB: not implemented"
	return 0, "", "", nil
}

// execBatchDescribeVolumesModifications executes a batched DescribeVolumesModifications API call.
func execBatchDescribeVolumesModifications(svc util.EC2API, input []string) (map[string]*types.VolumeModification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// batchDescribeVolumesModifications processes a DescribeVolumesModifications request by queuing the task and waiting for the result.
func (c *cloud) batchDescribeVolumesModifications(request *ec2.DescribeVolumesModificationsInput) (*types.VolumeModification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ModifyTags adds, updates, and deletes tags for the specified EBS volume.
func (c *cloud) ModifyTags(ctx context.Context, volumeID string, tagOptions ModifyTagsOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// ResizeOrModifyDisk resizes an EBS volume in GiB increments, rounding up to the next possible allocatable unit, and/or modifies an EBS
// volume with the parameters in ModifyDiskOptions.
// The resizing operation is performed only when newSizeBytes != 0.
// It returns the volume size after this call or an error if the size couldn't be determined or the volume couldn't be modified.
func (c *cloud) ResizeOrModifyDisk(ctx context.Context, volumeID string, newSizeBytes int64, options *ModifyDiskOptions) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// EBS doesn't handle empty outpost arn, so we have to include it only when it's non-empty

// Wrap error to preserve original message from AWS as to why this was an invalid argument

// If the volume modification isn't immediately completed, wait for it to finish

// Perform one final check on the volume

func (c *cloud) DeleteDisk(ctx context.Context, volumeID string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// execBatchDescribeInstances executes a batched DescribeInstances API call.
func execBatchDescribeInstances(svc util.EC2API, input []string, cache expiringcache.ExpiringCache[string, struct{}]) (map[string]*types.Instance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// batchDescribeInstances processes a DescribeInstances request by queuing the task and waiting for the result.
func (c *cloud) batchDescribeInstances(request *ec2.DescribeInstancesInput) (*types.Instance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getCardCount returns the number of EBS cards for a given instance type,
// using a cache to avoid repeated API calls. Falls back to the static table
// if the API call fails.
func (c *cloud) getCardCount(ctx context.Context, instanceType string) int {
	_ = "STUB: not implemented"
	return 0
}

func (c *cloud) AttachDisk(ctx context.Context, volumeID, nodeID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// If block device is "in use", that likely indicates a bad name that is in use by a block
// device that we do not know about (example: block devices attached in the AMI, which are
// not reported in DescribeInstance's block device map)
//
// Store such bad names in the "likely bad" map to be considered last in future attempts

// This is the only situation where we taint the device

// TODO: Check volume capability matches for ALREADY_EXISTS
// This could happen when request volume already attached to request node,
// but is incompatible with the specified volume_capability or readonly flag

func (c *cloud) attachDiskHyperPod(ctx context.Context, volumeID, nodeID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Construct real SageMaker AttachClusterNodeVolumeInput

// Wait for attachment completion

// HyperPod doesn't use card indexes

func (c *cloud) DetachDisk(ctx context.Context, volumeID, nodeID string) error {
	_ = "STUB: not implemented"
	return nil
}

// We expect it to be nil, it is (maybe) interesting if it is not

func (c *cloud) detachDiskHyperPod(ctx context.Context, volumeID, nodeID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Construct real SageMaker DetachClusterNodeVolumeInput

// Wait for detachment completion

type volumeInitialization struct {
	initialized                 bool
	estimatedInitializationTime time.Time
}

// IsVolumeInitialized calls EC2 DescribeVolumeStatus and returns whether the volume is initialized.
func (c *cloud) IsVolumeInitialized(ctx context.Context, volumeID string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Because volumes can take hours to initialize, we shouldn't poll DescribeVolumeStatus (DVS) as aggressively as we
// do for other EC2 APIs.
//
// We use a volumeInitializations cache to keep track of initializing volumes that we should poll at a slower rate.
//
// Furthermore, if initializationRate was set during volume creation, DVS returns an estimated initialization time.
// We cache that estimate and defer polling of DVS until we reach that time.
// We clamp to a minimum of 1 min because as of July 2025 it can take up to 5 min for volume initialization info to update.

// Check volumeInitializations cache to potentially delay EC2 DescribeVolumeStatus call

// Case 1: We've never called DVS for volume. Call DVS ASAP.

/* callASAP */
// Case 2: We already know volume is initialized. Don't call DVS.

// Case 3: We know volume is initializing, but there is no SLA. Call DVS eventually during next slow batch.

/* callASAP */
// Case 4: We have an estimated time for initialization. Wait to call DVS again until then unless RPC ctx is done.

/* callASAP */

// Parse volume status

// Update cache

// Clamp to a minimum of 1 min because as of July 2025 it can take up to 5 min for volume initialization info to update.

func isVolumeStatusInitializing(vsi types.VolumeStatusItem) bool {
	_ = "STUB: not implemented"
	return false
}

func execBatchDescribeVolumeStatus(svc util.EC2API, input []string) (map[string]*types.VolumeStatusItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// describeVolumeStatus will return the VolumeStatusItem associated with volumeID from EC2 DescribeVolumeStatus
// Set callASAP to true if you need status within seconds (Otherwise it may take minutes).
func (c *cloud) describeVolumeStatus(volumeID string, callASAP bool) (*types.VolumeStatusItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WaitForAttachmentState polls until the attachment status is the expected value.
func (c *cloud) WaitForAttachmentState(ctx context.Context, expectedState types.VolumeAttachmentState, volumeID string, expectedInstance string, expectedDevice string, alreadyAssigned bool, expectedCardIndex *int32) (*types.VolumeAttachment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The VolumeNotFound error is special -- we don't need to wait for it to repeat

// The disk doesn't exist, assume it's detached, log warning and stop waiting

// The disk doesn't exist, complain, give up waiting and report error

// Sometimes, a volume can get stuck attaching, for example when an instance is over the attachment limit
// or due to an EBS-side issue. When a volume has reached an extremely abnormal amount of time attaching,
// abort the attachment by calling DetachVolume and failing the ControllerPublishVolume RPC entirely to
// force a retry to occur with a fresh slate.

// Check card index if expected

// if we expected volume to be attached and it was reported as already attached via DescribeInstance call
// but DescribeVolume told us volume is detached, we will short-circuit this long wait loop and return error
// so as AttachDisk can be retried without waiting for 20 minutes.

// Attachment is in requested state, finish waiting

// But first, reset attachment to nil if expectedState equals volumeDetachedState.
// Caller will not expect an attachment to be returned for a detached volume if we're not also returning an error.

// continue waiting

func (c *cloud) GetDiskByName(ctx context.Context, name string, capacityBytes int64) (*Disk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *cloud) GetDiskByID(ctx context.Context, volumeID string) (*Disk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *cloud) GetVolumeIDByNodeAndDevice(ctx context.Context, nodeID string, deviceName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Only for hyperpod node, getInstanceIDFromHyperPodNode extracts the EC2 instance ID from a HyperPod node ID.
func getInstanceIDFromHyperPodNode(nodeID string) string { _ = "STUB: not implemented"; return "" }

// Only for hyperpod node, buildHyperPodClusterArn: arn:aws:sagemaker:region:account:cluster/clusterID.
func buildHyperPodClusterArn(nodeID string, region string, accountID string) string {
	_ = "STUB: not implemented"
	return ""
}

// For hyperpod node, AssociatedResource is in arn:aws:sagemaker:region:account:cluster/clusterID-instanceId format.
func getInstanceIDFromAssociatedResource(arn string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// execBatchDescribeSnapshots executes a batched DescribeSnapshots API call depending on the type of batcher.
func execBatchDescribeSnapshots(svc util.EC2API, input []string, batcher snapshotBatcherType, cache expiringcache.ExpiringCache[string, struct{}]) (map[string]*types.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// batchDescribeSnapshots processes a DescribeSnapshots request. Depending on the request,
// it determines the appropriate batcher to use, queues the task, and waits for the result.
func (c *cloud) batchDescribeSnapshots(request *ec2.DescribeSnapshotsInput) (*types.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// extractSnapshotKey retrieves the key associated with a given snapshot based on the batcher type.
// For the snapshotIDBatcher type, it returns the snapshot's ID.
// For other types, it searches for the SnapshotNameTagKey within the snapshot's tags.
func extractSnapshotKey(s *types.Snapshot, batcher snapshotBatcherType) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *cloud) CreateSnapshot(ctx context.Context, volumeID string, snapshotOptions *SnapshotOptions) (snapshot *Snapshot, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *cloud) LockSnapshot(ctx context.Context, lockOptions *SnapshotLockOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *cloud) DeleteSnapshot(ctx context.Context, snapshotID string) (success bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *cloud) GetSnapshotByName(ctx context.Context, name string) (snapshot *Snapshot, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *cloud) GetSnapshotByID(ctx context.Context, snapshotID string) (snapshot *Snapshot, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListSnapshots retrieves AWS EBS snapshots for an optionally specified volume ID.  If maxResults is set, it will return up to maxResults snapshots.  If there are more snapshots than maxResults,
// a next token value will be returned to the client as well.  They can use this token with subsequent calls to retrieve the next page of results.  If maxResults is not set (0),
// there will be no restriction up to 1000 results (https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#DescribeSnapshotsInput).
func (c *cloud) ListSnapshots(ctx context.Context, volumeID string, maxResults int32, nextToken string) (listSnapshotsResponse *ListSnapshotsResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Helper method converting EC2 snapshot type to the internal struct.
func (c *cloud) ec2SnapshotResponseToStruct(ec2Snapshot types.Snapshot) *Snapshot {
	_ = "STUB: not implemented"
	return nil
}

func (c *cloud) EnableFastSnapshotRestores(ctx context.Context, availabilityZones []string, snapshotID string) (*ec2.EnableFastSnapshotRestoresOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DryRun will make a dry-run EC2 API call. Nil return value means we successfully received EC2 DryRunOperation error code.
func (c *cloud) DryRun(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Rely on EC2 DAZ because it is required in ebs controller IAM role, but not in instance default role.

// Don't retry so we can catch network failures. CO should retry liveness check multiple times.
// Don't add our logging/metrics middleware because we expect errors.

func describeVolumes(ctx context.Context, svc util.EC2API, request *ec2.DescribeVolumesInput) ([]types.Volume, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *cloud) getVolume(ctx context.Context, request *ec2.DescribeVolumesInput) (*types.Volume, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func describeInstances(ctx context.Context, svc util.EC2API, request *ec2.DescribeInstancesInput) ([]types.Instance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *cloud) getInstance(ctx context.Context, nodeID string) (*types.Instance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetInstancesPatching returns the instance info associated with each node ID in `nodeIDs` and uses pagination
// to get instances for large clusters. The instances are also described in batches of size up to `maxInstancesDescribed`.
func (c *cloud) GetInstancesPatching(ctx context.Context, nodeIDs []string) ([]*types.Instance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *cloud) getInstancesPatchingBatch(ctx context.Context, nodeIDs []string) ([]*types.Instance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func describeSnapshots(ctx context.Context, svc util.EC2API, request *ec2.DescribeSnapshotsInput) ([]types.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *cloud) getSnapshot(ctx context.Context, request *ec2.DescribeSnapshotsInput) (*types.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// listSnapshots returns all snapshots based from a request.
func (c *cloud) listSnapshots(ctx context.Context, request *ec2.DescribeSnapshotsInput) (*ec2ListSnapshotsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// waitForVolume waits for volume to be in the "available" state.
func (c *cloud) waitForVolume(ctx context.Context, volumeID string) (*types.Volume, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getAccountID returns the account ID of the AWS Account for the IAM credentials in use.
//
// In the first call (or any calls made before the first call succeeds), getAccountID
// will attempt to determine the Account ID via sts:GetCallerIdentity.
// This attempt will retry indefinitely, however getAccountID will return when ctx is cancelled,
// leaving the account ID thread to run in the background.
//
// In subsequent calls (after the first success), getAccountID will use a cached value.
func (c *cloud) getAccountID(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Start background thread if it isn't already.
// Intentionally runs in the background until account ID is retrieved, so we don't pass the context.
//nolint:contextcheck

// Once.Do blocks until the function exits, even if we aren't the first caller.
// So the account ID must be available now.

// isAWSError returns a boolean indicating whether the error is AWS-related
// and has the given code. More information on AWS error codes at:
// https://docs.aws.amazon.com/AWSEC2/latest/APIReference/errors-overview.html
func isAWSError(err error, code string) bool { _ = "STUB: not implemented"; return false }

// isAWSErrorInstanceNotFound returns a boolean indicating whether the
// given error is an AWS InvalidInstanceID.NotFound error. This error is
// reported when the specified instance doesn't exist.
func isAWSErrorInstanceNotFound(err error) bool { _ = "STUB: not implemented"; return false }

// isAWSErrorVolumeNotFound returns a boolean indicating whether the
// given error is an AWS InvalidVolume.NotFound error. This error is
// reported when the specified volume doesn't exist.
func isAWSErrorVolumeNotFound(err error) bool { _ = "STUB: not implemented"; return false }

// isAWSErrorIncorrectState returns a boolean indicating whether the
// given error is an AWS IncorrectState error. This error is
// reported when the resource is not in a correct state for the request.
func isAWSErrorIncorrectState(err error) bool { _ = "STUB: not implemented"; return false }

// isAWSErrorInvalidAttachmentNotFound returns a boolean indicating whether the
// given error is an AWS InvalidAttachment.NotFound error. This error is reported
// when attempting to detach a volume from an instance to which it is not attached.
func isAWSErrorInvalidAttachmentNotFound(err error) bool { _ = "STUB: not implemented"; return false }

// isAWSErrorModificationNotFound returns a boolean indicating whether the given
// error is an AWS InvalidVolumeModification.NotFound error.
func isAWSErrorModificationNotFound(err error) bool { _ = "STUB: not implemented"; return false }

// isAWSErrorSnapshotNotFound returns a boolean indicating whether the
// given error is an AWS InvalidSnapshot.NotFound error. This error is
// reported when the specified snapshot doesn't exist.
func isAWSErrorSnapshotNotFound(err error) bool { _ = "STUB: not implemented"; return false }

// isAWSErrorIdempotentParameterMismatch returns a boolean indicating whether the
// given error is an AWS IdempotentParameterMismatch error.
// This error is reported when the two request contains same client-token but different parameters.
func isAWSErrorIdempotentParameterMismatch(err error) bool { _ = "STUB: not implemented"; return false }

// isAWSErrorInvalidParameterCombination returns a boolean indicating whether the
// given error is an AWS InvalidParameterCombination error.
// This error is reported when the combination of parameters passed to ec2 makes the request invalid.
func isAWSErrorInvalidParameterCombination(err error) bool { _ = "STUB: not implemented"; return false }

// isAWSErrorBlockDeviceInUse returns a boolean indicating whether the
// given error appears to be a block device name already in use error.
func isAWSErrorBlockDeviceInUse(err error) bool { _ = "STUB: not implemented"; return false }

// isAWSErrorAttachmentLimitExceeded checks if the error is an AttachmentLimitExceeded error.
// This error is reported when the maximum number of attachments for an instance is exceeded.
func isAWSErrorAttachmentLimitExceeded(err error) bool { _ = "STUB: not implemented"; return false }

// isAWSHyperPodErrorAttachmentLimitExceeded checks if the error is an AttachmentLimitExceeded error.
// This error is reported when the maximum number of attachments for an instance is exceeded.
func isAWSHyperPodErrorAttachmentLimitExceeded(err error) bool {
	_ = "STUB: not implemented"
	return false
}

// isAWSHyperPodErrorVolumeNotFound returns a boolean indicating whether the
// given error is a ValidationException error. This error is
// reported when the specified volume doesn't exist.
func isAWSHyperPodErrorVolumeNotFound(err error) bool { _ = "STUB: not implemented"; return false }

// isAWSHyperPodErrorIncorrectState returns a boolean indicating whether the
// given error is a ValidationException error. This error is
// reported when the resource is not in a correct state for the request.
func isAWSHyperPodErrorIncorrectState(err error) bool { _ = "STUB: not implemented"; return false }

// isAWSHyperPodErrorInvalidAttachmentNotFound returns a boolean indicating whether the
// given error is a ValidationException error. This error is reported
// when attempting to detach a volume from an instance to which it is not attached.
func isAWSHyperPodErrorInvalidAttachmentNotFound(err error) bool {
	_ = "STUB: not implemented"
	return false
}

// isAWSErrorModificationSizeLimitExceeded checks if the error is a VolumeModificationSizeLimitExceeded error.
// This error is reported when the limit on a volume modification storage in a region is exceeded.
func isAWSErrorVolumeModificationSizeLimitExceeded(err error) bool {
	_ = "STUB: not implemented"
	return false
}

// isAWSErrorVolumeLimitExceeded checks if the error is a VolumeLimitExceeded error.
// This error is reported when the limit on the amount of volume storage is exceeded.
func isAWSErrorVolumeLimitExceeded(err error) bool { _ = "STUB: not implemented"; return false }

// isAwsErrorMaxIOPSLimitExceeded checks if the error is a MaxIOPSLimitExceeded error.
// This error is reported when the limit on the IOPS usage for a region is exceeded.
func isAwsErrorMaxIOPSLimitExceeded(err error) bool { _ = "STUB: not implemented"; return false }

// isAwsErrorSnapshotLimitExceeded checks if the error is a SnapshotLimitExceeded error.
// This error is reported when the limit on the number of snapshots that can be created is exceeded.
func isAwsErrorSnapshotLimitExceeded(err error) bool { _ = "STUB: not implemented"; return false }

// isAWSErrorInvalidParameter returns a boolean indicating whether the
// given error is caused by invalid parameters in a EC2 API request.
func isAWSErrorInvalidParameter(err error) bool { _ = "STUB: not implemented"; return false }

// Checks for desired size on volume by also verifying volume size by describing volume.
// This is to get around potential eventual consistency problems with describing volume modifications
// objects and ensuring that we read two different objects to verify volume state.
func (c *cloud) checkDesiredState(ctx context.Context, volumeID string, desiredSizeGiB int32, options *ModifyDiskOptions) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// AWS resizes in chunks of GiB (not GB)

// Check if there is a mismatch between the requested modification and the current volume
// If there is, the volume is still modifying and we should not return a success

// waitForVolumeModification waits for a volume modification to finish.
func (c *cloud) waitForVolumeModification(ctx context.Context, volumeID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Consider volumes that have never been modified as done

func describeVolumesModifications(ctx context.Context, svc util.EC2API, request *ec2.DescribeVolumesModificationsInput) ([]types.VolumeModification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getLatestVolumeModification returns the last modification of the volume.
func (c *cloud) getLatestVolumeModification(ctx context.Context, volumeID string, isBatchable bool) (*types.VolumeModification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// randomAvailabilityZone returns a random zone from the given region
// the randomness relies on the response of DescribeAvailabilityZones.
func (c *cloud) randomAvailabilityZone(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// AvailabilityZones returns availability zones from the given region.
func (c *cloud) AvailabilityZones(ctx context.Context) (map[string]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func needsVolumeModification(volume types.Volume, newSizeGiB int32, req *ModifyDiskOptions) bool {
	_ = "STUB: not implemented"
	return false

	//nolint:staticcheck // staticcheck suggests merging all of the below conditionals into one line,
	// but that would be extremely difficult to read
}

func getVolumeAttachmentsList(volume types.Volume) []string { _ = "STUB: not implemented"; return nil }

// Checks if a volume's IOPS can be increased on expansion to adhere to IopsPerGB ratio.
func (c *cloud) checkIfIopsIncreaseOnExpansion(existingTags []types.Tag) (allowAutoIncreaseIsSet bool, iopsPerGbVal int32, err error) {
	_ = "STUB: not implemented"
	return false, 0, nil
}

func (c *cloud) getVolumeModificationState(ctx context.Context, volumeID string) (*types.VolumeModification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *cloud) validateVolumeState(ctx context.Context, volumeID string, newSizeGiB int32, oldSizeGiB int32, options *ModifyDiskOptions) (bool, int32, error) {
	_ = "STUB: not implemented"
	return false, 0, nil
}

// latestMod can be nil if the volume has never been modified

// If volume is already modifying, detour to waiting for it to modify

func (c *cloud) validateModifyVolume(ctx context.Context, volumeID string, newSizeGiB int32, options *ModifyDiskOptions, volume types.Volume) (bool, int32, error) {
	_ = "STUB: not implemented"
	return false, 0, nil
}

// At this point, we know we are starting a new volume modification
// If we're asked to modify a volume to its current state, ignore the request and immediately return a success
// This is because as of March 2024, EC2 ModifyVolume calls that don't change any parameters still modify the volume

// Wait for any existing modifications to prevent race conditions where DescribeVolume(s) returns the new
// state before the volume is actually finished modifying

func volumeModificationDone(state string) bool { _ = "STUB: not implemented"; return false }

// Calculate actual IOPS for a volume and cap it at supported AWS limits. Any limit of 0 is considered "infinite" (i.e. is not applied).
func capIOPS(volumeType string, requestedCapacityGiB int32, requestedIops int32, iopsLimits iopsLimits, allowIncrease bool) int32 {
	_ = "STUB: not implemented"
	// If requestedIops is zero the user did not request a specific amount, and the default will be used instead
	return 0
}

// Gets IOPS limits for a specific volume type in a specific Zone and caches it. If the limits are cached, simply return limits.
func (c *cloud) getVolumeLimits(ctx context.Context, volumeType string, azParams getVolumeLimitsParams) (iopsLimits iopsLimits) {
	_ = "STUB: not implemented"
	return *new(iopsLimits)
}

// Required by default EBS CSI Driver IAM policy.

// We only want either the Zone or ZoneID otherwise the DryRun call will fail with InvalidParameterCombination due to having both.

// Don't add our logging/metrics middleware because we expect errors.

// If DryRun unexpectedly succeeds, we use fallback values.

// Default To Hardcoded Limits if we can't get the max IOPS from the error message.

// Set minIops and maxIopsPerGb because we do not fetch these from DryRun Error, we can also catch invalid volume.

// Get what the maxIops is from DryRun error message.
func extractMaxIOPSFromError(errorMsg string, volumeType string) (int32, error) {
	_ = "STUB: not implemented"
	// Volume does not support IOPS, so return a limit of 0 (considered infinite in capIOPS).
	return 0, nil
}

// io1 and gp3 have the same error message but io2 has different one depending on the availability zone.

// No real overflow concern here but adding for safety.
