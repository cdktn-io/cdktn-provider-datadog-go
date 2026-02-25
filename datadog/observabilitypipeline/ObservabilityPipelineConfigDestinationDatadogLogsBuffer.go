// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationDatadogLogsBuffer struct {
	// disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/observability_pipeline#disk ObservabilityPipeline#disk}
	Disk interface{} `field:"optional" json:"disk" yaml:"disk"`
	// memory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/observability_pipeline#memory ObservabilityPipeline#memory}
	Memory interface{} `field:"optional" json:"memory" yaml:"memory"`
}

