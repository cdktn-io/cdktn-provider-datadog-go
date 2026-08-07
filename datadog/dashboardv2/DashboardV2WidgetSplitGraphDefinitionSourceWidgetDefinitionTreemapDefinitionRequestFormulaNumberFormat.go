// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestFormulaNumberFormat struct {
	// unit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#unit DashboardV2#unit}
	Unit *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestFormulaNumberFormatUnit `field:"required" json:"unit" yaml:"unit"`
	// unit_scale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#unit_scale DashboardV2#unit_scale}
	UnitScale *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestFormulaNumberFormatUnitScale `field:"optional" json:"unitScale" yaml:"unitScale"`
}

