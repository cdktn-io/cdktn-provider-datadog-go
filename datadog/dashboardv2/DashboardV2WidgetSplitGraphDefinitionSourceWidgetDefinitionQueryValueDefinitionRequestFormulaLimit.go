// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestFormulaLimit struct {
	// The number of results to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/dashboard_v2#count DashboardV2#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// The direction of the sort. Valid values are `asc`, `desc`. Defaults to `"desc"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/dashboard_v2#order DashboardV2#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}

