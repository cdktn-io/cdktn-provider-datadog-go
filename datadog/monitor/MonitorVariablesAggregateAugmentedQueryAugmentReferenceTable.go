// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitor


type MonitorVariablesAggregateAugmentedQueryAugmentReferenceTable struct {
	// Must be `reference_table`. Valid values are `reference_table`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/monitor#data_source Monitor#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// Name of the reference table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/monitor#table_name Monitor#table_name}
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/monitor#columns Monitor#columns}
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// Name of the augment sub-query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/monitor#name Monitor#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Optional filter expression for the reference table query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/monitor#query_filter Monitor#query_filter}
	QueryFilter *string `field:"optional" json:"queryFilter" yaml:"queryFilter"`
}

