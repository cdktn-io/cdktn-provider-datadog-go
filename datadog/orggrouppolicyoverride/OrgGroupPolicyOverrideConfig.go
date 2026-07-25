// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package orggrouppolicyoverride

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OrgGroupPolicyOverrideConfig struct {
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
	// The UUID of the org group that owns the policy. Must be a valid UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/org_group_policy_override#org_group_id OrgGroupPolicyOverride#org_group_id}
	OrgGroupId *string `field:"required" json:"orgGroupId" yaml:"orgGroupId"`
	// The short site name of the organization (e.g., `us1`, `eu1`, `us1-fed`). Part of the override's server-side identity; changing it replaces the resource. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/org_group_policy_override#org_site OrgGroupPolicyOverride#org_site}
	OrgSite *string `field:"required" json:"orgSite" yaml:"orgSite"`
	// The UUID of the organization being exempted from the policy. Must be a valid UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/org_group_policy_override#org_uuid OrgGroupPolicyOverride#org_uuid}
	OrgUuid *string `field:"required" json:"orgUuid" yaml:"orgUuid"`
	// The UUID of the org group policy the override applies to. Must be a valid UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/org_group_policy_override#policy_id OrgGroupPolicyOverride#policy_id}
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
}

