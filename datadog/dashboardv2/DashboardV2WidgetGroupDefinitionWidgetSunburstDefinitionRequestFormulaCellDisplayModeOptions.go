// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestFormulaCellDisplayModeOptions struct {
	// The type of trend line to display. Valid values are `area`, `line`, and `bars`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#trend_type DashboardV2#trend_type}
	TrendType *string `field:"optional" json:"trendType" yaml:"trendType"`
	// The scale of the y-axis. Valid values are `shared` and `independent`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#y_scale DashboardV2#y_scale}
	YScale *string `field:"optional" json:"yScale" yaml:"yScale"`
}

