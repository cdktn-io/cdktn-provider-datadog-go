// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationKafka struct {
	// Encoding format for log events. Valid values are `json`, `raw_message`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#encoding ObservabilityPipeline#encoding}
	Encoding *string `field:"required" json:"encoding" yaml:"encoding"`
	// The Kafka topic name to publish logs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#topic ObservabilityPipeline#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// Compression codec for Kafka messages. Valid values are `none`, `gzip`, `snappy`, `lz4`, `zstd`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#compression ObservabilityPipeline#compression}
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// The field name to use for Kafka message headers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#headers_key ObservabilityPipeline#headers_key}
	HeadersKey *string `field:"optional" json:"headersKey" yaml:"headersKey"`
	// The field name to use as the Kafka message key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#key_field ObservabilityPipeline#key_field}
	KeyField *string `field:"optional" json:"keyField" yaml:"keyField"`
	// librdkafka_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#librdkafka_option ObservabilityPipeline#librdkafka_option}
	LibrdkafkaOption interface{} `field:"optional" json:"librdkafkaOption" yaml:"librdkafkaOption"`
	// Maximum time in milliseconds to wait for message delivery confirmation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#message_timeout_ms ObservabilityPipeline#message_timeout_ms}
	MessageTimeoutMs *float64 `field:"optional" json:"messageTimeoutMs" yaml:"messageTimeoutMs"`
	// Duration in seconds for the rate limit window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#rate_limit_duration_secs ObservabilityPipeline#rate_limit_duration_secs}
	RateLimitDurationSecs *float64 `field:"optional" json:"rateLimitDurationSecs" yaml:"rateLimitDurationSecs"`
	// Maximum number of messages allowed per rate limit duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#rate_limit_num ObservabilityPipeline#rate_limit_num}
	RateLimitNum *float64 `field:"optional" json:"rateLimitNum" yaml:"rateLimitNum"`
	// sasl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#sasl ObservabilityPipeline#sasl}
	Sasl interface{} `field:"optional" json:"sasl" yaml:"sasl"`
	// Socket timeout in milliseconds for network requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#socket_timeout_ms ObservabilityPipeline#socket_timeout_ms}
	SocketTimeoutMs *float64 `field:"optional" json:"socketTimeoutMs" yaml:"socketTimeoutMs"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#tls ObservabilityPipeline#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

