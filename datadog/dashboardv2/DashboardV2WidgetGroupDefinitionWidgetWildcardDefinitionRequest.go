// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequest struct {
	// histogram_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#histogram_request DashboardV2#histogram_request}
	HistogramRequest *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestHistogramRequest `field:"optional" json:"histogramRequest" yaml:"histogramRequest"`
	// liststream_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#liststream_request DashboardV2#liststream_request}
	ListstreamRequest *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestListstreamRequest `field:"optional" json:"liststreamRequest" yaml:"liststreamRequest"`
	// timeseries_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#timeseries_request DashboardV2#timeseries_request}
	TimeseriesRequest *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequest `field:"optional" json:"timeseriesRequest" yaml:"timeseriesRequest"`
	// treemap_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#treemap_request DashboardV2#treemap_request}
	TreemapRequest *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequest `field:"optional" json:"treemapRequest" yaml:"treemapRequest"`
}

