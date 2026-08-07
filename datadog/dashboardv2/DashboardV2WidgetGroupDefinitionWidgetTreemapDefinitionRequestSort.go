// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestSort struct {
	// The number of items to limit the widget to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#count DashboardV2#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// order_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#order_by DashboardV2#order_by}
	OrderBy interface{} `field:"optional" json:"orderBy" yaml:"orderBy"`
}

