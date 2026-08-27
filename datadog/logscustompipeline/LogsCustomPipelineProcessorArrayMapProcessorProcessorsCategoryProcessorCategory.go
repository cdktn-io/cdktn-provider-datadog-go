// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessorArrayMapProcessorProcessorsCategoryProcessorCategory struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/logs_custom_pipeline#filter LogsCustomPipeline#filter}
	Filter *LogsCustomPipelineProcessorArrayMapProcessorProcessorsCategoryProcessorCategoryFilter `field:"required" json:"filter" yaml:"filter"`
	// Name of the category.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/logs_custom_pipeline#name LogsCustomPipeline#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

