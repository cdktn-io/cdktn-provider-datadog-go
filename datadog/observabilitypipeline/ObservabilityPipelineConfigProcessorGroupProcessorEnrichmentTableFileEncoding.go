// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorEnrichmentTableFileEncoding struct {
	// The `encoding` `delimiter`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#delimiter ObservabilityPipeline#delimiter}
	Delimiter *string `field:"required" json:"delimiter" yaml:"delimiter"`
	// File encoding format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#type ObservabilityPipeline#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The `encoding` `includes_headers`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#includes_headers ObservabilityPipeline#includes_headers}
	IncludesHeaders interface{} `field:"optional" json:"includesHeaders" yaml:"includesHeaders"`
}

