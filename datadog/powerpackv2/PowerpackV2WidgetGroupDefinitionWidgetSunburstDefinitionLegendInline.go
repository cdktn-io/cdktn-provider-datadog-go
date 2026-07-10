// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionLegendInline struct {
	// The type of legend (inline or automatic).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#type PowerpackV2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Whether to hide the percentages of the groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#hide_percent PowerpackV2#hide_percent}
	HidePercent interface{} `field:"optional" json:"hidePercent" yaml:"hidePercent"`
	// Whether to hide the values of the groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#hide_value PowerpackV2#hide_value}
	HideValue interface{} `field:"optional" json:"hideValue" yaml:"hideValue"`
}

