// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package costcustomforecast


type CostCustomForecastEntries struct {
	// The forecast override amount for the month. Value must be at least 0.000000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/cost_custom_forecast#amount CostCustomForecast#amount}
	Amount *float64 `field:"required" json:"amount" yaml:"amount"`
	// The month the entry applies to, in `YYYYMM` format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/cost_custom_forecast#month CostCustomForecast#month}
	Month *float64 `field:"required" json:"month" yaml:"month"`
	// tag_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/cost_custom_forecast#tag_filters CostCustomForecast#tag_filters}
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

