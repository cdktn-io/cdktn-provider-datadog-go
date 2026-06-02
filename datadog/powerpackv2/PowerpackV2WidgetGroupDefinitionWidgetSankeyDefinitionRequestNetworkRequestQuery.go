// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestNetworkRequestQuery struct {
	// The data source for the Sankey network query. Valid values are `network_device_flows`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#data_source PowerpackV2#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// Fields to group by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#group_by PowerpackV2#group_by}
	GroupBy *[]*string `field:"required" json:"groupBy" yaml:"groupBy"`
	// Maximum number of results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#limit PowerpackV2#limit}
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	// The search query string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#query_string PowerpackV2#query_string}
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#compute PowerpackV2#compute}
	Compute *PowerpackV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestNetworkRequestQueryCompute `field:"optional" json:"compute" yaml:"compute"`
	// The mode for the Sankey network query. Valid values are `target`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#mode PowerpackV2#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Whether to exclude missing values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#should_exclude_missing PowerpackV2#should_exclude_missing}
	ShouldExcludeMissing interface{} `field:"optional" json:"shouldExcludeMissing" yaml:"shouldExcludeMissing"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#sort PowerpackV2#sort}
	Sort *PowerpackV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestNetworkRequestQuerySort `field:"optional" json:"sort" yaml:"sort"`
}

