/*
Copyright 2023 The Kubernetes Authors.
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
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
)

func InitOtelTracing() (*otlptrace.Exporter, error) {
	_ = "STUB: not implemented"
	// Setup OTLP exporter
	return nil, nil
}

// Resource will auto populate spans with common attributes

// pull attributes from OTEL_RESOURCE_ATTRIBUTES and OTEL_SERVICE_NAME environment variables

// Create a trace provider with the exporter.
// Use propagator and sampler defined in environment variables.

// Register the trace provider as global.
