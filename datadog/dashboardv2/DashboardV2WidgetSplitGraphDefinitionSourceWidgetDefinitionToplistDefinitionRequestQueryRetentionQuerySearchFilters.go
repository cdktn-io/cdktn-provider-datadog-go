// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryRetentionQuerySearchFilters struct {
	// audience_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#audience_filters DashboardV2#audience_filters}
	AudienceFilters *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryRetentionQuerySearchFiltersAudienceFilters `field:"optional" json:"audienceFilters" yaml:"audienceFilters"`
	// String filter for the retention query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#string_filter DashboardV2#string_filter}
	StringFilter *string `field:"optional" json:"stringFilter" yaml:"stringFilter"`
}

