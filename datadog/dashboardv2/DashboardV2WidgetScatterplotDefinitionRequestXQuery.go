// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetScatterplotDefinitionRequestXQuery struct {
	// apm_dependency_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#apm_dependency_stats_query DashboardV2#apm_dependency_stats_query}
	ApmDependencyStatsQuery *DashboardV2WidgetScatterplotDefinitionRequestXQueryApmDependencyStatsQuery `field:"optional" json:"apmDependencyStatsQuery" yaml:"apmDependencyStatsQuery"`
	// apm_metrics_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#apm_metrics_query DashboardV2#apm_metrics_query}
	ApmMetricsQuery *DashboardV2WidgetScatterplotDefinitionRequestXQueryApmMetricsQuery `field:"optional" json:"apmMetricsQuery" yaml:"apmMetricsQuery"`
	// apm_resource_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#apm_resource_stats_query DashboardV2#apm_resource_stats_query}
	ApmResourceStatsQuery *DashboardV2WidgetScatterplotDefinitionRequestXQueryApmResourceStatsQuery `field:"optional" json:"apmResourceStatsQuery" yaml:"apmResourceStatsQuery"`
	// cloud_cost_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#cloud_cost_query DashboardV2#cloud_cost_query}
	CloudCostQuery *DashboardV2WidgetScatterplotDefinitionRequestXQueryCloudCostQuery `field:"optional" json:"cloudCostQuery" yaml:"cloudCostQuery"`
	// event_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#event_query DashboardV2#event_query}
	EventQuery *DashboardV2WidgetScatterplotDefinitionRequestXQueryEventQuery `field:"optional" json:"eventQuery" yaml:"eventQuery"`
	// metric_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#metric_query DashboardV2#metric_query}
	MetricQuery *DashboardV2WidgetScatterplotDefinitionRequestXQueryMetricQuery `field:"optional" json:"metricQuery" yaml:"metricQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#process_query DashboardV2#process_query}
	ProcessQuery *DashboardV2WidgetScatterplotDefinitionRequestXQueryProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// product_analytics_extended_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#product_analytics_extended_query DashboardV2#product_analytics_extended_query}
	ProductAnalyticsExtendedQuery *DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQuery `field:"optional" json:"productAnalyticsExtendedQuery" yaml:"productAnalyticsExtendedQuery"`
	// retention_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#retention_query DashboardV2#retention_query}
	RetentionQuery *DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuery `field:"optional" json:"retentionQuery" yaml:"retentionQuery"`
	// slo_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#slo_query DashboardV2#slo_query}
	SloQuery *DashboardV2WidgetScatterplotDefinitionRequestXQuerySloQuery `field:"optional" json:"sloQuery" yaml:"sloQuery"`
	// user_journey_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#user_journey_query DashboardV2#user_journey_query}
	UserJourneyQuery *DashboardV2WidgetScatterplotDefinitionRequestXQueryUserJourneyQuery `field:"optional" json:"userJourneyQuery" yaml:"userJourneyQuery"`
}

