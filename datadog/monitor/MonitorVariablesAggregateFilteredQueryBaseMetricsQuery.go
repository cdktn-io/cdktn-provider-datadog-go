// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitor


type MonitorVariablesAggregateFilteredQueryBaseMetricsQuery struct {
	// The data source for metrics queries. Valid values are `metrics`, `cloud_cost`, `datadog_usage`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/monitor#data_source Monitor#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// The metrics query definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/monitor#query Monitor#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// The aggregation method for metrics queries.
	//
	// Valid values are `avg`, `min`, `max`, `sum`, `last`, `mean`, `area`, `l2norm`, `percentile`, `stddev`, `count_unique`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/monitor#aggregator Monitor#aggregator}
	Aggregator *string `field:"optional" json:"aggregator" yaml:"aggregator"`
	// The name of the query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/monitor#name Monitor#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

