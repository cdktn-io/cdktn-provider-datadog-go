// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitor


type MonitorVariablesDataQualityQueryMonitorOptions struct {
	// Crontab expression to override the default schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/monitor#crontab_override Monitor#crontab_override}
	CrontabOverride *string `field:"optional" json:"crontabOverride" yaml:"crontabOverride"`
	// Custom SQL query for the monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/monitor#custom_sql Monitor#custom_sql}
	CustomSql *string `field:"optional" json:"customSql" yaml:"customSql"`
	// Custom WHERE clause for the query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/monitor#custom_where Monitor#custom_where}
	CustomWhere *string `field:"optional" json:"customWhere" yaml:"customWhere"`
	// Columns to group results by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/monitor#group_by_columns Monitor#group_by_columns}
	GroupByColumns *[]*string `field:"optional" json:"groupByColumns" yaml:"groupByColumns"`
	// Override for the model type. Valid values are `freshness`, `percentage`, `any`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/monitor#model_type_override Monitor#model_type_override}
	ModelTypeOverride *string `field:"optional" json:"modelTypeOverride" yaml:"modelTypeOverride"`
}

