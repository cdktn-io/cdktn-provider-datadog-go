// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationSplunkHec struct {
	// Encoding format for log events. Valid values: `json`, `raw_message`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#encoding ObservabilityPipeline#encoding}
	Encoding *string `field:"required" json:"encoding" yaml:"encoding"`
	// If `true`, Splunk tries to extract timestamps from incoming log events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#auto_extract_timestamp ObservabilityPipeline#auto_extract_timestamp}
	AutoExtractTimestamp interface{} `field:"optional" json:"autoExtractTimestamp" yaml:"autoExtractTimestamp"`
	// Optional name of the Splunk index where logs are written.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#index ObservabilityPipeline#index}
	Index *string `field:"optional" json:"index" yaml:"index"`
	// The Splunk sourcetype to assign to log events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#sourcetype ObservabilityPipeline#sourcetype}
	Sourcetype *string `field:"optional" json:"sourcetype" yaml:"sourcetype"`
}

