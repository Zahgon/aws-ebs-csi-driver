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

package cloud

import (
	"context"

	"github.com/aws/smithy-go/middleware"
)

// RecordRequestsMiddleware is added to the Complete chain; called after any request.
func RecordRequestsMiddleware(deprecatedMetrics bool) func(*middleware.Stack) error {
	_ = "STUB: not implemented"
	return nil
}

// LogServerErrorsMiddleware is a middleware that logs server errors received when attempting to contact the AWS API
// A specialized middleware is used instead of the SDK's built-in retry logging to allow for customizing the verbosity
// of throttle errors vs server/unknown errors, to prevent flooding the logs with throttle error.
func LogServerErrorsMiddleware() func(*middleware.Stack) error {
	_ = "STUB: not implemented"
	return nil
}

// Only log throttle errors under a high verbosity as we expect to see many of them
// under normal bursty/high-TPS workloads

func createLabels(ctx context.Context) map[string]string { _ = "STUB: not implemented"; return nil }
