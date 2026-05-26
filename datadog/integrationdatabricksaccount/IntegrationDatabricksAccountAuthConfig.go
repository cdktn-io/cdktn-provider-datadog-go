// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationdatabricksaccount


type IntegrationDatabricksAccountAuthConfig struct {
	// oauth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#oauth IntegrationDatabricksAccount#oauth}
	Oauth *IntegrationDatabricksAccountAuthConfigOauth `field:"optional" json:"oauth" yaml:"oauth"`
	// pat block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#pat IntegrationDatabricksAccount#pat}
	Pat *IntegrationDatabricksAccountAuthConfigPat `field:"optional" json:"pat" yaml:"pat"`
}

