// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitornotificationrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MonitorNotificationRuleConfig struct {
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
	// The name of the monitor notification rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/monitor_notification_rule#name MonitorNotificationRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// conditional_recipients block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/monitor_notification_rule#conditional_recipients MonitorNotificationRule#conditional_recipients}
	ConditionalRecipients *MonitorNotificationRuleConditionalRecipients `field:"optional" json:"conditionalRecipients" yaml:"conditionalRecipients"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/monitor_notification_rule#filter MonitorNotificationRule#filter}
	Filter *MonitorNotificationRuleFilter `field:"optional" json:"filter" yaml:"filter"`
	// List of recipients to notify. Cannot be used with `conditional_recipients`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/monitor_notification_rule#recipients MonitorNotificationRule#recipients}
	Recipients *[]*string `field:"optional" json:"recipients" yaml:"recipients"`
}

