// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatadogteamhierarchylinks

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatadogTeamHierarchyLinksConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Filter by parent team ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.82.0/docs/data-sources/team_hierarchy_links#filter_parent_team DataDatadogTeamHierarchyLinks#filter_parent_team}
	FilterParentTeam *string `field:"optional" json:"filterParentTeam" yaml:"filterParentTeam"`
	// Filter by sub team ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.82.0/docs/data-sources/team_hierarchy_links#filter_sub_team DataDatadogTeamHierarchyLinks#filter_sub_team}
	FilterSubTeam *string `field:"optional" json:"filterSubTeam" yaml:"filterSubTeam"`
	// The team hierarchy link’s identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.82.0/docs/data-sources/team_hierarchy_links#link_id DataDatadogTeamHierarchyLinks#link_id}
	LinkId *string `field:"optional" json:"linkId" yaml:"linkId"`
}

