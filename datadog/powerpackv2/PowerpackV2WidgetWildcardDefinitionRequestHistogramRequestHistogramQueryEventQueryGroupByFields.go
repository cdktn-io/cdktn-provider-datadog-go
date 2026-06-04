// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryEventQueryGroupByFields struct {
	// List of event facets to group by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/powerpack_v2#fields PowerpackV2#fields}
	Fields *[]*string `field:"required" json:"fields" yaml:"fields"`
	// The number of groups to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/powerpack_v2#limit PowerpackV2#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/powerpack_v2#sort PowerpackV2#sort}
	Sort *PowerpackV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryEventQueryGroupByFieldsSort `field:"optional" json:"sort" yaml:"sort"`
}

