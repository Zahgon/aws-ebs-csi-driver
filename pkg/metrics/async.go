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

package metrics

import (
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	metricAsyncDetachSeconds = namespace + "ec2_detach_pending_seconds_total"
	asyncCollectorScrapes    = namespace + "ec2_collector_scrapes_total"
	asyncCollectorDuration   = namespace + "ec2_collector_duration_seconds"
)

type attachment struct {
	volumeID   string
	instanceID string
}

type detachingVolume struct {
	detachStart          time.Time
	lastDetachStateCheck time.Time
	attachmentState      types.VolumeAttachmentState
}

// AsyncEC2Collector contains metrics related to async EC2 operations and utilities for tracking what metrics should be emitted.
type AsyncEC2Collector struct {
	// Metrics
	detachingDuration  *prometheus.Desc
	collectionDuration prometheus.Histogram
	scrapesTotal       prometheus.Counter

	// Utilities
	// detachingVolumes holds any attachment that the controller service expects to be detached but is not.
	// We manage concurrency and memory safety within the struct through mutex and ticker instead of relying
	// on an ExpiringCache because we require that getting a cached value doesn't reset expiration timer.
	detachingVolumes map[attachment]detachingVolume
	mutex            sync.Mutex
	ticker           *time.Ticker
	// lastCacheUpdate helps us not vend out-of-date metrics upon leader election change.
	lastCacheUpdate time.Time
	// minDurationThreshold for volume to not reach detached state for metric emission. Prevents cardinality bombs.
	minDurationThreshold time.Duration
}

// Describe sends the descriptor of each metric in the AsyncEC2Collector to Prometheus.
func (c *AsyncEC2Collector) Describe(ch chan<- *prometheus.Desc) { _ = "STUB: not implemented"; return }

// Collect is invoked by Prometheus at collection time for emitting AsyncEC2Collector metrics.
func (c *AsyncEC2Collector) Collect(ch chan<- prometheus.Metric) {
	_ = "STUB: not implemented"
	// Meta metrics for metric collection
	return
}

// TrackDetachment tracks the state of a volume that we expect to detach in our AsyncEC2Collector cache.
func (c *AsyncEC2Collector) TrackDetachment(volumeID, instanceID string, attachmentState types.VolumeAttachmentState) {
	_ = "STUB: not implemented"
	return
}

// Clear if detached

// Check if first time tracking this attachment

// ClearDetachMetric ensures AsyncEC2Collector is not emitting metrics for a given attachment.
func (c *AsyncEC2Collector) ClearDetachMetric(volumeID, instanceID string) {
	_ = "STUB: not implemented"
	return
}

// cleanupCache clears the detachingVolumes cache if no update has been made since minTimeSinceLastUpdate ago.
func (c *AsyncEC2Collector) cleanupCache(minTimeSinceLastUpdate time.Duration) {
	_ = "STUB: not implemented"
	return
}
