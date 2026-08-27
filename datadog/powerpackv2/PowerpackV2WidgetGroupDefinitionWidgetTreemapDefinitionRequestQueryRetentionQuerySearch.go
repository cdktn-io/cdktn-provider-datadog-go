// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestQueryRetentionQuerySearch struct {
	// cohort_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#cohort_criteria PowerpackV2#cohort_criteria}
	CohortCriteria *PowerpackV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestQueryRetentionQuerySearchCohortCriteria `field:"required" json:"cohortCriteria" yaml:"cohortCriteria"`
	// Entity tracked for retention. Valid values are `@usr.id`, `@account.id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#retention_entity PowerpackV2#retention_entity}
	RetentionEntity *string `field:"required" json:"retentionEntity" yaml:"retentionEntity"`
	// Condition for counting an entity as returned. Valid values are `conversion_on`, `conversion_on_or_after`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#return_condition PowerpackV2#return_condition}
	ReturnCondition *string `field:"required" json:"returnCondition" yaml:"returnCondition"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#filters PowerpackV2#filters}
	Filters *PowerpackV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestQueryRetentionQuerySearchFilters `field:"optional" json:"filters" yaml:"filters"`
	// return_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#return_criteria PowerpackV2#return_criteria}
	ReturnCriteria *PowerpackV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestQueryRetentionQuerySearchReturnCriteria `field:"optional" json:"returnCriteria" yaml:"returnCriteria"`
}

