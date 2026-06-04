// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitor


type MonitorVariablesAggregateAugmentedQueryAugmentReferenceTableColumns struct {
	// Reference table column name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/monitor#name Monitor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional alias for the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/monitor#alias Monitor#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

