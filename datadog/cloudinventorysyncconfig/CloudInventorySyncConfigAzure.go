// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudinventorysyncconfig


type CloudInventorySyncConfigAzure struct {
	// Azure Client ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/cloud_inventory_sync_config#client_id CloudInventorySyncConfig#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Azure Storage Container name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/cloud_inventory_sync_config#container CloudInventorySyncConfig#container}
	Container *string `field:"optional" json:"container" yaml:"container"`
	// Azure Resource Group name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/cloud_inventory_sync_config#resource_group CloudInventorySyncConfig#resource_group}
	ResourceGroup *string `field:"optional" json:"resourceGroup" yaml:"resourceGroup"`
	// Azure Storage Account name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/cloud_inventory_sync_config#storage_account CloudInventorySyncConfig#storage_account}
	StorageAccount *string `field:"optional" json:"storageAccount" yaml:"storageAccount"`
	// Azure Subscription ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/cloud_inventory_sync_config#subscription_id CloudInventorySyncConfig#subscription_id}
	SubscriptionId *string `field:"optional" json:"subscriptionId" yaml:"subscriptionId"`
	// Azure Tenant ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/cloud_inventory_sync_config#tenant_id CloudInventorySyncConfig#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

