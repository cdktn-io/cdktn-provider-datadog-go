// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSankeyDefinitionRequestRumRequestQuery struct {
	// The data source for the Sankey RUM query. Valid values are `rum`, `product_analytics`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#data_source DashboardV2#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// The mode for the Sankey RUM query. Valid values are `source`, `target`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#mode DashboardV2#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The search query string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#query_string DashboardV2#query_string}
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// Entries per step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#entries_per_step DashboardV2#entries_per_step}
	EntriesPerStep *float64 `field:"optional" json:"entriesPerStep" yaml:"entriesPerStep"`
	// Number of steps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#number_of_steps DashboardV2#number_of_steps}
	NumberOfSteps *float64 `field:"optional" json:"numberOfSteps" yaml:"numberOfSteps"`
	// Source field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#source DashboardV2#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Subquery ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#subquery_id DashboardV2#subquery_id}
	SubqueryId *string `field:"optional" json:"subqueryId" yaml:"subqueryId"`
	// Target field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#target DashboardV2#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
}

