// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinitionRequestQueryGroupBy struct {
	// Facet to group the user journey funnel by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#facet PowerpackV2#facet}
	Facet *string `field:"required" json:"facet" yaml:"facet"`
	// Maximum number of groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#limit PowerpackV2#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Whether to exclude missing values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#should_exclude_missing PowerpackV2#should_exclude_missing}
	ShouldExcludeMissing interface{} `field:"optional" json:"shouldExcludeMissing" yaml:"shouldExcludeMissing"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#sort PowerpackV2#sort}
	Sort *PowerpackV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinitionRequestQueryGroupBySort `field:"optional" json:"sort" yaml:"sort"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#target PowerpackV2#target}
	Target *PowerpackV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinitionRequestQueryGroupByTarget `field:"optional" json:"target" yaml:"target"`
}

