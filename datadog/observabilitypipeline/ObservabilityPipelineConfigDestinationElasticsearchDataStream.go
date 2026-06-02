// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationElasticsearchDataStream struct {
	// When `true`, automatically routes events to the appropriate data stream based on the event content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/observability_pipeline#auto_routing ObservabilityPipeline#auto_routing}
	AutoRouting interface{} `field:"optional" json:"autoRouting" yaml:"autoRouting"`
	// The data stream dataset. This groups events by their source or application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/observability_pipeline#dataset ObservabilityPipeline#dataset}
	Dataset *string `field:"optional" json:"dataset" yaml:"dataset"`
	// The data stream type. This determines how events are categorized within the data stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/observability_pipeline#dtype ObservabilityPipeline#dtype}
	Dtype *string `field:"optional" json:"dtype" yaml:"dtype"`
	// The data stream namespace. This separates events into different environments or domains.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/observability_pipeline#namespace ObservabilityPipeline#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// When `true`, synchronizes data stream fields with the Elasticsearch index mapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/observability_pipeline#sync_fields ObservabilityPipeline#sync_fields}
	SyncFields interface{} `field:"optional" json:"syncFields" yaml:"syncFields"`
}

