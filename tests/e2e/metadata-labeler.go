/*
Copyright 2025 The Kubernetes Authors.

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

package e2e

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	v1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/test/e2e/framework"
	admissionapi "k8s.io/pod-security-admission/api"
)

const (
	driverNamespace = "kube-system"
)

type instanceMetadata struct {
	ENIs             int
	Volumes          int
	InstanceType     string
	AllocatableCount int32
	NodeID           string
	AvailabilityZone string
}

var _ = framework.Describe("[ebs-csi-e2e] [Disruptive] Metadata Labeler Sidecar", framework.WithDisruptive(), func() {
	f := framework.NewDefaultFramework("ebs")
	f.NamespacePodSecurityEnforceLevel = admissionapi.LevelPrivileged

	var (
		ec2Client        *ec2.Client
		cs               clientset.Interface
		expectedMetadata map[string]*instanceMetadata
		labeledMetadata  map[string]*instanceMetadata
		cleanUp          []func()
	)

	BeforeEach(func() {
		cfg, err := config.LoadDefaultConfig(context.Background())
		Expect(err).NotTo(HaveOccurred(), "Failed to load AWS SDK config")
		ec2Client = ec2.NewFromConfig(cfg)

		cs = f.ClientSet

		labeledMetadata = make(map[string]*instanceMetadata)
	})

	AfterEach(func() {
		for i := len(cleanUp) - 1; i >= 0; i-- {
			cleanUp[i]()
		}
		deleteControllerPod(cs)
		checkLabelsUpdated(cs, labeledMetadata, expectedMetadata)
		By("Deleting the EBS CSI node pods to reset allocatable counts")
		for instance := range labeledMetadata {
			deleteNodePod(labeledMetadata[instance].NodeID, cs)
		}
		checkCSINodesUpdated(cs, labeledMetadata, expectedMetadata)
	})

	Describe("Node labeling volumes and ENIs", func() {
		It("should correctly label nodes with volume and ENI counts and have correct csinode allocatable counts", func() {
			By("Getting EC2 instance information")
			nodes, err := cs.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})

			clusterInstances := []string{}
			for _, clusterNode := range nodes.Items {
				clusterInstances = append(clusterInstances, parseProviderID(clusterNode.Spec.ProviderID))
			}

			Expect(err).NotTo(HaveOccurred(), "Failed to list nodes")
			resp, err := ec2Client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{
				InstanceIds: clusterInstances,
			})
			Expect(err).NotTo(HaveOccurred(), "Failed to describe EC2 instances")
			expectedMetadata = getVolENIs(resp)

			By("Checking initial node labels")
			checkVolENI(expectedMetadata, labeledMetadata, nodes)

			By("Checking CSI node allocatable counts")
			csiNodes, err := cs.StorageV1().CSINodes().List(context.TODO(), metav1.ListOptions{})
			Expect(err).NotTo(HaveOccurred(), "Failed to list CSI nodes")
			checkAllocatable(expectedMetadata, labeledMetadata, csiNodes)

			// For the instance that the new volume is attached to, the volume labels should increase by 1 and the allocatable count should decrease by 1
			By("Creating a new non CSI managed volume")
			firstChangedNonCSIVolumeInstance, firstCreatedNonCSIVolumeID := createNonCSIManagedVolume(ec2Client, expectedMetadata, "")
			cleanUp = append(cleanUp, func() {
				cleanUpVolume(firstCreatedNonCSIVolumeID, firstChangedNonCSIVolumeInstance, ec2Client, expectedMetadata)
			})

			By("Attaching the new non CSI managed volume")
			attachVolume(ec2Client, firstCreatedNonCSIVolumeID, firstChangedNonCSIVolumeInstance, "/dev/sdz", expectedMetadata)
			By("Deleting the EBS CSI controller pods to trigger Node label update")
			deleteControllerPod(cs)
			checkLabelsUpdated(cs, labeledMetadata, expectedMetadata)
			By("Deleting the EBS CSI node pod to trigger CSINode allocatable update")
			deleteNodePod(labeledMetadata[firstChangedNonCSIVolumeInstance].NodeID, cs)
			checkCSINodesUpdated(cs, labeledMetadata, expectedMetadata)

			// For the instance that the new volume is attached to, the volume labels and allocatable count should not change
			By("Creating and attaching a new CSI managed volume")
			createStorageClass(cs)
			cleanUp = append(cleanUp, func() { cleanUpStorageClass(cs, "ebs-sc") })
			pvc := createPVC(cs, f.Namespace.Name)
			cleanUp = append(cleanUp, func() { cleanUpPVC(cs, f.Namespace.Name, "ebs-claim") })
			pod := createPod(cs, f.Namespace.Name)
			cleanUp = append(cleanUp, func() { cleanUpPod(cs, f.Namespace.Name, "app") })
			changedCSIVolumeInstance := createCSIManagedVolume(cs, pvc, pod, f.Namespace.Name)

			// Because the previous step should not change volume labels/allocatable count, we add a non CSI managed volume to know that the
			// volume labels/allocatable count updated accordingly
			By("Creating a new non CSI managed volume")
			changedNonCSIVolumeInstance, createdNonCSIVolumeID := createNonCSIManagedVolume(ec2Client, expectedMetadata, changedCSIVolumeInstance)
			cleanUp = append(cleanUp, func() { cleanUpVolume(createdNonCSIVolumeID, changedNonCSIVolumeInstance, ec2Client, expectedMetadata) })
			By("Attaching the new non CSI managed volume")
			attachVolume(ec2Client, createdNonCSIVolumeID, changedNonCSIVolumeInstance, "/dev/sdy", expectedMetadata)
			By("Deleting the EBS CSI controller pods to trigger Node label update")
			deleteControllerPod(cs)
			checkLabelsUpdated(cs, labeledMetadata, expectedMetadata)
			By("Deleting the EBS CSI node pod to trigger CSINode allocatable update")
			deleteNodePod(labeledMetadata[changedNonCSIVolumeInstance].NodeID, cs)
			checkCSINodesUpdated(cs, labeledMetadata, expectedMetadata)

			By("Verifying updated node labels")
			updatedNodes, err := cs.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
			Expect(err).NotTo(HaveOccurred(), "Failed to list updated nodes")
			checkVolENI(expectedMetadata, labeledMetadata, updatedNodes)

			By("Verifying updated CSI node allocatable counts")
			updatedCsiNodes, err := cs.StorageV1().CSINodes().List(context.Background(), metav1.ListOptions{})
			Expect(err).NotTo(HaveOccurred(), "Failed to list updated CSI nodes")
			checkAllocatable(expectedMetadata, labeledMetadata, updatedCsiNodes)
		})
	})
})

// getAllocatableCount returns the limit of volumes that the node supports.
func getAllocatableCount(volumes, enis int) int32 { _ = "STUB: not implemented"; return 0 }

// getVolENIs gets the expected metadata of each instance from the ec2 API
func getVolENIs(resp *ec2.DescribeInstancesOutput) map[string]*instanceMetadata {
	_ = "STUB: not implemented"
	return nil
}

// we do not include the root volume in the expected number of volumes attached

// checkVolENI compares `expectedMetadata` and `labeledMetadata` to have the same number of volumes and ENIs attached to each node in `nodes`
func checkVolENI(expectedMetadata, labeledMetadata map[string]*instanceMetadata, nodes *corev1.NodeList) {
	_ = "STUB: not implemented"
	return
}

// checkAllocatable compares `expectedMetadata` and `labeledMetadata` to have the same allocatable count on each node in `nodes`
func checkAllocatable(expectedMetadata, labeledMetadata map[string]*instanceMetadata, csiNodes *storagev1.CSINodeList) {
	_ = "STUB: not implemented"
	return
}

func createCSIManagedVolume(cs kubernetes.Interface, pvc *corev1.PersistentVolumeClaim, pod *corev1.Pod, namespace string) string {
	_ = "STUB: not implemented"
	return ""
}

func checkLabelsUpdated(cs kubernetes.Interface, labeledMetadata, expectedMetadata map[string]*instanceMetadata) {
	_ = "STUB: not implemented"
	return
}

func checkCSINodesUpdated(cs kubernetes.Interface, labeledMetadata, expectedMetadata map[string]*instanceMetadata) {
	_ = "STUB: not implemented"
	return
}

func createNonCSIManagedVolume(ec2svc *ec2.Client, metadata map[string]*instanceMetadata, changedCSIVolumeInstance string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// a random instance is chosen for the test

func attachVolume(ec2svc *ec2.Client, volumeID, instanceID, device string, metadata map[string]*instanceMetadata) bool {
	_ = "STUB: not implemented"
	return false
}

func deleteNodePod(nodeID string, cs clientset.Interface) { _ = "STUB: not implemented"; return }

func deleteControllerPod(cs clientset.Interface) { _ = "STUB: not implemented"; return }

func parseProviderID(providerID string) string { _ = "STUB: not implemented"; return "" }

func createStorageClass(cs kubernetes.Interface) *v1.StorageClass {
	_ = "STUB: not implemented"
	return nil
}

func createPVC(cs kubernetes.Interface, namespace string) *corev1.PersistentVolumeClaim {
	_ = "STUB: not implemented"
	return nil
}

func createPod(cs kubernetes.Interface, namespace string) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func cleanUpPod(cs kubernetes.Interface, namespace, name string) { _ = "STUB: not implemented"; return }

func cleanUpVolume(volumeID, instanceID string, ec2Client *ec2.Client, expectedMetadata map[string]*instanceMetadata) {
	_ = "STUB: not implemented"
	return
}

func cleanUpPVC(cs kubernetes.Interface, namespace, name string) { _ = "STUB: not implemented"; return }

func cleanUpStorageClass(cs kubernetes.Interface, name string) { _ = "STUB: not implemented"; return }
