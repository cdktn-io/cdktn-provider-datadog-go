// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oncallusernotificationrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OnCallUserNotificationRuleConfig struct {
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
	// Notification category to associate the rule with. Valid values are `high_urgency`, `low_urgency`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/on_call_user_notification_rule#category OnCallUserNotificationRule#category}
	Category *string `field:"required" json:"category" yaml:"category"`
	// ID of the notification channel to associate the notification rule with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/on_call_user_notification_rule#channel_id OnCallUserNotificationRule#channel_id}
	ChannelId *string `field:"required" json:"channelId" yaml:"channelId"`
	// Number of minutes to elapse before this rule is evaluated.  `0` indicates immediate evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/on_call_user_notification_rule#delay_minutes OnCallUserNotificationRule#delay_minutes}
	DelayMinutes *float64 `field:"required" json:"delayMinutes" yaml:"delayMinutes"`
	// ID of the user to associate the notification rule with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/on_call_user_notification_rule#user_id OnCallUserNotificationRule#user_id}
	UserId *string `field:"required" json:"userId" yaml:"userId"`
	// phone block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/on_call_user_notification_rule#phone OnCallUserNotificationRule#phone}
	Phone *OnCallUserNotificationRulePhone `field:"optional" json:"phone" yaml:"phone"`
}

