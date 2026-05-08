// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetQueryTableDefinitionRequestTextFormatsTextFormat struct {
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/powerpack_v2#match PowerpackV2#match}
	Match *PowerpackV2WidgetQueryTableDefinitionRequestTextFormatsTextFormatMatch `field:"required" json:"match" yaml:"match"`
	// The custom color palette to apply to the background.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/powerpack_v2#custom_bg_color PowerpackV2#custom_bg_color}
	CustomBgColor *string `field:"optional" json:"customBgColor" yaml:"customBgColor"`
	// The custom color palette to apply to the foreground text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/powerpack_v2#custom_fg_color PowerpackV2#custom_fg_color}
	CustomFgColor *string `field:"optional" json:"customFgColor" yaml:"customFgColor"`
	// The color palette to apply. Valid values are `white_on_red`, `white_on_yellow`, `white_on_green`, `black_on_light_red`, `black_on_light_yellow`, `black_on_light_green`, `red_on_white`, `yellow_on_white`, `green_on_white`, `custom_bg`, `custom_text`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/powerpack_v2#palette PowerpackV2#palette}
	Palette *string `field:"optional" json:"palette" yaml:"palette"`
	// replace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/powerpack_v2#replace PowerpackV2#replace}
	Replace *PowerpackV2WidgetQueryTableDefinitionRequestTextFormatsTextFormatReplace `field:"optional" json:"replace" yaml:"replace"`
}

