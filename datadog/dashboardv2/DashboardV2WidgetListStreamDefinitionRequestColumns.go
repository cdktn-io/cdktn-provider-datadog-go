// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetListStreamDefinitionRequestColumns struct {
	// Widget column field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/dashboard_v2#field DashboardV2#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
	// Widget column width. Valid values are `auto`, `compact`, `full`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/dashboard_v2#width DashboardV2#width}
	Width *string `field:"optional" json:"width" yaml:"width"`
}

