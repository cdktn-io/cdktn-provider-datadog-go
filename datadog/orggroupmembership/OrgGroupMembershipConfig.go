// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package orggroupmembership

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OrgGroupMembershipConfig struct {
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
	// The UUID of the org group to assign the organization to. Must be a valid UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/org_group_membership#org_group_id OrgGroupMembership#org_group_id}
	OrgGroupId *string `field:"required" json:"orgGroupId" yaml:"orgGroupId"`
	// The UUID of the organization. Must be a valid UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/org_group_membership#org_uuid OrgGroupMembership#org_uuid}
	OrgUuid *string `field:"required" json:"orgUuid" yaml:"orgUuid"`
}

