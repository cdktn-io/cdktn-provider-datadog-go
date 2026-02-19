// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudinventorysyncconfig


type CloudInventorySyncConfigGcp struct {
	// Name of the GCS bucket holding the inventory files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/cloud_inventory_sync_config#destination_bucket_name CloudInventorySyncConfig#destination_bucket_name}
	DestinationBucketName *string `field:"optional" json:"destinationBucketName" yaml:"destinationBucketName"`
	// GCP Project ID of the project holding the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/cloud_inventory_sync_config#project_id CloudInventorySyncConfig#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Service account email used for reading the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/cloud_inventory_sync_config#service_account_email CloudInventorySyncConfig#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// Name of the source bucket the inventory report is generated for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/cloud_inventory_sync_config#source_bucket_name CloudInventorySyncConfig#source_bucket_name}
	SourceBucketName *string `field:"optional" json:"sourceBucketName" yaml:"sourceBucketName"`
}

