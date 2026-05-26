// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationdatabricksaccount


type IntegrationDatabricksAccountPrivateActionRunnerConfiguration struct {
	// Private Action Runner connection ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#connection_id IntegrationDatabricksAccount#connection_id}
	ConnectionId *string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// Path to the stored secret holding Databricks credentials inside the Private Action Runner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#secret_path IntegrationDatabricksAccount#secret_path}
	SecretPath *string `field:"optional" json:"secretPath" yaml:"secretPath"`
	// Service Account UUID used to execute Private Action Runner actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#user_uuid IntegrationDatabricksAccount#user_uuid}
	UserUuid *string `field:"optional" json:"userUuid" yaml:"userUuid"`
}

