// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourceSocketFraming struct {
	// The framing method. Valid values are `newline_delimited`, `bytes`, `character_delimited`, `octet_counting`, `chunked_gelf`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#method ObservabilityPipeline#method}
	Method *string `field:"required" json:"method" yaml:"method"`
	// character_delimited block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#character_delimited ObservabilityPipeline#character_delimited}
	CharacterDelimited interface{} `field:"optional" json:"characterDelimited" yaml:"characterDelimited"`
}

