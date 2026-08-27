// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetRetentionCurveDefinitionRequest struct {
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#query PowerpackV2#query}
	Query *PowerpackV2WidgetRetentionCurveDefinitionRequestQuery `field:"required" json:"query" yaml:"query"`
	// Request type for the retention curve widget. Valid values are `retention_curve`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#request_type PowerpackV2#request_type}
	RequestType *string `field:"required" json:"requestType" yaml:"requestType"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#style PowerpackV2#style}
	Style *PowerpackV2WidgetRetentionCurveDefinitionRequestStyle `field:"optional" json:"style" yaml:"style"`
}

