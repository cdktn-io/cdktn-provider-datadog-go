// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitor


type MonitorVariablesAggregateAugmentedQuery struct {
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/monitor#compute Monitor#compute}
	Compute interface{} `field:"required" json:"compute" yaml:"compute"`
	// The data source for aggregate-augmented composite queries. Must be `aggregate_augmented_query`. Valid values are `aggregate_augmented_query`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/monitor#data_source Monitor#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/monitor#group_by Monitor#group_by}
	GroupBy interface{} `field:"required" json:"groupBy" yaml:"groupBy"`
	// join_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/monitor#join_condition Monitor#join_condition}
	JoinCondition *MonitorVariablesAggregateAugmentedQueryJoinCondition `field:"required" json:"joinCondition" yaml:"joinCondition"`
	// augment_event_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/monitor#augment_event_query Monitor#augment_event_query}
	AugmentEventQuery *MonitorVariablesAggregateAugmentedQueryAugmentEventQuery `field:"optional" json:"augmentEventQuery" yaml:"augmentEventQuery"`
	// augment_reference_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/monitor#augment_reference_table Monitor#augment_reference_table}
	AugmentReferenceTable *MonitorVariablesAggregateAugmentedQueryAugmentReferenceTable `field:"optional" json:"augmentReferenceTable" yaml:"augmentReferenceTable"`
	// base_event_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/monitor#base_event_query Monitor#base_event_query}
	BaseEventQuery *MonitorVariablesAggregateAugmentedQueryBaseEventQuery `field:"optional" json:"baseEventQuery" yaml:"baseEventQuery"`
	// base_metrics_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/monitor#base_metrics_query Monitor#base_metrics_query}
	BaseMetricsQuery *MonitorVariablesAggregateAugmentedQueryBaseMetricsQuery `field:"optional" json:"baseMetricsQuery" yaml:"baseMetricsQuery"`
	// Name of the query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/monitor#name Monitor#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

