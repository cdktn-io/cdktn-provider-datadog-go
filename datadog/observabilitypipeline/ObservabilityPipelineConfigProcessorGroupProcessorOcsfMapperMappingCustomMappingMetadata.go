// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMetadata struct {
	// The OCSF event class name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#class ObservabilityPipeline#class}
	Class *string `field:"required" json:"class" yaml:"class"`
	// The OCSF schema version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#version ObservabilityPipeline#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// A list of OCSF profiles to apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#profiles ObservabilityPipeline#profiles}
	Profiles *[]*string `field:"optional" json:"profiles" yaml:"profiles"`
}

