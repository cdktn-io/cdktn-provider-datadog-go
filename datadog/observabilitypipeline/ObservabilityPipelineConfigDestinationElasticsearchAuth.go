// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationElasticsearchAuth struct {
	// The authentication strategy. Use `basic` for username/password. Valid values are `basic`, `aws`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#strategy ObservabilityPipeline#strategy}
	Strategy *string `field:"required" json:"strategy" yaml:"strategy"`
	// Name of the environment variable or secret that holds the Elasticsearch password (used when strategy is `basic`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#password_key ObservabilityPipeline#password_key}
	PasswordKey *string `field:"optional" json:"passwordKey" yaml:"passwordKey"`
	// Name of the environment variable or secret that holds the Elasticsearch username (used when strategy is `basic`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#username_key ObservabilityPipeline#username_key}
	UsernameKey *string `field:"optional" json:"usernameKey" yaml:"usernameKey"`
}

