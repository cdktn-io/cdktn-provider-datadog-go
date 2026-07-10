// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetHostmapDefinitionRequestFillQuery struct {
	// apm_dependency_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#apm_dependency_stats_query PowerpackV2#apm_dependency_stats_query}
	ApmDependencyStatsQuery *PowerpackV2WidgetHostmapDefinitionRequestFillQueryApmDependencyStatsQuery `field:"optional" json:"apmDependencyStatsQuery" yaml:"apmDependencyStatsQuery"`
	// apm_resource_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#apm_resource_stats_query PowerpackV2#apm_resource_stats_query}
	ApmResourceStatsQuery *PowerpackV2WidgetHostmapDefinitionRequestFillQueryApmResourceStatsQuery `field:"optional" json:"apmResourceStatsQuery" yaml:"apmResourceStatsQuery"`
	// cloud_cost_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#cloud_cost_query PowerpackV2#cloud_cost_query}
	CloudCostQuery *PowerpackV2WidgetHostmapDefinitionRequestFillQueryCloudCostQuery `field:"optional" json:"cloudCostQuery" yaml:"cloudCostQuery"`
	// event_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#event_query PowerpackV2#event_query}
	EventQuery *PowerpackV2WidgetHostmapDefinitionRequestFillQueryEventQuery `field:"optional" json:"eventQuery" yaml:"eventQuery"`
	// metric_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#metric_query PowerpackV2#metric_query}
	MetricQuery *PowerpackV2WidgetHostmapDefinitionRequestFillQueryMetricQuery `field:"optional" json:"metricQuery" yaml:"metricQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#process_query PowerpackV2#process_query}
	ProcessQuery *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// slo_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#slo_query PowerpackV2#slo_query}
	SloQuery *PowerpackV2WidgetHostmapDefinitionRequestFillQuerySloQuery `field:"optional" json:"sloQuery" yaml:"sloQuery"`
}

