// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2TemplateVariable struct {
	// The name of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#name DashboardV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of values that the template variable drop-down is limited to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#available_values DashboardV2#available_values}
	AvailableValues *[]*string `field:"optional" json:"availableValues" yaml:"availableValues"`
	// The default value for the template variable on dashboard load.
	//
	// Cannot be used in conjunction with `defaults`. **Deprecated.** Use `defaults` instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#default DashboardV2#default}
	Default *string `field:"optional" json:"default" yaml:"default"`
	// One or many default values for template variables on load.
	//
	// If more than one default is specified, they will be unioned together with `OR`. Cannot be used in conjunction with `default`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#defaults DashboardV2#defaults}
	Defaults *[]*string `field:"optional" json:"defaults" yaml:"defaults"`
	// The tag prefix associated with the variable. Only tags with this prefix appear in the variable dropdown.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#prefix DashboardV2#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The type of variable.
	//
	// This is to differentiate between filter variables (interpolated in query) and group by variables (interpolated into group by).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#type DashboardV2#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

