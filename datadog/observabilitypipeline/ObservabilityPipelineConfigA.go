// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigA struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#destination ObservabilityPipeline#destination}
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The type of data being ingested. Defaults to `logs` if not specified. Valid values are `logs`, `metrics`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#pipeline_type ObservabilityPipeline#pipeline_type}
	PipelineType *string `field:"optional" json:"pipelineType" yaml:"pipelineType"`
	// processor_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#processor_group ObservabilityPipeline#processor_group}
	ProcessorGroup interface{} `field:"optional" json:"processorGroup" yaml:"processorGroup"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#source ObservabilityPipeline#source}
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// Set to `true` to continue using the legacy search syntax while migrating filter queries.
	//
	// After migrating all queries to the new syntax, set to `false`. The legacy syntax is deprecated and will eventually be removed. Requires Observability Pipelines Worker 2.11 or later. See https://docs.datadoghq.com/observability_pipelines/guide/upgrade_your_filter_queries_to_the_new_search_syntax/ for more information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#use_legacy_search_syntax ObservabilityPipeline#use_legacy_search_syntax}
	UseLegacySearchSyntax interface{} `field:"optional" json:"useLegacySearchSyntax" yaml:"useLegacySearchSyntax"`
}

