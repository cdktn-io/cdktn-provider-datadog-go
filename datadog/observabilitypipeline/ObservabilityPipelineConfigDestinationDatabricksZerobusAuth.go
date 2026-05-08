// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationDatabricksZerobusAuth struct {
	// The OAuth client ID used to authenticate with Databricks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/observability_pipeline#client_id ObservabilityPipeline#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The name of the secret or environment variable holding the OAuth client secret. Defaults to `DESTINATION_DATABRICKS_ZEROBUS_OAUTH_CLIENT_SECRET`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/observability_pipeline#client_secret_key ObservabilityPipeline#client_secret_key}
	ClientSecretKey *string `field:"optional" json:"clientSecretKey" yaml:"clientSecretKey"`
}

