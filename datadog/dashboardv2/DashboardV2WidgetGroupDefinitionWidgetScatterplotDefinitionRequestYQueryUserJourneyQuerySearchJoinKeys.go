// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryUserJourneyQuerySearchJoinKeys struct {
	// Primary join key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#primary DashboardV2#primary}
	Primary *string `field:"required" json:"primary" yaml:"primary"`
	// Secondary join keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#secondary DashboardV2#secondary}
	Secondary *[]*string `field:"optional" json:"secondary" yaml:"secondary"`
}

