// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetRetentionCurveDefinitionRequestQuery struct {
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#compute DashboardV2#compute}
	Compute *DashboardV2WidgetRetentionCurveDefinitionRequestQueryCompute `field:"required" json:"compute" yaml:"compute"`
	// Data source for retention queries. Valid values are `product_analytics_retention`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#data_source DashboardV2#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// search block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#search DashboardV2#search}
	Search *DashboardV2WidgetRetentionCurveDefinitionRequestQuerySearch `field:"required" json:"search" yaml:"search"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#filters DashboardV2#filters}
	Filters *DashboardV2WidgetRetentionCurveDefinitionRequestQueryFilters `field:"optional" json:"filters" yaml:"filters"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#group_by DashboardV2#group_by}
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// Name of the retention query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#name DashboardV2#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

