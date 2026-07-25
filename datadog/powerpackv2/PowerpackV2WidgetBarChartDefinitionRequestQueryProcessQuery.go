// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetBarChartDefinitionRequestQueryProcessQuery struct {
	// The data source for process queries. Valid values are `process`, `container`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#data_source PowerpackV2#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// The process metric name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#metric PowerpackV2#metric}
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// The name of query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#name PowerpackV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The aggregation methods available for metrics queries. Valid values are `avg`, `min`, `max`, `sum`, `last`, `area`, `l2norm`, `percentile`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#aggregator PowerpackV2#aggregator}
	Aggregator *string `field:"optional" json:"aggregator" yaml:"aggregator"`
	// The source organization UUID for cross organization queries. Feature in Private Beta.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#cross_org_uuids PowerpackV2#cross_org_uuids}
	CrossOrgUuids *[]*string `field:"optional" json:"crossOrgUuids" yaml:"crossOrgUuids"`
	// Whether to normalize the CPU percentages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#is_normalized_cpu PowerpackV2#is_normalized_cpu}
	IsNormalizedCpu interface{} `field:"optional" json:"isNormalizedCpu" yaml:"isNormalizedCpu"`
	// The number of hits to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#limit PowerpackV2#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// The direction of the sort. Valid values are `asc`, `desc`. Defaults to `"desc"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#sort PowerpackV2#sort}
	Sort *string `field:"optional" json:"sort" yaml:"sort"`
	// An array of tags to filter by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#tag_filters PowerpackV2#tag_filters}
	TagFilters *[]*string `field:"optional" json:"tagFilters" yaml:"tagFilters"`
	// The text to use as a filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#text_filter PowerpackV2#text_filter}
	TextFilter *string `field:"optional" json:"textFilter" yaml:"textFilter"`
}

