// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetQueryValueDefinitionTimeseriesBackground struct {
	// Whether the Timeseries is made using an area or bars. Valid values are `area`, `bars`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#type PowerpackV2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// yaxis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#yaxis PowerpackV2#yaxis}
	Yaxis *PowerpackV2WidgetGroupDefinitionWidgetQueryValueDefinitionTimeseriesBackgroundYaxis `field:"optional" json:"yaxis" yaml:"yaxis"`
}

