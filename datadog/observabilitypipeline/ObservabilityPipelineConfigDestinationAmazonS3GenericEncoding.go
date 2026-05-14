// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationAmazonS3GenericEncoding struct {
	// The encoding type. Valid values are `json`, `parquet`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/observability_pipeline#type ObservabilityPipeline#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

