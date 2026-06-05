// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetToplistDefinitionTimeFixed struct {
	// Start time in seconds since epoch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/dashboard_v2#from DashboardV2#from}
	From *float64 `field:"required" json:"from" yaml:"from"`
	// End time in seconds since epoch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/dashboard_v2#to DashboardV2#to}
	To *float64 `field:"required" json:"to" yaml:"to"`
}

