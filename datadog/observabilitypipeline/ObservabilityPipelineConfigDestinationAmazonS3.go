// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationAmazonS3 struct {
	// S3 bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#bucket ObservabilityPipeline#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Prefix for object keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#key_prefix ObservabilityPipeline#key_prefix}
	KeyPrefix *string `field:"required" json:"keyPrefix" yaml:"keyPrefix"`
	// AWS region of the S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#region ObservabilityPipeline#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// S3 storage class. Valid values are `STANDARD`, `REDUCED_REDUNDANCY`, `INTELLIGENT_TIERING`, `STANDARD_IA`, `EXPRESS_ONEZONE`, `ONEZONE_IA`, `GLACIER`, `GLACIER_IR`, `DEEP_ARCHIVE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#storage_class ObservabilityPipeline#storage_class}
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
	// auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#auth ObservabilityPipeline#auth}
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// buffer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#buffer ObservabilityPipeline#buffer}
	Buffer interface{} `field:"optional" json:"buffer" yaml:"buffer"`
}

