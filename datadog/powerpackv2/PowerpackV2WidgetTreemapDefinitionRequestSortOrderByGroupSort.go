// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetTreemapDefinitionRequestSortOrderByGroupSort struct {
	// The name of the group tag to sort by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#name PowerpackV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Widget sorting direction. Valid values are `asc`, `desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#order PowerpackV2#order}
	Order *string `field:"required" json:"order" yaml:"order"`
}

