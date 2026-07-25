// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetBarChartDefinitionRequestQuery struct {
	// apm_dependency_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#apm_dependency_stats_query DashboardV2#apm_dependency_stats_query}
	ApmDependencyStatsQuery *DashboardV2WidgetBarChartDefinitionRequestQueryApmDependencyStatsQuery `field:"optional" json:"apmDependencyStatsQuery" yaml:"apmDependencyStatsQuery"`
	// apm_resource_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#apm_resource_stats_query DashboardV2#apm_resource_stats_query}
	ApmResourceStatsQuery *DashboardV2WidgetBarChartDefinitionRequestQueryApmResourceStatsQuery `field:"optional" json:"apmResourceStatsQuery" yaml:"apmResourceStatsQuery"`
	// cloud_cost_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#cloud_cost_query DashboardV2#cloud_cost_query}
	CloudCostQuery *DashboardV2WidgetBarChartDefinitionRequestQueryCloudCostQuery `field:"optional" json:"cloudCostQuery" yaml:"cloudCostQuery"`
	// event_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#event_query DashboardV2#event_query}
	EventQuery *DashboardV2WidgetBarChartDefinitionRequestQueryEventQuery `field:"optional" json:"eventQuery" yaml:"eventQuery"`
	// metric_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#metric_query DashboardV2#metric_query}
	MetricQuery *DashboardV2WidgetBarChartDefinitionRequestQueryMetricQuery `field:"optional" json:"metricQuery" yaml:"metricQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#process_query DashboardV2#process_query}
	ProcessQuery *DashboardV2WidgetBarChartDefinitionRequestQueryProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// slo_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#slo_query DashboardV2#slo_query}
	SloQuery *DashboardV2WidgetBarChartDefinitionRequestQuerySloQuery `field:"optional" json:"sloQuery" yaml:"sloQuery"`
}

