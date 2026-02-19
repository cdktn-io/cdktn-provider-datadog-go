// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datastoreitem

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatastoreItemConfig struct {
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
	// The unique identifier of the datastore containing this item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/datastore_item#datastore_id DatastoreItem#datastore_id}
	DatastoreId *string `field:"required" json:"datastoreId" yaml:"datastoreId"`
	// The primary key value that identifies this item. Cannot exceed 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/datastore_item#item_key DatastoreItem#item_key}
	ItemKey *string `field:"required" json:"itemKey" yaml:"itemKey"`
	// The data content (as key-value pairs) of the datastore item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/datastore_item#value DatastoreItem#value}
	Value *map[string]*string `field:"required" json:"value" yaml:"value"`
}

