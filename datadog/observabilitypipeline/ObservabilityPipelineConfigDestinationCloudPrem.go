// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationCloudPrem struct {
	// Name of the environment variable or secret that holds the endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#endpoint_url_key ObservabilityPipeline#endpoint_url_key}
	EndpointUrlKey *string `field:"optional" json:"endpointUrlKey" yaml:"endpointUrlKey"`
}

