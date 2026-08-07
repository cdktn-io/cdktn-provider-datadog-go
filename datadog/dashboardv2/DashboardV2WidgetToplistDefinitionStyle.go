// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetToplistDefinitionStyle struct {
	// display block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#display DashboardV2#display}
	Display *DashboardV2WidgetToplistDefinitionStyleDisplay `field:"optional" json:"display" yaml:"display"`
	// The color palette for the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#palette DashboardV2#palette}
	Palette *string `field:"optional" json:"palette" yaml:"palette"`
	// The scaling mode for the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#scaling DashboardV2#scaling}
	Scaling *string `field:"optional" json:"scaling" yaml:"scaling"`
}

