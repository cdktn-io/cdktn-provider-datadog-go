// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionEvent struct {
	// The event query to use in the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#q DashboardV2#q}
	Q *string `field:"required" json:"q" yaml:"q"`
	// The execution method for multi-value filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#tags_execution DashboardV2#tags_execution}
	TagsExecution *string `field:"optional" json:"tagsExecution" yaml:"tagsExecution"`
}

