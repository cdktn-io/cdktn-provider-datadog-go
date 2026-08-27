// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilter struct {
	// Graph filter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#name PowerpackV2#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Graph filter operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#operator PowerpackV2#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#target PowerpackV2#target}
	Target *PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterTarget `field:"optional" json:"target" yaml:"target"`
	// Graph filter value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#value PowerpackV2#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

