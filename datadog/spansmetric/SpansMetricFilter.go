// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package spansmetric


type SpansMetricFilter struct {
	// The search query - following the span search syntax. Defaults to `"*"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/spans_metric#query SpansMetric#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
}

