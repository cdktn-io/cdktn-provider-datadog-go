// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetListStreamDefinitionRequest struct {
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/dashboard_v2#columns DashboardV2#columns}
	Columns interface{} `field:"required" json:"columns" yaml:"columns"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/dashboard_v2#query DashboardV2#query}
	Query *DashboardV2WidgetGroupDefinitionWidgetListStreamDefinitionRequestQuery `field:"required" json:"query" yaml:"query"`
	// Widget response format. Valid values are `event_list`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/dashboard_v2#response_format DashboardV2#response_format}
	ResponseFormat *string `field:"required" json:"responseFormat" yaml:"responseFormat"`
}

