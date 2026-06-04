// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetBarChartDefinitionStyle struct {
	// display block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#display DashboardV2#display}
	Display *DashboardV2WidgetGroupDefinitionWidgetBarChartDefinitionStyleDisplay `field:"optional" json:"display" yaml:"display"`
	// Color palette for the bar chart.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#palette DashboardV2#palette}
	Palette *string `field:"optional" json:"palette" yaml:"palette"`
	// Scaling definition for the bar chart. Valid values are `absolute`, `relative`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#scaling DashboardV2#scaling}
	Scaling *string `field:"optional" json:"scaling" yaml:"scaling"`
}

