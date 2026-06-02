// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetWildcardDefinitionRequest struct {
	// histogram_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#histogram_request PowerpackV2#histogram_request}
	HistogramRequest *PowerpackV2WidgetWildcardDefinitionRequestHistogramRequest `field:"optional" json:"histogramRequest" yaml:"histogramRequest"`
	// liststream_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#liststream_request PowerpackV2#liststream_request}
	ListstreamRequest *PowerpackV2WidgetWildcardDefinitionRequestListstreamRequest `field:"optional" json:"liststreamRequest" yaml:"liststreamRequest"`
	// timeseries_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#timeseries_request PowerpackV2#timeseries_request}
	TimeseriesRequest *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequest `field:"optional" json:"timeseriesRequest" yaml:"timeseriesRequest"`
	// treemap_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#treemap_request PowerpackV2#treemap_request}
	TreemapRequest *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequest `field:"optional" json:"treemapRequest" yaml:"treemapRequest"`
}

