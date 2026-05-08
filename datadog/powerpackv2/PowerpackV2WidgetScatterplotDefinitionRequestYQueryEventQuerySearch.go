// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetScatterplotDefinitionRequestYQueryEventQuerySearch struct {
	// The events search string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/powerpack_v2#query PowerpackV2#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

