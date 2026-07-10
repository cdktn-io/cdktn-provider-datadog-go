// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourceHttpServerValidToken struct {
	// Name of the environment variable or secret that holds the expected token value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#token_key ObservabilityPipeline#token_key}
	TokenKey *string `field:"required" json:"tokenKey" yaml:"tokenKey"`
	// Whether this token is currently accepted. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#enabled ObservabilityPipeline#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// field_to_add block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#field_to_add ObservabilityPipeline#field_to_add}
	FieldToAdd interface{} `field:"optional" json:"fieldToAdd" yaml:"fieldToAdd"`
	// path_to_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#path_to_token ObservabilityPipeline#path_to_token}
	PathToToken interface{} `field:"optional" json:"pathToToken" yaml:"pathToToken"`
}

