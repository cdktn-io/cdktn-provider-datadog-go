// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatadogorganizationsettings


type DataDatadogOrganizationSettingsSettings struct {
	// saml block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/data-sources/organization_settings#saml DataDatadogOrganizationSettings#saml}
	Saml interface{} `field:"optional" json:"saml" yaml:"saml"`
	// saml_autocreate_users_domains block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/data-sources/organization_settings#saml_autocreate_users_domains DataDatadogOrganizationSettings#saml_autocreate_users_domains}
	SamlAutocreateUsersDomains interface{} `field:"optional" json:"samlAutocreateUsersDomains" yaml:"samlAutocreateUsersDomains"`
	// saml_idp_initiated_login block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/data-sources/organization_settings#saml_idp_initiated_login DataDatadogOrganizationSettings#saml_idp_initiated_login}
	SamlIdpInitiatedLogin interface{} `field:"optional" json:"samlIdpInitiatedLogin" yaml:"samlIdpInitiatedLogin"`
	// saml_strict_mode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/data-sources/organization_settings#saml_strict_mode DataDatadogOrganizationSettings#saml_strict_mode}
	SamlStrictMode interface{} `field:"optional" json:"samlStrictMode" yaml:"samlStrictMode"`
}

