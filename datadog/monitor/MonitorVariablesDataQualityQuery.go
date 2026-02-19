// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitor


type MonitorVariablesDataQualityQuery struct {
	// The data source for data quality queries. Valid value is `data_quality_metrics`. Valid values are `data_quality_metrics`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/monitor#data_source Monitor#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// Filter expression used to match on data entities. Uses AAstra query syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/monitor#filter Monitor#filter}
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// The measure to query.
	//
	// Common values include `bytes`, `cardinality`, `custom`, `freshness`, `max`, `mean`, `min`, `nullness`, `percent_negative`, `percent_zero`, `row_count`, `stddev`, `sum`, `uniqueness`. Additional values may be supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/monitor#measure Monitor#measure}
	Measure *string `field:"required" json:"measure" yaml:"measure"`
	// The name of the query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/monitor#name Monitor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional grouping fields for aggregation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/monitor#group_by Monitor#group_by}
	GroupBy *[]*string `field:"optional" json:"groupBy" yaml:"groupBy"`
	// monitor_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/monitor#monitor_options Monitor#monitor_options}
	MonitorOptions *MonitorVariablesDataQualityQueryMonitorOptions `field:"optional" json:"monitorOptions" yaml:"monitorOptions"`
	// Schema version for the data quality query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/monitor#schema_version Monitor#schema_version}
	SchemaVersion *string `field:"optional" json:"schemaVersion" yaml:"schemaVersion"`
	// Optional scoping expression to further filter metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/monitor#scope Monitor#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

