// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customallocationrule


type CustomAllocationRuleStrategyBasedOnTimeseries struct {
	// The timeseries query that determines the allocation proportions, encoded as a JSON object.
	//
	// Required when `method` is `proportional_timeseries`. Uses Datadog's formulas-and-functions request format with `queries`, `formulas`, and `response_format` keys. Build it with `jsonencode()`. The set of supported `data_source` values is defined by the API, not by this provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/custom_allocation_rule#json CustomAllocationRule#json}
	Json *string `field:"optional" json:"json" yaml:"json"`
}

