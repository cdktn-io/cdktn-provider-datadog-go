// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oncallteamroutingrules


type OnCallTeamRoutingRulesRuleAction struct {
	// escalation_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/on_call_team_routing_rules#escalation_policy OnCallTeamRoutingRules#escalation_policy}
	EscalationPolicy *OnCallTeamRoutingRulesRuleActionEscalationPolicy `field:"optional" json:"escalationPolicy" yaml:"escalationPolicy"`
	// send_slack_message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/on_call_team_routing_rules#send_slack_message OnCallTeamRoutingRules#send_slack_message}
	SendSlackMessage *OnCallTeamRoutingRulesRuleActionSendSlackMessage `field:"optional" json:"sendSlackMessage" yaml:"sendSlackMessage"`
	// send_teams_message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/on_call_team_routing_rules#send_teams_message OnCallTeamRoutingRules#send_teams_message}
	SendTeamsMessage *OnCallTeamRoutingRulesRuleActionSendTeamsMessage `field:"optional" json:"sendTeamsMessage" yaml:"sendTeamsMessage"`
	// trigger_workflow_automation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/on_call_team_routing_rules#trigger_workflow_automation OnCallTeamRoutingRules#trigger_workflow_automation}
	TriggerWorkflowAutomation *OnCallTeamRoutingRulesRuleActionTriggerWorkflowAutomation `field:"optional" json:"triggerWorkflowAutomation" yaml:"triggerWorkflowAutomation"`
}

