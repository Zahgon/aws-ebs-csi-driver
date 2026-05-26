/*
Copyright 2024 The Kubernetes Authors.

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
	"time"

	flag "github.com/spf13/pflag"
)

// Options contains options and configuration settings for the driver.
type Options struct {
	Mode Mode

	// Kubeconfig is an absolute path to a kubeconfig file.
	// If empty, the in-cluster config will be loaded.
	Kubeconfig string

	// #### Server options ####

	// Endpoint is the endpoint for the CSI driver server
	Endpoint string
	// HTTPEndpoint is the TCP network address where the HTTP server for metrics will listen
	HTTPEndpoint string
	// MetricsCertFile is the location of the certificate for serving the metrics server over HTTPS
	MetricsCertFile string
	// MetricsKeyFile is the location of the key for serving the metrics server over HTTPS
	MetricsKeyFile string
	// EnableOtelTracing is a flag to enable opentelemetry tracing for the driver
	EnableOtelTracing bool

	// #### Controller options ####

	// ExtraTags is a map of tags that will be attached to each dynamically provisioned
	// resource.
	ExtraTags map[string]string
	// ExtraVolumeTags is a map of tags that will be attached to each dynamically provisioned
	// volume.
	//
	// Deprecated: Use ExtraTags instead.
	ExtraVolumeTags map[string]string
	// ID of the kubernetes cluster.
	KubernetesClusterID string
	// flag to enable sdk debug log
	AwsSdkDebugLog bool
	// flag to warn on invalid tag, instead of returning an error
	WarnOnInvalidTag bool
	// flag to set user agent
	UserAgentExtra string
	// flag to enable batching of API calls
	Batching bool
	// flag to set the timeout for volume modification requests to be coalesced into a single
	// volume modification call to AWS.
	ModifyVolumeRequestHandlerTimeout time.Duration
	// flag to enable deprecated metrics
	DeprecatedMetrics bool
	// flag to enable node-local volume support
	EnableNodeLocalVolumes bool

	// #### Node options #####

	// VolumeAttachLimit specifies the value that shall be reported as "maximum number of attachable volumes"
	// in CSINode objects. It is similar to https://kubernetes.io/docs/concepts/storage/storage-limits/#custom-limits
	// which allowed administrators to specify custom volume limits by configuring the kube-scheduler. Also, each AWS
	// machine type has different volume limits. By default, the EBS CSI driver parses the machine type name and then
	// decides the volume limit. However, this is only a rough approximation and not good enough in most cases.
	// Specifying the volume attach limit via command line is the alternative until a more sophisticated solution presents
	// itself (dynamically discovering the maximum number of attachable volume per EC2 machine type, see also
	// https://github.com/kubernetes-sigs/aws-ebs-csi-driver/issues/347).
	VolumeAttachLimit int64
	// ReservedVolumeAttachments specifies number of volume attachments reserved for system use.
	// Typically 1 for the root disk, but may be larger when more system disks are attached to nodes.
	// This option is not used when --volume-attach-limit is specified.
	// When -1, the amount of reserved attachments is loaded from instance metadata that captured state at node boot
	// and may include not only system disks but also CSI volumes (and therefore it may be wrong).
	ReservedVolumeAttachments int
	// ALPHA: WindowsHostProcess indicates whether the driver is running in a Windows privileged container
	WindowsHostProcess bool
	// LegacyXFSProgs formats XFS volumes with `bigtime=0,inobtcount=0,reflink=0,nrext64=0`, so that they can be mounted onto nodes with linux kernel ≤ v5.4. Volumes formatted with this option may experience issues after 2038, and will be unable to use some XFS features (for example, reflinks).
	LegacyXFSProgs bool
	// CsiMountPointPath is the path where CSI volumes are expected to be mounted on the node.
	CsiMountPointPath string
	// MetadataSources dictates which sources are used to retrieve instance metadata.
	// The driver will attempt to rely on each source in order until one succeeds.
	// Valid options include 'imds' and 'kubernetes'.
	MetadataSources []string
}

func (o *Options) AddFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

// Server options

// AWS SDK options, shared by all modes that create a cloud client

// Controller options

// Node options

func (o *Options) Validate() error { _ = "STUB: not implemented"; return nil }
