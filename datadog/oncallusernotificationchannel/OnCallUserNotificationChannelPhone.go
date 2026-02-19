// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oncallusernotificationchannel


type OnCallUserNotificationChannelPhone struct {
	// The E-164 formatted phone number (for example, +3371234567).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/on_call_user_notification_channel#number OnCallUserNotificationChannel#number}
	Number *string `field:"optional" json:"number" yaml:"number"`
}

