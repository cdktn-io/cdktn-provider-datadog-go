// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetHeatmapDefinitionRequestHistogramRequestHistogramQuery struct {
	// metric_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#metric_query DashboardV2#metric_query}
	MetricQuery *DashboardV2WidgetHeatmapDefinitionRequestHistogramRequestHistogramQueryMetricQuery `field:"required" json:"metricQuery" yaml:"metricQuery"`
}

