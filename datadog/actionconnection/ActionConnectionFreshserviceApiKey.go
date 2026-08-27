// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionFreshserviceApiKey struct {
	// Freshservice API key. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/action_connection#api_key ActionConnection#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Freshservice domain. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/action_connection#domain ActionConnection#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
}

