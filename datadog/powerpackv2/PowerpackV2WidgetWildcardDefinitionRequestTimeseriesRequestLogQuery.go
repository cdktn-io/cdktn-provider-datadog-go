// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestLogQuery struct {
	// A comma separated-list of index names. Use `*` to query all indexes at once. [Multiple Indexes](https://docs.datadoghq.com/logs/indexes/#multiple-indexes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#index PowerpackV2#index}
	Index *string `field:"required" json:"index" yaml:"index"`
	// compute_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#compute_query PowerpackV2#compute_query}
	ComputeQuery *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestLogQueryComputeQuery `field:"optional" json:"computeQuery" yaml:"computeQuery"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#group_by PowerpackV2#group_by}
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// multi_compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#multi_compute PowerpackV2#multi_compute}
	MultiCompute interface{} `field:"optional" json:"multiCompute" yaml:"multiCompute"`
	// The search query to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#search_query PowerpackV2#search_query}
	SearchQuery *string `field:"optional" json:"searchQuery" yaml:"searchQuery"`
}

