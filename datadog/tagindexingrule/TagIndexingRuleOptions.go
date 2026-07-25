// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tagindexingrule


type TagIndexingRuleOptions struct {
	// Behavioral options for how the rule applies to metrics, including backfill and override behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/tag_indexing_rule#data TagIndexingRule#data}
	Data *TagIndexingRuleOptionsData `field:"required" json:"data" yaml:"data"`
	// Options schema version. Only `1` is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/tag_indexing_rule#version TagIndexingRule#version}
	Version *float64 `field:"required" json:"version" yaml:"version"`
}

