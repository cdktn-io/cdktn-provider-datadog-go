// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryEventQueryGroupBy struct {
	// The event facet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/dashboard_v2#facet DashboardV2#facet}
	Facet *string `field:"required" json:"facet" yaml:"facet"`
	// The number of groups to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/dashboard_v2#limit DashboardV2#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/dashboard_v2#sort DashboardV2#sort}
	Sort *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryEventQueryGroupBySort `field:"optional" json:"sort" yaml:"sort"`
}

