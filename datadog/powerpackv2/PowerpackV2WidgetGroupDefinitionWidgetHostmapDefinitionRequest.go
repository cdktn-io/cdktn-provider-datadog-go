// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequest struct {
	// fill block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#fill PowerpackV2#fill}
	Fill *PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFill `field:"optional" json:"fill" yaml:"fill"`
	// size block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#size PowerpackV2#size}
	Size *PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSize `field:"optional" json:"size" yaml:"size"`
}

