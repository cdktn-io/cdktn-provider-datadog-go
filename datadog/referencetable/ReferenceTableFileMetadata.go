// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package referencetable


type ReferenceTableFileMetadata struct {
	// Whether this table should automatically sync with the cloud storage source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.82.0/docs/resources/reference_table#sync_enabled ReferenceTable#sync_enabled}
	SyncEnabled interface{} `field:"required" json:"syncEnabled" yaml:"syncEnabled"`
	// access_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.82.0/docs/resources/reference_table#access_details ReferenceTable#access_details}
	AccessDetails *ReferenceTableFileMetadataAccessDetails `field:"optional" json:"accessDetails" yaml:"accessDetails"`
}

