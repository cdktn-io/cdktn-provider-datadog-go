// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package orggrouppolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OrgGroupPolicyConfig struct {
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
	// The policy content as a JSON-encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/org_group_policy#content OrgGroupPolicy#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The UUID of the org group this policy belongs to. Must be a valid UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/org_group_policy#org_group_id OrgGroupPolicy#org_group_id}
	OrgGroupId *string `field:"required" json:"orgGroupId" yaml:"orgGroupId"`
	// The name of the policy. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/org_group_policy#policy_name OrgGroupPolicy#policy_name}
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// The enforcement tier of the policy.
	//
	// `DEFAULT` means the policy is set but member orgs may mutate it. `ENFORCE` means the policy is strictly controlled and mutations are blocked for affected orgs. `DELEGATE` means each member org controls its own value. Valid values are `DEFAULT`, `ENFORCE`, `DELEGATE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/org_group_policy#enforcement_tier OrgGroupPolicy#enforcement_tier}
	EnforcementTier *string `field:"optional" json:"enforcementTier" yaml:"enforcementTier"`
	// The type of the policy. Valid values are `org_config`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/org_group_policy#policy_type OrgGroupPolicy#policy_type}
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
}

