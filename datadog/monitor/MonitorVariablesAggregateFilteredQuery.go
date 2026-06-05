// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitor


type MonitorVariablesAggregateFilteredQuery struct {
	// The data source for aggregate-filtered composite queries. Must be `aggregate_filtered_query`. Valid values are `aggregate_filtered_query`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/monitor#data_source Monitor#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/monitor#filters Monitor#filters}
	Filters interface{} `field:"required" json:"filters" yaml:"filters"`
	// base_event_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/monitor#base_event_query Monitor#base_event_query}
	BaseEventQuery *MonitorVariablesAggregateFilteredQueryBaseEventQuery `field:"optional" json:"baseEventQuery" yaml:"baseEventQuery"`
	// base_metrics_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/monitor#base_metrics_query Monitor#base_metrics_query}
	BaseMetricsQuery *MonitorVariablesAggregateFilteredQueryBaseMetricsQuery `field:"optional" json:"baseMetricsQuery" yaml:"baseMetricsQuery"`
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/monitor#compute Monitor#compute}
	Compute interface{} `field:"optional" json:"compute" yaml:"compute"`
	// filter_event_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/monitor#filter_event_query Monitor#filter_event_query}
	FilterEventQuery *MonitorVariablesAggregateFilteredQueryFilterEventQuery `field:"optional" json:"filterEventQuery" yaml:"filterEventQuery"`
	// filter_reference_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/monitor#filter_reference_table Monitor#filter_reference_table}
	FilterReferenceTable *MonitorVariablesAggregateFilteredQueryFilterReferenceTable `field:"optional" json:"filterReferenceTable" yaml:"filterReferenceTable"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/monitor#group_by Monitor#group_by}
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// Name of the query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/monitor#name Monitor#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

