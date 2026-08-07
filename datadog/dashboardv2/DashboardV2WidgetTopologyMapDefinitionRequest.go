// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetTopologyMapDefinitionRequest struct {
	// The request type for the Topology request ('topology').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#request_type DashboardV2#request_type}
	RequestType *string `field:"required" json:"requestType" yaml:"requestType"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#query DashboardV2#query}
	Query *DashboardV2WidgetTopologyMapDefinitionRequestQuery `field:"optional" json:"query" yaml:"query"`
}

