// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionDatadogApiKey struct {
	// Datadog API key. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#api_key ActionConnection#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Datadog application key. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#app_key ActionConnection#app_key}
	AppKey *string `field:"optional" json:"appKey" yaml:"appKey"`
	// Datadog site data center. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#datacenter ActionConnection#datacenter}
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// Custom subdomain used for URLs generated with this connection. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#subdomain ActionConnection#subdomain}
	Subdomain *string `field:"optional" json:"subdomain" yaml:"subdomain"`
}

