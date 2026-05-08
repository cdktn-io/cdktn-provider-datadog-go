// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestFormula struct {
	// A string expression built from queries, formulas, and functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#formula_expression DashboardV2#formula_expression}
	FormulaExpression *string `field:"required" json:"formulaExpression" yaml:"formulaExpression"`
	// An expression alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#alias DashboardV2#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// A list of display modes for each table cell. Valid values are `number`, `bar`, and `trend`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#cell_display_mode DashboardV2#cell_display_mode}
	CellDisplayMode *string `field:"optional" json:"cellDisplayMode" yaml:"cellDisplayMode"`
	// cell_display_mode_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#cell_display_mode_options DashboardV2#cell_display_mode_options}
	CellDisplayModeOptions *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestFormulaCellDisplayModeOptions `field:"optional" json:"cellDisplayModeOptions" yaml:"cellDisplayModeOptions"`
	// conditional_formats block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#conditional_formats DashboardV2#conditional_formats}
	ConditionalFormats interface{} `field:"optional" json:"conditionalFormats" yaml:"conditionalFormats"`
	// limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#limit DashboardV2#limit}
	Limit *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestFormulaLimit `field:"optional" json:"limit" yaml:"limit"`
	// number_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#number_format DashboardV2#number_format}
	NumberFormat *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestFormulaNumberFormat `field:"optional" json:"numberFormat" yaml:"numberFormat"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#style DashboardV2#style}
	Style *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestFormulaStyle `field:"optional" json:"style" yaml:"style"`
}

