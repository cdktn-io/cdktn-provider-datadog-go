// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package teamnotificationrule


type TeamNotificationRuleEmail struct {
	// Whether to send email notifications to team members when alerts are triggered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/team_notification_rule#enabled TeamNotificationRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

