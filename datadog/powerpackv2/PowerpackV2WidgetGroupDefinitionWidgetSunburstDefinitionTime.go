// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionTime struct {
	// fixed block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/powerpack_v2#fixed PowerpackV2#fixed}
	Fixed *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionTimeFixed `field:"optional" json:"fixed" yaml:"fixed"`
	// live block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/powerpack_v2#live PowerpackV2#live}
	Live *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionTimeLive `field:"optional" json:"live" yaml:"live"`
}

