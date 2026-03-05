// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourceHttpServer struct {
	// HTTP authentication method. Valid values are `none`, `plain`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#auth_strategy ObservabilityPipeline#auth_strategy}
	AuthStrategy *string `field:"required" json:"authStrategy" yaml:"authStrategy"`
	// The decoding format used to interpret incoming logs. Valid values are `json`, `gelf`, `syslog`, `bytes`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#decoding ObservabilityPipeline#decoding}
	Decoding *string `field:"required" json:"decoding" yaml:"decoding"`
	// Name of the environment variable or secret that holds the listen address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#address_key ObservabilityPipeline#address_key}
	AddressKey *string `field:"optional" json:"addressKey" yaml:"addressKey"`
	// Name of the environment variable or secret that holds the password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#password_key ObservabilityPipeline#password_key}
	PasswordKey *string `field:"optional" json:"passwordKey" yaml:"passwordKey"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#tls ObservabilityPipeline#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
	// Name of the environment variable or secret that holds the username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#username_key ObservabilityPipeline#username_key}
	UsernameKey *string `field:"optional" json:"usernameKey" yaml:"usernameKey"`
}

