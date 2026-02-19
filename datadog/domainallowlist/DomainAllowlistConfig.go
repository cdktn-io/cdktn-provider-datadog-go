// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package domainallowlist

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DomainAllowlistConfig struct {
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
	// The domains within the domain allowlist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/domain_allowlist#domains DomainAllowlist#domains}
	Domains *[]*string `field:"required" json:"domains" yaml:"domains"`
	// Whether the Email Domain Allowlist is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/domain_allowlist#enabled DomainAllowlist#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

