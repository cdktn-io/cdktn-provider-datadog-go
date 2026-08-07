// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequest struct {
	// formula block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#formula DashboardV2#formula}
	Formula interface{} `field:"optional" json:"formula" yaml:"formula"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#query DashboardV2#query}
	Query interface{} `field:"optional" json:"query" yaml:"query"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#sort DashboardV2#sort}
	Sort *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestSort `field:"optional" json:"sort" yaml:"sort"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#style DashboardV2#style}
	Style *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestStyle `field:"optional" json:"style" yaml:"style"`
}

