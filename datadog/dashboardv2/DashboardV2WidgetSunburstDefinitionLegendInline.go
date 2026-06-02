// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSunburstDefinitionLegendInline struct {
	// The type of legend (inline or automatic).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#type DashboardV2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Whether to hide the percentages of the groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#hide_percent DashboardV2#hide_percent}
	HidePercent interface{} `field:"optional" json:"hidePercent" yaml:"hidePercent"`
	// Whether to hide the values of the groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#hide_value DashboardV2#hide_value}
	HideValue interface{} `field:"optional" json:"hideValue" yaml:"hideValue"`
}

