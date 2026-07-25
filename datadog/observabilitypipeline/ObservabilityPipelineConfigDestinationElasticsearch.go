// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationElasticsearch struct {
	// The Elasticsearch API version to use. Set to `auto` to auto-detect. Valid values are `auto`, `v6`, `v7`, `v8`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#api_version ObservabilityPipeline#api_version}
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#auth ObservabilityPipeline#auth}
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// buffer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#buffer ObservabilityPipeline#buffer}
	Buffer interface{} `field:"optional" json:"buffer" yaml:"buffer"`
	// The name of the index to write events to in Elasticsearch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#bulk_index ObservabilityPipeline#bulk_index}
	BulkIndex *string `field:"optional" json:"bulkIndex" yaml:"bulkIndex"`
	// compression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#compression ObservabilityPipeline#compression}
	Compression interface{} `field:"optional" json:"compression" yaml:"compression"`
	// data_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#data_stream ObservabilityPipeline#data_stream}
	DataStream interface{} `field:"optional" json:"dataStream" yaml:"dataStream"`
	// Name of the environment variable or secret that holds the Elasticsearch endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#endpoint_url_key ObservabilityPipeline#endpoint_url_key}
	EndpointUrlKey *string `field:"optional" json:"endpointUrlKey" yaml:"endpointUrlKey"`
	// The name of the field used as the document ID in Elasticsearch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#id_key ObservabilityPipeline#id_key}
	IdKey *string `field:"optional" json:"idKey" yaml:"idKey"`
	// The name of an Elasticsearch ingest pipeline to apply to events before indexing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#pipeline ObservabilityPipeline#pipeline}
	Pipeline *string `field:"optional" json:"pipeline" yaml:"pipeline"`
	// When `true`, retries failed partial bulk requests when some events in a batch fail while others succeed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#request_retry_partial ObservabilityPipeline#request_retry_partial}
	RequestRetryPartial interface{} `field:"optional" json:"requestRetryPartial" yaml:"requestRetryPartial"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#tls ObservabilityPipeline#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

