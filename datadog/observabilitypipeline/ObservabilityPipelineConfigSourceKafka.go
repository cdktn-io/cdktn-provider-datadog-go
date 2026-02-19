// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourceKafka struct {
	// The Kafka consumer group ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#group_id ObservabilityPipeline#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// A list of Kafka topic names to subscribe to. The source ingests messages from each topic specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#topics ObservabilityPipeline#topics}
	Topics *[]*string `field:"required" json:"topics" yaml:"topics"`
	// librdkafka_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#librdkafka_option ObservabilityPipeline#librdkafka_option}
	LibrdkafkaOption interface{} `field:"optional" json:"librdkafkaOption" yaml:"librdkafkaOption"`
	// sasl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#sasl ObservabilityPipeline#sasl}
	Sasl interface{} `field:"optional" json:"sasl" yaml:"sasl"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#tls ObservabilityPipeline#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

