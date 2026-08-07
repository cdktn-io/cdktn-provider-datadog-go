// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityfindingsmuterulesorder

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityFindingsMuteRulesOrderConfig struct {
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
	// A unique identifier for the order resource. This field has no server-side equivalent; Datadog recommends matching the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/security_findings_mute_rules_order#name SecurityFindingsMuteRulesOrder#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ordered list of all mute rule IDs.
	//
	// The order of IDs in this attribute defines the evaluation order of the mute rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/security_findings_mute_rules_order#rule_ids SecurityFindingsMuteRulesOrder#rule_ids}
	RuleIds *[]*string `field:"required" json:"ruleIds" yaml:"ruleIds"`
}

