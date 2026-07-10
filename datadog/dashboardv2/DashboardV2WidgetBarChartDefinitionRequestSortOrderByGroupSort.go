// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetBarChartDefinitionRequestSortOrderByGroupSort struct {
	// The name of the group tag to sort by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#name DashboardV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Widget sorting direction. Valid values are `asc`, `desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#order DashboardV2#order}
	Order *string `field:"required" json:"order" yaml:"order"`
}

