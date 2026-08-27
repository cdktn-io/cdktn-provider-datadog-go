// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinitionRequest struct {
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#query PowerpackV2#query}
	Query *PowerpackV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinitionRequestQuery `field:"required" json:"query" yaml:"query"`
	// Request type for the Product Analytics funnel. Valid values are `user_journey_funnel`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#request_type PowerpackV2#request_type}
	RequestType *string `field:"required" json:"requestType" yaml:"requestType"`
	// Segments to compare in the funnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#comparison_segments PowerpackV2#comparison_segments}
	ComparisonSegments *[]*string `field:"optional" json:"comparisonSegments" yaml:"comparisonSegments"`
	// comparison_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#comparison_time PowerpackV2#comparison_time}
	ComparisonTime *PowerpackV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinitionRequestComparisonTime `field:"optional" json:"comparisonTime" yaml:"comparisonTime"`
}

