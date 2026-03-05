// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticssuite

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SyntheticsSuiteConfig struct {
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
	// Name of the Synthetics suite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/synthetics_suite#name SyntheticsSuite#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Message of the Synthetics suite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/synthetics_suite#message SyntheticsSuite#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/synthetics_suite#options SyntheticsSuite#options}
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// A set of tags to associate with your synthetics suite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/synthetics_suite#tags SyntheticsSuite#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// tests block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/synthetics_suite#tests SyntheticsSuite#tests}
	Tests interface{} `field:"optional" json:"tests" yaml:"tests"`
}

