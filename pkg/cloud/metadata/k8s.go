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

package metadata

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	LabelSageMakerENICount                 = "sagemaker.amazonaws.com/num-eni-attachments"
	LabelSageMakerBlockDeviceMappingsCount = "sagemaker.amazonaws.com/num-block-device-mappings"
)

type KubernetesAPIClient func() (kubernetes.Interface, error)

func DefaultKubernetesAPIClient(kubeconfig string) KubernetesAPIClient {
	_ = "STUB: not implemented"
	return *new(KubernetesAPIClient)
}

// creates the in-cluster config

// CONTAINER_SANDBOX_MOUNT_POINT env is set upon container creation in containerd v1.6+
// it provides the absolute host path to the container volume.

// #nosec G703 -- tokenFile is built via filepath.Join with controlled path components

// creates the clientset

func KubernetesAPIInstanceInfo(clientset kubernetes.Interface, metadataLabeler bool) (*Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get node with k8s API

// Default: All nodes have at least 1 attached ENI
// Default: 0

// Initialize ENIsLabel and VolumesLabel globals.

//nolint: nilerr // Want to catch retry all errs until context times out

// sagemaker instance type has 'ml.' prefix, remove 'ml.' prefix

// Only let metadata.UpdateMetadata work for metadataLabeler data source

func getEC2ENIsVolumes(node *corev1.Node) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func getLabelAsInt(node *corev1.Node, label string, defaultValue int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func parseProviderID(node *corev1.Node) (string, error) { _ = "STUB: not implemented"; return "", nil }
