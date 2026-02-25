// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package referencetable

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ReferenceTableConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The source type for the reference table. Valid values are `S3`, `GCS`, `AZURE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/reference_table#source ReferenceTable#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// The name of the reference table. This must be unique within your organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/reference_table#table_name ReferenceTable#table_name}
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// The description of the reference table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/reference_table#description ReferenceTable#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// file_metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/reference_table#file_metadata ReferenceTable#file_metadata}
	FileMetadata *ReferenceTableFileMetadata `field:"optional" json:"fileMetadata" yaml:"fileMetadata"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/reference_table#schema ReferenceTable#schema}
	Schema *ReferenceTableSchema `field:"optional" json:"schema" yaml:"schema"`
	// A list of tags to associate with the reference table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/reference_table#tags ReferenceTable#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

