// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourceKafkaSasl struct {
	// SASL mechanism to use (e.g., PLAIN, SCRAM-SHA-256, SCRAM-SHA-512). Valid values are `PLAIN`, `SCRAM-SHA-256`, `SCRAM-SHA-512`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#mechanism ObservabilityPipeline#mechanism}
	Mechanism *string `field:"required" json:"mechanism" yaml:"mechanism"`
	// Name of the environment variable or secret that holds the SASL password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#password_key ObservabilityPipeline#password_key}
	PasswordKey *string `field:"optional" json:"passwordKey" yaml:"passwordKey"`
	// Name of the environment variable or secret that holds the SASL username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#username_key ObservabilityPipeline#username_key}
	UsernameKey *string `field:"optional" json:"usernameKey" yaml:"usernameKey"`
}

