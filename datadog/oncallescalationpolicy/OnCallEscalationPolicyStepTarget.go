// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oncallescalationpolicy


type OnCallEscalationPolicyStepTarget struct {
	// For schedule targets, specifies which on-call user to page.
	//
	// Valid values: `current` (default), `previous`, `next`. Valid values are `current`, `previous`, `next`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.82.0/docs/resources/on_call_escalation_policy#position OnCallEscalationPolicy#position}
	Position *string `field:"optional" json:"position" yaml:"position"`
	// Targeted schedule ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.82.0/docs/resources/on_call_escalation_policy#schedule OnCallEscalationPolicy#schedule}
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// Targeted team ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.82.0/docs/resources/on_call_escalation_policy#team OnCallEscalationPolicy#team}
	Team *string `field:"optional" json:"team" yaml:"team"`
	// Targeted user ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.82.0/docs/resources/on_call_escalation_policy#user OnCallEscalationPolicy#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

