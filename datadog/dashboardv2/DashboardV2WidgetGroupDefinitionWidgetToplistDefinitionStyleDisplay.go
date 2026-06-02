// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionStyleDisplay struct {
	// flat block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#flat DashboardV2#flat}
	Flat *DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionStyleDisplayFlat `field:"optional" json:"flat" yaml:"flat"`
	// stacked block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#stacked DashboardV2#stacked}
	Stacked *DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionStyleDisplayStacked `field:"optional" json:"stacked" yaml:"stacked"`
}

