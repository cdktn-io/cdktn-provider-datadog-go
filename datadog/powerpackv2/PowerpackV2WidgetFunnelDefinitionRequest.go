// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetFunnelDefinitionRequest struct {
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#query PowerpackV2#query}
	Query *PowerpackV2WidgetFunnelDefinitionRequestQuery `field:"required" json:"query" yaml:"query"`
}

