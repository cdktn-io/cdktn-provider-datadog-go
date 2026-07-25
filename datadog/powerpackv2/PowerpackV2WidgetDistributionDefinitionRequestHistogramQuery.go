// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetDistributionDefinitionRequestHistogramQuery struct {
	// apm_resource_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#apm_resource_stats_query PowerpackV2#apm_resource_stats_query}
	ApmResourceStatsQuery *PowerpackV2WidgetDistributionDefinitionRequestHistogramQueryApmResourceStatsQuery `field:"optional" json:"apmResourceStatsQuery" yaml:"apmResourceStatsQuery"`
	// event_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#event_query PowerpackV2#event_query}
	EventQuery *PowerpackV2WidgetDistributionDefinitionRequestHistogramQueryEventQuery `field:"optional" json:"eventQuery" yaml:"eventQuery"`
	// metric_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#metric_query PowerpackV2#metric_query}
	MetricQuery *PowerpackV2WidgetDistributionDefinitionRequestHistogramQueryMetricQuery `field:"optional" json:"metricQuery" yaml:"metricQuery"`
}

