// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinitionRequestQueryGroupBy struct {
	// Facet to group the user journey funnel by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#facet DashboardV2#facet}
	Facet *string `field:"required" json:"facet" yaml:"facet"`
	// Maximum number of groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#limit DashboardV2#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Whether to exclude missing values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#should_exclude_missing DashboardV2#should_exclude_missing}
	ShouldExcludeMissing interface{} `field:"optional" json:"shouldExcludeMissing" yaml:"shouldExcludeMissing"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#sort DashboardV2#sort}
	Sort *DashboardV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinitionRequestQueryGroupBySort `field:"optional" json:"sort" yaml:"sort"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#target DashboardV2#target}
	Target *DashboardV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinitionRequestQueryGroupByTarget `field:"optional" json:"target" yaml:"target"`
}

