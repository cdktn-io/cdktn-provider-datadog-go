// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetQueryValueDefinitionRequestComparison struct {
	// duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#duration PowerpackV2#duration}
	Duration *PowerpackV2WidgetQueryValueDefinitionRequestComparisonDuration `field:"required" json:"duration" yaml:"duration"`
	// Which direction of change is considered an improvement. Valid values are `increase_better`, `decrease_better`, `neutral`. Defaults to `"neutral"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#directionality PowerpackV2#directionality}
	Directionality *string `field:"optional" json:"directionality" yaml:"directionality"`
	// How the delta is expressed. Valid values are `absolute`, `relative`, `both`. Defaults to `"absolute"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#type PowerpackV2#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

