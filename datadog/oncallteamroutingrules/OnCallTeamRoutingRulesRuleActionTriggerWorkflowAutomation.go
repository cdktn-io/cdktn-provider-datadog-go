// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oncallteamroutingrules


type OnCallTeamRoutingRulesRuleActionTriggerWorkflowAutomation struct {
	// The handle of the Workflow Automation to trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/on_call_team_routing_rules#handle OnCallTeamRoutingRules#handle}
	Handle *string `field:"optional" json:"handle" yaml:"handle"`
}

