// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryUserJourneyQueryGroupBy struct {
	// Facet name to group by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#facet PowerpackV2#facet}
	Facet *string `field:"required" json:"facet" yaml:"facet"`
	// Maximum number of groups to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#limit PowerpackV2#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Whether to exclude events missing the group-by facet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#should_exclude_missing PowerpackV2#should_exclude_missing}
	ShouldExcludeMissing interface{} `field:"optional" json:"shouldExcludeMissing" yaml:"shouldExcludeMissing"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#sort PowerpackV2#sort}
	Sort *PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryUserJourneyQueryGroupBySort `field:"optional" json:"sort" yaml:"sort"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#target PowerpackV2#target}
	Target *PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryUserJourneyQueryGroupByTarget `field:"optional" json:"target" yaml:"target"`
}

