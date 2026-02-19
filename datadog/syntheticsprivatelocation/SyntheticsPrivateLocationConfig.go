// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticsprivatelocation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SyntheticsPrivateLocationConfig struct {
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
	// Synthetics private location name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/synthetics_private_location#name SyntheticsPrivateLocation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// API key used to generate the private location configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/synthetics_private_location#api_key SyntheticsPrivateLocation#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Description of the private location. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/synthetics_private_location#description SyntheticsPrivateLocation#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/synthetics_private_location#metadata SyntheticsPrivateLocation#metadata}
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// A list of tags to associate with your synthetics private location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/synthetics_private_location#tags SyntheticsPrivateLocation#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

