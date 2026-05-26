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

package template

import (
	"text/template"
)

type PVProps struct {
	PVCName      string
	PVCNamespace string
	PVName       string
}

type VolumeSnapshotProps struct {
	VolumeSnapshotName        string
	VolumeSnapshotNamespace   string
	VolumeSnapshotContentName string
}

func Evaluate(tm []string, props any, warnOnly bool) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func execTemplate(value string, props any, t *template.Template) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
