// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatadogcostbudget


type DataDatadogCostBudgetBudgetLine struct {
	// child_tag_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/data-sources/cost_budget#child_tag_filters DataDatadogCostBudget#child_tag_filters}
	ChildTagFilters interface{} `field:"optional" json:"childTagFilters" yaml:"childTagFilters"`
	// parent_tag_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/data-sources/cost_budget#parent_tag_filters DataDatadogCostBudget#parent_tag_filters}
	ParentTagFilters interface{} `field:"optional" json:"parentTagFilters" yaml:"parentTagFilters"`
	// tag_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/data-sources/cost_budget#tag_filters DataDatadogCostBudget#tag_filters}
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

