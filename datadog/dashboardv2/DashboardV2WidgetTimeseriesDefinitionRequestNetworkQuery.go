// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetTimeseriesDefinitionRequestNetworkQuery struct {
	// A comma separated-list of index names. Use `*` to query all indexes at once. [Multiple Indexes](https://docs.datadoghq.com/logs/indexes/#multiple-indexes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#index DashboardV2#index}
	Index *string `field:"required" json:"index" yaml:"index"`
	// compute_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#compute_query DashboardV2#compute_query}
	ComputeQuery *DashboardV2WidgetTimeseriesDefinitionRequestNetworkQueryComputeQuery `field:"optional" json:"computeQuery" yaml:"computeQuery"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#group_by DashboardV2#group_by}
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// multi_compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#multi_compute DashboardV2#multi_compute}
	MultiCompute interface{} `field:"optional" json:"multiCompute" yaml:"multiCompute"`
	// The search query to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#search_query DashboardV2#search_query}
	SearchQuery *string `field:"optional" json:"searchQuery" yaml:"searchQuery"`
}

