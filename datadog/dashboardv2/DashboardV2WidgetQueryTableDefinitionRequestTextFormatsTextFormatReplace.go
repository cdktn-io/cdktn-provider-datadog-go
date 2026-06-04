// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetQueryTableDefinitionRequestTextFormatsTextFormatReplace struct {
	// Table widget text format replace all type. Valid values are `all`, `substring`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#type DashboardV2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Table Widget Match String.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#with DashboardV2#with}
	With *string `field:"required" json:"with" yaml:"with"`
	// Text that will be replaced. Must be used with type `substring`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#substring DashboardV2#substring}
	Substring *string `field:"optional" json:"substring" yaml:"substring"`
}

