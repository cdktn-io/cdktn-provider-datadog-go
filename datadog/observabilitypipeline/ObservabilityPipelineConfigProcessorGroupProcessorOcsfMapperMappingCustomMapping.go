// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMapping struct {
	// The version of the custom mapping configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#version ObservabilityPipeline#version}
	Version *float64 `field:"required" json:"version" yaml:"version"`
	// mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#mapping ObservabilityPipeline#mapping}
	Mapping interface{} `field:"optional" json:"mapping" yaml:"mapping"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#metadata ObservabilityPipeline#metadata}
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
}

