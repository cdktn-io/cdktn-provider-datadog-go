// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormula struct {
	// A string expression built from queries, formulas, and functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#formula_expression PowerpackV2#formula_expression}
	FormulaExpression *string `field:"required" json:"formulaExpression" yaml:"formulaExpression"`
	// An expression alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#alias PowerpackV2#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// A list of display modes for each table cell. Valid values are `number`, `bar`, and `trend`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#cell_display_mode PowerpackV2#cell_display_mode}
	CellDisplayMode *string `field:"optional" json:"cellDisplayMode" yaml:"cellDisplayMode"`
	// cell_display_mode_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#cell_display_mode_options PowerpackV2#cell_display_mode_options}
	CellDisplayModeOptions *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaCellDisplayModeOptions `field:"optional" json:"cellDisplayModeOptions" yaml:"cellDisplayModeOptions"`
	// conditional_formats block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#conditional_formats PowerpackV2#conditional_formats}
	ConditionalFormats interface{} `field:"optional" json:"conditionalFormats" yaml:"conditionalFormats"`
	// limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#limit PowerpackV2#limit}
	Limit *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaLimit `field:"optional" json:"limit" yaml:"limit"`
	// number_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#number_format PowerpackV2#number_format}
	NumberFormat *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormat `field:"optional" json:"numberFormat" yaml:"numberFormat"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#style PowerpackV2#style}
	Style *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaStyle `field:"optional" json:"style" yaml:"style"`
}

