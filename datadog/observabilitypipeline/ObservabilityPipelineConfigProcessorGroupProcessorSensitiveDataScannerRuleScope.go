// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorSensitiveDataScannerRuleScope struct {
	// Scan all fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/observability_pipeline#all ObservabilityPipeline#all}
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// exclude block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/observability_pipeline#exclude ObservabilityPipeline#exclude}
	Exclude interface{} `field:"optional" json:"exclude" yaml:"exclude"`
	// include block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/observability_pipeline#include ObservabilityPipeline#include}
	Include interface{} `field:"optional" json:"include" yaml:"include"`
}

