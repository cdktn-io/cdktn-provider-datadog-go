// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessor struct {
	// Whether this processor is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#enabled ObservabilityPipeline#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The unique identifier for this processor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#id ObservabilityPipeline#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// A Datadog search query used to determine which logs this processor targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#include ObservabilityPipeline#include}
	Include *string `field:"required" json:"include" yaml:"include"`
	// add_env_vars block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#add_env_vars ObservabilityPipeline#add_env_vars}
	AddEnvVars interface{} `field:"optional" json:"addEnvVars" yaml:"addEnvVars"`
	// add_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#add_fields ObservabilityPipeline#add_fields}
	AddFields interface{} `field:"optional" json:"addFields" yaml:"addFields"`
	// add_hostname block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#add_hostname ObservabilityPipeline#add_hostname}
	AddHostname interface{} `field:"optional" json:"addHostname" yaml:"addHostname"`
	// custom_processor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#custom_processor ObservabilityPipeline#custom_processor}
	CustomProcessor interface{} `field:"optional" json:"customProcessor" yaml:"customProcessor"`
	// datadog_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#datadog_tags ObservabilityPipeline#datadog_tags}
	DatadogTags interface{} `field:"optional" json:"datadogTags" yaml:"datadogTags"`
	// dedupe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#dedupe ObservabilityPipeline#dedupe}
	Dedupe interface{} `field:"optional" json:"dedupe" yaml:"dedupe"`
	// A human-friendly name for this processor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#display_name ObservabilityPipeline#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// enrichment_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#enrichment_table ObservabilityPipeline#enrichment_table}
	EnrichmentTable interface{} `field:"optional" json:"enrichmentTable" yaml:"enrichmentTable"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#filter ObservabilityPipeline#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// generate_datadog_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#generate_datadog_metrics ObservabilityPipeline#generate_datadog_metrics}
	GenerateDatadogMetrics interface{} `field:"optional" json:"generateDatadogMetrics" yaml:"generateDatadogMetrics"`
	// metric_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#metric_tags ObservabilityPipeline#metric_tags}
	MetricTags interface{} `field:"optional" json:"metricTags" yaml:"metricTags"`
	// ocsf_mapper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#ocsf_mapper ObservabilityPipeline#ocsf_mapper}
	OcsfMapper interface{} `field:"optional" json:"ocsfMapper" yaml:"ocsfMapper"`
	// parse_grok block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#parse_grok ObservabilityPipeline#parse_grok}
	ParseGrok interface{} `field:"optional" json:"parseGrok" yaml:"parseGrok"`
	// parse_json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#parse_json ObservabilityPipeline#parse_json}
	ParseJson interface{} `field:"optional" json:"parseJson" yaml:"parseJson"`
	// parse_xml block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#parse_xml ObservabilityPipeline#parse_xml}
	ParseXml interface{} `field:"optional" json:"parseXml" yaml:"parseXml"`
	// quota block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#quota ObservabilityPipeline#quota}
	Quota interface{} `field:"optional" json:"quota" yaml:"quota"`
	// reduce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#reduce ObservabilityPipeline#reduce}
	Reduce interface{} `field:"optional" json:"reduce" yaml:"reduce"`
	// remove_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#remove_fields ObservabilityPipeline#remove_fields}
	RemoveFields interface{} `field:"optional" json:"removeFields" yaml:"removeFields"`
	// rename_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#rename_fields ObservabilityPipeline#rename_fields}
	RenameFields interface{} `field:"optional" json:"renameFields" yaml:"renameFields"`
	// sample block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#sample ObservabilityPipeline#sample}
	Sample interface{} `field:"optional" json:"sample" yaml:"sample"`
	// sensitive_data_scanner block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#sensitive_data_scanner ObservabilityPipeline#sensitive_data_scanner}
	SensitiveDataScanner interface{} `field:"optional" json:"sensitiveDataScanner" yaml:"sensitiveDataScanner"`
	// split_array block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#split_array ObservabilityPipeline#split_array}
	SplitArray interface{} `field:"optional" json:"splitArray" yaml:"splitArray"`
	// throttle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#throttle ObservabilityPipeline#throttle}
	Throttle interface{} `field:"optional" json:"throttle" yaml:"throttle"`
}

