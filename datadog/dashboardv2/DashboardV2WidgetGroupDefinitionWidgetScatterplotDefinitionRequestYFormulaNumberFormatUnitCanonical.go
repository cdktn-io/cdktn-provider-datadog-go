// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYFormulaNumberFormatUnitCanonical struct {
	// per unit name. If you want to represent megabytes/s, you set 'unit_name' = 'megabyte' and 'per_unit_name = 'second'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#per_unit_name DashboardV2#per_unit_name}
	PerUnitName *string `field:"optional" json:"perUnitName" yaml:"perUnitName"`
	// Unit name. It should be in singular form ('megabyte' and not 'megabytes').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#unit_name DashboardV2#unit_name}
	UnitName *string `field:"optional" json:"unitName" yaml:"unitName"`
}

