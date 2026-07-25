// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetHeatmapDefinitionRequestQueryEventQueryGroupByFields struct {
	// List of event facets to group by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack#fields Powerpack#fields}
	Fields *[]*string `field:"required" json:"fields" yaml:"fields"`
	// The number of groups to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack#limit Powerpack#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack#sort Powerpack#sort}
	Sort *PowerpackWidgetHeatmapDefinitionRequestQueryEventQueryGroupByFieldsSort `field:"optional" json:"sort" yaml:"sort"`
}

