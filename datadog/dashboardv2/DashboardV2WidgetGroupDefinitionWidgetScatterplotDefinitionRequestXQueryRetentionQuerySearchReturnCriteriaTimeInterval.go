// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchReturnCriteriaTimeInterval struct {
	// Type of return interval. Valid values are `fixed`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#type DashboardV2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Unit of the return interval. Valid values are `day`, `week`, `month`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#unit DashboardV2#unit}
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Value of the return interval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#value DashboardV2#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

