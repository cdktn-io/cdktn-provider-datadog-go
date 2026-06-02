// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSplitConfig struct {
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#sort DashboardV2#sort}
	Sort *DashboardV2WidgetSplitGraphDefinitionSplitConfigSort `field:"required" json:"sort" yaml:"sort"`
	// split_dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#split_dimensions DashboardV2#split_dimensions}
	SplitDimensions interface{} `field:"required" json:"splitDimensions" yaml:"splitDimensions"`
	// Maximum number of graphs to display in the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#limit DashboardV2#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// static_splits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#static_splits DashboardV2#static_splits}
	StaticSplits interface{} `field:"optional" json:"staticSplits" yaml:"staticSplits"`
}

