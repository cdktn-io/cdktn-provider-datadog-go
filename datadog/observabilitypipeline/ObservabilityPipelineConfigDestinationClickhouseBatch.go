// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationClickhouseBatch struct {
	// Maximum number of events per batch. Value must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#max_events ObservabilityPipeline#max_events}
	MaxEvents *float64 `field:"optional" json:"maxEvents" yaml:"maxEvents"`
	// Maximum time in seconds before a partial batch is flushed. Value must be between 1 and 65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#timeout_secs ObservabilityPipeline#timeout_secs}
	TimeoutSecs *float64 `field:"optional" json:"timeoutSecs" yaml:"timeoutSecs"`
}

