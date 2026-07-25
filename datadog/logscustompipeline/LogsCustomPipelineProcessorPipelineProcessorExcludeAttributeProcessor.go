// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessorPipelineProcessorExcludeAttributeProcessor struct {
	// Name of the log attribute to remove from the log event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/logs_custom_pipeline#attribute_to_exclude LogsCustomPipeline#attribute_to_exclude}
	AttributeToExclude *string `field:"required" json:"attributeToExclude" yaml:"attributeToExclude"`
	// If the processor is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/logs_custom_pipeline#is_enabled LogsCustomPipeline#is_enabled}
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Name of the processor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/logs_custom_pipeline#name LogsCustomPipeline#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

