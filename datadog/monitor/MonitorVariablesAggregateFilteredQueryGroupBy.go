// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitor


type MonitorVariablesAggregateFilteredQueryGroupBy struct {
	// The facet to group by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/monitor#facet Monitor#facet}
	Facet *string `field:"required" json:"facet" yaml:"facet"`
	// The number of groups to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/monitor#limit Monitor#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/monitor#sort Monitor#sort}
	Sort *MonitorVariablesAggregateFilteredQueryGroupBySort `field:"optional" json:"sort" yaml:"sort"`
	// Identifies which sub-query this facet refers to (for example `filter_query`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/monitor#source Monitor#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
}

