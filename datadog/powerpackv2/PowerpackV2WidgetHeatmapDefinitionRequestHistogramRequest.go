// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetHeatmapDefinitionRequestHistogramRequest struct {
	// histogram_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#histogram_query PowerpackV2#histogram_query}
	HistogramQuery *PowerpackV2WidgetHeatmapDefinitionRequestHistogramRequestHistogramQuery `field:"required" json:"histogramQuery" yaml:"histogramQuery"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#style PowerpackV2#style}
	Style *PowerpackV2WidgetHeatmapDefinitionRequestHistogramRequestStyle `field:"optional" json:"style" yaml:"style"`
}

