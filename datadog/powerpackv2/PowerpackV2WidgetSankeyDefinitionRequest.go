// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetSankeyDefinitionRequest struct {
	// network_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/powerpack_v2#network_request PowerpackV2#network_request}
	NetworkRequest *PowerpackV2WidgetSankeyDefinitionRequestNetworkRequest `field:"optional" json:"networkRequest" yaml:"networkRequest"`
	// rum_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/powerpack_v2#rum_request PowerpackV2#rum_request}
	RumRequest *PowerpackV2WidgetSankeyDefinitionRequestRumRequest `field:"optional" json:"rumRequest" yaml:"rumRequest"`
}

