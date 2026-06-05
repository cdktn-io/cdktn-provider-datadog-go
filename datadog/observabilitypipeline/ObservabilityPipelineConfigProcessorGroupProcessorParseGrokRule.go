// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorParseGrokRule struct {
	// The value of the source field in log events which should be processed by the Grok rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/observability_pipeline#source ObservabilityPipeline#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// match_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/observability_pipeline#match_rule ObservabilityPipeline#match_rule}
	MatchRule interface{} `field:"optional" json:"matchRule" yaml:"matchRule"`
	// support_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/observability_pipeline#support_rule ObservabilityPipeline#support_rule}
	SupportRule interface{} `field:"optional" json:"supportRule" yaml:"supportRule"`
}

