// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorEnrichmentTableFileKey struct {
	// The `items` `column`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#column ObservabilityPipeline#column}
	Column *string `field:"optional" json:"column" yaml:"column"`
	// The comparison method (e.g. equals).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#comparison ObservabilityPipeline#comparison}
	Comparison *string `field:"optional" json:"comparison" yaml:"comparison"`
	// The `items` `field`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#field ObservabilityPipeline#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
}

