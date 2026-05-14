// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryTableDefinitionTimeLive struct {
	// Unit of the time span. Valid values are `minute`, `hour`, `day`, `week`, `month`, `year`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#unit DashboardV2#unit}
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Value of the time span.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#value DashboardV2#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

