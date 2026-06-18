// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetTopologyMapDefinitionRequestQuery struct {
	// The data source for the Topology request ('service_map' or 'data_streams').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#data_source PowerpackV2#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// Your environment and primary tag (or `*` if enabled for your account).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#filters PowerpackV2#filters}
	Filters *[]*string `field:"required" json:"filters" yaml:"filters"`
	// Name of the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#service PowerpackV2#service}
	Service *string `field:"required" json:"service" yaml:"service"`
}

