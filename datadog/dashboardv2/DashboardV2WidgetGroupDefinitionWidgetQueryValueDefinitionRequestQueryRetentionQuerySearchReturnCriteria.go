// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryRetentionQuerySearchReturnCriteria struct {
	// base_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#base_query DashboardV2#base_query}
	BaseQuery *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryRetentionQuerySearchReturnCriteriaBaseQuery `field:"required" json:"baseQuery" yaml:"baseQuery"`
	// time_interval block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#time_interval DashboardV2#time_interval}
	TimeInterval *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryRetentionQuerySearchReturnCriteriaTimeInterval `field:"optional" json:"timeInterval" yaml:"timeInterval"`
}

