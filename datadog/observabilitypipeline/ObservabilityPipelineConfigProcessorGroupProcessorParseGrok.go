// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorParseGrok struct {
	// If set to `true`, disables the default Grok rules provided by Datadog. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/observability_pipeline#disable_library_rules ObservabilityPipeline#disable_library_rules}
	DisableLibraryRules interface{} `field:"optional" json:"disableLibraryRules" yaml:"disableLibraryRules"`
	// The log field to parse with the Grok rules. Defaults to `"message"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/observability_pipeline#field ObservabilityPipeline#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
	// include_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/observability_pipeline#include_rule ObservabilityPipeline#include_rule}
	IncludeRule interface{} `field:"optional" json:"includeRule" yaml:"includeRule"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/observability_pipeline#rule ObservabilityPipeline#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
}

