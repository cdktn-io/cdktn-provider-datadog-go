// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oncallusernotificationrule


type OnCallUserNotificationRulePhone struct {
	// Specifies the method in which a phone is used in a notification rule. Valid values are `sms`, `voice`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/on_call_user_notification_rule#method OnCallUserNotificationRule#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
}

