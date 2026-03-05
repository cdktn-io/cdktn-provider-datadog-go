// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationCrowdstrikeNextGenSiem struct {
	// Encoding format for log events. Valid values are `json`, `raw_message`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#encoding ObservabilityPipeline#encoding}
	Encoding *string `field:"required" json:"encoding" yaml:"encoding"`
	// buffer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#buffer ObservabilityPipeline#buffer}
	Buffer interface{} `field:"optional" json:"buffer" yaml:"buffer"`
	// compression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#compression ObservabilityPipeline#compression}
	Compression interface{} `field:"optional" json:"compression" yaml:"compression"`
	// Name of the environment variable or secret that holds the endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#endpoint_url_key ObservabilityPipeline#endpoint_url_key}
	EndpointUrlKey *string `field:"optional" json:"endpointUrlKey" yaml:"endpointUrlKey"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#tls ObservabilityPipeline#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
	// Name of the environment variable or secret that holds the authentication token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#token_key ObservabilityPipeline#token_key}
	TokenKey *string `field:"optional" json:"tokenKey" yaml:"tokenKey"`
}

