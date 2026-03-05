// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMapping struct {
	// Search query for selecting which logs the mapping applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#include ObservabilityPipeline#include}
	Include *string `field:"required" json:"include" yaml:"include"`
	// custom_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#custom_mapping ObservabilityPipeline#custom_mapping}
	CustomMapping interface{} `field:"optional" json:"customMapping" yaml:"customMapping"`
	// Predefined library mapping for log transformation. Use this or custom_mapping, not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#library_mapping ObservabilityPipeline#library_mapping}
	LibraryMapping *string `field:"optional" json:"libraryMapping" yaml:"libraryMapping"`
}

