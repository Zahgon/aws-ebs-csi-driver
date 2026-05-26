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
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
)

const (
	// OutpostArnEndpoint is the IMDS endpoint to query to get the outpost arn.
	OutpostArnEndpoint string = "outpost-arn"

	// EnisEndpoint is the IMDS endpoint to query the number of attached ENIs.
	EnisEndpoint string = "network/interfaces/macs"

	// BlockDevicesEndpoint is the IMDS endpoint to query the number of attached block devices.
	BlockDevicesEndpoint string = "block-device-mapping"
)

type IMDSClient func() (IMDS, error)

var DefaultIMDSClient = func() (IMDS, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}
	svc := imds.NewFromConfig(cfg)
	return svc, nil
}

func IMDSInstanceInfo(svc IMDS) (*Metadata, error) { _ = "STUB: not implemented"; return nil, nil }

// "outpust-arn" returns 404 for non-outpost instances. note that the request is made to a link-local address.
// it's guaranteed to be in the form `arn:<partition>:outposts:<region>:<account>:outpost/<outpost-id>`
// There's a case to be made here to ignore the error so a failure here wouldn't affect non-outpost calls.

func getAttachedENIs(svc IMDS) (int, error) { _ = "STUB: not implemented"; return 0, nil }
