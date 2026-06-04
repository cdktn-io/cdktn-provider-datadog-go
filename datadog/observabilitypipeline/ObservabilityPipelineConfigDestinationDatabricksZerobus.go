// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationDatabricksZerobus struct {
	// The name of the Databricks table to ingest logs into.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/observability_pipeline#table_name ObservabilityPipeline#table_name}
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/observability_pipeline#auth ObservabilityPipeline#auth}
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// The name of the secret or environment variable holding the Databricks Zerobus ingestion endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/observability_pipeline#ingestion_endpoint_key ObservabilityPipeline#ingestion_endpoint_key}
	IngestionEndpointKey *string `field:"optional" json:"ingestionEndpointKey" yaml:"ingestionEndpointKey"`
	// The name of the secret or environment variable holding the Databricks Unity Catalog endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/observability_pipeline#unity_catalog_endpoint_key ObservabilityPipeline#unity_catalog_endpoint_key}
	UnityCatalogEndpointKey *string `field:"optional" json:"unityCatalogEndpointKey" yaml:"unityCatalogEndpointKey"`
}

