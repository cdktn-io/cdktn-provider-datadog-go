// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringcriticalasset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityMonitoringCriticalAssetConfig struct {
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
	// The query used to match a critical asset and the associated signals.
	//
	// Uses the same syntax as the search bar in the Security Signals Explorer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/security_monitoring_critical_asset#query SecurityMonitoringCriticalAsset#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// The rule query to filter which detection rules this critical asset applies to.
	//
	// Uses the same syntax as the search bar for detection rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/security_monitoring_critical_asset#rule_query SecurityMonitoringCriticalAsset#rule_query}
	RuleQuery *string `field:"required" json:"ruleQuery" yaml:"ruleQuery"`
	// The severity change applied to signals matching this critical asset.
	//
	// Valid values are `critical`, `high`, `medium`, `low`, `info`, `increase`, `decrease`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/security_monitoring_critical_asset#severity SecurityMonitoringCriticalAsset#severity}
	Severity *string `field:"required" json:"severity" yaml:"severity"`
	// Whether the critical asset is enabled. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/security_monitoring_critical_asset#enabled SecurityMonitoringCriticalAsset#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of tags associated with the critical asset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/security_monitoring_critical_asset#tags SecurityMonitoringCriticalAsset#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

