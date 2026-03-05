// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatasetConfig struct {
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
	// The name of the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/dataset#name Dataset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// An array of principals.
	//
	// A principal is a subject or group of subjects. Each principal is formatted as `type:id`. Supported types: `role` and `team`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/dataset#principals Dataset#principals}
	Principals *[]*string `field:"required" json:"principals" yaml:"principals"`
	// product_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/dataset#product_filters Dataset#product_filters}
	ProductFilters interface{} `field:"optional" json:"productFilters" yaml:"productFilters"`
}

