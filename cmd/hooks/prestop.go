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

package hooks

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

/*
When a node is terminated, persistent workflows using EBS volumes can take 6+ minutes to start up again.
This happens when a volume is not cleanly unmounted, which causes the Attach/Detach controller (in kube-controller-manager)
to wait for 6 minutes before issuing a force detach and allowing the volume to be attached to another node.

This PreStop lifecycle hook aims to ensure that before the node (and the CSI driver node pod running on it) is shut down,
all VolumeAttachment objects associated with that node are removed, thereby indicating that all volumes have been successfully unmounted and detached.

No unnecessary delay is added to the termination workflow, as the PreStop hook logic is only executed when the node is being drained
(thus preventing delays in termination where the node pod is killed due to a rolling restart, or during driver upgrades, but the workload pods are expected to be running).
If the PreStop hook hangs during its execution, the driver node pod will be forcefully terminated after terminationGracePeriodSeconds, defined in the pod spec.
*/

const clusterAutoscalerTaint = "ToBeDeletedByClusterAutoscaler"
const v1KarpenterTaint = "karpenter.sh/disrupted"
const v1beta1KarpenterTaint = "karpenter.sh/disruption"

// drainTaints includes taints used by K8s or autoscalers that signify node draining or pod eviction.
var drainTaints = map[string]struct{}{
	v1.TaintNodeUnschedulable: {}, // Kubernetes common eviction taint (kubectl drain)
	clusterAutoscalerTaint:    {},
	v1KarpenterTaint:          {},
	v1beta1KarpenterTaint:     {},
}

func PreStop(clientset kubernetes.Interface) error { _ = "STUB: not implemented"; return nil }

func fetchNode(clientset kubernetes.Interface, nodeName string) (*v1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// isNodeBeingDrained returns true if node resource has a known drain/eviction taint.
func isNodeBeingDrained(node *v1.Node) bool { _ = "STUB: not implemented"; return false }

func waitForVolumeAttachments(clientset kubernetes.Interface, nodeName string) error {
	_ = "STUB: not implemented"
	return nil
}

func checkVolumeAttachments(clientset kubernetes.Interface, nodeName string, allAttachmentsDeleted chan struct{}) error {
	_ = "STUB: not implemented"
	return nil
}
