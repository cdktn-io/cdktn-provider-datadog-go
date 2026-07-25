// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionOktaApiToken struct {
	// Okta API token. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#api_token ActionConnection#api_token}
	ApiToken *string `field:"optional" json:"apiToken" yaml:"apiToken"`
	// Okta domain. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#domain ActionConnection#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
}

