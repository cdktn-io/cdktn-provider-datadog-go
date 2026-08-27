// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetProductAnalyticsFunnelDefinitionRequestQuery struct {
	// Data source for Product Analytics funnel queries. Valid values are `product_analytics_journey`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#data_source DashboardV2#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// search block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#search DashboardV2#search}
	Search *DashboardV2WidgetProductAnalyticsFunnelDefinitionRequestQuerySearch `field:"required" json:"search" yaml:"search"`
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#compute DashboardV2#compute}
	Compute *DashboardV2WidgetProductAnalyticsFunnelDefinitionRequestQueryCompute `field:"optional" json:"compute" yaml:"compute"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#group_by DashboardV2#group_by}
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// Subquery identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#subquery_id DashboardV2#subquery_id}
	SubqueryId *string `field:"optional" json:"subqueryId" yaml:"subqueryId"`
}

