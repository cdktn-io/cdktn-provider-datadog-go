// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetToplistDefinitionStyleDisplay struct {
	// flat block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/powerpack_v2#flat PowerpackV2#flat}
	Flat *PowerpackV2WidgetToplistDefinitionStyleDisplayFlat `field:"optional" json:"flat" yaml:"flat"`
	// stacked block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/powerpack_v2#stacked PowerpackV2#stacked}
	Stacked *PowerpackV2WidgetToplistDefinitionStyleDisplayStacked `field:"optional" json:"stacked" yaml:"stacked"`
}

