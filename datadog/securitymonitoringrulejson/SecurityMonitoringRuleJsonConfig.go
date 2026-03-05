// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrulejson

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityMonitoringRuleJsonConfig struct {
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
	// The JSON definition of the Security Monitoring Rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/security_monitoring_rule_json#json SecurityMonitoringRuleJson#json}
	Json *string `field:"required" json:"json" yaml:"json"`
}

