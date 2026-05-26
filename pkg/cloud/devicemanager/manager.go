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

package devicemanager

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type Device struct {
	Instance          *types.Instance
	Path              string
	VolumeID          string
	IsAlreadyAssigned bool
	CardIndex         *int32

	isTainted   bool
	releaseFunc func() error
}

func (d *Device) Release(force bool) { _ = "STUB: not implemented"; return }

// Taint marks the device as no longer reusable.
func (d *Device) Taint() { _ = "STUB: not implemented"; return }

type DeviceManager interface {
	// NewDevice retrieves the device if the device is already assigned.
	// Otherwise it creates a new device with next available device name
	// and mark it as unassigned device.
	// numCards is the number of EBS cards on the instance (1 for single-card instances).
	NewDevice(instance *types.Instance, volumeID string, likelyBadNames *sync.Map, numCards int) (device *Device, err error)

	// GetDevice returns the device already assigned to the volume.
	GetDevice(instance *types.Instance, volumeID string) (device *Device, err error)
}

type deviceManager struct {
	// nameAllocator assigns new device name
	nameAllocator NameAllocator

	// We keep an active list of devices we have assigned but not yet
	// attached, to avoid a race condition where we assign a device mapping
	// and then get a second request before we attach the volume.
	mux      sync.Mutex
	inFlight inFlightAttaching
}

var _ DeviceManager = &deviceManager{}

// inFlightEntry represents a volume attachment in progress.
type inFlightEntry struct {
	DeviceName string
	CardIndex  *int32
}

// inFlightAttaching represents the volumes being currently attached to nodes.
// A valid pseudo-representation of it would be {"nodeID": {"volumeID": {deviceName, cardIndex}}}.
type inFlightAttaching map[string]map[string]inFlightEntry

func (i inFlightAttaching) Add(nodeID, volumeID, deviceName string, cardIndex *int32) {
	_ = "STUB: not implemented"
	return
}

func (i inFlightAttaching) Del(nodeID, volumeID string) { _ = "STUB: not implemented"; return }

func (i inFlightAttaching) GetNames(nodeID string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (i inFlightAttaching) GetEntry(nodeID, volumeID string) (inFlightEntry, bool) {
	_ = "STUB: not implemented"
	return *new(inFlightEntry), false
}

func (i inFlightAttaching) GetEntries(nodeID string) map[string]inFlightEntry {
	_ = "STUB: not implemented"
	return nil
}

func NewDeviceManager() DeviceManager { _ = "STUB: not implemented"; return *new(DeviceManager) }

func (d *deviceManager) NewDevice(instance *types.Instance, volumeID string, likelyBadNames *sync.Map, numCards int) (*Device, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get device names being attached and already attached to this instance

// Check if this volume is already assigned a device on this machine

// Calculate card index for new volume

// Add the chosen device and volume to the "attachments in progress" map

// getCardCounts returns a map of card index to volume count, accounting for both
// volumes on the instance device mapping and volumes in the inflight map.
// It ensures volumes are not double counted if they appear in both.
func (d *deviceManager) getCardCounts(instance *types.Instance) map[int32]int {
	_ = "STUB: not implemented"
	return nil
}

// Track volume IDs we've already counted to avoid double counting

// Count volumes per card from existing block device mappings

// Count volumes from inflight map, avoiding double counting

// getNextCardIndex determines the card index to use for a new volume attachment.
// It implements a "least occupied" load balancing strategy.
// Returns nil when the instance has only 1 card. For instances with multiple cards,
// returns the card with the fewest volumes. When counts are equal, prefers lower card index.
func getNextCardIndex(numCards int, cardCounts map[int32]int) *int32 {
	_ = "STUB: not implemented"
	// If instance has only 1 card, return nil (no index needed)
	return nil
}

// Initialize card counts for all cards (ensure all cards are represented)

// Find the card with the fewest devices (prefer lower index on tie)

// getCardIndexForExistingVolume finds the card index for an already attached volume.
func (d *deviceManager) getCardIndexForExistingVolume(instance *types.Instance, volumeID string) *int32 {
	_ = "STUB: not implemented"
	return nil
}

func (d *deviceManager) GetDevice(instance *types.Instance, volumeID string) (*Device, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *deviceManager) newBlockDevice(instance *types.Instance, volumeID string, path string, isAlreadyAssigned bool, cardIndex *int32) *Device {
	_ = "STUB: not implemented"
	return nil
}

func (d *deviceManager) release(device *Device) error { _ = "STUB: not implemented"; return nil }

// Attaching is not in progress, so there's nothing to release

// This actually can happen, because GetNext combines the inFlightAttaching map with the volumes
// attached to the instance (as reported by the EC2 API).  So if release comes after
// a 10 second poll delay, we might as well have had a concurrent request to allocate a mountpoint,
// which because we allocate sequentially is very likely to get the immediately freed volume.

// getDeviceNamesInUse returns the device to volume ID mapping
// the mapping includes both already attached and being attached volumes.
func (d *deviceManager) getDeviceNamesInUse(instance *types.Instance) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (d *deviceManager) getPath(inUse map[string]string, volumeID string) string {
	_ = "STUB: not implemented"
	return ""
}

func getInstanceID(instance *types.Instance) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
