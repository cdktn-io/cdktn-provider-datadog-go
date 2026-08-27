// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetQueryValueDefinitionRequestComparisonDuration struct {
	// The comparison window type. Valid values are `previous_timeframe`, `custom_timeframe`, `previous_day`, `previous_week`, `previous_month`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#type DashboardV2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// custom_timeframe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#custom_timeframe DashboardV2#custom_timeframe}
	CustomTimeframe *DashboardV2WidgetQueryValueDefinitionRequestComparisonDurationCustomTimeframe `field:"optional" json:"customTimeframe" yaml:"customTimeframe"`
}

