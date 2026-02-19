// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitornotificationrule


type MonitorNotificationRuleConditionalRecipientsConditions struct {
	// A list of recipients to notify.
	//
	// Uses the same format as the monitor message field. Must not start with an '@'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/monitor_notification_rule#recipients MonitorNotificationRule#recipients}
	Recipients *[]*string `field:"required" json:"recipients" yaml:"recipients"`
	// Defines the condition under which the recipients are notified.
	//
	// Supported formats: Monitor status condition using `transition_type:<status>` (for example `transition_type:is_alert`) or a single tag `key:value pair` (for example `env:prod`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/monitor_notification_rule#scope MonitorNotificationRule#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

