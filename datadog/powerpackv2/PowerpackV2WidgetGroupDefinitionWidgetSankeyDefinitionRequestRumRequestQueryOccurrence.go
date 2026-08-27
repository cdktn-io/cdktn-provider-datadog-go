// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOccurrence struct {
	// The comparison operator used for the occurrence filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#operator PowerpackV2#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// The threshold value to compare occurrence counts against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#value PowerpackV2#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

