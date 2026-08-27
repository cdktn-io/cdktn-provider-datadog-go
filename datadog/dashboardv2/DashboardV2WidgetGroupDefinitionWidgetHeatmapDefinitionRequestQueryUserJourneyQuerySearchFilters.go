// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestQueryUserJourneyQuerySearchFilters struct {
	// audience_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#audience_filters DashboardV2#audience_filters}
	AudienceFilters *DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters `field:"optional" json:"audienceFilters" yaml:"audienceFilters"`
	// graph_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#graph_filter DashboardV2#graph_filter}
	GraphFilter interface{} `field:"optional" json:"graphFilter" yaml:"graphFilter"`
	// String filter for the user journey search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#string_filter DashboardV2#string_filter}
	StringFilter *string `field:"optional" json:"stringFilter" yaml:"stringFilter"`
}

