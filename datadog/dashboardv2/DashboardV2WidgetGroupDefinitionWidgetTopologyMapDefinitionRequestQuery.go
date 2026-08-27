// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetTopologyMapDefinitionRequestQuery struct {
	// The data source for the Topology request. Valid values are `service_map`, `data_streams`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#data_source DashboardV2#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// Your environment and primary tag (or `*` if enabled for your account).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#filters DashboardV2#filters}
	Filters *[]*string `field:"required" json:"filters" yaml:"filters"`
	// Name of the service. Leave this empty and use `query_string` instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#service DashboardV2#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// A search string for filtering services. When set, this replaces the `service` field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#query_string DashboardV2#query_string}
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
}

