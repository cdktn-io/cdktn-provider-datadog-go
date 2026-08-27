// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetToplistDefinitionRequestQueryRetentionQuerySearchCohortCriteriaTimeInterval struct {
	// Type of cohort time interval. Valid values are `calendar`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#type DashboardV2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#value DashboardV2#value}
	Value *DashboardV2WidgetToplistDefinitionRequestQueryRetentionQuerySearchCohortCriteriaTimeIntervalValue `field:"required" json:"value" yaml:"value"`
}

