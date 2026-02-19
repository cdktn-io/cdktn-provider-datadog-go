// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationAzureStorage struct {
	// The name of the Azure Blob Storage container to store logs in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#container_name ObservabilityPipeline#container_name}
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Optional prefix for blobs written to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#blob_prefix ObservabilityPipeline#blob_prefix}
	BlobPrefix *string `field:"optional" json:"blobPrefix" yaml:"blobPrefix"`
}

