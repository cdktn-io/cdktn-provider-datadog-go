// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logsindex


type LogsIndexExclusionFilterFilter struct {
	// Only logs matching the filter criteria and the query of the parent index will be considered for this exclusion filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/logs_index#query LogsIndex#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
	// The log attribute used as the sampling key.
	//
	// When present, logs sharing the same value are excluded or kept together at the configured sample rate (a single attribute path, e.g. `@lambda.request_id`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/logs_index#sample_attribute LogsIndex#sample_attribute}
	SampleAttribute *string `field:"optional" json:"sampleAttribute" yaml:"sampleAttribute"`
	// The fraction of logs excluded by the exclusion filter, when active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/logs_index#sample_rate LogsIndex#sample_rate}
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}

