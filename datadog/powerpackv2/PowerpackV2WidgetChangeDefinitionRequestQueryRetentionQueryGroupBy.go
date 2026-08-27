// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetChangeDefinitionRequestQueryRetentionQueryGroupBy struct {
	// Facet to group by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#facet PowerpackV2#facet}
	Facet *string `field:"required" json:"facet" yaml:"facet"`
	// Target for the retention group by. Valid values are `cohort`, `return_period`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#target PowerpackV2#target}
	Target *string `field:"required" json:"target" yaml:"target"`
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
	Sort *PowerpackV2WidgetChangeDefinitionRequestQueryRetentionQueryGroupBySort `field:"optional" json:"sort" yaml:"sort"`
	// Source field for the retention group by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#source PowerpackV2#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
}

