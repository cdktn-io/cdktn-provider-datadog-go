// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourceSyslogNgTls struct {
	// Path to the TLS server certificate file used to identify the pipeline component to connecting clients.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#crt_file ObservabilityPipeline#crt_file}
	CrtFile *string `field:"required" json:"crtFile" yaml:"crtFile"`
	// Path to the Certificate Authority (CA) file used to validate connecting clients' TLS certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#ca_file ObservabilityPipeline#ca_file}
	CaFile *string `field:"optional" json:"caFile" yaml:"caFile"`
	// Path to the private key file associated with the TLS server certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#key_file ObservabilityPipeline#key_file}
	KeyFile *string `field:"optional" json:"keyFile" yaml:"keyFile"`
	// Name of the environment variable or secret that holds the passphrase for the private key file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#key_pass_key ObservabilityPipeline#key_pass_key}
	KeyPassKey *string `field:"optional" json:"keyPassKey" yaml:"keyPassKey"`
	// When `true`, requires client connections to present a valid certificate, enabling mutual TLS authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#verify_certificate ObservabilityPipeline#verify_certificate}
	VerifyCertificate interface{} `field:"optional" json:"verifyCertificate" yaml:"verifyCertificate"`
}

