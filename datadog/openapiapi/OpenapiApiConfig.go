// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package openapiapi

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OpenapiApiConfig struct {
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
	// The textual content of the OpenAPI specification.
	//
	// Use [`file()`](https://developer.hashicorp.com/terraform/language/functions/file) in order to reference another file in the repository (see exmaple).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/openapi_api#spec OpenapiApi#spec}
	Spec *string `field:"required" json:"spec" yaml:"spec"`
}

