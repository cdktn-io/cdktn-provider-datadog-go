// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryEventQueryGroupBy struct {
	// The event facet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#facet PowerpackV2#facet}
	Facet *string `field:"required" json:"facet" yaml:"facet"`
	// The number of groups to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#limit PowerpackV2#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#sort PowerpackV2#sort}
	Sort *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryEventQueryGroupBySort `field:"optional" json:"sort" yaml:"sort"`
}

