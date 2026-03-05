// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorEnrichmentTableReferenceTable struct {
	// Path to the field in the log event to match against the reference table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#key_field ObservabilityPipeline#key_field}
	KeyField *string `field:"required" json:"keyField" yaml:"keyField"`
	// The unique identifier of the reference table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#table_id ObservabilityPipeline#table_id}
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
	// Name of the environment variable or secret that holds the Datadog application key for the reference table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#app_key_key ObservabilityPipeline#app_key_key}
	AppKeyKey *string `field:"optional" json:"appKeyKey" yaml:"appKeyKey"`
	// List of column names to include from the reference table. If not provided, all columns are included.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#columns ObservabilityPipeline#columns}
	Columns *[]*string `field:"optional" json:"columns" yaml:"columns"`
}

