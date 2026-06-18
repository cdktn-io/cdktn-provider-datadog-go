// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSankeyDefinitionRequest struct {
	// network_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/dashboard_v2#network_request DashboardV2#network_request}
	NetworkRequest *DashboardV2WidgetSankeyDefinitionRequestNetworkRequest `field:"optional" json:"networkRequest" yaml:"networkRequest"`
	// rum_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/dashboard_v2#rum_request DashboardV2#rum_request}
	RumRequest *DashboardV2WidgetSankeyDefinitionRequestRumRequest `field:"optional" json:"rumRequest" yaml:"rumRequest"`
}

