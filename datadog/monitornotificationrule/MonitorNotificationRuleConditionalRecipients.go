// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitornotificationrule


type MonitorNotificationRuleConditionalRecipients struct {
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/monitor_notification_rule#conditions MonitorNotificationRule#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// If none of the `conditions` applied, `fallback_recipients` will get notified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/monitor_notification_rule#fallback_recipients MonitorNotificationRule#fallback_recipients}
	FallbackRecipients *[]*string `field:"optional" json:"fallbackRecipients" yaml:"fallbackRecipients"`
}

