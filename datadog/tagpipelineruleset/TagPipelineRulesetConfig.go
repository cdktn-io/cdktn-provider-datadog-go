// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tagpipelineruleset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TagPipelineRulesetConfig struct {
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
	// The name of the ruleset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/tag_pipeline_ruleset#name TagPipelineRuleset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether the ruleset is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/tag_pipeline_ruleset#enabled TagPipelineRuleset#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/tag_pipeline_ruleset#rules TagPipelineRuleset#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

