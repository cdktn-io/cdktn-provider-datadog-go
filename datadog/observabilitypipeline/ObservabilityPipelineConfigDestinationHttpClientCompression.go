// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationHttpClientCompression struct {
	// Compression algorithm. Valid values are `gzip`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#algorithm ObservabilityPipeline#algorithm}
	Algorithm *string `field:"required" json:"algorithm" yaml:"algorithm"`
}

