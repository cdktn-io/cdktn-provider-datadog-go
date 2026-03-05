// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorSample struct {
	// The percentage of logs to sample.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#percentage ObservabilityPipeline#percentage}
	Percentage *float64 `field:"required" json:"percentage" yaml:"percentage"`
	// Optional list of fields to group events by. Each group is sampled independently.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#group_by ObservabilityPipeline#group_by}
	GroupBy *[]*string `field:"optional" json:"groupBy" yaml:"groupBy"`
}

