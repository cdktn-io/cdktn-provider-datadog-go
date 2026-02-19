// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oncallusernotificationchannel


type OnCallUserNotificationChannelEmail struct {
	// The e-mail address to be notified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/on_call_user_notification_channel#address OnCallUserNotificationChannel#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Preferred content formats for notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/on_call_user_notification_channel#formats OnCallUserNotificationChannel#formats}
	Formats *[]*string `field:"optional" json:"formats" yaml:"formats"`
}

