// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionAzureTenant struct {
	// Azure application client ID. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#app_client_id ActionConnection#app_client_id}
	AppClientId *string `field:"optional" json:"appClientId" yaml:"appClientId"`
	// Azure application client secret. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#client_secret ActionConnection#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Custom scope requested when acquiring an OAuth 2 access token. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#custom_scopes ActionConnection#custom_scopes}
	CustomScopes *string `field:"optional" json:"customScopes" yaml:"customScopes"`
	// Azure Active Directory tenant ID. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#tenant_id ActionConnection#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

