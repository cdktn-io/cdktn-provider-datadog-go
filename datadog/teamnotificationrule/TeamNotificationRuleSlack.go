// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package teamnotificationrule


type TeamNotificationRuleSlack struct {
	// Slack channel name for notifications (for example, #alerts or #team-notifications).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/team_notification_rule#channel TeamNotificationRule#channel}
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// Slack workspace name where the channel is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/team_notification_rule#workspace TeamNotificationRule#workspace}
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
}

