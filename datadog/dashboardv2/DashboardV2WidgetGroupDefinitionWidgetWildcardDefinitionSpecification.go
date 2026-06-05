// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionSpecification struct {
	// The Vega or Vega-Lite specification as a JSON string. Use `jsonencode()` to encode the specification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/dashboard_v2#contents DashboardV2#contents}
	Contents *string `field:"required" json:"contents" yaml:"contents"`
	// The type of specification (Vega or Vega-Lite). Valid values are `vega`, `vega-lite`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/dashboard_v2#type DashboardV2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

