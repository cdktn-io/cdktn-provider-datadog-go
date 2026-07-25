// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetPointPlotDefinitionLegend struct {
	// Type of legend to show for the point plot widget. Valid values are `automatic`, `none`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#type PowerpackV2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

