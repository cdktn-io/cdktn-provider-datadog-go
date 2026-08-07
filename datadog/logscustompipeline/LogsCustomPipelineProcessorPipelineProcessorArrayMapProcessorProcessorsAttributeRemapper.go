// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessorPipelineProcessorArrayMapProcessorProcessorsAttributeRemapper struct {
	// List of source attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/logs_custom_pipeline#sources LogsCustomPipeline#sources}
	Sources *[]*string `field:"required" json:"sources" yaml:"sources"`
	// Target attribute path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/logs_custom_pipeline#target LogsCustomPipeline#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// Name of the sub-processor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/logs_custom_pipeline#name LogsCustomPipeline#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Override the target element if already set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/logs_custom_pipeline#override_on_conflict LogsCustomPipeline#override_on_conflict}
	OverrideOnConflict interface{} `field:"optional" json:"overrideOnConflict" yaml:"overrideOnConflict"`
	// Remove or preserve the remapped source element. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/logs_custom_pipeline#preserve_source LogsCustomPipeline#preserve_source}
	PreserveSource interface{} `field:"optional" json:"preserveSource" yaml:"preserveSource"`
	// If the target type is an attribute, cast the value to a new type (auto, string, integer, double).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/logs_custom_pipeline#target_format LogsCustomPipeline#target_format}
	TargetFormat *string `field:"optional" json:"targetFormat" yaml:"targetFormat"`
}

