// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessorArrayMapProcessorProcessorsStringBuilderProcessor struct {
	// Target attribute path for the result.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/logs_custom_pipeline#target LogsCustomPipeline#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// Formula with one or more attributes and raw text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/logs_custom_pipeline#template LogsCustomPipeline#template}
	Template *string `field:"required" json:"template" yaml:"template"`
	// Replace missing attributes with an empty string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/logs_custom_pipeline#is_replace_missing LogsCustomPipeline#is_replace_missing}
	IsReplaceMissing interface{} `field:"optional" json:"isReplaceMissing" yaml:"isReplaceMissing"`
	// Name of the sub-processor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/logs_custom_pipeline#name LogsCustomPipeline#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

