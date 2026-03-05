// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatadogmetricactivetagsandaggregations

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatadogMetricActiveTagsAndAggregationsConfig struct {
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
	// The metric for which to fetch tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/data-sources/metric_active_tags_and_aggregations#metric DataDatadogMetricActiveTagsAndAggregations#metric}
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// The number of seconds to look back from now.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/data-sources/metric_active_tags_and_aggregations#window DataDatadogMetricActiveTagsAndAggregations#window}
	Window *float64 `field:"optional" json:"window" yaml:"window"`
}

