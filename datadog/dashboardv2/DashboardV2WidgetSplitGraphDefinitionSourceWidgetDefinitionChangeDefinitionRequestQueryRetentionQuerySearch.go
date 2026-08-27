// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestQueryRetentionQuerySearch struct {
	// cohort_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#cohort_criteria DashboardV2#cohort_criteria}
	CohortCriteria *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestQueryRetentionQuerySearchCohortCriteria `field:"required" json:"cohortCriteria" yaml:"cohortCriteria"`
	// Entity tracked for retention. Valid values are `@usr.id`, `@account.id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#retention_entity DashboardV2#retention_entity}
	RetentionEntity *string `field:"required" json:"retentionEntity" yaml:"retentionEntity"`
	// Condition for counting an entity as returned. Valid values are `conversion_on`, `conversion_on_or_after`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#return_condition DashboardV2#return_condition}
	ReturnCondition *string `field:"required" json:"returnCondition" yaml:"returnCondition"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#filters DashboardV2#filters}
	Filters *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestQueryRetentionQuerySearchFilters `field:"optional" json:"filters" yaml:"filters"`
	// return_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#return_criteria DashboardV2#return_criteria}
	ReturnCriteria *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestQueryRetentionQuerySearchReturnCriteria `field:"optional" json:"returnCriteria" yaml:"returnCriteria"`
}

