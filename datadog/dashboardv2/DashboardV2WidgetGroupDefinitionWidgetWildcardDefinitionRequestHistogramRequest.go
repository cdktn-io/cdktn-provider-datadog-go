// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestHistogramRequest struct {
	// histogram_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/dashboard_v2#histogram_query DashboardV2#histogram_query}
	HistogramQuery *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestHistogramRequestHistogramQuery `field:"required" json:"histogramQuery" yaml:"histogramQuery"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/dashboard_v2#style DashboardV2#style}
	Style *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestHistogramRequestStyle `field:"optional" json:"style" yaml:"style"`
}

