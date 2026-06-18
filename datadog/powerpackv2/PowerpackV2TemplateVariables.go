// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2TemplateVariables struct {
	// The name of the powerpack template variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#name PowerpackV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// One or many default values for powerpack template variables on load.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/powerpack_v2#defaults PowerpackV2#defaults}
	Defaults *[]*string `field:"optional" json:"defaults" yaml:"defaults"`
}

