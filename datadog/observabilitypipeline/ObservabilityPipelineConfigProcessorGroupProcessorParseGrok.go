// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorParseGrok struct {
	// If set to `true`, disables the default Grok rules provided by Datadog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#disable_library_rules ObservabilityPipeline#disable_library_rules}
	DisableLibraryRules interface{} `field:"optional" json:"disableLibraryRules" yaml:"disableLibraryRules"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#rule ObservabilityPipeline#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
}

