// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package restrictionpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RestrictionPolicyConfig struct {
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
	// Identifier for the resource, formatted as resource_type:resource_id.
	//
	// Resources to define `resource_type` :
	// * [List of supported resources](https://docs.datadoghq.com/account_management/rbac/granular_access)
	// * [Resource type definition](https://docs.datadoghq.com/api/latest/restriction-policies/#supported-resources)
	//
	// Restrictions :
	// * Dashboards : support is in private beta. Reach out to your Datadog contact or support to enable this.
	// * Monitors : Management of restriction policy through terraform is now available in Preview. Please request access via https://docs.datadoghq.com/help
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/restriction_policy#resource_id RestrictionPolicy#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// bindings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/restriction_policy#bindings RestrictionPolicy#bindings}
	Bindings interface{} `field:"optional" json:"bindings" yaml:"bindings"`
}

