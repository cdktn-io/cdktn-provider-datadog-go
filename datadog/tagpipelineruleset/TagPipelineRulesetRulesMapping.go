// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tagpipelineruleset


type TagPipelineRulesetRulesMapping struct {
	// The destination key for the mapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/tag_pipeline_ruleset#destination_key TagPipelineRuleset#destination_key}
	DestinationKey *string `field:"optional" json:"destinationKey" yaml:"destinationKey"`
	// Whether to apply the mapping only if the destination key doesn't exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/tag_pipeline_ruleset#if_not_exists TagPipelineRuleset#if_not_exists}
	IfNotExists interface{} `field:"optional" json:"ifNotExists" yaml:"ifNotExists"`
	// Behavior when the tag already exists.
	//
	// Valid values: `append` (append to the existing tag value), `replace` (replace existing tag value), `do_not_apply` (never apply if tag already exists). Valid values are `append`, `replace`, `do_not_apply`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/tag_pipeline_ruleset#if_tag_exists TagPipelineRuleset#if_tag_exists}
	IfTagExists *string `field:"optional" json:"ifTagExists" yaml:"ifTagExists"`
	// The source keys for the mapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/tag_pipeline_ruleset#source_keys TagPipelineRuleset#source_keys}
	SourceKeys *[]*string `field:"optional" json:"sourceKeys" yaml:"sourceKeys"`
}

