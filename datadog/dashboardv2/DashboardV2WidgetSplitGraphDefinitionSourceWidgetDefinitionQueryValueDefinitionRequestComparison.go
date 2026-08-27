// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestComparison struct {
	// duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#duration DashboardV2#duration}
	Duration *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestComparisonDuration `field:"required" json:"duration" yaml:"duration"`
	// Which direction of change is considered an improvement. Valid values are `increase_better`, `decrease_better`, `neutral`. Defaults to `"neutral"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#directionality DashboardV2#directionality}
	Directionality *string `field:"optional" json:"directionality" yaml:"directionality"`
	// How the delta is expressed. Valid values are `absolute`, `relative`, `both`. Defaults to `"absolute"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#type DashboardV2#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

