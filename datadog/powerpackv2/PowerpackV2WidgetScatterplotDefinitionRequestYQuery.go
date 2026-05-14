// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetScatterplotDefinitionRequestYQuery struct {
	// apm_dependency_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/powerpack_v2#apm_dependency_stats_query PowerpackV2#apm_dependency_stats_query}
	ApmDependencyStatsQuery *PowerpackV2WidgetScatterplotDefinitionRequestYQueryApmDependencyStatsQuery `field:"optional" json:"apmDependencyStatsQuery" yaml:"apmDependencyStatsQuery"`
	// apm_resource_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/powerpack_v2#apm_resource_stats_query PowerpackV2#apm_resource_stats_query}
	ApmResourceStatsQuery *PowerpackV2WidgetScatterplotDefinitionRequestYQueryApmResourceStatsQuery `field:"optional" json:"apmResourceStatsQuery" yaml:"apmResourceStatsQuery"`
	// cloud_cost_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/powerpack_v2#cloud_cost_query PowerpackV2#cloud_cost_query}
	CloudCostQuery *PowerpackV2WidgetScatterplotDefinitionRequestYQueryCloudCostQuery `field:"optional" json:"cloudCostQuery" yaml:"cloudCostQuery"`
	// event_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/powerpack_v2#event_query PowerpackV2#event_query}
	EventQuery *PowerpackV2WidgetScatterplotDefinitionRequestYQueryEventQuery `field:"optional" json:"eventQuery" yaml:"eventQuery"`
	// metric_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/powerpack_v2#metric_query PowerpackV2#metric_query}
	MetricQuery *PowerpackV2WidgetScatterplotDefinitionRequestYQueryMetricQuery `field:"optional" json:"metricQuery" yaml:"metricQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/powerpack_v2#process_query PowerpackV2#process_query}
	ProcessQuery *PowerpackV2WidgetScatterplotDefinitionRequestYQueryProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// slo_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/powerpack_v2#slo_query PowerpackV2#slo_query}
	SloQuery *PowerpackV2WidgetScatterplotDefinitionRequestYQuerySloQuery `field:"optional" json:"sloQuery" yaml:"sloQuery"`
}

