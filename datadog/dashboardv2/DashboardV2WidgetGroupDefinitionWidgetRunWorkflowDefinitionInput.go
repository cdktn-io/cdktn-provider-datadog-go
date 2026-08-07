// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetRunWorkflowDefinitionInput struct {
	// Name of the workflow input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#name DashboardV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Dashboard template variable. Can be suffixed with `.value` or `.key`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#value DashboardV2#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

