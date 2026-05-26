// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetFunnelDefinitionRequestQueryStep struct {
	// The facet of the step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/dashboard_v2#facet DashboardV2#facet}
	Facet *string `field:"required" json:"facet" yaml:"facet"`
	// The value of the step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/dashboard_v2#value DashboardV2#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

