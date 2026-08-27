// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetHostmapDefinitionRequestSizeQueryProductAnalyticsExtendedQueryCompute struct {
	// Aggregation method for the Product Analytics Extended query.
	//
	// Valid values are `count`, `cardinality`, `median`, `pc75`, `pc90`, `pc95`, `pc98`, `pc99`, `sum`, `min`, `max`, `avg`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#aggregation DashboardV2#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// Fixed-width time bucket interval in milliseconds. Mutually exclusive with `rollup`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#interval DashboardV2#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Measurable attribute to compute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#metric DashboardV2#metric}
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
	// Name of the compute for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#name DashboardV2#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// rollup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#rollup DashboardV2#rollup}
	Rollup *DashboardV2WidgetHostmapDefinitionRequestSizeQueryProductAnalyticsExtendedQueryComputeRollup `field:"optional" json:"rollup" yaml:"rollup"`
}

