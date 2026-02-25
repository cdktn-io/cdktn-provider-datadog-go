// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudinventorysyncconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudInventorySyncConfigConfig struct {
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
	// The cloud provider type. Valid values are `aws`, `azure`, `gcp`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/cloud_inventory_sync_config#cloud_provider CloudInventorySyncConfig#cloud_provider}
	CloudProvider *string `field:"required" json:"cloudProvider" yaml:"cloudProvider"`
	// aws block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/cloud_inventory_sync_config#aws CloudInventorySyncConfig#aws}
	Aws *CloudInventorySyncConfigAws `field:"optional" json:"aws" yaml:"aws"`
	// azure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/cloud_inventory_sync_config#azure CloudInventorySyncConfig#azure}
	Azure *CloudInventorySyncConfigAzure `field:"optional" json:"azure" yaml:"azure"`
	// gcp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/cloud_inventory_sync_config#gcp CloudInventorySyncConfig#gcp}
	Gcp *CloudInventorySyncConfigGcp `field:"optional" json:"gcp" yaml:"gcp"`
}

