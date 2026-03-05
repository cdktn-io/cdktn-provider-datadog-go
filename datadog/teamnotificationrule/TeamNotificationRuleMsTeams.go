// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package teamnotificationrule


type TeamNotificationRuleMsTeams struct {
	// MS Teams connector name used to route notifications to the appropriate channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/team_notification_rule#connector_name TeamNotificationRule#connector_name}
	ConnectorName *string `field:"optional" json:"connectorName" yaml:"connectorName"`
}

