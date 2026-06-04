// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestListstreamRequestQuery struct {
	// Source from which to query items to display in the stream.
	//
	// Valid values are `logs_stream`, `audit_stream`, `ci_pipeline_stream`, `ci_test_stream`, `rum_issue_stream`, `apm_issue_stream`, `trace_stream`, `logs_issue_stream`, `logs_pattern_stream`, `logs_transaction_stream`, `event_stream`, `rum_stream`, `llm_observability_stream`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/powerpack_v2#data_source PowerpackV2#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// Specifies the field for logs pattern clustering. Can only be used with `logs_pattern_stream`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/powerpack_v2#clustering_pattern_field_path PowerpackV2#clustering_pattern_field_path}
	ClusteringPatternFieldPath *string `field:"optional" json:"clusteringPatternFieldPath" yaml:"clusteringPatternFieldPath"`
	// Size of events displayed in widget. Required if `data_source` is `event_stream`. Valid values are `s`, `l`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/powerpack_v2#event_size PowerpackV2#event_size}
	EventSize *string `field:"optional" json:"eventSize" yaml:"eventSize"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/powerpack_v2#group_by PowerpackV2#group_by}
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// List of indexes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/powerpack_v2#indexes PowerpackV2#indexes}
	Indexes *[]*string `field:"optional" json:"indexes" yaml:"indexes"`
	// Widget query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/powerpack_v2#query_string PowerpackV2#query_string}
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/powerpack_v2#sort PowerpackV2#sort}
	Sort *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestListstreamRequestQuerySort `field:"optional" json:"sort" yaml:"sort"`
	// Storage location (private beta).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/powerpack_v2#storage PowerpackV2#storage}
	Storage *string `field:"optional" json:"storage" yaml:"storage"`
}

