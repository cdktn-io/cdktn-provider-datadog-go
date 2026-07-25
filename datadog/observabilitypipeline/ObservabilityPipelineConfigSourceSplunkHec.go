// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourceSplunkHec struct {
	// Name of the environment variable or secret that holds the listen address for the HEC API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#address_key ObservabilityPipeline#address_key}
	AddressKey *string `field:"optional" json:"addressKey" yaml:"addressKey"`
	// When `true`, the Splunk HEC token from the incoming request is stored in the event, allowing downstream components to forward it to other Splunk HEC destinations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#store_hec_token ObservabilityPipeline#store_hec_token}
	StoreHecToken interface{} `field:"optional" json:"storeHecToken" yaml:"storeHecToken"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#tls ObservabilityPipeline#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
	// valid_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#valid_token ObservabilityPipeline#valid_token}
	ValidToken interface{} `field:"optional" json:"validToken" yaml:"validToken"`
}

