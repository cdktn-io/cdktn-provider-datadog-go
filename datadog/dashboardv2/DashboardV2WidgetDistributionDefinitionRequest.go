// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetDistributionDefinitionRequest struct {
	// apm_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#apm_query DashboardV2#apm_query}
	ApmQuery *DashboardV2WidgetDistributionDefinitionRequestApmQuery `field:"optional" json:"apmQuery" yaml:"apmQuery"`
	// apm_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#apm_stats_query DashboardV2#apm_stats_query}
	ApmStatsQuery *DashboardV2WidgetDistributionDefinitionRequestApmStatsQuery `field:"optional" json:"apmStatsQuery" yaml:"apmStatsQuery"`
	// formula block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#formula DashboardV2#formula}
	Formula interface{} `field:"optional" json:"formula" yaml:"formula"`
	// histogram_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#histogram_query DashboardV2#histogram_query}
	HistogramQuery *DashboardV2WidgetDistributionDefinitionRequestHistogramQuery `field:"optional" json:"histogramQuery" yaml:"histogramQuery"`
	// log_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#log_query DashboardV2#log_query}
	LogQuery *DashboardV2WidgetDistributionDefinitionRequestLogQuery `field:"optional" json:"logQuery" yaml:"logQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#process_query DashboardV2#process_query}
	ProcessQuery *DashboardV2WidgetDistributionDefinitionRequestProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// The metric query to use for this widget. **Deprecated.** Use queries and formulas instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#q DashboardV2#q}
	Q *string `field:"optional" json:"q" yaml:"q"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#query DashboardV2#query}
	Query interface{} `field:"optional" json:"query" yaml:"query"`
	// Set to 'histogram' for distribution-of-point-values requests. Valid values are `histogram`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#request_type DashboardV2#request_type}
	RequestType *string `field:"optional" json:"requestType" yaml:"requestType"`
	// rum_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#rum_query DashboardV2#rum_query}
	RumQuery *DashboardV2WidgetDistributionDefinitionRequestRumQuery `field:"optional" json:"rumQuery" yaml:"rumQuery"`
	// security_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#security_query DashboardV2#security_query}
	SecurityQuery *DashboardV2WidgetDistributionDefinitionRequestSecurityQuery `field:"optional" json:"securityQuery" yaml:"securityQuery"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#style DashboardV2#style}
	Style *DashboardV2WidgetDistributionDefinitionRequestStyle `field:"optional" json:"style" yaml:"style"`
}

