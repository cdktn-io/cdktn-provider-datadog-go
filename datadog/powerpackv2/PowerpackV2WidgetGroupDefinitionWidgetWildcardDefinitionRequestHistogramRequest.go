// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestHistogramRequest struct {
	// histogram_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#histogram_query PowerpackV2#histogram_query}
	HistogramQuery *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestHistogramRequestHistogramQuery `field:"required" json:"histogramQuery" yaml:"histogramQuery"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#style PowerpackV2#style}
	Style *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestHistogramRequestStyle `field:"optional" json:"style" yaml:"style"`
}

