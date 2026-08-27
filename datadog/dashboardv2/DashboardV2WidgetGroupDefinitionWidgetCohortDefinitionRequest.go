// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetCohortDefinitionRequest struct {
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#query DashboardV2#query}
	Query *DashboardV2WidgetGroupDefinitionWidgetCohortDefinitionRequestQuery `field:"required" json:"query" yaml:"query"`
	// Request type for the retention grid widget. Valid values are `retention_grid`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#request_type DashboardV2#request_type}
	RequestType *string `field:"required" json:"requestType" yaml:"requestType"`
}

