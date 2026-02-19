// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customallocationrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CustomAllocationRuleConfig struct {
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
	// Whether the custom allocation rule is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/custom_allocation_rule#enabled CustomAllocationRule#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// List of cloud providers the rule applies to. Valid values include `aws`, `azure`, and `gcp`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/custom_allocation_rule#providernames CustomAllocationRule#providernames}
	Providernames *[]*string `field:"required" json:"providernames" yaml:"providernames"`
	// The name of the custom allocation rule.
	//
	// This field is immutable - changing it will force replacement of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/custom_allocation_rule#rule_name CustomAllocationRule#rule_name}
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// costs_to_allocate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/custom_allocation_rule#costs_to_allocate CustomAllocationRule#costs_to_allocate}
	CostsToAllocate interface{} `field:"optional" json:"costsToAllocate" yaml:"costsToAllocate"`
	// strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/custom_allocation_rule#strategy CustomAllocationRule#strategy}
	Strategy *CustomAllocationRuleStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

