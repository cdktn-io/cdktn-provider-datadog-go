// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinitionRequestQuerySearch struct {
	// Expression describing the journey between nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#expression PowerpackV2#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// JSON object mapping journey node names to Product Analytics base queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#node_objects PowerpackV2#node_objects}
	NodeObjects *string `field:"required" json:"nodeObjects" yaml:"nodeObjects"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#filters PowerpackV2#filters}
	Filters *PowerpackV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinitionRequestQuerySearchFilters `field:"optional" json:"filters" yaml:"filters"`
	// join_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#join_keys PowerpackV2#join_keys}
	JoinKeys *PowerpackV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinitionRequestQuerySearchJoinKeys `field:"optional" json:"joinKeys" yaml:"joinKeys"`
	// JSON object mapping journey step names to display aliases.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#step_aliases PowerpackV2#step_aliases}
	StepAliases *string `field:"optional" json:"stepAliases" yaml:"stepAliases"`
}

