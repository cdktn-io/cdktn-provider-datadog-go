// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSunburstDefinitionRequestQueryEventQueryCompute struct {
	// The aggregation methods for event platform queries.
	//
	// Valid values are `count`, `cardinality`, `median`, `pc75`, `pc90`, `pc95`, `pc98`, `pc99`, `sum`, `min`, `max`, `avg`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#aggregation DashboardV2#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// A time interval in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#interval DashboardV2#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The measurable attribute to compute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#metric DashboardV2#metric}
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
}

