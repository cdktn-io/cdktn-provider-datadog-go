// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package referencetable


type ReferenceTableFileMetadataAccessDetailsAwsDetail struct {
	// The ID of the AWS account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.82.0/docs/resources/reference_table#aws_account_id ReferenceTable#aws_account_id}
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// The name of the AWS S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.82.0/docs/resources/reference_table#aws_bucket_name ReferenceTable#aws_bucket_name}
	AwsBucketName *string `field:"optional" json:"awsBucketName" yaml:"awsBucketName"`
	// The relative file path from the AWS S3 bucket root to the CSV file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.82.0/docs/resources/reference_table#file_path ReferenceTable#file_path}
	FilePath *string `field:"optional" json:"filePath" yaml:"filePath"`
}

