// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourceSplunkHecValidTokenFieldToAdd struct {
	// The metadata field name to add to incoming events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/observability_pipeline#key ObservabilityPipeline#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The metadata field value to add to incoming events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/observability_pipeline#value ObservabilityPipeline#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

