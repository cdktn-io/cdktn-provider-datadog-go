// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetToplistDefinitionStyle struct {
	// display block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/powerpack_v2#display PowerpackV2#display}
	Display *PowerpackV2WidgetToplistDefinitionStyleDisplay `field:"optional" json:"display" yaml:"display"`
	// The color palette for the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/powerpack_v2#palette PowerpackV2#palette}
	Palette *string `field:"optional" json:"palette" yaml:"palette"`
	// The scaling mode for the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/powerpack_v2#scaling PowerpackV2#scaling}
	Scaling *string `field:"optional" json:"scaling" yaml:"scaling"`
}

