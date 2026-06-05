// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionYaxis struct {
	// Set to `true` to include zero.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/dashboard_v2#include_zero DashboardV2#include_zero}
	IncludeZero interface{} `field:"optional" json:"includeZero" yaml:"includeZero"`
	// The label of the axis to display on the graph. Only usable on Scatterplot Widgets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/dashboard_v2#label DashboardV2#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Specifies the maximum numeric value to show on the axis. Defaults to `auto`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/dashboard_v2#max DashboardV2#max}
	Max *string `field:"optional" json:"max" yaml:"max"`
	// Specifies the minimum numeric value to show on the axis. Defaults to `auto`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/dashboard_v2#min DashboardV2#min}
	Min *string `field:"optional" json:"min" yaml:"min"`
	// Specifies the scale type. Possible values are `linear`, `log`, `sqrt`, and `pow##` (for example `pow2` or `pow0.5`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/dashboard_v2#scale DashboardV2#scale}
	Scale *string `field:"optional" json:"scale" yaml:"scale"`
}

