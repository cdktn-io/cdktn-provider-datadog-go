// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationNewRelic struct {
	// The New Relic region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#region ObservabilityPipeline#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// Name of the environment variable or secret that holds the New Relic account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#account_id_key ObservabilityPipeline#account_id_key}
	AccountIdKey *string `field:"optional" json:"accountIdKey" yaml:"accountIdKey"`
	// buffer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#buffer ObservabilityPipeline#buffer}
	Buffer interface{} `field:"optional" json:"buffer" yaml:"buffer"`
	// Name of the environment variable or secret that holds the New Relic license key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#license_key_key ObservabilityPipeline#license_key_key}
	LicenseKeyKey *string `field:"optional" json:"licenseKeyKey" yaml:"licenseKeyKey"`
}

