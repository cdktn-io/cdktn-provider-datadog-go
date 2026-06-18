// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetWildcardDefinitionRequest struct {
	// histogram_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/dashboard_v2#histogram_request DashboardV2#histogram_request}
	HistogramRequest *DashboardV2WidgetWildcardDefinitionRequestHistogramRequest `field:"optional" json:"histogramRequest" yaml:"histogramRequest"`
	// liststream_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/dashboard_v2#liststream_request DashboardV2#liststream_request}
	ListstreamRequest *DashboardV2WidgetWildcardDefinitionRequestListstreamRequest `field:"optional" json:"liststreamRequest" yaml:"liststreamRequest"`
	// timeseries_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/dashboard_v2#timeseries_request DashboardV2#timeseries_request}
	TimeseriesRequest *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequest `field:"optional" json:"timeseriesRequest" yaml:"timeseriesRequest"`
	// treemap_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/dashboard_v2#treemap_request DashboardV2#treemap_request}
	TreemapRequest *DashboardV2WidgetWildcardDefinitionRequestTreemapRequest `field:"optional" json:"treemapRequest" yaml:"treemapRequest"`
}

