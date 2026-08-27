// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinitionRequestQueryCompute struct {
	// Aggregation type for the user journey funnel computation. Valid values are `cardinality`, `count`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#aggregation DashboardV2#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// Metric for the user journey funnel computation. Valid values are `__dd.conversion`, `__dd.conversion_rate`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#metric DashboardV2#metric}
	Metric *string `field:"required" json:"metric" yaml:"metric"`
}

