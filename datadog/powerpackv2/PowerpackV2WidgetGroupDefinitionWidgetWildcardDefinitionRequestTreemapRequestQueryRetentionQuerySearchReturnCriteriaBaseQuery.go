// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaBaseQuery struct {
	// Data source for the Product Analytics event query. Valid values are `product_analytics`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#data_source PowerpackV2#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// search block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#search PowerpackV2#search}
	Search *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaBaseQuerySearch `field:"required" json:"search" yaml:"search"`
}

