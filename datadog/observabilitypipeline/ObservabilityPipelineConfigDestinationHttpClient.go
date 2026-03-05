// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationHttpClient struct {
	// Encoding format for events. Valid values are `json`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#encoding ObservabilityPipeline#encoding}
	Encoding *string `field:"required" json:"encoding" yaml:"encoding"`
	// HTTP authentication strategy. Valid values are `none`, `basic`, `bearer`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#auth_strategy ObservabilityPipeline#auth_strategy}
	AuthStrategy *string `field:"optional" json:"authStrategy" yaml:"authStrategy"`
	// compression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#compression ObservabilityPipeline#compression}
	Compression interface{} `field:"optional" json:"compression" yaml:"compression"`
	// Name of the environment variable or secret that holds the password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#password_key ObservabilityPipeline#password_key}
	PasswordKey *string `field:"optional" json:"passwordKey" yaml:"passwordKey"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#tls ObservabilityPipeline#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
	// Name of the environment variable or secret that holds the authentication token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#token_key ObservabilityPipeline#token_key}
	TokenKey *string `field:"optional" json:"tokenKey" yaml:"tokenKey"`
	// Name of the environment variable or secret that holds the request URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#uri_key ObservabilityPipeline#uri_key}
	UriKey *string `field:"optional" json:"uriKey" yaml:"uriKey"`
	// Name of the environment variable or secret that holds the username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#username_key ObservabilityPipeline#username_key}
	UsernameKey *string `field:"optional" json:"usernameKey" yaml:"usernameKey"`
}

