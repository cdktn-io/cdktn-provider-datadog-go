// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetNoteDefinitionTimeLive struct {
	// Unit of the time span. Valid values are `minute`, `hour`, `day`, `week`, `month`, `year`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#unit PowerpackV2#unit}
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Value of the time span.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#value PowerpackV2#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

