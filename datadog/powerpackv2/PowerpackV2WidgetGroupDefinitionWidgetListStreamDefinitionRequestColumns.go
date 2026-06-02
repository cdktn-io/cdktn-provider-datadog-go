// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetListStreamDefinitionRequestColumns struct {
	// Widget column field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#field PowerpackV2#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
	// Widget column width. Valid values are `auto`, `compact`, `full`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#width PowerpackV2#width}
	Width *string `field:"optional" json:"width" yaml:"width"`
}

