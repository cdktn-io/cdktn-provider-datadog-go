// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationmsteamstenantbasedhandle

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IntegrationMsTeamsTenantBasedHandleConfig struct {
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
	// Your channel name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/integration_ms_teams_tenant_based_handle#channel_name IntegrationMsTeamsTenantBasedHandle#channel_name}
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// Your tenant-based handle name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/integration_ms_teams_tenant_based_handle#name IntegrationMsTeamsTenantBasedHandle#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Your team name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/integration_ms_teams_tenant_based_handle#team_name IntegrationMsTeamsTenantBasedHandle#team_name}
	TeamName *string `field:"required" json:"teamName" yaml:"teamName"`
	// Your tenant name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/integration_ms_teams_tenant_based_handle#tenant_name IntegrationMsTeamsTenantBasedHandle#tenant_name}
	TenantName *string `field:"required" json:"tenantName" yaml:"tenantName"`
}

