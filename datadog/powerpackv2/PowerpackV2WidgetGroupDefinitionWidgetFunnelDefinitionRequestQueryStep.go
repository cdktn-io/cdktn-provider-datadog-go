// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetFunnelDefinitionRequestQueryStep struct {
	// The facet of the step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#facet PowerpackV2#facet}
	Facet *string `field:"required" json:"facet" yaml:"facet"`
	// The value of the step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#value PowerpackV2#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

