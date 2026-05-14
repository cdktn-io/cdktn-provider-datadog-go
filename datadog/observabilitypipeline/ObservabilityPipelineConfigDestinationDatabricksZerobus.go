// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationDatabricksZerobus struct {
	// The Databricks Zerobus ingestion endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/observability_pipeline#ingestion_endpoint ObservabilityPipeline#ingestion_endpoint}
	IngestionEndpoint *string `field:"required" json:"ingestionEndpoint" yaml:"ingestionEndpoint"`
	// The name of the Databricks table to ingest logs into.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/observability_pipeline#table_name ObservabilityPipeline#table_name}
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// The Databricks Unity Catalog endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/observability_pipeline#unity_catalog_endpoint ObservabilityPipeline#unity_catalog_endpoint}
	UnityCatalogEndpoint *string `field:"required" json:"unityCatalogEndpoint" yaml:"unityCatalogEndpoint"`
	// auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/observability_pipeline#auth ObservabilityPipeline#auth}
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
}

