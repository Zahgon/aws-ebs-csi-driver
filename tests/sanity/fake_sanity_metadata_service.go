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
	"github.com/aws/aws-sdk-go-v2/aws/arn"
)

type fakeMetadataService struct {
	instanceID       string
	region           string
	availabilityZone string
	outpostArn       arn.ARN
}

func newFakeMetadataService(id string, r string, az string, oa arn.ARN) *fakeMetadataService {
	_ = "STUB: not implemented"
	return nil
}

func (m *fakeMetadataService) UpdateMetadata() error { _ = "STUB: not implemented"; return nil }

func (m *fakeMetadataService) GetInstanceID() string { _ = "STUB: not implemented"; return "" }

func (m *fakeMetadataService) GetInstanceType() string { _ = "STUB: not implemented"; return "" }

func (m *fakeMetadataService) GetRegion() string { _ = "STUB: not implemented"; return "" }

func (m *fakeMetadataService) GetAvailabilityZone() string { _ = "STUB: not implemented"; return "" }

func (m *fakeMetadataService) GetNumAttachedENIs() int { _ = "STUB: not implemented"; return 0 }

func (m *fakeMetadataService) GetNumBlockDeviceMappings() int { _ = "STUB: not implemented"; return 0 }

func (m *fakeMetadataService) GetOutpostArn() arn.ARN {
	_ = "STUB: not implemented"
	return *new(arn.ARN)
}
