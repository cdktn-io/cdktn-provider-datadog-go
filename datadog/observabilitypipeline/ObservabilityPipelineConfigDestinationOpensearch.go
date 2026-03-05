// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationOpensearch struct {
	// buffer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#buffer ObservabilityPipeline#buffer}
	Buffer interface{} `field:"optional" json:"buffer" yaml:"buffer"`
	// The index or datastream to write logs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#bulk_index ObservabilityPipeline#bulk_index}
	BulkIndex *string `field:"optional" json:"bulkIndex" yaml:"bulkIndex"`
	// data_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#data_stream ObservabilityPipeline#data_stream}
	DataStream interface{} `field:"optional" json:"dataStream" yaml:"dataStream"`
}

