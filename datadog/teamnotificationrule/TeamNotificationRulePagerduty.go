// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package teamnotificationrule


type TeamNotificationRulePagerduty struct {
	// PagerDuty service name to send incident notifications to. The service name can be found in your PagerDuty service settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/team_notification_rule#service_name TeamNotificationRule#service_name}
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

