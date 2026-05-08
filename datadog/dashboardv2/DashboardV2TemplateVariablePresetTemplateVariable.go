// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2TemplateVariablePresetTemplateVariable struct {
	// The name of the template variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#name DashboardV2#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value that should be assumed by the template variable in this preset.
	//
	// Cannot be used in conjunction with `values`. **Deprecated.** Use `values` instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#value DashboardV2#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// One or many template variable values within the saved view, which will be unioned together using `OR` if more than one is specified.
	//
	// Cannot be used in conjunction with `value`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#values DashboardV2#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

