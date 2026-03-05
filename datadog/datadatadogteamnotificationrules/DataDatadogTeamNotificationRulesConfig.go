// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatadogteamnotificationrules

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatadogTeamNotificationRulesConfig struct {
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
	// The team ID to fetch notification rules for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/data-sources/team_notification_rules#team_id DataDatadogTeamNotificationRules#team_id}
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// notification_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/data-sources/team_notification_rules#notification_rules DataDatadogTeamNotificationRules#notification_rules}
	NotificationRules interface{} `field:"optional" json:"notificationRules" yaml:"notificationRules"`
}

