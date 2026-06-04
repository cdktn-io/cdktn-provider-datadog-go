// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationSplunkHec struct {
	// Encoding format for log events. Valid values are `json`, `raw_message`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/observability_pipeline#encoding ObservabilityPipeline#encoding}
	Encoding *string `field:"required" json:"encoding" yaml:"encoding"`
	// If `true`, Splunk tries to extract timestamps from incoming log events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/observability_pipeline#auto_extract_timestamp ObservabilityPipeline#auto_extract_timestamp}
	AutoExtractTimestamp interface{} `field:"optional" json:"autoExtractTimestamp" yaml:"autoExtractTimestamp"`
	// buffer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/observability_pipeline#buffer ObservabilityPipeline#buffer}
	Buffer interface{} `field:"optional" json:"buffer" yaml:"buffer"`
	// Name of the environment variable or secret that holds the Splunk HEC endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/observability_pipeline#endpoint_url_key ObservabilityPipeline#endpoint_url_key}
	EndpointUrlKey *string `field:"optional" json:"endpointUrlKey" yaml:"endpointUrlKey"`
	// Optional name of the Splunk index where logs are written.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/observability_pipeline#index ObservabilityPipeline#index}
	Index *string `field:"optional" json:"index" yaml:"index"`
	// List of log field names to send as indexed fields to Splunk HEC. Available only when `encoding` is `json`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/observability_pipeline#indexed_fields ObservabilityPipeline#indexed_fields}
	IndexedFields *[]*string `field:"optional" json:"indexedFields" yaml:"indexedFields"`
	// The Splunk sourcetype to assign to log events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/observability_pipeline#sourcetype ObservabilityPipeline#sourcetype}
	Sourcetype *string `field:"optional" json:"sourcetype" yaml:"sourcetype"`
	// Name of the environment variable or secret that holds the Splunk HEC token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/observability_pipeline#token_key ObservabilityPipeline#token_key}
	TokenKey *string `field:"optional" json:"tokenKey" yaml:"tokenKey"`
	// Controls how the Splunk HEC token is supplied.
	//
	// Use `custom` to provide a token via `token_key`, or `from_source` to forward the token received from an upstream Splunk HEC source. Valid values are `custom`, `from_source`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/observability_pipeline#token_strategy ObservabilityPipeline#token_strategy}
	TokenStrategy *string `field:"optional" json:"tokenStrategy" yaml:"tokenStrategy"`
}

