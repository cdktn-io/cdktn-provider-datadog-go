// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetHostmapDefinitionStyle struct {
	// The max value to use to color the map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/powerpack_v2#fill_max PowerpackV2#fill_max}
	FillMax *string `field:"optional" json:"fillMax" yaml:"fillMax"`
	// The min value to use to color the map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/powerpack_v2#fill_min PowerpackV2#fill_min}
	FillMin *string `field:"optional" json:"fillMin" yaml:"fillMin"`
	// A color palette to apply to the widget. The available options are available at: https://docs.datadoghq.com/dashboards/widgets/timeseries/#appearance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/powerpack_v2#palette PowerpackV2#palette}
	Palette *string `field:"optional" json:"palette" yaml:"palette"`
	// A Boolean indicating whether to flip the palette tones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/powerpack_v2#palette_flip PowerpackV2#palette_flip}
	PaletteFlip interface{} `field:"optional" json:"paletteFlip" yaml:"paletteFlip"`
}

