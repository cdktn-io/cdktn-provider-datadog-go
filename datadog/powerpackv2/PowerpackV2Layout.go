// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2Layout struct {
	// The height of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#height PowerpackV2#height}
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// The width of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#width PowerpackV2#width}
	Width *float64 `field:"optional" json:"width" yaml:"width"`
	// The position of the widget on the x (horizontal) axis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#x PowerpackV2#x}
	X *float64 `field:"optional" json:"x" yaml:"x"`
	// The position of the widget on the y (vertical) axis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#y PowerpackV2#y}
	Y *float64 `field:"optional" json:"y" yaml:"y"`
}

