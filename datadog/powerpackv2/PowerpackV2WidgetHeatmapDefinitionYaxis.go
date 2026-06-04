// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetHeatmapDefinitionYaxis struct {
	// Set to `true` to include zero.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/powerpack_v2#include_zero PowerpackV2#include_zero}
	IncludeZero interface{} `field:"optional" json:"includeZero" yaml:"includeZero"`
	// The label of the axis to display on the graph. Only usable on Scatterplot Widgets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/powerpack_v2#label PowerpackV2#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Specifies the maximum numeric value to show on the axis. Defaults to `auto`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/powerpack_v2#max PowerpackV2#max}
	Max *string `field:"optional" json:"max" yaml:"max"`
	// Specifies the minimum numeric value to show on the axis. Defaults to `auto`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/powerpack_v2#min PowerpackV2#min}
	Min *string `field:"optional" json:"min" yaml:"min"`
	// Specifies the scale type. Possible values are `linear`, `log`, `sqrt`, and `pow##` (for example `pow2` or `pow0.5`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/powerpack_v2#scale PowerpackV2#scale}
	Scale *string `field:"optional" json:"scale" yaml:"scale"`
}

