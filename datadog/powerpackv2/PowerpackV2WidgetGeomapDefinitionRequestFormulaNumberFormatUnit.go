// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnit struct {
	// canonical block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#canonical PowerpackV2#canonical}
	Canonical *PowerpackV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonical `field:"optional" json:"canonical" yaml:"canonical"`
	// custom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#custom PowerpackV2#custom}
	Custom *PowerpackV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitCustom `field:"optional" json:"custom" yaml:"custom"`
}

