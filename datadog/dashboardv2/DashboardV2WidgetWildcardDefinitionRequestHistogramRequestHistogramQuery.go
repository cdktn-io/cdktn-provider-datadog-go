// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQuery struct {
	// apm_resource_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#apm_resource_stats_query DashboardV2#apm_resource_stats_query}
	ApmResourceStatsQuery *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryApmResourceStatsQuery `field:"optional" json:"apmResourceStatsQuery" yaml:"apmResourceStatsQuery"`
	// event_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#event_query DashboardV2#event_query}
	EventQuery *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryEventQuery `field:"optional" json:"eventQuery" yaml:"eventQuery"`
	// metric_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#metric_query DashboardV2#metric_query}
	MetricQuery *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryMetricQuery `field:"optional" json:"metricQuery" yaml:"metricQuery"`
}

