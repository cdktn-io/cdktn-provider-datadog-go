// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetPointPlotDefinitionRequest struct {
	// projection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#projection PowerpackV2#projection}
	Projection *PowerpackV2WidgetGroupDefinitionWidgetPointPlotDefinitionRequestProjection `field:"required" json:"projection" yaml:"projection"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#query PowerpackV2#query}
	Query *PowerpackV2WidgetGroupDefinitionWidgetPointPlotDefinitionRequestQuery `field:"required" json:"query" yaml:"query"`
	// The type of data request. Must be `data_projection`. Valid values are `data_projection`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#request_type PowerpackV2#request_type}
	RequestType *string `field:"required" json:"requestType" yaml:"requestType"`
	// Maximum number of data points to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#limit PowerpackV2#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
}

