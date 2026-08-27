// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetSunburstDefinitionRequestQuery struct {
	// apm_dependency_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#apm_dependency_stats_query PowerpackV2#apm_dependency_stats_query}
	ApmDependencyStatsQuery *PowerpackV2WidgetSunburstDefinitionRequestQueryApmDependencyStatsQuery `field:"optional" json:"apmDependencyStatsQuery" yaml:"apmDependencyStatsQuery"`
	// apm_metrics_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#apm_metrics_query PowerpackV2#apm_metrics_query}
	ApmMetricsQuery *PowerpackV2WidgetSunburstDefinitionRequestQueryApmMetricsQuery `field:"optional" json:"apmMetricsQuery" yaml:"apmMetricsQuery"`
	// apm_resource_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#apm_resource_stats_query PowerpackV2#apm_resource_stats_query}
	ApmResourceStatsQuery *PowerpackV2WidgetSunburstDefinitionRequestQueryApmResourceStatsQuery `field:"optional" json:"apmResourceStatsQuery" yaml:"apmResourceStatsQuery"`
	// cloud_cost_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#cloud_cost_query PowerpackV2#cloud_cost_query}
	CloudCostQuery *PowerpackV2WidgetSunburstDefinitionRequestQueryCloudCostQuery `field:"optional" json:"cloudCostQuery" yaml:"cloudCostQuery"`
	// event_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#event_query PowerpackV2#event_query}
	EventQuery *PowerpackV2WidgetSunburstDefinitionRequestQueryEventQuery `field:"optional" json:"eventQuery" yaml:"eventQuery"`
	// metric_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#metric_query PowerpackV2#metric_query}
	MetricQuery *PowerpackV2WidgetSunburstDefinitionRequestQueryMetricQuery `field:"optional" json:"metricQuery" yaml:"metricQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#process_query PowerpackV2#process_query}
	ProcessQuery *PowerpackV2WidgetSunburstDefinitionRequestQueryProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// product_analytics_extended_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#product_analytics_extended_query PowerpackV2#product_analytics_extended_query}
	ProductAnalyticsExtendedQuery *PowerpackV2WidgetSunburstDefinitionRequestQueryProductAnalyticsExtendedQuery `field:"optional" json:"productAnalyticsExtendedQuery" yaml:"productAnalyticsExtendedQuery"`
	// retention_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#retention_query PowerpackV2#retention_query}
	RetentionQuery *PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuery `field:"optional" json:"retentionQuery" yaml:"retentionQuery"`
	// slo_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#slo_query PowerpackV2#slo_query}
	SloQuery *PowerpackV2WidgetSunburstDefinitionRequestQuerySloQuery `field:"optional" json:"sloQuery" yaml:"sloQuery"`
	// user_journey_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#user_journey_query PowerpackV2#user_journey_query}
	UserJourneyQuery *PowerpackV2WidgetSunburstDefinitionRequestQueryUserJourneyQuery `field:"optional" json:"userJourneyQuery" yaml:"userJourneyQuery"`
}

