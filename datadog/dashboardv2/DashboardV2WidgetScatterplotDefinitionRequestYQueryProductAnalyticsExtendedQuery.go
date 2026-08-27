// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetScatterplotDefinitionRequestYQueryProductAnalyticsExtendedQuery struct {
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#compute DashboardV2#compute}
	Compute *DashboardV2WidgetScatterplotDefinitionRequestYQueryProductAnalyticsExtendedQueryCompute `field:"required" json:"compute" yaml:"compute"`
	// Data source for Product Analytics Extended queries. Valid values are `product_analytics_extended`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#data_source DashboardV2#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// Name of the query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#name DashboardV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#query DashboardV2#query}
	Query *DashboardV2WidgetScatterplotDefinitionRequestYQueryProductAnalyticsExtendedQueryQuery `field:"required" json:"query" yaml:"query"`
	// audience_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#audience_filters DashboardV2#audience_filters}
	AudienceFilters *DashboardV2WidgetScatterplotDefinitionRequestYQueryProductAnalyticsExtendedQueryAudienceFilters `field:"optional" json:"audienceFilters" yaml:"audienceFilters"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#group_by DashboardV2#group_by}
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// Event indexes to query. Use `*` to query all indexes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#indexes DashboardV2#indexes}
	Indexes *[]*string `field:"optional" json:"indexes" yaml:"indexes"`
}

