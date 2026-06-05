// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetTimeseriesDefinitionRequestFormulaStyle struct {
	// The color palette used to display the formula.
	//
	// A guide to the available color palettes can be found at https://docs.datadoghq.com/dashboards/guide/widget_colors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/dashboard_v2#palette DashboardV2#palette}
	Palette *string `field:"optional" json:"palette" yaml:"palette"`
	// Index specifying which color to use within the palette.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/dashboard_v2#palette_index DashboardV2#palette_index}
	PaletteIndex *float64 `field:"optional" json:"paletteIndex" yaml:"paletteIndex"`
}

