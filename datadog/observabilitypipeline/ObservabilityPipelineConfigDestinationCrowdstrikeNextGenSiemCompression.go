// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationCrowdstrikeNextGenSiemCompression struct {
	// Compression algorithm for log events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#algorithm ObservabilityPipeline#algorithm}
	Algorithm *string `field:"required" json:"algorithm" yaml:"algorithm"`
	// Compression level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#level ObservabilityPipeline#level}
	Level *float64 `field:"optional" json:"level" yaml:"level"`
}

