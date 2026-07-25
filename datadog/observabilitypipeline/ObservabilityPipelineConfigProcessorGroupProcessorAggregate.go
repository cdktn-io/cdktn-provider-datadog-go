// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorAggregate struct {
	// The interval, in seconds, over which metrics are aggregated.
	//
	// Must be between 1 and 60. Value must be between 1 and 60.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#interval_secs ObservabilityPipeline#interval_secs}
	IntervalSecs *float64 `field:"required" json:"intervalSecs" yaml:"intervalSecs"`
	// The aggregation mode.
	//
	// One of `auto`, `sum`, `latest`, `count`, `max`, `min`, `mean`. Valid values are `auto`, `sum`, `latest`, `count`, `max`, `min`, `mean`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#mode ObservabilityPipeline#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

