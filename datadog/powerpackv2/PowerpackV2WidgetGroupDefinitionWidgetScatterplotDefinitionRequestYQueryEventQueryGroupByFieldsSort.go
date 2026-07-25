// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryGroupByFieldsSort struct {
	// The aggregation methods for the event platform queries.
	//
	// Valid values are `count`, `cardinality`, `median`, `pc75`, `pc90`, `pc95`, `pc98`, `pc99`, `sum`, `min`, `max`, `avg`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#aggregation PowerpackV2#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// The metric used for sorting group by results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#metric PowerpackV2#metric}
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
	// Direction of sort. Valid values are `asc`, `desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#order PowerpackV2#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}

