// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tagindexingrule


type TagIndexingRuleOptionsDataDynamicTags struct {
	// Lookback window, in seconds, for excluding tags that were not queried in that period.
	//
	// Requires `exclude_tags_mode` to be `true`. Value must be between 1 and 7776000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/tag_indexing_rule#exclude_not_queried_window_seconds TagIndexingRule#exclude_not_queried_window_seconds}
	ExcludeNotQueriedWindowSeconds *float64 `field:"optional" json:"excludeNotQueriedWindowSeconds" yaml:"excludeNotQueriedWindowSeconds"`
	// When true, excludes tags not used in any dashboards or monitors. Requires `exclude_tags_mode` to be `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/tag_indexing_rule#exclude_not_used_in_assets TagIndexingRule#exclude_not_used_in_assets}
	ExcludeNotUsedInAssets interface{} `field:"optional" json:"excludeNotUsedInAssets" yaml:"excludeNotUsedInAssets"`
}

