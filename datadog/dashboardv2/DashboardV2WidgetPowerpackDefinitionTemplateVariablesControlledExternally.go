// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetPowerpackDefinitionTemplateVariablesControlledExternally struct {
	// The name of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#name DashboardV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// One or many template variable values within the saved view, which will be unioned together using `OR` if more than one is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#values DashboardV2#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// The tag prefix associated with the variable. Only tags with this prefix appear in the variable dropdown.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#prefix DashboardV2#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

