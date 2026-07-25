// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequest struct {
	// apm_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#apm_query PowerpackV2#apm_query}
	ApmQuery *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestApmQuery `field:"optional" json:"apmQuery" yaml:"apmQuery"`
	// apm_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#apm_stats_query PowerpackV2#apm_stats_query}
	ApmStatsQuery *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestApmStatsQuery `field:"optional" json:"apmStatsQuery" yaml:"apmStatsQuery"`
	// formula block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#formula PowerpackV2#formula}
	Formula interface{} `field:"optional" json:"formula" yaml:"formula"`
	// histogram_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#histogram_query PowerpackV2#histogram_query}
	HistogramQuery *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestHistogramQuery `field:"optional" json:"histogramQuery" yaml:"histogramQuery"`
	// log_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#log_query PowerpackV2#log_query}
	LogQuery *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestLogQuery `field:"optional" json:"logQuery" yaml:"logQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#process_query PowerpackV2#process_query}
	ProcessQuery *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// The metric query to use for this widget. **Deprecated.** Use queries and formulas instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#q PowerpackV2#q}
	Q *string `field:"optional" json:"q" yaml:"q"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#query PowerpackV2#query}
	Query interface{} `field:"optional" json:"query" yaml:"query"`
	// Set to 'histogram' for distribution-of-point-values requests. Valid values are `histogram`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#request_type PowerpackV2#request_type}
	RequestType *string `field:"optional" json:"requestType" yaml:"requestType"`
	// rum_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#rum_query PowerpackV2#rum_query}
	RumQuery *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestRumQuery `field:"optional" json:"rumQuery" yaml:"rumQuery"`
	// security_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#security_query PowerpackV2#security_query}
	SecurityQuery *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestSecurityQuery `field:"optional" json:"securityQuery" yaml:"securityQuery"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#style PowerpackV2#style}
	Style *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestStyle `field:"optional" json:"style" yaml:"style"`
}

