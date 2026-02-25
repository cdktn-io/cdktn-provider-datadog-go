// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitor


type MonitorVariables struct {
	// cloud_cost_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/monitor#cloud_cost_query Monitor#cloud_cost_query}
	CloudCostQuery interface{} `field:"optional" json:"cloudCostQuery" yaml:"cloudCostQuery"`
	// data_quality_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/monitor#data_quality_query Monitor#data_quality_query}
	DataQualityQuery interface{} `field:"optional" json:"dataQualityQuery" yaml:"dataQualityQuery"`
	// event_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/monitor#event_query Monitor#event_query}
	EventQuery interface{} `field:"optional" json:"eventQuery" yaml:"eventQuery"`
}

