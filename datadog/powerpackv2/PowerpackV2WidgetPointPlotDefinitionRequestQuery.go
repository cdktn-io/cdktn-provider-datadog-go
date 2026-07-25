// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetPointPlotDefinitionRequestQuery struct {
	// Data source for the query (for example, `logs`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#data_source PowerpackV2#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// The query string to filter events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#query_string PowerpackV2#query_string}
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// List of indexes to query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#indexes PowerpackV2#indexes}
	Indexes *[]*string `field:"optional" json:"indexes" yaml:"indexes"`
	// Storage location for the query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#storage PowerpackV2#storage}
	Storage *string `field:"optional" json:"storage" yaml:"storage"`
}

