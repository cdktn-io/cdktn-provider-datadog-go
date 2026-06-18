// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tagindexingrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TagIndexingRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Metric name prefixes (glob patterns) this rule applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/tag_indexing_rule#metric_name_matches TagIndexingRule#metric_name_matches}
	MetricNameMatches *[]*string `field:"required" json:"metricNameMatches" yaml:"metricNameMatches"`
	// Human-readable name for the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/tag_indexing_rule#name TagIndexingRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// When true, the rule excludes the listed tags and indexes all others.
	//
	// When false (default), the rule includes only the listed tags. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/tag_indexing_rule#exclude_tags_mode TagIndexingRule#exclude_tags_mode}
	ExcludeTagsMode interface{} `field:"optional" json:"excludeTagsMode" yaml:"excludeTagsMode"`
	// Metric name prefixes excluded from the rule's scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/tag_indexing_rule#ignored_metric_name_matches TagIndexingRule#ignored_metric_name_matches}
	IgnoredMetricNameMatches *[]*string `field:"optional" json:"ignoredMetricNameMatches" yaml:"ignoredMetricNameMatches"`
	// Versioned configuration options for the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/tag_indexing_rule#options TagIndexingRule#options}
	Options *TagIndexingRuleOptions `field:"optional" json:"options" yaml:"options"`
	// Tag keys this rule includes or excludes, depending on exclude_tags_mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/tag_indexing_rule#tags TagIndexingRule#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

