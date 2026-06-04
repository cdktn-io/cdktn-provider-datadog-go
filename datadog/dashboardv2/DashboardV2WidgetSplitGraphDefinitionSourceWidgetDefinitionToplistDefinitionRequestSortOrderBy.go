// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestSortOrderBy struct {
	// formula_sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#formula_sort DashboardV2#formula_sort}
	FormulaSort *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestSortOrderByFormulaSort `field:"optional" json:"formulaSort" yaml:"formulaSort"`
	// group_sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#group_sort DashboardV2#group_sort}
	GroupSort *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestSortOrderByGroupSort `field:"optional" json:"groupSort" yaml:"groupSort"`
}

