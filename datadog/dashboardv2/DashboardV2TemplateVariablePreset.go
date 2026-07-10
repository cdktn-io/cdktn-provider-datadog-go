// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2TemplateVariablePreset struct {
	// The name of the preset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#name DashboardV2#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// template_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#template_variable DashboardV2#template_variable}
	TemplateVariable interface{} `field:"optional" json:"templateVariable" yaml:"templateVariable"`
}

