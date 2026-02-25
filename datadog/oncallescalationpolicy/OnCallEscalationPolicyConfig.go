// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oncallescalationpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OnCallEscalationPolicyConfig struct {
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
	// A human-readable name for the escalation policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/on_call_escalation_policy#name OnCallEscalationPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// If true, pages will be automatically resolved if unacknowledged after the final step. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/on_call_escalation_policy#resolve_page_on_policy_end OnCallEscalationPolicy#resolve_page_on_policy_end}
	ResolvePageOnPolicyEnd interface{} `field:"optional" json:"resolvePageOnPolicyEnd" yaml:"resolvePageOnPolicyEnd"`
	// If set, policy will be retried this many times after the final step.
	//
	// Must be in the range 0-10. Value must be between 0 and 10. Defaults to `0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/on_call_escalation_policy#retries OnCallEscalationPolicy#retries}
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/on_call_escalation_policy#step OnCallEscalationPolicy#step}
	Step interface{} `field:"optional" json:"step" yaml:"step"`
	// A list of team ids associated with the escalation policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/on_call_escalation_policy#teams OnCallEscalationPolicy#teams}
	Teams *[]*string `field:"optional" json:"teams" yaml:"teams"`
}

