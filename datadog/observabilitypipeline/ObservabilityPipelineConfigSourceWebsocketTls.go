// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourceWebsocketTls struct {
	// The TLS mode.
	//
	// Use `enabled` for server-only TLS, or `with_client_cert` for mutual TLS with a client certificate. Valid values are `enabled`, `with_client_cert`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#mode ObservabilityPipeline#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Path to the Certificate Authority (CA) file used to validate the server's TLS certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#ca_file ObservabilityPipeline#ca_file}
	CaFile *string `field:"optional" json:"caFile" yaml:"caFile"`
	// Path to the client certificate file. Required when `mode` is `with_client_cert`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#crt_file ObservabilityPipeline#crt_file}
	CrtFile *string `field:"optional" json:"crtFile" yaml:"crtFile"`
	// Path to the private key file associated with the client certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#key_file ObservabilityPipeline#key_file}
	KeyFile *string `field:"optional" json:"keyFile" yaml:"keyFile"`
	// Name of the environment variable or secret that holds the passphrase for the private key file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#key_pass_key ObservabilityPipeline#key_pass_key}
	KeyPassKey *string `field:"optional" json:"keyPassKey" yaml:"keyPassKey"`
}

