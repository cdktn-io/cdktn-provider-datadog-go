// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetHeatmapDefinitionRequestHistogramRequestHistogramQuery struct {
	// metric_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#metric_query PowerpackV2#metric_query}
	MetricQuery *PowerpackV2WidgetHeatmapDefinitionRequestHistogramRequestHistogramQueryMetricQuery `field:"required" json:"metricQuery" yaml:"metricQuery"`
}

