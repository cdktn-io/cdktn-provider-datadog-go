// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetServicemapDefinitionTime struct {
	// fixed block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/powerpack_v2#fixed PowerpackV2#fixed}
	Fixed *PowerpackV2WidgetGroupDefinitionWidgetServicemapDefinitionTimeFixed `field:"optional" json:"fixed" yaml:"fixed"`
	// live block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/powerpack_v2#live PowerpackV2#live}
	Live *PowerpackV2WidgetGroupDefinitionWidgetServicemapDefinitionTimeLive `field:"optional" json:"live" yaml:"live"`
}

