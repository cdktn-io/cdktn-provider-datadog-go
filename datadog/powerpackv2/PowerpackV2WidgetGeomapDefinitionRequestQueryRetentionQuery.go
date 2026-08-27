// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGeomapDefinitionRequestQueryRetentionQuery struct {
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#compute PowerpackV2#compute}
	Compute *PowerpackV2WidgetGeomapDefinitionRequestQueryRetentionQueryCompute `field:"required" json:"compute" yaml:"compute"`
	// Data source for retention queries. Valid values are `product_analytics_retention`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#data_source PowerpackV2#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// Name of the query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#name PowerpackV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// search block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#search PowerpackV2#search}
	Search *PowerpackV2WidgetGeomapDefinitionRequestQueryRetentionQuerySearch `field:"required" json:"search" yaml:"search"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#group_by PowerpackV2#group_by}
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
}

