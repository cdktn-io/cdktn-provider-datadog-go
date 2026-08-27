// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQueryCompute struct {
	// Aggregation method for the User Journey query.
	//
	// Valid values are `count`, `cardinality`, `median`, `pc75`, `pc90`, `pc95`, `pc98`, `pc99`, `sum`, `min`, `max`, `avg`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#aggregation PowerpackV2#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// Time bucket interval in milliseconds for timeseries queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#interval PowerpackV2#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Metric for the User Journey computation. Valid values are `__dd.conversion`, `__dd.conversion_rate`, `__dd.time_to_convert`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#metric PowerpackV2#metric}
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#target PowerpackV2#target}
	Target *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQueryComputeTarget `field:"optional" json:"target" yaml:"target"`
}

