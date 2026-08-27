// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryRetentionQuerySearchCohortCriteria struct {
	// base_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#base_query PowerpackV2#base_query}
	BaseQuery *PowerpackV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryRetentionQuerySearchCohortCriteriaBaseQuery `field:"required" json:"baseQuery" yaml:"baseQuery"`
	// time_interval block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#time_interval PowerpackV2#time_interval}
	TimeInterval *PowerpackV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryRetentionQuerySearchCohortCriteriaTimeInterval `field:"required" json:"timeInterval" yaml:"timeInterval"`
}

