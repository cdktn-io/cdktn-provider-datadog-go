// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorSensitiveDataScannerRuleScopeInclude struct {
	// The fields to include in scanning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#fields ObservabilityPipeline#fields}
	Fields *[]*string `field:"optional" json:"fields" yaml:"fields"`
}

