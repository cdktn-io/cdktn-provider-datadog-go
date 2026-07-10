// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitor


type MonitorVariablesDataJobsQuery struct {
	// Filter expression used to select the jobs to monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/monitor#jobs_query Monitor#jobs_query}
	JobsQuery *string `field:"required" json:"jobsQuery" yaml:"jobsQuery"`
	// The type of job being monitored.
	//
	// Valid values include `databricks.job`, `spark.application`, `airflow.dag`, `dbt.job`, `dbt.model`, `dbt.test`, `glue.job`. Custom job types are supported with the `custom.ol.` prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/monitor#job_type Monitor#job_type}
	JobType *string `field:"required" json:"jobType" yaml:"jobType"`
	// Name of the query for use in formulas. Must be `run_query`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/monitor#name Monitor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Query dialect for data jobs queries. Currently only `metric` is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/monitor#query_dialect Monitor#query_dialect}
	QueryDialect *string `field:"required" json:"queryDialect" yaml:"queryDialect"`
}

