// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequest struct {
	// scatterplot_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#scatterplot_table PowerpackV2#scatterplot_table}
	ScatterplotTable interface{} `field:"optional" json:"scatterplotTable" yaml:"scatterplotTable"`
	// x block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#x PowerpackV2#x}
	X *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestX `field:"optional" json:"x" yaml:"x"`
	// y block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#y PowerpackV2#y}
	Y *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestY `field:"optional" json:"y" yaml:"y"`
}

