// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationSplunkHecMetrics struct {
	// buffer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/observability_pipeline#buffer ObservabilityPipeline#buffer}
	Buffer interface{} `field:"optional" json:"buffer" yaml:"buffer"`
	// Compression algorithm applied when sending metrics to Splunk HEC. Valid values are `none`, `gzip`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/observability_pipeline#compression ObservabilityPipeline#compression}
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// Optional default namespace for metrics sent to Splunk HEC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/observability_pipeline#default_namespace ObservabilityPipeline#default_namespace}
	DefaultNamespace *string `field:"optional" json:"defaultNamespace" yaml:"defaultNamespace"`
	// Name of the environment variable or secret that holds the Splunk HEC endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/observability_pipeline#endpoint_url_key ObservabilityPipeline#endpoint_url_key}
	EndpointUrlKey *string `field:"optional" json:"endpointUrlKey" yaml:"endpointUrlKey"`
	// Optional name of the Splunk index where metrics are written.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/observability_pipeline#index ObservabilityPipeline#index}
	Index *string `field:"optional" json:"index" yaml:"index"`
	// The Splunk source field value for metric events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/observability_pipeline#source ObservabilityPipeline#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The Splunk sourcetype to assign to metric events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/observability_pipeline#sourcetype ObservabilityPipeline#sourcetype}
	Sourcetype *string `field:"optional" json:"sourcetype" yaml:"sourcetype"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/observability_pipeline#tls ObservabilityPipeline#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
	// Name of the environment variable or secret that holds the Splunk HEC token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/observability_pipeline#token_key ObservabilityPipeline#token_key}
	TokenKey *string `field:"optional" json:"tokenKey" yaml:"tokenKey"`
}

