// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package referencetable


type ReferenceTableFileMetadataAccessDetailsAzureDetail struct {
	// The Azure client ID (application ID).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/reference_table#azure_client_id ReferenceTable#azure_client_id}
	AzureClientId *string `field:"optional" json:"azureClientId" yaml:"azureClientId"`
	// The name of the Azure container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/reference_table#azure_container_name ReferenceTable#azure_container_name}
	AzureContainerName *string `field:"optional" json:"azureContainerName" yaml:"azureContainerName"`
	// The name of the Azure storage account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/reference_table#azure_storage_account_name ReferenceTable#azure_storage_account_name}
	AzureStorageAccountName *string `field:"optional" json:"azureStorageAccountName" yaml:"azureStorageAccountName"`
	// The ID of the Azure tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/reference_table#azure_tenant_id ReferenceTable#azure_tenant_id}
	AzureTenantId *string `field:"optional" json:"azureTenantId" yaml:"azureTenantId"`
	// The relative file path from the Azure container root to the CSV file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/reference_table#file_path ReferenceTable#file_path}
	FilePath *string `field:"optional" json:"filePath" yaml:"filePath"`
}

