// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinitionRequestQuery struct {
	// Data source for Product Analytics funnel queries. Valid values are `product_analytics_journey`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#data_source PowerpackV2#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// search block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#search PowerpackV2#search}
	Search *PowerpackV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinitionRequestQuerySearch `field:"required" json:"search" yaml:"search"`
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#compute PowerpackV2#compute}
	Compute *PowerpackV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinitionRequestQueryCompute `field:"optional" json:"compute" yaml:"compute"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#group_by PowerpackV2#group_by}
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// Subquery identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#subquery_id PowerpackV2#subquery_id}
	SubqueryId *string `field:"optional" json:"subqueryId" yaml:"subqueryId"`
}

