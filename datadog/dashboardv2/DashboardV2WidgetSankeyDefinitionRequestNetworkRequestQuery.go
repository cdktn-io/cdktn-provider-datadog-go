// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSankeyDefinitionRequestNetworkRequestQuery struct {
	// The data source for the Sankey network query. Valid values are `network_device_flows`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/dashboard_v2#data_source DashboardV2#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// Fields to group by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/dashboard_v2#group_by DashboardV2#group_by}
	GroupBy *[]*string `field:"required" json:"groupBy" yaml:"groupBy"`
	// Maximum number of results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/dashboard_v2#limit DashboardV2#limit}
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	// The search query string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/dashboard_v2#query_string DashboardV2#query_string}
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/dashboard_v2#compute DashboardV2#compute}
	Compute *DashboardV2WidgetSankeyDefinitionRequestNetworkRequestQueryCompute `field:"optional" json:"compute" yaml:"compute"`
	// The mode for the Sankey network query. Valid values are `target`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/dashboard_v2#mode DashboardV2#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Whether to exclude missing values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/dashboard_v2#should_exclude_missing DashboardV2#should_exclude_missing}
	ShouldExcludeMissing interface{} `field:"optional" json:"shouldExcludeMissing" yaml:"shouldExcludeMissing"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/dashboard_v2#sort DashboardV2#sort}
	Sort *DashboardV2WidgetSankeyDefinitionRequestNetworkRequestQuerySort `field:"optional" json:"sort" yaml:"sort"`
}

