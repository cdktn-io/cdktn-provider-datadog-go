// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSplitConfigStaticSplitsSplitVector struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#tag_key DashboardV2#tag_key}.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#tag_values DashboardV2#tag_values}.
	TagValues *[]*string `field:"required" json:"tagValues" yaml:"tagValues"`
}

