// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationSocketBufferMemory struct {
	// Maximum events for the memory buffer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#max_events ObservabilityPipeline#max_events}
	MaxEvents *float64 `field:"optional" json:"maxEvents" yaml:"maxEvents"`
	// Maximum size of the memory buffer (in bytes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#max_size ObservabilityPipeline#max_size}
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// Behavior when the buffer is full. Valid values are `block` or `drop_newest`. Defaults to `"block"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#when_full ObservabilityPipeline#when_full}
	WhenFull *string `field:"optional" json:"whenFull" yaml:"whenFull"`
}

