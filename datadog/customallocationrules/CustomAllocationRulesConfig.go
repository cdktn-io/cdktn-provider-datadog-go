// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customallocationrules

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CustomAllocationRulesConfig struct {
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
	// The list of Custom Allocation Rule IDs, in order.
	//
	// Rules are executed in the order specified in this list. Comes from the `id` field on a `datadog_custom_allocation_rule` resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/custom_allocation_rules#rule_ids CustomAllocationRules#rule_ids}
	RuleIds *[]*string `field:"required" json:"ruleIds" yaml:"ruleIds"`
	// Whether to override UI-defined rules.
	//
	// When set to true, any rules created via the UI that are not defined in Terraform will be deleted and Terraform will be used as the source of truth for rules and their ordering. When set to false, any rules created via the UI that are at the end of order will be kept but will be warned, otherwise an error will be thrown in terraform plan phase. Default is false
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/custom_allocation_rules#override_ui_defined_resources CustomAllocationRules#override_ui_defined_resources}
	OverrideUiDefinedResources interface{} `field:"optional" json:"overrideUiDefinedResources" yaml:"overrideUiDefinedResources"`
}

