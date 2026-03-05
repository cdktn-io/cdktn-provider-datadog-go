// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookup struct {
	// The default value to use if no lookup match is found.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#default ObservabilityPipeline#default}
	Default *string `field:"optional" json:"default" yaml:"default"`
	// table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#table ObservabilityPipeline#table}
	Table interface{} `field:"optional" json:"table" yaml:"table"`
}

