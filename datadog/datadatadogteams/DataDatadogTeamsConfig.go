// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatadogteams

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatadogTeamsConfig struct {
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
	// Search query. Can be team name, team handle, or email of team member.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/data-sources/teams#filter_keyword DataDatadogTeams#filter_keyword}
	FilterKeyword *string `field:"optional" json:"filterKeyword" yaml:"filterKeyword"`
	// When true, only returns teams the current user belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/data-sources/teams#filter_me DataDatadogTeams#filter_me}
	FilterMe interface{} `field:"optional" json:"filterMe" yaml:"filterMe"`
	// teams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/data-sources/teams#teams DataDatadogTeams#teams}
	Teams interface{} `field:"optional" json:"teams" yaml:"teams"`
}

