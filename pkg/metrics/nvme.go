//go:build linux

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

package metrics

import (
	"errors"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Counter metrics.
	metricReadOps         = namespace + "read_ops_total"
	metricWriteOps        = namespace + "write_ops_total"
	metricReadBytes       = namespace + "read_bytes_total"
	metricWriteBytes      = namespace + "write_bytes_total"
	metricReadOpsSeconds  = namespace + "read_seconds_total"
	metricWriteOpsSeconds = namespace + "write_seconds_total"
	metricExceededIOPS    = namespace + "exceeded_iops_seconds_total"
	metricExceededTP      = namespace + "exceeded_tp_seconds_total"
	metricEC2ExceededIOPS = namespace + "ec2_exceeded_iops_seconds_total"
	metricEC2ExceededTP   = namespace + "ec2_exceeded_tp_seconds_total"
	nvmeCollectorScrapes  = namespace + "nvme_collector_scrapes_total"
	nvmeCollectorErrors   = namespace + "nvme_collector_errors_total"

	// Gauge metrics.
	metricVolumeQueueLength = namespace + "volume_queue_length"

	// Histogram metrics.
	metricReadLatency     = namespace + "read_io_latency_seconds"
	metricWriteLatency    = namespace + "write_io_latency_seconds"
	nvmeCollectorDuration = namespace + "nvme_collector_duration_seconds"

	// Conversion factor.
	microsecondsInSeconds = 1e6
)

// EBSMetrics represents the parsed metrics from the NVMe log page.
type EBSMetrics struct {
	EBSMagic              uint64
	ReadOps               uint64
	WriteOps              uint64
	ReadBytes             uint64
	WriteBytes            uint64
	TotalReadTime         uint64
	TotalWriteTime        uint64
	EBSIOPSExceeded       uint64
	EBSThroughputExceeded uint64
	EC2IOPSExceeded       uint64
	EC2ThroughputExceeded uint64
	QueueLength           uint64
	ReservedArea          [416]byte
	ReadLatency           Histogram
	WriteLatency          Histogram
}

type Histogram struct {
	BinCount uint64
	Bins     [64]HistogramBin
}

type HistogramBin struct {
	Lower uint64
	Upper uint64
	Count uint64
}

// As defined in <linux/nvme_ioctl.h>.
type nvmePassthruCommand struct {
	opcode      uint8
	flags       uint8
	rsvd1       uint16
	nsid        uint32
	cdw2        uint32
	cdw3        uint32
	metadata    uint64
	addr        uint64
	metadataLen uint32
	dataLen     uint32
	cdw10       uint32
	cdw11       uint32
	cdw12       uint32
	cdw13       uint32
	cdw14       uint32
	cdw15       uint32
	timeoutMs   uint32
	result      uint32
}

type NVMECollector struct {
	metrics            map[string]*prometheus.Desc
	csiMountPointPath  string
	instanceID         string
	collectionDuration prometheus.Histogram
	scrapesTotal       prometheus.Counter
	scrapeErrorsTotal  prometheus.Counter
}

var (
	ErrInvalidEBSMagic = errors.New("invalid EBS magic number")
	ErrParseLogPage    = errors.New("failed to parse log page")
)

// NewNVMECollector creates a new instance of NVMECollector.
func NewNVMECollector(path, instanceID string) *NVMECollector {
	_ = "STUB: not implemented"
	return nil
}

// Clean CSI mount point path to normalize path
// Add trailing slash back that Clean prunes

func registerNVMECollector(r *MetricRecorder, csiMountPointPath, instanceID string) {
	_ = "STUB: not implemented"
	return
}

// Describe sends the descriptor of each metric in the NVMECollector to Prometheus.
func (c *NVMECollector) Describe(ch chan<- *prometheus.Desc) { _ = "STUB: not implemented"; return }

// Collect is invoked by Prometheus at collection time.
func (c *NVMECollector) Collect(ch chan<- prometheus.Metric) { _ = "STUB: not implemented"; return }

// Send all collected metrics to Prometheus

// Read Latency Histogram

// Write Latency Histogram

// convertHistogram converts the Histogram structure to a format suitable for Prometheus histogram metrics.
func convertHistogram(hist Histogram) (uint64, map[float64]uint64) {
	_ = "STUB: not implemented"
	return 0, nil
}

// getNVMEMetrics retrieves NVMe metrics by reading the log page from the NVMe device at the given path.
func getNVMEMetrics(devicePath string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Allocate before opening the device so we hold it open for as little time as possible

// Write handle is not needed to call ioctl on linux, thus open RDONLY

// parseLogPage parses the binary data from an EBS log page into EBSMetrics.
func parseLogPage(data []byte) (EBSMetrics, error) {
	_ = "STUB: not implemented"
	return *new(EBSMetrics), nil
}

// getCSIManagedDevices returns a slice of unique device paths for NVMe devices mounted under the given path.
func getCSIManagedDevices(path string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read /proc/self/mountinfo to identify NVMe devices

// https://man7.org/linux/man-pages/man5/proc.5.html

// Skip lines with insufficient fields

// Check mount source (field 3) for directly mounted NVMe devices

// Check root (field 9) for block devices

type BlockDevice struct {
	Name   string `json:"name"`
	Serial string `json:"serial"`
}

type LsblkOutput struct {
	BlockDevices []BlockDevice `json:"blockdevices"`
}

// mapDevicePathsToVolumeIDs takes a list of device paths and lsblk output, and returns a map of device paths to volume IDs.
func mapDevicePathsToVolumeIDs(devicePaths []string, lsblkOutput []byte) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func executeLsblk() ([]byte, error) {
	_ = "STUB: not implemented"
	// TODO: Pass context down from Prometheus handler
	return nil, nil
}

func fetchDevicePathToVolumeIDMapping(devicePaths []string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
