// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessorPipelineProcessorArrayProcessorOperationKeyValue struct {
	// Key of the attribute in each array element that holds the name to use for the extracted attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/logs_custom_pipeline#key_to_extract LogsCustomPipeline#key_to_extract}
	KeyToExtract *string `field:"required" json:"keyToExtract" yaml:"keyToExtract"`
	// Attribute path of the array to extract key-value pairs from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/logs_custom_pipeline#source LogsCustomPipeline#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// Key of the attribute in each array element that holds the value to use for the extracted attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/logs_custom_pipeline#value_to_extract LogsCustomPipeline#value_to_extract}
	ValueToExtract *string `field:"required" json:"valueToExtract" yaml:"valueToExtract"`
	// Whether to override the target element if it's already set. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/logs_custom_pipeline#override_on_conflict LogsCustomPipeline#override_on_conflict}
	OverrideOnConflict interface{} `field:"optional" json:"overrideOnConflict" yaml:"overrideOnConflict"`
	// Attribute that receives the extracted key-value pairs.
	//
	// If not specified, the extracted attributes are added at the root level of the log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/logs_custom_pipeline#target LogsCustomPipeline#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
}

