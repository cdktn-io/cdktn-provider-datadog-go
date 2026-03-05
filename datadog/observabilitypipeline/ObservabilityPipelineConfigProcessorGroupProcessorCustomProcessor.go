// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorCustomProcessor struct {
	// remap block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#remap ObservabilityPipeline#remap}
	Remap interface{} `field:"optional" json:"remap" yaml:"remap"`
}

