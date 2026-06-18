// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitor


type MonitorVariablesAggregateFilteredQueryFilters struct {
	// Attribute from the base query to filter on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/monitor#base_attribute Monitor#base_attribute}
	BaseAttribute *string `field:"required" json:"baseAttribute" yaml:"baseAttribute"`
	// Attribute from the filter query to match against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/monitor#filter_attribute Monitor#filter_attribute}
	FilterAttribute *string `field:"required" json:"filterAttribute" yaml:"filterAttribute"`
	// When true, exclude matching records instead of including them.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/monitor#exclude Monitor#exclude}
	Exclude interface{} `field:"optional" json:"exclude" yaml:"exclude"`
}

