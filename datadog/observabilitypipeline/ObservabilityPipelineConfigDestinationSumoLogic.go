// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationSumoLogic struct {
	// buffer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#buffer ObservabilityPipeline#buffer}
	Buffer interface{} `field:"optional" json:"buffer" yaml:"buffer"`
	// The output encoding format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#encoding ObservabilityPipeline#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Name of the environment variable or secret that holds the Sumo Logic endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#endpoint_url_key ObservabilityPipeline#endpoint_url_key}
	EndpointUrlKey *string `field:"optional" json:"endpointUrlKey" yaml:"endpointUrlKey"`
	// header_custom_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#header_custom_field ObservabilityPipeline#header_custom_field}
	HeaderCustomField interface{} `field:"optional" json:"headerCustomField" yaml:"headerCustomField"`
	// Optional override for the host name header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#header_host_name ObservabilityPipeline#header_host_name}
	HeaderHostName *string `field:"optional" json:"headerHostName" yaml:"headerHostName"`
	// Optional override for the source category header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#header_source_category ObservabilityPipeline#header_source_category}
	HeaderSourceCategory *string `field:"optional" json:"headerSourceCategory" yaml:"headerSourceCategory"`
	// Optional override for the source name header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#header_source_name ObservabilityPipeline#header_source_name}
	HeaderSourceName *string `field:"optional" json:"headerSourceName" yaml:"headerSourceName"`
}

