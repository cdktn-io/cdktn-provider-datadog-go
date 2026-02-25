// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudinventorysyncconfig


type CloudInventorySyncConfigAws struct {
	// AWS Account ID of the account holding the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/cloud_inventory_sync_config#aws_account_id CloudInventorySyncConfig#aws_account_id}
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Name of the S3 bucket holding the inventory files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/cloud_inventory_sync_config#destination_bucket_name CloudInventorySyncConfig#destination_bucket_name}
	DestinationBucketName *string `field:"optional" json:"destinationBucketName" yaml:"destinationBucketName"`
	// AWS Region of the bucket holding the inventory files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/cloud_inventory_sync_config#destination_bucket_region CloudInventorySyncConfig#destination_bucket_region}
	DestinationBucketRegion *string `field:"optional" json:"destinationBucketRegion" yaml:"destinationBucketRegion"`
	// Prefix path within the bucket for inventory files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/cloud_inventory_sync_config#destination_prefix CloudInventorySyncConfig#destination_prefix}
	DestinationPrefix *string `field:"optional" json:"destinationPrefix" yaml:"destinationPrefix"`
}

