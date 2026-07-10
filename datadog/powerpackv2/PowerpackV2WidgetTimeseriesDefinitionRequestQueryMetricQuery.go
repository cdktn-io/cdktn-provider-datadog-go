// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetTimeseriesDefinitionRequestQueryMetricQuery struct {
	// The name of the query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#name PowerpackV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The metrics query definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#query PowerpackV2#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// The aggregation methods available for metrics queries. Valid values are `avg`, `min`, `max`, `sum`, `last`, `area`, `l2norm`, `percentile`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#aggregator PowerpackV2#aggregator}
	Aggregator *string `field:"optional" json:"aggregator" yaml:"aggregator"`
	// The source organization UUID for cross organization queries. Feature in Private Beta.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#cross_org_uuids PowerpackV2#cross_org_uuids}
	CrossOrgUuids *[]*string `field:"optional" json:"crossOrgUuids" yaml:"crossOrgUuids"`
	// The data source for metrics queries. Defaults to `"metrics"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#data_source PowerpackV2#data_source}
	DataSource *string `field:"optional" json:"dataSource" yaml:"dataSource"`
	// Semantic mode for metrics queries.
	//
	// This determines how metrics from different sources are combined or displayed. Valid values are `combined`, `native`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#semantic_mode PowerpackV2#semantic_mode}
	SemanticMode *string `field:"optional" json:"semanticMode" yaml:"semanticMode"`
}

