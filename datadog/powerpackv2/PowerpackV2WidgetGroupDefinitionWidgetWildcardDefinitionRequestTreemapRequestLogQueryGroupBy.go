// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestLogQueryGroupBy struct {
	// The facet name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#facet PowerpackV2#facet}
	Facet *string `field:"optional" json:"facet" yaml:"facet"`
	// The maximum number of items in the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#limit PowerpackV2#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// sort_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#sort_query PowerpackV2#sort_query}
	SortQuery *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestLogQueryGroupBySortQuery `field:"optional" json:"sortQuery" yaml:"sortQuery"`
}

