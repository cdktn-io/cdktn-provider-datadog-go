// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitor


type MonitorVariablesAggregateAugmentedQueryJoinCondition struct {
	// Attribute from the augment query to join on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/monitor#augment_attribute Monitor#augment_attribute}
	AugmentAttribute *string `field:"required" json:"augmentAttribute" yaml:"augmentAttribute"`
	// Attribute from the base query to join on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/monitor#base_attribute Monitor#base_attribute}
	BaseAttribute *string `field:"required" json:"baseAttribute" yaml:"baseAttribute"`
	// Join type (for example `inner`). Valid values are `inner`, `left`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/monitor#join_type Monitor#join_type}
	JoinType *string `field:"required" json:"joinType" yaml:"joinType"`
}

