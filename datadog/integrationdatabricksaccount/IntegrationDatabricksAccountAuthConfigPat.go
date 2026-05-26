// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationdatabricksaccount


type IntegrationDatabricksAccountAuthConfigPat struct {
	// Databricks Personal Access Token (PAT).
	//
	// Generate from Settings > Developer > Access tokens. This value is write-only; changes made outside of Terraform will not be drift-detected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#token IntegrationDatabricksAccount#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
}

