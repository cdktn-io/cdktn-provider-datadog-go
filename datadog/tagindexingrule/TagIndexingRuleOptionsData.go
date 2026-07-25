// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tagindexingrule


type TagIndexingRuleOptionsData struct {
	// Configuration for including dynamically queried tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/tag_indexing_rule#dynamic_tags TagIndexingRule#dynamic_tags}
	DynamicTags *TagIndexingRuleOptionsDataDynamicTags `field:"optional" json:"dynamicTags" yaml:"dynamicTags"`
	// When true, the rule applies to metrics ingested before the rule was created. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/tag_indexing_rule#manage_preexisting_metrics TagIndexingRule#manage_preexisting_metrics}
	ManagePreexistingMetrics interface{} `field:"optional" json:"managePreexistingMetrics" yaml:"managePreexistingMetrics"`
	// Criteria for matching metrics based on query state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/tag_indexing_rule#metric_match TagIndexingRule#metric_match}
	MetricMatch *TagIndexingRuleOptionsDataMetricMatch `field:"optional" json:"metricMatch" yaml:"metricMatch"`
	// When true, this rule's tag list overrides tags configured by earlier rules for the same metric. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/tag_indexing_rule#override_previous_rules TagIndexingRule#override_previous_rules}
	OverridePreviousRules interface{} `field:"optional" json:"overridePreviousRules" yaml:"overridePreviousRules"`
}

