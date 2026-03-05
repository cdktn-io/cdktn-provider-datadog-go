// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTable struct {
	// The substring to match in the source value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#contains ObservabilityPipeline#contains}
	Contains *string `field:"optional" json:"contains" yaml:"contains"`
	// The source field to match against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#equals_source ObservabilityPipeline#equals_source}
	EqualsSource *string `field:"optional" json:"equalsSource" yaml:"equalsSource"`
	// The exact value to match in the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#equals ObservabilityPipeline#equals}
	EqualTo *string `field:"optional" json:"equalTo" yaml:"equalTo"`
	// A regex pattern to match in the source value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#matches ObservabilityPipeline#matches}
	Matches *string `field:"optional" json:"matches" yaml:"matches"`
	// A regex pattern that must not match the source value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#not_matches ObservabilityPipeline#not_matches}
	NotMatches *string `field:"optional" json:"notMatches" yaml:"notMatches"`
	// The value to use when a match is found.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#value ObservabilityPipeline#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

