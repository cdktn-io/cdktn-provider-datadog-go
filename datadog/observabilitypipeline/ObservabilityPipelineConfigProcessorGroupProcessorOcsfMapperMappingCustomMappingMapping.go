// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMapping struct {
	// The destination OCSF field path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#dest ObservabilityPipeline#dest}
	Dest *string `field:"required" json:"dest" yaml:"dest"`
	// The default value to use if the source field is missing or empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#default ObservabilityPipeline#default}
	Default *string `field:"optional" json:"default" yaml:"default"`
	// lookup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#lookup ObservabilityPipeline#lookup}
	Lookup interface{} `field:"optional" json:"lookup" yaml:"lookup"`
	// The source field path from the log event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#source ObservabilityPipeline#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Multiple source field paths for combined mapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#sources ObservabilityPipeline#sources}
	Sources *[]*string `field:"optional" json:"sources" yaml:"sources"`
	// A static value to use for the destination field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#value ObservabilityPipeline#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

