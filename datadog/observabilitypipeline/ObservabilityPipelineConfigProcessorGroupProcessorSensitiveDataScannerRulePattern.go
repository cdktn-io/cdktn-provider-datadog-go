// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorSensitiveDataScannerRulePattern struct {
	// custom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/observability_pipeline#custom ObservabilityPipeline#custom}
	Custom interface{} `field:"optional" json:"custom" yaml:"custom"`
	// library block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/observability_pipeline#library ObservabilityPipeline#library}
	Library interface{} `field:"optional" json:"library" yaml:"library"`
}

