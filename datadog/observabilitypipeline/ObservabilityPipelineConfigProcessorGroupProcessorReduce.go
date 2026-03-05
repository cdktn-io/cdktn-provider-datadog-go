// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorReduce struct {
	// A list of fields used to group log events for merging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#group_by ObservabilityPipeline#group_by}
	GroupBy *[]*string `field:"required" json:"groupBy" yaml:"groupBy"`
	// merge_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#merge_strategy ObservabilityPipeline#merge_strategy}
	MergeStrategy interface{} `field:"optional" json:"mergeStrategy" yaml:"mergeStrategy"`
}

