// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package costbudget


type CostBudgetBudgetLine struct {
	// Map of month (YYYYMM) to budget amount. Example: {"202601": 1000.0, "202602": 1200.0}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/cost_budget#amounts CostBudget#amounts}
	Amounts *map[string]*float64 `field:"required" json:"amounts" yaml:"amounts"`
	// child_tag_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/cost_budget#child_tag_filters CostBudget#child_tag_filters}
	ChildTagFilters interface{} `field:"optional" json:"childTagFilters" yaml:"childTagFilters"`
	// parent_tag_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/cost_budget#parent_tag_filters CostBudget#parent_tag_filters}
	ParentTagFilters interface{} `field:"optional" json:"parentTagFilters" yaml:"parentTagFilters"`
	// tag_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/cost_budget#tag_filters CostBudget#tag_filters}
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

