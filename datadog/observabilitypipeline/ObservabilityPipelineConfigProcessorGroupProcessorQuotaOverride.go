// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorQuotaOverride struct {
	// field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#field ObservabilityPipeline#field}
	Field interface{} `field:"optional" json:"field" yaml:"field"`
	// limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#limit ObservabilityPipeline#limit}
	Limit interface{} `field:"optional" json:"limit" yaml:"limit"`
}

