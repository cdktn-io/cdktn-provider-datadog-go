// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatadogcustomallocationrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatadogCustomAllocationRuleConfig struct {
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
	// costs_to_allocate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/data-sources/custom_allocation_rule#costs_to_allocate DataDatadogCustomAllocationRule#costs_to_allocate}
	CostsToAllocate interface{} `field:"optional" json:"costsToAllocate" yaml:"costsToAllocate"`
	// The ID of the custom allocation rule to retrieve.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/data-sources/custom_allocation_rule#rule_id DataDatadogCustomAllocationRule#rule_id}
	RuleId *float64 `field:"optional" json:"ruleId" yaml:"ruleId"`
	// strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/data-sources/custom_allocation_rule#strategy DataDatadogCustomAllocationRule#strategy}
	Strategy *DataDatadogCustomAllocationRuleStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

