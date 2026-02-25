// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package spansmetric

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SpansMetricConfig struct {
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
	// The name of the span-based metric. This field can't be updated after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/spans_metric#name SpansMetric#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/spans_metric#compute SpansMetric#compute}
	Compute *SpansMetricCompute `field:"optional" json:"compute" yaml:"compute"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/spans_metric#filter SpansMetric#filter}
	Filter *SpansMetricFilter `field:"optional" json:"filter" yaml:"filter"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/spans_metric#group_by SpansMetric#group_by}
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
}

