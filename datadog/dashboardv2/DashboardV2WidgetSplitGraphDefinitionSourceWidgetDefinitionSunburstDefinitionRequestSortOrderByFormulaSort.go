// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestSortOrderByFormulaSort struct {
	// The index of the formula to sort by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#index DashboardV2#index}
	Index *float64 `field:"required" json:"index" yaml:"index"`
	// Widget sorting direction. Valid values are `asc`, `desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#order DashboardV2#order}
	Order *string `field:"required" json:"order" yaml:"order"`
}

