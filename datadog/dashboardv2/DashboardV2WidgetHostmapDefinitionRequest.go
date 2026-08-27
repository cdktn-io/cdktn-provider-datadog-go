// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetHostmapDefinitionRequest struct {
	// fill block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#fill DashboardV2#fill}
	Fill *DashboardV2WidgetHostmapDefinitionRequestFill `field:"optional" json:"fill" yaml:"fill"`
	// size block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#size DashboardV2#size}
	Size *DashboardV2WidgetHostmapDefinitionRequestSize `field:"optional" json:"size" yaml:"size"`
}

