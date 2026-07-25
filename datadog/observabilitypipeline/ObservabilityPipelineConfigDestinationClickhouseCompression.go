// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationClickhouseCompression struct {
	// Compression algorithm. Valid values are `gzip` and `none`. Valid values are `gzip`, `none`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#algorithm ObservabilityPipeline#algorithm}
	Algorithm *string `field:"required" json:"algorithm" yaml:"algorithm"`
	// Compression level (1–9). Only valid when `algorithm` is `gzip`. Value must be between 1 and 9.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#level ObservabilityPipeline#level}
	Level *float64 `field:"optional" json:"level" yaml:"level"`
}

