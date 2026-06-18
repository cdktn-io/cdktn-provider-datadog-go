// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetTimeseriesDefinitionRequest struct {
	// apm_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#apm_query PowerpackV2#apm_query}
	ApmQuery *PowerpackV2WidgetTimeseriesDefinitionRequestApmQuery `field:"optional" json:"apmQuery" yaml:"apmQuery"`
	// audit_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#audit_query PowerpackV2#audit_query}
	AuditQuery *PowerpackV2WidgetTimeseriesDefinitionRequestAuditQuery `field:"optional" json:"auditQuery" yaml:"auditQuery"`
	// How to display the marker lines. Valid values are `area`, `bars`, `line`, `overlay`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#display_type PowerpackV2#display_type}
	DisplayType *string `field:"optional" json:"displayType" yaml:"displayType"`
	// event_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#event_query PowerpackV2#event_query}
	EventQuery *PowerpackV2WidgetTimeseriesDefinitionRequestEventQuery `field:"optional" json:"eventQuery" yaml:"eventQuery"`
	// formula block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#formula PowerpackV2#formula}
	Formula interface{} `field:"optional" json:"formula" yaml:"formula"`
	// log_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#log_query PowerpackV2#log_query}
	LogQuery *PowerpackV2WidgetTimeseriesDefinitionRequestLogQuery `field:"optional" json:"logQuery" yaml:"logQuery"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#metadata PowerpackV2#metadata}
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// network_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#network_query PowerpackV2#network_query}
	NetworkQuery *PowerpackV2WidgetTimeseriesDefinitionRequestNetworkQuery `field:"optional" json:"networkQuery" yaml:"networkQuery"`
	// A Boolean indicating whether the request uses the right or left Y-Axis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#on_right_yaxis PowerpackV2#on_right_yaxis}
	OnRightYaxis interface{} `field:"optional" json:"onRightYaxis" yaml:"onRightYaxis"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#process_query PowerpackV2#process_query}
	ProcessQuery *PowerpackV2WidgetTimeseriesDefinitionRequestProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// profile_metrics_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#profile_metrics_query PowerpackV2#profile_metrics_query}
	ProfileMetricsQuery *PowerpackV2WidgetTimeseriesDefinitionRequestProfileMetricsQuery `field:"optional" json:"profileMetricsQuery" yaml:"profileMetricsQuery"`
	// The metric query to use for this widget. **Deprecated.** Use queries and formulas instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#q PowerpackV2#q}
	Q *string `field:"optional" json:"q" yaml:"q"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#query PowerpackV2#query}
	Query interface{} `field:"optional" json:"query" yaml:"query"`
	// rum_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#rum_query PowerpackV2#rum_query}
	RumQuery *PowerpackV2WidgetTimeseriesDefinitionRequestRumQuery `field:"optional" json:"rumQuery" yaml:"rumQuery"`
	// security_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#security_query PowerpackV2#security_query}
	SecurityQuery *PowerpackV2WidgetTimeseriesDefinitionRequestSecurityQuery `field:"optional" json:"securityQuery" yaml:"securityQuery"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#style PowerpackV2#style}
	Style *PowerpackV2WidgetTimeseriesDefinitionRequestStyle `field:"optional" json:"style" yaml:"style"`
}

