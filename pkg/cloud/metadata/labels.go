// Copyright 2025 The Kubernetes Authors.
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

package metadata

import (
	"context"
	"sync"
	"time"

	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/cloud"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

const (
	// ControllerMetadataLabelerInterval is the interval metadata-labeler mode refreshes node labels with volume and ENI count.
	ControllerMetadataLabelerInterval = 60 * time.Minute

	// patchFails is the number of nodes we fail to patch before returning an error.
	patchFails = 5

	// numWorkersPatchLabels is the number of worker threads patching node labels.
	numWorkersPatchLabels = 10
)

// Initialized in ContinuousUpdateLabelsLeaderElection (depends on driver name).
var (
	// Prevent races in initialization.
	once sync.Once

	// VolumesLabel is the label name for the number of volumes on a node.
	VolumesLabel string

	// ENIsLabel is the label name for the number of ENIs on a node.
	ENIsLabel string
)

type enisVolumes struct {
	ENIs    int
	Volumes int
}

// initVariables initializes variables that depend on driver name.
// Separated into a spearate function from ContinuousUpdateLabelsLeaderElection so it can be called in tests.
func initVariables() { _ = "STUB: not implemented"; return }

// ContinuousUpdateLabelsLeaderElection uses leader election so that only one controller pod calls continuousUpdateLabels().
func ContinuousUpdateLabelsLeaderElection(clientset kubernetes.Interface, cloud cloud.Cloud, updateTime time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// continuousUpdateLabels is a go routine that updates the metadata labels of each node once every
// `updateTime` minutes and uses an informer to update the labels of new nodes that join the cluster.
// A PV informer is also used to keep track of CSI managed volumes when updating labels to avoid
// double counting.
func continuousUpdateLabels(ctx context.Context, k8sClient kubernetes.Interface, cloud cloud.Cloud, updateTime time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func volumeIDIndexFunc(obj any) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func getNonCSIManagedVolumes(pvInformer cache.SharedIndexInformer, volumes []ec2types.InstanceBlockDeviceMapping) int {
	_ = "STUB: not implemented"
	return 0
}

// patchNewNodes patches metadata labels for new nodes that join the cluster.
func patchNewNodes(ctx context.Context, clientset kubernetes.Interface, cloud cloud.Cloud, nodesInformer, pvInformer cache.SharedIndexInformer) error {
	_ = "STUB: not implemented"
	return nil
}

func updateLabels(ctx context.Context, k8sClient kubernetes.Interface, cloud cloud.Cloud, pvCache cache.SharedIndexInformer) error {
	_ = "STUB: not implemented"
	return nil
}

func getNodes(ctx context.Context, kubeclient kubernetes.Interface) (*v1.NodeList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func updateMetadataEC2(ctx context.Context, kubeclient kubernetes.Interface, cloud cloud.Cloud, nodes *v1.NodeList, pvInformer cache.SharedIndexInformer) error {
	_ = "STUB: not implemented"
	return nil
}

// getMetadata calls the EC2 API to get the number of ENIs and non-CSI managed volumes attached to each node.
func getMetadata(ctx context.Context, cloud cloud.Cloud, nodes *v1.NodeList, pvInformer cache.SharedIndexInformer) (map[string]enisVolumes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// -1 for root volume because we eventually add this back in when calculating allocatable count in getVolumesLimit()

// patchNodes patches the labels of each node to have the number of ENIs and non-CSI managed volumes attached to each node.
func patchNodes(ctx context.Context, nodes *v1.NodeList, enisVolumeMap map[string]enisVolumes, clientset kubernetes.Interface, patchFails int) error {
	_ = "STUB: not implemented"
	return nil
}

func patchSingleNode(ctx context.Context, node v1.Node, enisVolumeMap map[string]enisVolumes, clientset kubernetes.Interface) error {
	_ = "STUB: not implemented"
	return nil
}
