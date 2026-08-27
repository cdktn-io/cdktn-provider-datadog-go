// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeRollup struct {
	// Type of calendar interval. Valid values are `day`, `week`, `month`, `year`, `quarter`, `minute`, `hour`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#type DashboardV2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Alignment of the calendar interval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#alignment DashboardV2#alignment}
	Alignment *string `field:"optional" json:"alignment" yaml:"alignment"`
	// Quantity of the calendar interval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#quantity DashboardV2#quantity}
	Quantity *float64 `field:"optional" json:"quantity" yaml:"quantity"`
	// Timezone for the calendar interval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#timezone DashboardV2#timezone}
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

