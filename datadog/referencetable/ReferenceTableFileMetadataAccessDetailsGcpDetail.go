// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package referencetable


type ReferenceTableFileMetadataAccessDetailsGcpDetail struct {
	// The relative file path from the GCS bucket root to the CSV file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/reference_table#file_path ReferenceTable#file_path}
	FilePath *string `field:"optional" json:"filePath" yaml:"filePath"`
	// The name of the GCP bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/reference_table#gcp_bucket_name ReferenceTable#gcp_bucket_name}
	GcpBucketName *string `field:"optional" json:"gcpBucketName" yaml:"gcpBucketName"`
	// The ID of the GCP project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/reference_table#gcp_project_id ReferenceTable#gcp_project_id}
	GcpProjectId *string `field:"optional" json:"gcpProjectId" yaml:"gcpProjectId"`
	// The email of the GCP service account used to access the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/reference_table#gcp_service_account_email ReferenceTable#gcp_service_account_email}
	GcpServiceAccountEmail *string `field:"optional" json:"gcpServiceAccountEmail" yaml:"gcpServiceAccountEmail"`
}

