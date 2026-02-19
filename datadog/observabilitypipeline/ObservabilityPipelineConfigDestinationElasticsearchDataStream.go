// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationElasticsearchDataStream struct {
	// The data stream dataset for your logs. This groups logs by their source or application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#dataset ObservabilityPipeline#dataset}
	Dataset *string `field:"optional" json:"dataset" yaml:"dataset"`
	// The data stream type for your logs. This determines how logs are categorized within the data stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#dtype ObservabilityPipeline#dtype}
	Dtype *string `field:"optional" json:"dtype" yaml:"dtype"`
	// The data stream namespace for your logs. This separates logs into different environments or domains.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#namespace ObservabilityPipeline#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

