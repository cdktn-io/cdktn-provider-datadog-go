// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryCompute struct {
	// Aggregation method for the Product Analytics Extended query.
	//
	// Valid values are `count`, `cardinality`, `median`, `pc75`, `pc90`, `pc95`, `pc98`, `pc99`, `sum`, `min`, `max`, `avg`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#aggregation PowerpackV2#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// Fixed-width time bucket interval in milliseconds. Mutually exclusive with `rollup`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#interval PowerpackV2#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Measurable attribute to compute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#metric PowerpackV2#metric}
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
	// Name of the compute for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#name PowerpackV2#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// rollup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#rollup PowerpackV2#rollup}
	Rollup *PowerpackV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryComputeRollup `field:"optional" json:"rollup" yaml:"rollup"`
}

