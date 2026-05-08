// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionStyle struct {
	// The color palette to apply to the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/powerpack_v2#palette PowerpackV2#palette}
	Palette *string `field:"required" json:"palette" yaml:"palette"`
	// A Boolean indicating whether to flip the palette tones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/powerpack_v2#palette_flip PowerpackV2#palette_flip}
	PaletteFlip interface{} `field:"required" json:"paletteFlip" yaml:"paletteFlip"`
}

