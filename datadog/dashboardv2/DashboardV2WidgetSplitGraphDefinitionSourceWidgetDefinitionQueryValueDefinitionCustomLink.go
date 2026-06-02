// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionCustomLink struct {
	// The flag for toggling context menu link visibility.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#is_hidden DashboardV2#is_hidden}
	IsHidden interface{} `field:"optional" json:"isHidden" yaml:"isHidden"`
	// The label for the custom link URL. Keep the label short and descriptive. Use metrics and tags as variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#label DashboardV2#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The URL of the custom link. URL must include `http` or `https`. A relative URL must start with `/`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#link DashboardV2#link}
	Link *string `field:"optional" json:"link" yaml:"link"`
	// The label ID that refers to a context menu link.
	//
	// Can be `logs`, `hosts`, `traces`, `profiles`, `processes`, `containers`, or `rum`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#override_label DashboardV2#override_label}
	OverrideLabel *string `field:"optional" json:"overrideLabel" yaml:"overrideLabel"`
}

