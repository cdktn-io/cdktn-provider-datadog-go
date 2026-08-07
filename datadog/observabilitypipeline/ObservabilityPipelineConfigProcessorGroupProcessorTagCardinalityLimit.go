// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimit struct {
	// The default action to take when the cardinality limit is exceeded.
	//
	// One of `drop_tag`, `drop_event`. Valid values are `drop_tag`, `drop_event`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/observability_pipeline#limit_exceeded_action ObservabilityPipeline#limit_exceeded_action}
	LimitExceededAction *string `field:"required" json:"limitExceededAction" yaml:"limitExceededAction"`
	// The default maximum number of distinct tag value combinations allowed per metric.
	//
	// Between 0 and 1000000. Value must be between 0 and 1000000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/observability_pipeline#value_limit ObservabilityPipeline#value_limit}
	ValueLimit *float64 `field:"required" json:"valueLimit" yaml:"valueLimit"`
	// per_metric_limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/observability_pipeline#per_metric_limit ObservabilityPipeline#per_metric_limit}
	PerMetricLimit interface{} `field:"optional" json:"perMetricLimit" yaml:"perMetricLimit"`
	// tracking_mode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/observability_pipeline#tracking_mode ObservabilityPipeline#tracking_mode}
	TrackingMode interface{} `field:"optional" json:"trackingMode" yaml:"trackingMode"`
}

