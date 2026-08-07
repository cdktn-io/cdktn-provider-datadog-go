// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringdefaultrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityMonitoringDefaultRuleConfig struct {
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
	// case block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/security_monitoring_default_rule#case SecurityMonitoringDefaultRule#case}
	Case interface{} `field:"optional" json:"case" yaml:"case"`
	// Custom Message (will override default message) for generated signals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/security_monitoring_default_rule#custom_message SecurityMonitoringDefaultRule#custom_message}
	CustomMessage *string `field:"optional" json:"customMessage" yaml:"customMessage"`
	// The name (will override default name) of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/security_monitoring_default_rule#custom_name SecurityMonitoringDefaultRule#custom_name}
	CustomName *string `field:"optional" json:"customName" yaml:"customName"`
	// Custom tags for generated signals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/security_monitoring_default_rule#custom_tags SecurityMonitoringDefaultRule#custom_tags}
	CustomTags *[]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// Enable the rule. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/security_monitoring_default_rule#enabled SecurityMonitoringDefaultRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/security_monitoring_default_rule#filter SecurityMonitoringDefaultRule#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/security_monitoring_default_rule#options SecurityMonitoringDefaultRule#options}
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/security_monitoring_default_rule#query SecurityMonitoringDefaultRule#query}
	Query interface{} `field:"optional" json:"query" yaml:"query"`
}

