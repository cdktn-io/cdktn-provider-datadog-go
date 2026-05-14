// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapper struct {
	// Whether to keep an event that does not match any of the mapping filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/observability_pipeline#keep_unmatched ObservabilityPipeline#keep_unmatched}
	KeepUnmatched interface{} `field:"optional" json:"keepUnmatched" yaml:"keepUnmatched"`
	// mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/observability_pipeline#mapping ObservabilityPipeline#mapping}
	Mapping interface{} `field:"optional" json:"mapping" yaml:"mapping"`
}

