// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorEnrichmentTableFile struct {
	// encoding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#encoding ObservabilityPipeline#encoding}
	Encoding interface{} `field:"optional" json:"encoding" yaml:"encoding"`
	// key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#key ObservabilityPipeline#key}
	Key interface{} `field:"optional" json:"key" yaml:"key"`
	// Path to the CSV file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#path ObservabilityPipeline#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

