// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationClickhouseBatchEncoding struct {
	// Batch encoding codec. Must be `arrow_stream`. Valid values are `arrow_stream`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#codec ObservabilityPipeline#codec}
	Codec *string `field:"required" json:"codec" yaml:"codec"`
	// If `true`, allows null values for non-nullable fields in the ClickHouse schema. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/observability_pipeline#allow_nullable_fields ObservabilityPipeline#allow_nullable_fields}
	AllowNullableFields interface{} `field:"optional" json:"allowNullableFields" yaml:"allowNullableFields"`
}

