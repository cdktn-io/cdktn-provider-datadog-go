// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tagindexingrule


type TagIndexingRuleOptionsData struct {
	// Configuration for excluding tags based on dynamic usage signals. Only applies when `exclude_tags_mode` is `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/tag_indexing_rule#dynamic_tags TagIndexingRule#dynamic_tags}
	DynamicTags *TagIndexingRuleOptionsDataDynamicTags `field:"optional" json:"dynamicTags" yaml:"dynamicTags"`
	// When true, the rule applies to metrics ingested before the rule was created. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/tag_indexing_rule#manage_preexisting_metrics TagIndexingRule#manage_preexisting_metrics}
	ManagePreexistingMetrics interface{} `field:"optional" json:"managePreexistingMetrics" yaml:"managePreexistingMetrics"`
	// When true, this rule's tag list overrides tags configured by earlier rules for the same metric. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/tag_indexing_rule#override_previous_rules TagIndexingRule#override_previous_rules}
	OverridePreviousRules interface{} `field:"optional" json:"overridePreviousRules" yaml:"overridePreviousRules"`
}

