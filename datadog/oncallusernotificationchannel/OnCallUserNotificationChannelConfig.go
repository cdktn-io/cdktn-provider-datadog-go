// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oncallusernotificationchannel

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OnCallUserNotificationChannelConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// ID of the user to associate the notification channel with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/on_call_user_notification_channel#user_id OnCallUserNotificationChannel#user_id}
	UserId *string `field:"required" json:"userId" yaml:"userId"`
	// email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/on_call_user_notification_channel#email OnCallUserNotificationChannel#email}
	Email *OnCallUserNotificationChannelEmail `field:"optional" json:"email" yaml:"email"`
	// phone block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/on_call_user_notification_channel#phone OnCallUserNotificationChannel#phone}
	Phone *OnCallUserNotificationChannelPhone `field:"optional" json:"phone" yaml:"phone"`
}

