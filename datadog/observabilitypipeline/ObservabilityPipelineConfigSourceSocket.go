// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourceSocket struct {
	// The protocol used to receive logs. Valid values are `tcp`, `udp`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#mode ObservabilityPipeline#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Name of the environment variable or secret that holds the listen address for the socket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#address_key ObservabilityPipeline#address_key}
	AddressKey *string `field:"optional" json:"addressKey" yaml:"addressKey"`
	// framing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#framing ObservabilityPipeline#framing}
	Framing interface{} `field:"optional" json:"framing" yaml:"framing"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#tls ObservabilityPipeline#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

