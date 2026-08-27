// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetHeatmapDefinitionRequestHistogramRequest struct {
	// histogram_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#histogram_query DashboardV2#histogram_query}
	HistogramQuery *DashboardV2WidgetHeatmapDefinitionRequestHistogramRequestHistogramQuery `field:"required" json:"histogramQuery" yaml:"histogramQuery"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#style DashboardV2#style}
	Style *DashboardV2WidgetHeatmapDefinitionRequestHistogramRequestStyle `field:"optional" json:"style" yaml:"style"`
}

