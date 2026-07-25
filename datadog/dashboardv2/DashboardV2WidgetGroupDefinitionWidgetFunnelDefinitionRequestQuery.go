// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetFunnelDefinitionRequestQuery struct {
	// The data source for funnel queries. Valid values are `rum`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#data_source DashboardV2#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// The widget query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#query_string DashboardV2#query_string}
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#step DashboardV2#step}
	Step interface{} `field:"optional" json:"step" yaml:"step"`
}

