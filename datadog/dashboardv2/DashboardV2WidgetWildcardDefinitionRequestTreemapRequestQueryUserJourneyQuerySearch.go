// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryUserJourneyQuerySearch struct {
	// Expression describing the journey between nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#expression DashboardV2#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// JSON object mapping journey node names to Product Analytics base queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#node_objects DashboardV2#node_objects}
	NodeObjects *string `field:"required" json:"nodeObjects" yaml:"nodeObjects"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#filters DashboardV2#filters}
	Filters *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryUserJourneyQuerySearchFilters `field:"optional" json:"filters" yaml:"filters"`
	// join_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#join_keys DashboardV2#join_keys}
	JoinKeys *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryUserJourneyQuerySearchJoinKeys `field:"optional" json:"joinKeys" yaml:"joinKeys"`
	// JSON object mapping journey step names to display aliases.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#step_aliases DashboardV2#step_aliases}
	StepAliases *string `field:"optional" json:"stepAliases" yaml:"stepAliases"`
}

