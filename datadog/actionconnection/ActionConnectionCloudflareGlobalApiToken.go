// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionCloudflareGlobalApiToken struct {
	// Email address associated with the Cloudflare account. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/action_connection#auth_email ActionConnection#auth_email}
	AuthEmail *string `field:"optional" json:"authEmail" yaml:"authEmail"`
	// Cloudflare global API key. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/action_connection#global_api_key ActionConnection#global_api_key}
	GlobalApiKey *string `field:"optional" json:"globalApiKey" yaml:"globalApiKey"`
}

