// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSplitConfigSplitDimensions struct {
	// The system interprets this attribute differently depending on the data source of the query being split.
	//
	// For metrics, it's a tag. For the events platform, it's an attribute or tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/dashboard_v2#one_graph_per DashboardV2#one_graph_per}
	OneGraphPer *string `field:"required" json:"oneGraphPer" yaml:"oneGraphPer"`
}

