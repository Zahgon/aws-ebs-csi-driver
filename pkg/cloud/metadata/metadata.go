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

package metadata

import (
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"k8s.io/client-go/kubernetes"
)

// Metadata is info about the ec2 instance on which the driver is running.
type Metadata struct {
	InstanceID             string
	InstanceType           string
	Region                 string
	AvailabilityZone       string
	NumAttachedENIs        int
	NumBlockDeviceMappings int
	OutpostArn             arn.ARN
	IMDSClient             IMDS
	K8sAPIClient           kubernetes.Interface
}

type MetadataServiceConfig struct {
	MetadataSources []string
	IMDSClient      IMDSClient
	K8sAPIClient    KubernetesAPIClient
}

const (
	SourceIMDS            = "imds"
	SourceMetadataLabeler = "metadata-labeler"
	SourceK8s             = "kubernetes"
)

var (
	// DefaultMetadataSources lists the default fallback order of driver Metadata sources.
	DefaultMetadataSources = []string{SourceIMDS, SourceK8s}
)

var _ MetadataService = &Metadata{}

// NewMetadataService retrieves instance Metadata from one of the clients in MetadataServiceConfig.
// It tries each client included in MetadataServiceConfig.MetadataSources in order until one succeeds.
func NewMetadataService(cfg MetadataServiceConfig, region string) (MetadataService, error) {
	_ = "STUB: not implemented"
	return *new(MetadataService), nil
}

// Unexpected cases should have been caught during driver option validation

// UpdateMetadata refreshes metadata cache based upon driver startup metadata source.
func (m *Metadata) UpdateMetadata() error { _ = "STUB: not implemented"; return nil }

// We do not refresh blockDeviceMappings because IMDS only reports data from instance start (As of April 2025)

/* metadataLabeler */

func retrieveIMDSMetadata(imdsClient IMDSClient) (*Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func retrieveK8sMetadata(k8sAPIClient KubernetesAPIClient, metadataLabeler bool) (*Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Override the region on a Metadata object if it is non-empty.
func (m *Metadata) overrideRegion(region string) *Metadata { _ = "STUB: not implemented"; return nil }

// GetInstanceID returns the instance identification.
func (m *Metadata) GetInstanceID() string { _ = "STUB: not implemented"; return "" }

// GetInstanceType returns the instance type.
func (m *Metadata) GetInstanceType() string { _ = "STUB: not implemented"; return "" }

// GetRegion returns the region which the instance is in.
func (m *Metadata) GetRegion() string {
	_ = "STUB: not implemented"

	// GetAvailabilityZone returns the Availability Zone which the instance is in.
	return ""
}

func (m *Metadata) GetAvailabilityZone() string { _ = "STUB: not implemented"; return "" }

// GetNumAttachedENIs returns the number of attached ENIs.
func (m *Metadata) GetNumAttachedENIs() int { _ = "STUB: not implemented"; return 0 }

// GetNumBlockDeviceMappings returns the number of block device mappings.
func (m *Metadata) GetNumBlockDeviceMappings() int { _ = "STUB: not implemented"; return 0 }

// GetOutpostArn returns outpost arn if instance is running on an outpost. empty otherwise.
func (m *Metadata) GetOutpostArn() arn.ARN {
	_ = "STUB: not implemented"
	return *

	// InvalidSourceErr returns an error message when a metadata source is invalid.
	new(arn.ARN)
}

func InvalidSourceErr(sources []string, invalidSource string) error {
	_ = "STUB: not implemented"
	return nil
}

func sourcesUnavailableErr(metadataSources []string) error { _ = "STUB: not implemented"; return nil }
