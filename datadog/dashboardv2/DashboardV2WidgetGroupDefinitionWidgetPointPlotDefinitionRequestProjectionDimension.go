// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetPointPlotDefinitionRequestProjectionDimension struct {
	// Source column name from the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#column DashboardV2#column}
	Column *string `field:"required" json:"column" yaml:"column"`
	// Dimension of the point plot. Valid values are `group`, `time`, `y`, `radius`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#dimension DashboardV2#dimension}
	Dimension *string `field:"required" json:"dimension" yaml:"dimension"`
	// Alias for the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#alias DashboardV2#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

