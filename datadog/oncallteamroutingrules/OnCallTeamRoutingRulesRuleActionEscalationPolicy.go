// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oncallteamroutingrules


type OnCallTeamRoutingRulesRuleActionEscalationPolicy struct {
	// Number of minutes before an acknowledged page is re-triggered. Value must be between 30 and 4320.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/on_call_team_routing_rules#ack_timeout_minutes OnCallTeamRoutingRules#ack_timeout_minutes}
	AckTimeoutMinutes *float64 `field:"optional" json:"ackTimeoutMinutes" yaml:"ackTimeoutMinutes"`
	// Escalation policy ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/on_call_team_routing_rules#policy_id OnCallTeamRoutingRules#policy_id}
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
	// support_hours block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/on_call_team_routing_rules#support_hours OnCallTeamRoutingRules#support_hours}
	SupportHours *OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHours `field:"optional" json:"supportHours" yaml:"supportHours"`
	// Urgency for pages created via this action. Valid values are `high`, `low`, `dynamic`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/on_call_team_routing_rules#urgency OnCallTeamRoutingRules#urgency}
	Urgency *string `field:"optional" json:"urgency" yaml:"urgency"`
}

