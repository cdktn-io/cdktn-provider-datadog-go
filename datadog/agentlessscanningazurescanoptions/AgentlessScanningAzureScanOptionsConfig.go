// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentlessscanningazurescanoptions

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AgentlessScanningAzureScanOptionsConfig struct {
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
	// The Azure subscription ID for which agentless scanning is configured. Must be a valid Azure subscription ID (UUID format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/agentless_scanning_azure_scan_options#azure_subscription_id AgentlessScanningAzureScanOptions#azure_subscription_id}
	AzureSubscriptionId *string `field:"required" json:"azureSubscriptionId" yaml:"azureSubscriptionId"`
	// Indicates if scanning for vulnerabilities in containers is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/agentless_scanning_azure_scan_options#vuln_containers_os AgentlessScanningAzureScanOptions#vuln_containers_os}
	VulnContainersOs interface{} `field:"required" json:"vulnContainersOs" yaml:"vulnContainersOs"`
	// Indicates if scanning for vulnerabilities in hosts is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/agentless_scanning_azure_scan_options#vuln_host_os AgentlessScanningAzureScanOptions#vuln_host_os}
	VulnHostOs interface{} `field:"required" json:"vulnHostOs" yaml:"vulnHostOs"`
	// Indicates whether host compliance scanning is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/agentless_scanning_azure_scan_options#compliance_host AgentlessScanningAzureScanOptions#compliance_host}
	ComplianceHost interface{} `field:"optional" json:"complianceHost" yaml:"complianceHost"`
}

