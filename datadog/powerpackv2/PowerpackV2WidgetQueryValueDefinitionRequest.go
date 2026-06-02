// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetQueryValueDefinitionRequest struct {
	// The aggregator to use for time aggregation. Valid values are `avg`, `min`, `max`, `sum`, `last`, `area`, `l2norm`, `percentile`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#aggregator PowerpackV2#aggregator}
	Aggregator *string `field:"optional" json:"aggregator" yaml:"aggregator"`
	// apm_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#apm_query PowerpackV2#apm_query}
	ApmQuery *PowerpackV2WidgetQueryValueDefinitionRequestApmQuery `field:"optional" json:"apmQuery" yaml:"apmQuery"`
	// audit_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#audit_query PowerpackV2#audit_query}
	AuditQuery *PowerpackV2WidgetQueryValueDefinitionRequestAuditQuery `field:"optional" json:"auditQuery" yaml:"auditQuery"`
	// conditional_formats block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#conditional_formats PowerpackV2#conditional_formats}
	ConditionalFormats interface{} `field:"optional" json:"conditionalFormats" yaml:"conditionalFormats"`
	// formula block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#formula PowerpackV2#formula}
	Formula interface{} `field:"optional" json:"formula" yaml:"formula"`
	// log_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#log_query PowerpackV2#log_query}
	LogQuery *PowerpackV2WidgetQueryValueDefinitionRequestLogQuery `field:"optional" json:"logQuery" yaml:"logQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#process_query PowerpackV2#process_query}
	ProcessQuery *PowerpackV2WidgetQueryValueDefinitionRequestProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
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
	RumQuery *PowerpackV2WidgetQueryValueDefinitionRequestRumQuery `field:"optional" json:"rumQuery" yaml:"rumQuery"`
	// security_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#security_query PowerpackV2#security_query}
	SecurityQuery *PowerpackV2WidgetQueryValueDefinitionRequestSecurityQuery `field:"optional" json:"securityQuery" yaml:"securityQuery"`
}

