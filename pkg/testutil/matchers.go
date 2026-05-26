// Copyright 2026 The Kubernetes Authors.
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

package testutil

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

type contextMatcher struct{}

func (m contextMatcher) Matches(x any) bool { _ = "STUB: not implemented"; return false }

func (m contextMatcher) String() string { _ = "STUB: not implemented"; return "" }

func AnyContext() gomock.Matcher { _ = "STUB: not implemented"; return *new(gomock.Matcher) }

type typeMatcher struct {
	t reflect.Type
}

func (m typeMatcher) Matches(x any) bool { _ = "STUB: not implemented"; return false }

func (m typeMatcher) String() string { _ = "STUB: not implemented"; return "" }

func OfType(example any) gomock.Matcher { _ = "STUB: not implemented"; return *new(gomock.Matcher) }

type ec2OptionsMatcher struct{}

func (m ec2OptionsMatcher) Matches(x any) bool {
	_ = "STUB: not implemented"
	// Check if it's a single function
	return false
}

// Check if it's a slice of functions

func (m ec2OptionsMatcher) String() string { _ = "STUB: not implemented"; return "" }

func EC2Options() gomock.Matcher { _ = "STUB: not implemented"; return *new(gomock.Matcher) }

type ec2InputMatcher struct {
	expectedType reflect.Type
}

func (m ec2InputMatcher) Matches(x any) bool { _ = "STUB: not implemented"; return false }

func (m ec2InputMatcher) String() string { _ = "STUB: not implemented"; return "" }

func EC2Input(example any) gomock.Matcher { _ = "STUB: not implemented"; return *new(gomock.Matcher) }
