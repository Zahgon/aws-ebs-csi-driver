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

// PLUGIN AUTHORS: DO NOT MODIFY THIS FILE
// This file contains the driver-maintained plugin logic and interfaces.
//
// In order to create a new plugin, create a new .go file in this package
// that implements your plugin logic, see plugin.go.sample and docs/plugins.md.

//nolint:unused // Functions in this file are only used if a plugin is loaded
package plugin

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/pflag"
)

// Plugin stores the currently loaded plugin.
var plugin EbsCsiPlugin = nil

// GetPlugin returns the currently loaded plugin. It will return nil if no plugin is loaded.
func GetPlugin() EbsCsiPlugin {
	_ = "STUB: not implemented"

	// loadPlugin loads a plugin into memory.
	return *new(EbsCsiPlugin)
}

func loadPlugin(pluginToLoad EbsCsiPlugin) {
	_ = "STUB: not implemented"

	// Multiple plugins are not currently supported
	// Thus, exit as quickly as possible
	return
}

// EbsCsiPlugin is the common plugin interface all plugins implement.
type EbsCsiPlugin interface {
	// Init is called extremely early to get CLI flags, before Init
	// Do not use for any non-CLI initialization - do that in Init()
	InitFlags(fs *pflag.FlagSet)
	// Init is called on startup, use to perform initialization
	// Returning an error will cause the driver process to exit
	//
	// NOTE: registry can be nil if the user did not enable metrics,
	// plugins must handle this case gracefully
	Init(region string, registry *prometheus.Registry) error

	// GetEC2Client replaces the AWS EC2 client the driver uses
	GetEC2Client(cfg aws.Config, optFns ...func(*ec2.Options)) util.EC2API
	// GetSageMakerClient replaces the AWS SageMaker client the driver uses
	GetSageMakerClient(cfg aws.Config, optFns ...func(*sagemaker.Options)) util.SageMakerAPI

	// GetDriverName replaces the driver name in use (normally "ebs.csi.aws.com")
	// This function can be called before Init and should not depend on it
	GetDriverName() string

	// GetDiskTopologySegments provides a set of additional topology segments that should be applied to volumes created by
	// the driver. These segments are merged with the existing topology segments applied by the base driver. Segments
	// provided by the plugin take precedent.
	GetDiskTopologySegments() map[string]string
	// GetNodeTopologySegments provides a set of additional topology segments that should be resolved for nodes running
	// the driver. These segments are merged with the existing topology segments resolved by the base driver. Segments
	// provided by the plugin take precedent.
	GetNodeTopologySegments() map[string]string
}

// ebsCsiPluginBase implements stub functionality of all plugin methods except Init().
// It is strongly recommended to embed into plugin implementations to prevent bulid failures
// if/when new functions are added to the EbsCsiPlugin interface.
type ebsCsiPluginBase struct{}

func (p *ebsCsiPluginBase) InitFlags(_ *pflag.FlagSet) {
	_ = "STUB: not implemented"
	// Do nothing intentionally.
	return
}

func (p *ebsCsiPluginBase) GetEC2Client(_ aws.Config, _ ...func(o *ec2.Options)) util.EC2API {
	_ = "STUB: not implemented"
	return *new(util.EC2API)
}

func (p *ebsCsiPluginBase) GetSageMakerClient(_ aws.Config, _ ...func(o *sagemaker.Options)) util.SageMakerAPI {
	_ = "STUB: not implemented"
	return *new(util.SageMakerAPI)
}

func (p *ebsCsiPluginBase) GetDriverName() string { _ = "STUB: not implemented"; return "" }

func (p *ebsCsiPluginBase) GetDiskTopologySegments() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (p *ebsCsiPluginBase) GetNodeTopologySegments() map[string]string {
	_ = "STUB: not implemented"
	return nil
}
