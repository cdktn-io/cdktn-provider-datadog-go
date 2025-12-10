// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package referencetable


type ReferenceTableFileMetadataAccessDetails struct {
	// aws_detail block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.82.0/docs/resources/reference_table#aws_detail ReferenceTable#aws_detail}
	AwsDetail *ReferenceTableFileMetadataAccessDetailsAwsDetail `field:"optional" json:"awsDetail" yaml:"awsDetail"`
	// azure_detail block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.82.0/docs/resources/reference_table#azure_detail ReferenceTable#azure_detail}
	AzureDetail *ReferenceTableFileMetadataAccessDetailsAzureDetail `field:"optional" json:"azureDetail" yaml:"azureDetail"`
	// gcp_detail block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.82.0/docs/resources/reference_table#gcp_detail ReferenceTable#gcp_detail}
	GcpDetail *ReferenceTableFileMetadataAccessDetailsGcpDetail `field:"optional" json:"gcpDetail" yaml:"gcpDetail"`
}

