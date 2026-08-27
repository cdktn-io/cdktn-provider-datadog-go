// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorRenameMetricTagsTag struct {
	// The new tag key to assign in place of the original.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/observability_pipeline#rename_to ObservabilityPipeline#rename_to}
	RenameTo *string `field:"required" json:"renameTo" yaml:"renameTo"`
	// The original tag key on the metric event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/observability_pipeline#tag ObservabilityPipeline#tag}
	Tag *string `field:"required" json:"tag" yaml:"tag"`
}

