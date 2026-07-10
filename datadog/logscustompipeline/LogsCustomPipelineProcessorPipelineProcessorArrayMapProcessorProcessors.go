// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessorPipelineProcessorArrayMapProcessorProcessors struct {
	// arithmetic_processor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/logs_custom_pipeline#arithmetic_processor LogsCustomPipeline#arithmetic_processor}
	ArithmeticProcessor *LogsCustomPipelineProcessorPipelineProcessorArrayMapProcessorProcessorsArithmeticProcessor `field:"optional" json:"arithmeticProcessor" yaml:"arithmeticProcessor"`
	// attribute_remapper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/logs_custom_pipeline#attribute_remapper LogsCustomPipeline#attribute_remapper}
	AttributeRemapper *LogsCustomPipelineProcessorPipelineProcessorArrayMapProcessorProcessorsAttributeRemapper `field:"optional" json:"attributeRemapper" yaml:"attributeRemapper"`
	// category_processor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/logs_custom_pipeline#category_processor LogsCustomPipeline#category_processor}
	CategoryProcessor *LogsCustomPipelineProcessorPipelineProcessorArrayMapProcessorProcessorsCategoryProcessor `field:"optional" json:"categoryProcessor" yaml:"categoryProcessor"`
	// string_builder_processor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/logs_custom_pipeline#string_builder_processor LogsCustomPipeline#string_builder_processor}
	StringBuilderProcessor *LogsCustomPipelineProcessorPipelineProcessorArrayMapProcessorProcessorsStringBuilderProcessor `field:"optional" json:"stringBuilderProcessor" yaml:"stringBuilderProcessor"`
}

