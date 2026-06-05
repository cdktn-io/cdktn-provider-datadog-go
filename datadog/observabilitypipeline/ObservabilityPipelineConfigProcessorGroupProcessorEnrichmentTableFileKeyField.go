// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorEnrichmentTableFileKeyField struct {
	// The path to the field in the log event to use as the lookup key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/observability_pipeline#event ObservabilityPipeline#event}
	Event *string `field:"optional" json:"event" yaml:"event"`
	// The name of the secret containing the lookup key value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/observability_pipeline#secret ObservabilityPipeline#secret}
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
	// A plain field path in the log event (for example, `log.user.id`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/observability_pipeline#string_path ObservabilityPipeline#string_path}
	StringPath *string `field:"optional" json:"stringPath" yaml:"stringPath"`
	// A VRL expression that returns the value to use as the lookup key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/observability_pipeline#vrl ObservabilityPipeline#vrl}
	Vrl *string `field:"optional" json:"vrl" yaml:"vrl"`
}

