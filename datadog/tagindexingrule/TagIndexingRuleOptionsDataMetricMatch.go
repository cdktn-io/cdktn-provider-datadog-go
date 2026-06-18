// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tagindexingrule


type TagIndexingRuleOptionsDataMetricMatch struct {
	// Match metrics that are being queried.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/tag_indexing_rule#is_queried TagIndexingRule#is_queried}
	IsQueried interface{} `field:"optional" json:"isQueried" yaml:"isQueried"`
	// Match metrics that are not being queried.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/tag_indexing_rule#not_queried TagIndexingRule#not_queried}
	NotQueried interface{} `field:"optional" json:"notQueried" yaml:"notQueried"`
	// Match metrics not used in any dashboards or monitors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/tag_indexing_rule#not_used_in_assets TagIndexingRule#not_used_in_assets}
	NotUsedInAssets interface{} `field:"optional" json:"notUsedInAssets" yaml:"notUsedInAssets"`
	// Window in seconds for evaluating query state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/tag_indexing_rule#queried_window_seconds TagIndexingRule#queried_window_seconds}
	QueriedWindowSeconds *float64 `field:"optional" json:"queriedWindowSeconds" yaml:"queriedWindowSeconds"`
	// Match metrics used in dashboards or monitors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/tag_indexing_rule#used_in_assets TagIndexingRule#used_in_assets}
	UsedInAssets interface{} `field:"optional" json:"usedInAssets" yaml:"usedInAssets"`
}

