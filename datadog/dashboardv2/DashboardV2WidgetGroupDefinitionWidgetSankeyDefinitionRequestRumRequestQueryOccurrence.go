// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOccurrence struct {
	// The comparison operator used for the occurrence filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#operator DashboardV2#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// The threshold value to compare occurrence counts against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#value DashboardV2#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

