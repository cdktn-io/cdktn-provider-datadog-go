// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestListstreamRequest struct {
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/powerpack_v2#columns PowerpackV2#columns}
	Columns interface{} `field:"required" json:"columns" yaml:"columns"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/powerpack_v2#query PowerpackV2#query}
	Query *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestListstreamRequestQuery `field:"required" json:"query" yaml:"query"`
}

