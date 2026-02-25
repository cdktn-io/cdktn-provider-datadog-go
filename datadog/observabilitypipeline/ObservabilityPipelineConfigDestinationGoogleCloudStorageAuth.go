// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationGoogleCloudStorageAuth struct {
	// Path to the GCP service account key file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/observability_pipeline#credentials_file ObservabilityPipeline#credentials_file}
	CredentialsFile *string `field:"required" json:"credentialsFile" yaml:"credentialsFile"`
}

