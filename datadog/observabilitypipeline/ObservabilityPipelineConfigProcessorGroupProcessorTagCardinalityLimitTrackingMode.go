// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitTrackingMode struct {
	// The cardinality tracking algorithm to use. One of `exact_fingerprint`, `probabilistic`. Valid values are `exact_fingerprint`, `probabilistic`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#mode ObservabilityPipeline#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

