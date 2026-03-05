// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourceOpentelemetry struct {
	// Environment variable name containing the gRPC server address for receiving OTLP data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#grpc_address_key ObservabilityPipeline#grpc_address_key}
	GrpcAddressKey *string `field:"optional" json:"grpcAddressKey" yaml:"grpcAddressKey"`
	// Environment variable name containing the HTTP server address for receiving OTLP data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#http_address_key ObservabilityPipeline#http_address_key}
	HttpAddressKey *string `field:"optional" json:"httpAddressKey" yaml:"httpAddressKey"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#tls ObservabilityPipeline#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

