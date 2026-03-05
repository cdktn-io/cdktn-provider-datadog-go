// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationGoogleCloudStorage struct {
	// Name of the GCS bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#bucket ObservabilityPipeline#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Storage class used for objects stored in GCS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#storage_class ObservabilityPipeline#storage_class}
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
	// Access control list setting for objects written to the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#acl ObservabilityPipeline#acl}
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#auth ObservabilityPipeline#auth}
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// buffer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#buffer ObservabilityPipeline#buffer}
	Buffer interface{} `field:"optional" json:"buffer" yaml:"buffer"`
	// Optional prefix for object keys within the GCS bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#key_prefix ObservabilityPipeline#key_prefix}
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#metadata ObservabilityPipeline#metadata}
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
}

