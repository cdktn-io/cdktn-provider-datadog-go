// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationClickhouseAuth struct {
	// Authentication strategy. Must be `basic`. Valid values are `basic`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/observability_pipeline#strategy ObservabilityPipeline#strategy}
	Strategy *string `field:"required" json:"strategy" yaml:"strategy"`
	// Name of the environment variable or secret that holds the ClickHouse password. Defaults to `DESTINATION_CLICKHOUSE_PASSWORD`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/observability_pipeline#password_key ObservabilityPipeline#password_key}
	PasswordKey *string `field:"optional" json:"passwordKey" yaml:"passwordKey"`
	// Name of the environment variable or secret that holds the ClickHouse username. Defaults to `DESTINATION_CLICKHOUSE_USERNAME`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/observability_pipeline#username_key ObservabilityPipeline#username_key}
	UsernameKey *string `field:"optional" json:"usernameKey" yaml:"usernameKey"`
}

