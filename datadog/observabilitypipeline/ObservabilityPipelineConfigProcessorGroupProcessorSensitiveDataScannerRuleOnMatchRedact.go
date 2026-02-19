// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorSensitiveDataScannerRuleOnMatchRedact struct {
	// Replacement string for redacted values (e.g., `***`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#replace ObservabilityPipeline#replace}
	Replace *string `field:"optional" json:"replace" yaml:"replace"`
}

