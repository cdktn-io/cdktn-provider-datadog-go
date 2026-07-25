// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationClickhouse struct {
	// Target ClickHouse table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#table ObservabilityPipeline#table}
	Table *string `field:"required" json:"table" yaml:"table"`
	// auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#auth ObservabilityPipeline#auth}
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// batch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#batch ObservabilityPipeline#batch}
	Batch interface{} `field:"optional" json:"batch" yaml:"batch"`
	// batch_encoding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#batch_encoding ObservabilityPipeline#batch_encoding}
	BatchEncoding interface{} `field:"optional" json:"batchEncoding" yaml:"batchEncoding"`
	// buffer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#buffer ObservabilityPipeline#buffer}
	Buffer interface{} `field:"optional" json:"buffer" yaml:"buffer"`
	// compression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#compression ObservabilityPipeline#compression}
	Compression interface{} `field:"optional" json:"compression" yaml:"compression"`
	// Optional name of the ClickHouse database to write to. When omitted, the user's default database is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#database ObservabilityPipeline#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// If `true`, enables flexible DateTime parsing on the server side.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#date_time_best_effort ObservabilityPipeline#date_time_best_effort}
	DateTimeBestEffort interface{} `field:"optional" json:"dateTimeBestEffort" yaml:"dateTimeBestEffort"`
	// Name of the environment variable or secret that holds the ClickHouse HTTP endpoint URL. Defaults to `DESTINATION_CLICKHOUSE_ENDPOINT_URL`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#endpoint_url_key ObservabilityPipeline#endpoint_url_key}
	EndpointUrlKey *string `field:"optional" json:"endpointUrlKey" yaml:"endpointUrlKey"`
	// Insert format for events.
	//
	// `json_each_row` maps event fields to columns by name. `json_as_object` and `json_as_string` insert each event into a single JSON or String column. `arrow_stream` batches events with Apache Arrow IPC streaming and requires `batch_encoding`. Valid values are `json_each_row`, `json_as_object`, `json_as_string`, `arrow_stream`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#format ObservabilityPipeline#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// If `true`, fields not present in the target table schema are dropped instead of causing insert errors.
	//
	// When unset, the ClickHouse server's own `input_format_skip_unknown_fields` setting applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#skip_unknown_fields ObservabilityPipeline#skip_unknown_fields}
	SkipUnknownFields interface{} `field:"optional" json:"skipUnknownFields" yaml:"skipUnknownFields"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/observability_pipeline#tls ObservabilityPipeline#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

