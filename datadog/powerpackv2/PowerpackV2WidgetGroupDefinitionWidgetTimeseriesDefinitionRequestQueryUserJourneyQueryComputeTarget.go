// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQueryComputeTarget struct {
	// Target type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#type PowerpackV2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// End node of the target range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#end PowerpackV2#end}
	End *string `field:"optional" json:"end" yaml:"end"`
	// Start node of the target range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#start PowerpackV2#start}
	Start *string `field:"optional" json:"start" yaml:"start"`
	// Target node value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#value PowerpackV2#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

