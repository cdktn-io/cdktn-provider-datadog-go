// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequest struct {
	// apm_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#apm_query DashboardV2#apm_query}
	ApmQuery *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestApmQuery `field:"optional" json:"apmQuery" yaml:"apmQuery"`
	// How the data points are displayed on the graph.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#display_type DashboardV2#display_type}
	DisplayType *string `field:"optional" json:"displayType" yaml:"displayType"`
	// formula block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#formula DashboardV2#formula}
	Formula interface{} `field:"optional" json:"formula" yaml:"formula"`
	// log_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#log_query DashboardV2#log_query}
	LogQuery *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestLogQuery `field:"optional" json:"logQuery" yaml:"logQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#process_query DashboardV2#process_query}
	ProcessQuery *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// The metric query to use for this widget. **Deprecated.** Use queries and formulas instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#q DashboardV2#q}
	Q *string `field:"optional" json:"q" yaml:"q"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#query DashboardV2#query}
	Query interface{} `field:"optional" json:"query" yaml:"query"`
	// rum_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#rum_query DashboardV2#rum_query}
	RumQuery *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestRumQuery `field:"optional" json:"rumQuery" yaml:"rumQuery"`
	// security_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#security_query DashboardV2#security_query}
	SecurityQuery *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestSecurityQuery `field:"optional" json:"securityQuery" yaml:"securityQuery"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#sort DashboardV2#sort}
	Sort *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestSort `field:"optional" json:"sort" yaml:"sort"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#style DashboardV2#style}
	Style *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestStyle `field:"optional" json:"style" yaml:"style"`
}

