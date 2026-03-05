// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorEnrichmentTable struct {
	// Path where enrichment results should be stored in the log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#target ObservabilityPipeline#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#file ObservabilityPipeline#file}
	File interface{} `field:"optional" json:"file" yaml:"file"`
	// geoip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#geoip ObservabilityPipeline#geoip}
	Geoip interface{} `field:"optional" json:"geoip" yaml:"geoip"`
	// reference_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#reference_table ObservabilityPipeline#reference_table}
	ReferenceTable interface{} `field:"optional" json:"referenceTable" yaml:"referenceTable"`
}

