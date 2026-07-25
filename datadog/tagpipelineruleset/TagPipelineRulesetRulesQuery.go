// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tagpipelineruleset


type TagPipelineRulesetRulesQuery struct {
	// addition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/tag_pipeline_ruleset#addition TagPipelineRuleset#addition}
	Addition *TagPipelineRulesetRulesQueryAddition `field:"optional" json:"addition" yaml:"addition"`
	// Whether the query matching is case insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/tag_pipeline_ruleset#case_insensitivity TagPipelineRuleset#case_insensitivity}
	CaseInsensitivity interface{} `field:"optional" json:"caseInsensitivity" yaml:"caseInsensitivity"`
	// Whether to apply the query only if the key doesn't exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/tag_pipeline_ruleset#if_not_exists TagPipelineRuleset#if_not_exists}
	IfNotExists interface{} `field:"optional" json:"ifNotExists" yaml:"ifNotExists"`
	// Behavior when the tag already exists.
	//
	// Valid values: `append` (append to the existing tag value), `replace` (replace existing tag value), `do_not_apply` (never apply if tag already exists). Valid values are `append`, `replace`, `do_not_apply`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/tag_pipeline_ruleset#if_tag_exists TagPipelineRuleset#if_tag_exists}
	IfTagExists *string `field:"optional" json:"ifTagExists" yaml:"ifTagExists"`
	// The query string.
	//
	// Datadog normalizes queries to a canonical form (operator casing, spacing, redundant parentheses, quoting, and similar formatting differences). The provider validates the configured value against this canonical form during planning. If the value doesn't match, the plan fails with an error showing the canonical query to use in the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/tag_pipeline_ruleset#query TagPipelineRuleset#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
}

