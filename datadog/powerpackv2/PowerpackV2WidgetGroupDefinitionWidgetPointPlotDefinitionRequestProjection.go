// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetPointPlotDefinitionRequestProjection struct {
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#dimension PowerpackV2#dimension}
	Dimension interface{} `field:"required" json:"dimension" yaml:"dimension"`
	// Type of the projection. Must be `point_plot`. Valid values are `point_plot`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#type PowerpackV2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Additional columns to include in the projection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#extra_columns PowerpackV2#extra_columns}
	ExtraColumns *[]*string `field:"optional" json:"extraColumns" yaml:"extraColumns"`
}

