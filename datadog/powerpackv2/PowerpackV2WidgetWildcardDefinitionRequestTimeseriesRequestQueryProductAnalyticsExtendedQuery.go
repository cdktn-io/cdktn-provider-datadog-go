// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQuery struct {
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#compute PowerpackV2#compute}
	Compute *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQueryCompute `field:"required" json:"compute" yaml:"compute"`
	// Data source for Product Analytics Extended queries. Valid values are `product_analytics_extended`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#data_source PowerpackV2#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// Name of the query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#name PowerpackV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#query PowerpackV2#query}
	Query *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQueryQuery `field:"required" json:"query" yaml:"query"`
	// audience_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#audience_filters PowerpackV2#audience_filters}
	AudienceFilters *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQueryAudienceFilters `field:"optional" json:"audienceFilters" yaml:"audienceFilters"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#group_by PowerpackV2#group_by}
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// Event indexes to query. Use `*` to query all indexes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#indexes PowerpackV2#indexes}
	Indexes *[]*string `field:"optional" json:"indexes" yaml:"indexes"`
}

