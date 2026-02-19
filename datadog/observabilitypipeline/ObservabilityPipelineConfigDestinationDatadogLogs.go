// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationDatadogLogs struct {
	// routes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#routes ObservabilityPipeline#routes}
	Routes interface{} `field:"optional" json:"routes" yaml:"routes"`
}

