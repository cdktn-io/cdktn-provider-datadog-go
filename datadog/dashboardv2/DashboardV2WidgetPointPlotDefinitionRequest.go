// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetPointPlotDefinitionRequest struct {
	// projection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#projection DashboardV2#projection}
	Projection *DashboardV2WidgetPointPlotDefinitionRequestProjection `field:"required" json:"projection" yaml:"projection"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#query DashboardV2#query}
	Query *DashboardV2WidgetPointPlotDefinitionRequestQuery `field:"required" json:"query" yaml:"query"`
	// The type of data request. Must be `data_projection`. Valid values are `data_projection`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#request_type DashboardV2#request_type}
	RequestType *string `field:"required" json:"requestType" yaml:"requestType"`
	// Maximum number of data points to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#limit DashboardV2#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
}

