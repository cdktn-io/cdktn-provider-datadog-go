// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestStyle struct {
	// How to order series. Valid values are `tags`, `values`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#order_by DashboardV2#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// A color palette to apply to the widget. The available options are available at: https://docs.datadoghq.com/dashboards/widgets/timeseries/#appearance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#palette DashboardV2#palette}
	Palette *string `field:"optional" json:"palette" yaml:"palette"`
}

