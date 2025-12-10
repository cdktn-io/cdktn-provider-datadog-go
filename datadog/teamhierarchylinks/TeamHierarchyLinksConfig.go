// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teamhierarchylinks

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TeamHierarchyLinksConfig struct {
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
	// ID of the parent team the team hierarchy link is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.82.0/docs/resources/team_hierarchy_links#parent_team_id TeamHierarchyLinks#parent_team_id}
	ParentTeamId *string `field:"required" json:"parentTeamId" yaml:"parentTeamId"`
	// ID of the sub team the team hierarchy link is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.82.0/docs/resources/team_hierarchy_links#sub_team_id TeamHierarchyLinks#sub_team_id}
	SubTeamId *string `field:"required" json:"subTeamId" yaml:"subTeamId"`
}

