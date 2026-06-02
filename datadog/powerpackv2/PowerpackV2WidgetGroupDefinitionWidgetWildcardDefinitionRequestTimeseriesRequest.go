// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequest struct {
	// apm_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#apm_query PowerpackV2#apm_query}
	ApmQuery *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestApmQuery `field:"optional" json:"apmQuery" yaml:"apmQuery"`
	// How the data points are displayed on the graph.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#display_type PowerpackV2#display_type}
	DisplayType *string `field:"optional" json:"displayType" yaml:"displayType"`
	// formula block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#formula PowerpackV2#formula}
	Formula interface{} `field:"optional" json:"formula" yaml:"formula"`
	// log_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#log_query PowerpackV2#log_query}
	LogQuery *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestLogQuery `field:"optional" json:"logQuery" yaml:"logQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#process_query PowerpackV2#process_query}
	ProcessQuery *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// The metric query to use for this widget. **Deprecated.** Use queries and formulas instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#q PowerpackV2#q}
	Q *string `field:"optional" json:"q" yaml:"q"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#query PowerpackV2#query}
	Query interface{} `field:"optional" json:"query" yaml:"query"`
	// rum_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#rum_query PowerpackV2#rum_query}
	RumQuery *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestRumQuery `field:"optional" json:"rumQuery" yaml:"rumQuery"`
	// security_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#security_query PowerpackV2#security_query}
	SecurityQuery *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestSecurityQuery `field:"optional" json:"securityQuery" yaml:"securityQuery"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#sort PowerpackV2#sort}
	Sort *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestSort `field:"optional" json:"sort" yaml:"sort"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#style PowerpackV2#style}
	Style *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestStyle `field:"optional" json:"style" yaml:"style"`
}

