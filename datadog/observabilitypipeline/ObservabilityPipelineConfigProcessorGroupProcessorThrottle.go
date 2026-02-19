// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorThrottle struct {
	// The number of events to allow before throttling is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#threshold ObservabilityPipeline#threshold}
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// The time window in seconds over which the threshold applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#window ObservabilityPipeline#window}
	Window *float64 `field:"required" json:"window" yaml:"window"`
	// Optional list of fields used to group events before applying throttling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#group_by ObservabilityPipeline#group_by}
	GroupBy *[]*string `field:"optional" json:"groupBy" yaml:"groupBy"`
}

