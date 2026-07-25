// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitPerTagLimit struct {
	// How the per-tag override is applied. One of `limit_override`, `excluded`. Valid values are `limit_override`, `excluded`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#override_type ObservabilityPipeline#override_type}
	OverrideType *string `field:"required" json:"overrideType" yaml:"overrideType"`
	// The tag key this override applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#tag_key ObservabilityPipeline#tag_key}
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// The cardinality cap for this tag.
	//
	// Required when `override_type` is `limit_override`; must be omitted when `override_type` is `excluded`. Value must be between 0 and 1000000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#value_limit ObservabilityPipeline#value_limit}
	ValueLimit *float64 `field:"optional" json:"valueLimit" yaml:"valueLimit"`
}

