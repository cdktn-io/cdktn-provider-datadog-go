// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionTime struct {
	// fixed block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/dashboard_v2#fixed DashboardV2#fixed}
	Fixed *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionTimeFixed `field:"optional" json:"fixed" yaml:"fixed"`
	// live block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/dashboard_v2#live DashboardV2#live}
	Live *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionTimeLive `field:"optional" json:"live" yaml:"live"`
}

