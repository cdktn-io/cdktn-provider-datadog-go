// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorSensitiveDataScannerRulePatternCustom struct {
	// Human-readable description providing context about a sensitive data scanner rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#description ObservabilityPipeline#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A regular expression used to detect sensitive values. Must be a valid regex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#rule ObservabilityPipeline#rule}
	Rule *string `field:"optional" json:"rule" yaml:"rule"`
}

