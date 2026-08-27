// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimit struct {
	// The metric name this override applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/observability_pipeline#metric_name ObservabilityPipeline#metric_name}
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// How the per-metric override is applied. One of `limit_override`, `excluded`. Valid values are `limit_override`, `excluded`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/observability_pipeline#override_type ObservabilityPipeline#override_type}
	OverrideType *string `field:"required" json:"overrideType" yaml:"overrideType"`
	// The action to take on this metric when the limit is exceeded.
	//
	// Required when `override_type` is `limit_override`; must be omitted when `override_type` is `excluded`. Valid values are `drop_tag`, `drop_event`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/observability_pipeline#limit_exceeded_action ObservabilityPipeline#limit_exceeded_action}
	LimitExceededAction *string `field:"optional" json:"limitExceededAction" yaml:"limitExceededAction"`
	// per_tag_limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/observability_pipeline#per_tag_limit ObservabilityPipeline#per_tag_limit}
	PerTagLimit interface{} `field:"optional" json:"perTagLimit" yaml:"perTagLimit"`
	// The cardinality cap for this metric.
	//
	// Required when `override_type` is `limit_override`; must be omitted when `override_type` is `excluded`. Value must be between 0 and 1000000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/observability_pipeline#value_limit ObservabilityPipeline#value_limit}
	ValueLimit *float64 `field:"optional" json:"valueLimit" yaml:"valueLimit"`
}

