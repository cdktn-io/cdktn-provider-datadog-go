// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableFormula struct {
	// Dimension of the Scatterplot. Valid values are `x`, `y`, `radius`, `color`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/dashboard_v2#dimension DashboardV2#dimension}
	Dimension *string `field:"required" json:"dimension" yaml:"dimension"`
	// A string expression built from queries, formulas, and functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/dashboard_v2#formula_expression DashboardV2#formula_expression}
	FormulaExpression *string `field:"required" json:"formulaExpression" yaml:"formulaExpression"`
	// An expression alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/dashboard_v2#alias DashboardV2#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

