// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetTreemapDefinitionRequest struct {
	// formula block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#formula PowerpackV2#formula}
	Formula interface{} `field:"optional" json:"formula" yaml:"formula"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#query PowerpackV2#query}
	Query interface{} `field:"optional" json:"query" yaml:"query"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#sort PowerpackV2#sort}
	Sort *PowerpackV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestSort `field:"optional" json:"sort" yaml:"sort"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#style PowerpackV2#style}
	Style *PowerpackV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestStyle `field:"optional" json:"style" yaml:"style"`
}

