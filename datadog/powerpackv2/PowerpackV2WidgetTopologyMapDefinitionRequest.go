// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetTopologyMapDefinitionRequest struct {
	// The request type for the Topology request ('topology').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#request_type PowerpackV2#request_type}
	RequestType *string `field:"required" json:"requestType" yaml:"requestType"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#query PowerpackV2#query}
	Query *PowerpackV2WidgetTopologyMapDefinitionRequestQuery `field:"optional" json:"query" yaml:"query"`
}

