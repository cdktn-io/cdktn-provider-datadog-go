// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestNetworkRequestQuerySort struct {
	// Field to sort by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#field DashboardV2#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
	// Sort direction. Valid values are `asc`, `desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#order DashboardV2#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}

