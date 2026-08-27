// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilter struct {
	// Graph filter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#name DashboardV2#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Graph filter operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#operator DashboardV2#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#target DashboardV2#target}
	Target *DashboardV2WidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterTarget `field:"optional" json:"target" yaml:"target"`
	// Graph filter value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#value DashboardV2#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

