// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryRetentionQuerySearchFiltersAudienceFilters struct {
	// account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#account PowerpackV2#account}
	Account interface{} `field:"optional" json:"account" yaml:"account"`
	// An optional filter condition applied to the audience subquery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#filter_condition PowerpackV2#filter_condition}
	FilterCondition *string `field:"optional" json:"filterCondition" yaml:"filterCondition"`
	// segment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#segment PowerpackV2#segment}
	Segment interface{} `field:"optional" json:"segment" yaml:"segment"`
	// user block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#user PowerpackV2#user}
	User interface{} `field:"optional" json:"user" yaml:"user"`
}

