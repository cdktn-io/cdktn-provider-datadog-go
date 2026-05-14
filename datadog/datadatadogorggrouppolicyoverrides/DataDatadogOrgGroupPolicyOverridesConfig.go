// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatadogorggrouppolicyoverrides

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatadogOrgGroupPolicyOverridesConfig struct {
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
	// The UUID of the org group whose overrides to list. Must be a valid UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/data-sources/org_group_policy_overrides#org_group_id DataDatadogOrgGroupPolicyOverrides#org_group_id}
	OrgGroupId *string `field:"required" json:"orgGroupId" yaml:"orgGroupId"`
	// Filter overrides to those for the given organization.
	//
	// Applied client-side after the List call because the API does not accept an `org_uuid` filter on this endpoint. Must be a valid UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/data-sources/org_group_policy_overrides#org_uuid DataDatadogOrgGroupPolicyOverrides#org_uuid}
	OrgUuid *string `field:"optional" json:"orgUuid" yaml:"orgUuid"`
	// Filter overrides to those on the given policy. Must be a valid UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/data-sources/org_group_policy_overrides#policy_id DataDatadogOrgGroupPolicyOverrides#policy_id}
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
}

