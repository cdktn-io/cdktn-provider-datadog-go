// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSplitConfigSortCompute struct {
	// The metric to use for sorting graphs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/dashboard#metric Dashboard#metric}
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// How to aggregate the sort metric for the purposes of ordering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/dashboard#aggregation Dashboard#aggregation}
	Aggregation *string `field:"optional" json:"aggregation" yaml:"aggregation"`
}

