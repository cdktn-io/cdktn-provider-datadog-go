// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetQueryValueDefinitionRequestSecurityQueryComputeQuery struct {
	// The aggregation method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#aggregation DashboardV2#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// The facet name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#facet DashboardV2#facet}
	Facet *string `field:"optional" json:"facet" yaml:"facet"`
	// Define the time interval in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#interval DashboardV2#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
}

