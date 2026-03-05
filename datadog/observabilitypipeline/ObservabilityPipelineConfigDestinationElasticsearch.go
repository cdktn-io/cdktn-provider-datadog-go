// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationElasticsearch struct {
	// The Elasticsearch API version to use. Set to `auto` to auto-detect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#api_version ObservabilityPipeline#api_version}
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#auth ObservabilityPipeline#auth}
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// buffer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#buffer ObservabilityPipeline#buffer}
	Buffer interface{} `field:"optional" json:"buffer" yaml:"buffer"`
	// The index or datastream to write logs to in Elasticsearch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#bulk_index ObservabilityPipeline#bulk_index}
	BulkIndex *string `field:"optional" json:"bulkIndex" yaml:"bulkIndex"`
	// data_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#data_stream ObservabilityPipeline#data_stream}
	DataStream interface{} `field:"optional" json:"dataStream" yaml:"dataStream"`
	// Name of the environment variable or secret that holds the Elasticsearch endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#endpoint_url_key ObservabilityPipeline#endpoint_url_key}
	EndpointUrlKey *string `field:"optional" json:"endpointUrlKey" yaml:"endpointUrlKey"`
}

