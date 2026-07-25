// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetTreemapDefinitionRequestSortOrderBy struct {
	// formula_sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#formula_sort PowerpackV2#formula_sort}
	FormulaSort *PowerpackV2WidgetTreemapDefinitionRequestSortOrderByFormulaSort `field:"optional" json:"formulaSort" yaml:"formulaSort"`
	// group_sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#group_sort PowerpackV2#group_sort}
	GroupSort *PowerpackV2WidgetTreemapDefinitionRequestSortOrderByGroupSort `field:"optional" json:"groupSort" yaml:"groupSort"`
}

