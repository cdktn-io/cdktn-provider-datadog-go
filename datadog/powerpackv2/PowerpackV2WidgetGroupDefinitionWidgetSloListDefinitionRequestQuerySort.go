// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetSloListDefinitionRequestQuerySort struct {
	// The facet path for the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/powerpack_v2#column PowerpackV2#column}
	Column *string `field:"required" json:"column" yaml:"column"`
	// Widget sorting methods. Valid values are `asc`, `desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/powerpack_v2#order PowerpackV2#order}
	Order *string `field:"required" json:"order" yaml:"order"`
}

