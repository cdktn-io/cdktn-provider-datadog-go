// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorSensitiveDataScannerRule struct {
	// A name identifying the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#name ObservabilityPipeline#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Tags assigned to this rule for filtering and classification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#tags ObservabilityPipeline#tags}
	Tags *[]*string `field:"required" json:"tags" yaml:"tags"`
	// keyword_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#keyword_options ObservabilityPipeline#keyword_options}
	KeywordOptions interface{} `field:"optional" json:"keywordOptions" yaml:"keywordOptions"`
	// on_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#on_match ObservabilityPipeline#on_match}
	OnMatch interface{} `field:"optional" json:"onMatch" yaml:"onMatch"`
	// pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#pattern ObservabilityPipeline#pattern}
	Pattern interface{} `field:"optional" json:"pattern" yaml:"pattern"`
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#scope ObservabilityPipeline#scope}
	Scope interface{} `field:"optional" json:"scope" yaml:"scope"`
}

