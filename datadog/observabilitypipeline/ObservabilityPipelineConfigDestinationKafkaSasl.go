// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationKafkaSasl struct {
	// SASL authentication mechanism. Valid values are `PLAIN`, `SCRAM-SHA-256`, `SCRAM-SHA-512`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#mechanism ObservabilityPipeline#mechanism}
	Mechanism *string `field:"required" json:"mechanism" yaml:"mechanism"`
}

