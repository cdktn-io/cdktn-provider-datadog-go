// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationElasticsearchCompression struct {
	// The compression algorithm applied when sending data to Elasticsearch. Valid values are `none`, `gzip`, `zlib`, `zstd`, `snappy`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/observability_pipeline#algorithm ObservabilityPipeline#algorithm}
	Algorithm *string `field:"required" json:"algorithm" yaml:"algorithm"`
	// The compression level. Only applicable for `gzip`, `zlib`, and `zstd` algorithms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/observability_pipeline#level ObservabilityPipeline#level}
	Level *float64 `field:"optional" json:"level" yaml:"level"`
}

