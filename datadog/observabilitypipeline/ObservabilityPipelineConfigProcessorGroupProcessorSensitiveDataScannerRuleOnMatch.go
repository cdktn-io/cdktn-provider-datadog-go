// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorSensitiveDataScannerRuleOnMatch struct {
	// hash block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#hash ObservabilityPipeline#hash}
	Hash interface{} `field:"optional" json:"hash" yaml:"hash"`
	// partial_redact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#partial_redact ObservabilityPipeline#partial_redact}
	PartialRedact interface{} `field:"optional" json:"partialRedact" yaml:"partialRedact"`
	// redact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#redact ObservabilityPipeline#redact}
	Redact interface{} `field:"optional" json:"redact" yaml:"redact"`
}

