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

package util

import (
	"context"
	"regexp"
	"time"

	csi "github.com/container-storage-interface/spec/lib/go/csi"
)

const (
	GiB              = int64(1024 * 1024 * 1024)
	DefaultBlockSize = 4096

	// AttachmentShared volume attachment type constant.
	AttachmentShared    = "shared"
	AttachmentDedicated = "dedicated"

	VolumeIDRegex   = "vol-[a-z0-9]+"
	InstanceIDRegex = "i-[a-z0-9]+"
	SnapshotIDRegex = "snap-[a-z0-9]+"
)

var (
	isAlphanumericRegex = regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString
	// MAC Address Regex Source: https://stackoverflow.com/a/4260512
	isMACAddressRegex = regexp.MustCompile(`([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})`)

	// DriverName is the domain for all EBS CSI Driver related components.
	// Variable instead of constant to allow initialization from plugin.
	driverName = ""
)

func SetDriverName(name string) { _ = "STUB: not implemented"; return }

func GetDriverName() string { _ = "STUB: not implemented"; return "" }

// Return a hardcoded value in unit tests, as main.go won't be called to setup the name

// SetDriverName hasn't been initialized in main.go yet - this will result in a very
// difficult to debug bug, so immediately exit if code is added that calls this too early

// RoundUpBytes rounds up the volume size in bytes up to multiplications of GiB.
func RoundUpBytes(volumeSizeBytes int64) int64 { _ = "STUB: not implemented"; return 0 }

// RoundUpGiB rounds up the volume size in bytes upto multiplications of GiB
// in the unit of GiB.
func RoundUpGiB(volumeSizeBytes int64) (int32, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:gosec // Integer overflow handled

// BytesToGiB converts Bytes to GiB.
func BytesToGiB(volumeSizeBytes int64) int32 { _ = "STUB: not implemented"; return 0 }

// Handle overflow

//nolint:gosec // Integer overflow handled

// GiBToBytes converts GiB to Bytes.
func GiBToBytes(volumeSizeGiB int32) int64 { _ = "STUB: not implemented"; return 0 }

func ParseEndpoint(endpoint string, hostprocess bool) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// Remove the socket file if it already exists

// #nosec G703 -- addr is derived from a parsed URL path, not direct user input

func roundUpSize(volumeSizeBytes int64, allocationUnitBytes int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

// Avoid division by zero

// GetAccessModes returns a slice containing all of the access modes defined
// in the passed in VolumeCapabilities.
func GetAccessModes(caps []*csi.VolumeCapability) *[]string { _ = "STUB: not implemented"; return nil }

// StringIsAlphanumeric returns true if a given string contains only English letters or numbers.
func StringIsAlphanumeric(s string) bool { _ = "STUB: not implemented"; return false }

// CountMACAddresses returns the amount of MAC addresses within a string.
func CountMACAddresses(s string) int { _ = "STUB: not implemented"; return 0 }

// NormalizeWindowsPath normalizes a Windows path.
func NormalizeWindowsPath(path string) string { _ = "STUB: not implemented"; return "" }

// SanitizeRequest takes a request object and returns a copy of the request with
// the "Secrets" field cleared.
func SanitizeRequest(req any) any { _ = "STUB: not implemented"; return *new(any) }

// WaitUntilTimeOrContext returns once time wakeup has elapsed or ctx is done.
func WaitUntilTimeOrContext(ctx context.Context, wakeup time.Time) {
	_ = "STUB: not implemented"
	return
}

func IsHyperPodNode(nodeID string) bool { _ = "STUB: not implemented"; return false }
