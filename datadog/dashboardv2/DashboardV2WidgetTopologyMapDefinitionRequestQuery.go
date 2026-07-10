// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetTopologyMapDefinitionRequestQuery struct {
	// The data source for the Topology request ('service_map' or 'data_streams').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#data_source DashboardV2#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// Your environment and primary tag (or `*` if enabled for your account).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#filters DashboardV2#filters}
	Filters *[]*string `field:"required" json:"filters" yaml:"filters"`
	// Name of the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/dashboard_v2#service DashboardV2#service}
	Service *string `field:"required" json:"service" yaml:"service"`
}

