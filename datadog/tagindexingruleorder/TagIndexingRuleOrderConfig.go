// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tagindexingruleorder

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TagIndexingRuleOrderConfig struct {
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
	// A unique name for the order resource.
	//
	// Recommended to match the resource name. No corresponding field exists in the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/tag_indexing_rule_order#name TagIndexingRuleOrder#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Ordered list of EVERY active tag indexing rule UUID in the org.
	//
	// The server assigns each rule a rule_order (1, 2, 3, ...) by its position in this list. This resource claims full ownership of the org's evaluation order: rules created outside Terraform (e.g. via the UI) appear as drift on the next plan and must be added here. The list must be the COMPLETE set of active rules and contain each UUID exactly once — omitting an existing rule or repeating a UUID is rejected by the API with a 400; listing a UUID that does not exist returns 404.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/tag_indexing_rule_order#rule_ids TagIndexingRuleOrder#rule_ids}
	RuleIds *[]*string `field:"required" json:"ruleIds" yaml:"ruleIds"`
}

