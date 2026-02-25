// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tagpipelinerulesets

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TagPipelineRulesetsConfig struct {
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
	// The list of Tag Pipeline Ruleset IDs, in order. Rulesets are executed in the order specified in this list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/tag_pipeline_rulesets#ruleset_ids TagPipelineRulesets#ruleset_ids}
	RulesetIds *[]*string `field:"required" json:"rulesetIds" yaml:"rulesetIds"`
	// Whether to override UI-defined rulesets.
	//
	// When set to true, any rulesets created via the UI that are not defined in Terraform will be deleted and Terraform will be used as the source of truth for rules and their ordering. When set to false, any rulesets created via the UI that are at the end of order will be kept but will be warned, otherwise an error will be thrown in terraform plan phase. Default is false
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/tag_pipeline_rulesets#override_ui_defined_resources TagPipelineRulesets#override_ui_defined_resources}
	OverrideUiDefinedResources interface{} `field:"optional" json:"overrideUiDefinedResources" yaml:"overrideUiDefinedResources"`
}

