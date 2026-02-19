// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apikey

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApiKeyConfig struct {
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
	// Name for API Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/api_key#name ApiKey#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether the API key is used for remote config.
	//
	// Set to true only if remote config is enabled in `/organization-settings/remote-config`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/api_key#remote_config_read_enabled ApiKey#remote_config_read_enabled}
	RemoteConfigReadEnabled interface{} `field:"optional" json:"remoteConfigReadEnabled" yaml:"remoteConfigReadEnabled"`
}

