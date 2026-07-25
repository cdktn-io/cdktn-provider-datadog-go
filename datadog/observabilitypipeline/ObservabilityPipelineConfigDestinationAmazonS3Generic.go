// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationAmazonS3Generic struct {
	// S3 bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#bucket ObservabilityPipeline#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// AWS region of the S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#region ObservabilityPipeline#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// S3 storage class. Valid values are `STANDARD`, `REDUCED_REDUNDANCY`, `INTELLIGENT_TIERING`, `STANDARD_IA`, `EXPRESS_ONEZONE`, `ONEZONE_IA`, `GLACIER`, `GLACIER_IR`, `DEEP_ARCHIVE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#storage_class ObservabilityPipeline#storage_class}
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
	// auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#auth ObservabilityPipeline#auth}
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// batch_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#batch_settings ObservabilityPipeline#batch_settings}
	BatchSettings interface{} `field:"optional" json:"batchSettings" yaml:"batchSettings"`
	// buffer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#buffer ObservabilityPipeline#buffer}
	Buffer interface{} `field:"optional" json:"buffer" yaml:"buffer"`
	// compression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#compression ObservabilityPipeline#compression}
	Compression interface{} `field:"optional" json:"compression" yaml:"compression"`
	// encoding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#encoding ObservabilityPipeline#encoding}
	Encoding interface{} `field:"optional" json:"encoding" yaml:"encoding"`
	// Optional prefix for object keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#key_prefix ObservabilityPipeline#key_prefix}
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
	// The server-side encryption algorithm used when storing objects in S3. Valid values: `aws:kms`, `AES256`. Valid values are `aws:kms`, `AES256`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#server_side_encryption ObservabilityPipeline#server_side_encryption}
	ServerSideEncryption *string `field:"optional" json:"serverSideEncryption" yaml:"serverSideEncryption"`
	// ID of the AWS KMS key to use for SSE-KMS encryption. Only applies when `server_side_encryption` is `aws:kms`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#ssekms_key_id ObservabilityPipeline#ssekms_key_id}
	SsekmsKeyId *string `field:"optional" json:"ssekmsKeyId" yaml:"ssekmsKeyId"`
}

