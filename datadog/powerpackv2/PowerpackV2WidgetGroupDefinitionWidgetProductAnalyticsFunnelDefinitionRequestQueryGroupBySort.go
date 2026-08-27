// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinitionRequestQueryGroupBySort struct {
	// Aggregation used to sort the funnel groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#aggregation PowerpackV2#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// Metric to sort by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#metric PowerpackV2#metric}
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
	// Sort order for funnel groups. Valid values are `asc`, `desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#order PowerpackV2#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}

