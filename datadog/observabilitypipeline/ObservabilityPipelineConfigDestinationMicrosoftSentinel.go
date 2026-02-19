// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationMicrosoftSentinel struct {
	// Azure AD client ID used for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#client_id ObservabilityPipeline#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The immutable ID of the Data Collection Rule (DCR).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#dcr_immutable_id ObservabilityPipeline#dcr_immutable_id}
	DcrImmutableId *string `field:"required" json:"dcrImmutableId" yaml:"dcrImmutableId"`
	// The name of the Log Analytics table where logs will be sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#table ObservabilityPipeline#table}
	Table *string `field:"required" json:"table" yaml:"table"`
	// Azure AD tenant ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#tenant_id ObservabilityPipeline#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

