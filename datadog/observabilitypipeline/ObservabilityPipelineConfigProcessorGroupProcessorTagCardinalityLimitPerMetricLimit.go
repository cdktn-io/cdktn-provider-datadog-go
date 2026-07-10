// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimit struct {
	// The metric name this override applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#metric_name ObservabilityPipeline#metric_name}
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// How the per-metric override is applied. One of `tracked`, `excluded`. Valid values are `tracked`, `excluded`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#mode ObservabilityPipeline#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The action to take on this metric when the limit is exceeded.
	//
	// Required when `mode` is `tracked`; must be omitted when `mode` is `excluded`. Valid values are `drop_tag`, `drop_event`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#limit_exceeded_action ObservabilityPipeline#limit_exceeded_action}
	LimitExceededAction *string `field:"optional" json:"limitExceededAction" yaml:"limitExceededAction"`
	// per_tag_limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#per_tag_limit ObservabilityPipeline#per_tag_limit}
	PerTagLimit interface{} `field:"optional" json:"perTagLimit" yaml:"perTagLimit"`
	// The cardinality cap for this metric.
	//
	// Required when `mode` is `tracked`; must be omitted when `mode` is `excluded`. Value must be between 0 and 1000000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#value_limit ObservabilityPipeline#value_limit}
	ValueLimit *float64 `field:"optional" json:"valueLimit" yaml:"valueLimit"`
}

