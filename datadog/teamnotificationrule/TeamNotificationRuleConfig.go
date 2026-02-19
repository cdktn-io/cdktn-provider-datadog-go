// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package teamnotificationrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TeamNotificationRuleConfig struct {
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
	// The ID of the team that this notification rule belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/team_notification_rule#team_id TeamNotificationRule#team_id}
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/team_notification_rule#email TeamNotificationRule#email}
	Email *TeamNotificationRuleEmail `field:"optional" json:"email" yaml:"email"`
	// ms_teams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/team_notification_rule#ms_teams TeamNotificationRule#ms_teams}
	MsTeams *TeamNotificationRuleMsTeams `field:"optional" json:"msTeams" yaml:"msTeams"`
	// pagerduty block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/team_notification_rule#pagerduty TeamNotificationRule#pagerduty}
	Pagerduty *TeamNotificationRulePagerduty `field:"optional" json:"pagerduty" yaml:"pagerduty"`
	// slack block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/team_notification_rule#slack TeamNotificationRule#slack}
	Slack *TeamNotificationRuleSlack `field:"optional" json:"slack" yaml:"slack"`
}

