// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationSocketFramingCharacterDelimited struct {
	// A single ASCII character used as a delimiter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#delimiter ObservabilityPipeline#delimiter}
	Delimiter *string `field:"required" json:"delimiter" yaml:"delimiter"`
}

