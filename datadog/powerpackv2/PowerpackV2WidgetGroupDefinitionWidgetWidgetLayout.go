// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetWidgetLayout struct {
	// The height of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#height PowerpackV2#height}
	Height *float64 `field:"required" json:"height" yaml:"height"`
	// The width of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#width PowerpackV2#width}
	Width *float64 `field:"required" json:"width" yaml:"width"`
	// The position of the widget on the x (horizontal) axis. Must be greater than or equal to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#x PowerpackV2#x}
	X *float64 `field:"required" json:"x" yaml:"x"`
	// The position of the widget on the y (vertical) axis. Must be greater than or equal to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#y PowerpackV2#y}
	Y *float64 `field:"required" json:"y" yaml:"y"`
	// Whether the widget should be the first one on the second column in high density or not.
	//
	// Only one widget in the dashboard should have this property set to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#is_column_break PowerpackV2#is_column_break}
	IsColumnBreak interface{} `field:"optional" json:"isColumnBreak" yaml:"isColumnBreak"`
}

