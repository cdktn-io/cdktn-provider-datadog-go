// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorParseGrokIncludeRule struct {
	// A Datadog search query used to determine which logs this Grok rule targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#include ObservabilityPipeline#include}
	Include *string `field:"required" json:"include" yaml:"include"`
	// match_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#match_rule ObservabilityPipeline#match_rule}
	MatchRule interface{} `field:"optional" json:"matchRule" yaml:"matchRule"`
	// support_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#support_rule ObservabilityPipeline#support_rule}
	SupportRule interface{} `field:"optional" json:"supportRule" yaml:"supportRule"`
}

