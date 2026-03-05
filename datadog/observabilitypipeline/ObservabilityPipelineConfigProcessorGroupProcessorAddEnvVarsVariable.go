// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorAddEnvVarsVariable struct {
	// The target field in the log event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#field ObservabilityPipeline#field}
	Field *string `field:"required" json:"field" yaml:"field"`
	// The name of the environment variable to read.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#name ObservabilityPipeline#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

