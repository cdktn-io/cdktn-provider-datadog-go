// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorRemoveFields struct {
	// List of fields to remove from the events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/observability_pipeline#fields ObservabilityPipeline#fields}
	Fields *[]*string `field:"required" json:"fields" yaml:"fields"`
}

