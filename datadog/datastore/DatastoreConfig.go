// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datastore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatastoreConfig struct {
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
	// The display name for the new datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/datastore#name Datastore#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the primary key column for this datastore.
	//
	// Primary column names:   - Must abide by both [PostgreSQL naming conventions](https://www.postgresql.org/docs/7.0/syntax525.htm)   - Cannot exceed 63 characters
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/datastore#primary_column_name Datastore#primary_column_name}
	PrimaryColumnName *string `field:"required" json:"primaryColumnName" yaml:"primaryColumnName"`
	// A human-readable description about the datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/datastore#description Datastore#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The organization access level for the datastore. For example, 'contributor'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/datastore#org_access Datastore#org_access}
	OrgAccess *string `field:"optional" json:"orgAccess" yaml:"orgAccess"`
	// Can be set to `uuid` to automatically generate primary keys when new items are added.
	//
	// Default value is `none`, which requires you to supply a primary key for each new item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/datastore#primary_key_generation_strategy Datastore#primary_key_generation_strategy}
	PrimaryKeyGenerationStrategy *string `field:"optional" json:"primaryKeyGenerationStrategy" yaml:"primaryKeyGenerationStrategy"`
}

