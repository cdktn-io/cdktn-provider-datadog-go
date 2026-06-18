// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tagindexingrule


type TagIndexingRuleOptionsDataDynamicTags struct {
	// Lookback window for determining which tags were recently queried.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/tag_indexing_rule#queried_tags_window_seconds TagIndexingRule#queried_tags_window_seconds}
	QueriedTagsWindowSeconds *float64 `field:"optional" json:"queriedTagsWindowSeconds" yaml:"queriedTagsWindowSeconds"`
	// When true, tags from related assets are included.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/tag_indexing_rule#related_asset_tags TagIndexingRule#related_asset_tags}
	RelatedAssetTags interface{} `field:"optional" json:"relatedAssetTags" yaml:"relatedAssetTags"`
}

