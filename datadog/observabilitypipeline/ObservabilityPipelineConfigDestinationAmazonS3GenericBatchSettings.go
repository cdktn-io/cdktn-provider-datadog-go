// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationAmazonS3GenericBatchSettings struct {
	// Maximum batch size in bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#batch_size ObservabilityPipeline#batch_size}
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// Maximum number of seconds to wait before flushing the batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#timeout_secs ObservabilityPipeline#timeout_secs}
	TimeoutSecs *float64 `field:"optional" json:"timeoutSecs" yaml:"timeoutSecs"`
}

