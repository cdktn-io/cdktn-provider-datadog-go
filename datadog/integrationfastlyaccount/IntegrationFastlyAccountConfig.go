// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationfastlyaccount

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IntegrationFastlyAccountConfig struct {
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
	// The name of the Fastly account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/integration_fastly_account#name IntegrationFastlyAccount#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The API key for the Fastly account. Exactly one of `api_key` or `api_key_wo` must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/integration_fastly_account#api_key IntegrationFastlyAccount#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Write-only API key for the Fastly account.
	//
	// Exactly one of `api_key` or `api_key_wo` must be set. Must be used with `api_key_wo_version`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/integration_fastly_account#api_key_wo IntegrationFastlyAccount#api_key_wo}
	ApiKeyWo *string `field:"optional" json:"apiKeyWo" yaml:"apiKeyWo"`
	// Version for `api_key_wo` rotation. Changing this triggers an update. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/integration_fastly_account#api_key_wo_version IntegrationFastlyAccount#api_key_wo_version}
	ApiKeyWoVersion *string `field:"optional" json:"apiKeyWoVersion" yaml:"apiKeyWoVersion"`
}

