// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitornotificationrule


type MonitorNotificationRuleFilter struct {
	// A scope expression composed of `key:value` pairs (such as `env:prod`) with boolean operators (AND, OR, NOT) and parentheses for grouping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/monitor_notification_rule#scope MonitorNotificationRule#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// A list of tag key:value pairs (e.g. team:product). All tags must match (AND semantics).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/monitor_notification_rule#tags MonitorNotificationRule#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

