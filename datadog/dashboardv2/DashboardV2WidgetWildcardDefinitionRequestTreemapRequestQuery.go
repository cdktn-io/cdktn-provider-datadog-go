// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQuery struct {
	// apm_dependency_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#apm_dependency_stats_query DashboardV2#apm_dependency_stats_query}
	ApmDependencyStatsQuery *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryApmDependencyStatsQuery `field:"optional" json:"apmDependencyStatsQuery" yaml:"apmDependencyStatsQuery"`
	// apm_metrics_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#apm_metrics_query DashboardV2#apm_metrics_query}
	ApmMetricsQuery *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryApmMetricsQuery `field:"optional" json:"apmMetricsQuery" yaml:"apmMetricsQuery"`
	// apm_resource_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#apm_resource_stats_query DashboardV2#apm_resource_stats_query}
	ApmResourceStatsQuery *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryApmResourceStatsQuery `field:"optional" json:"apmResourceStatsQuery" yaml:"apmResourceStatsQuery"`
	// cloud_cost_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#cloud_cost_query DashboardV2#cloud_cost_query}
	CloudCostQuery *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryCloudCostQuery `field:"optional" json:"cloudCostQuery" yaml:"cloudCostQuery"`
	// event_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#event_query DashboardV2#event_query}
	EventQuery *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryEventQuery `field:"optional" json:"eventQuery" yaml:"eventQuery"`
	// metric_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#metric_query DashboardV2#metric_query}
	MetricQuery *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryMetricQuery `field:"optional" json:"metricQuery" yaml:"metricQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#process_query DashboardV2#process_query}
	ProcessQuery *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// product_analytics_extended_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#product_analytics_extended_query DashboardV2#product_analytics_extended_query}
	ProductAnalyticsExtendedQuery *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryProductAnalyticsExtendedQuery `field:"optional" json:"productAnalyticsExtendedQuery" yaml:"productAnalyticsExtendedQuery"`
	// retention_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#retention_query DashboardV2#retention_query}
	RetentionQuery *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuery `field:"optional" json:"retentionQuery" yaml:"retentionQuery"`
	// slo_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#slo_query DashboardV2#slo_query}
	SloQuery *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQuerySloQuery `field:"optional" json:"sloQuery" yaml:"sloQuery"`
	// user_journey_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#user_journey_query DashboardV2#user_journey_query}
	UserJourneyQuery *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryUserJourneyQuery `field:"optional" json:"userJourneyQuery" yaml:"userJourneyQuery"`
}

