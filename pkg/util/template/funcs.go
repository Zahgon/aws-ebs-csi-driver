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
	"html/template"
)

// Disable functions.
func html(...any) (string, error) { _ = "STUB: not implemented"; return "", nil }

func js(...any) (string, error) { _ = "STUB: not implemented"; return "", nil }

func call(...any) (string, error) { _ = "STUB: not implemented"; return "", nil }

func urlquery(...any) (string, error) { _ = "STUB: not implemented"; return "", nil }

func contains(arg1, arg2 string) bool { _ = "STUB: not implemented"; return false }

func substring(start, end int, arg string) string { _ = "STUB: not implemented"; return "" }

func field(delim string, idx int, arg string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func index(arg1, arg2 string) int { _ = "STUB: not implemented"; return 0 }

func lastIndex(arg1, arg2 string) int { _ = "STUB: not implemented"; return 0 }

func newFuncMap() template.FuncMap { _ = "STUB: not implemented"; return *new(template.FuncMap) }
