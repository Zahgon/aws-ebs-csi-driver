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
	csi "github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/mounter"
	"google.golang.org/grpc"
	"k8s.io/client-go/kubernetes"
)

// Mode is the operating mode of the CSI driver.
type Mode string

const (
	// ControllerMode is the mode that only starts the controller service.
	ControllerMode Mode = "controller"
	// NodeMode is the mode that only starts the node service.
	NodeMode Mode = "node"
	// AllMode is the mode that only starts both the controller and the node service.
	AllMode Mode = "all"

	// MetadataLabelerMode is the mode that starts the metadata labeler.
	MetadataLabelerMode Mode = "metadataLabeler"
)

const (
	WellKnownZoneTopologyKey = "topology.kubernetes.io/zone"
	// ZoneIDTopologyKey name is purposefully consistent with the CCM's ZoneID topology key.
	// This key is only used for provisioning by az-id and will not be used for node topology
	// to prevent any backwards compatibility issues.
	ZoneIDTopologyKey = "topology.k8s.aws/zone-id"
	OSTopologyKey     = "kubernetes.io/os"
)

// Initialized in NewDriver (depend on driver name).
var (
	AgentNotReadyNodeTaintKey string
	AwsPartitionKey           string
	AwsAccountIDKey           string
	AwsRegionKey              string
	AwsOutpostIDKey           string
	// Deprecated: Use the WellKnownZoneTopologyKey instead.
	ZoneTopologyKey string
)

type Driver struct {
	controller *ControllerService
	node       *NodeService
	srv        *grpc.Server
	options    *Options
	csi.UnimplementedIdentityServer
}

// initVariables initializes variables that depend on driver name.
// Separated into a spearate function from NewDriver so it can be called in tests.
func initVariables() { _ = "STUB: not implemented"; return }

// Deprecated: Use the WellKnownZoneTopologyKey instead.

func NewDriver(c cloud.Cloud, o *Options, m mounter.Mounter, md metadata.MetadataService, k kubernetes.Interface) (*Driver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Driver) Run() error { _ = "STUB: not implemented"; return nil }

func (d *Driver) Stop() { _ = "STUB: not implemented"; return }
