// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package costcustomforecast


type CostCustomForecastEntriesTagFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/cost_custom_forecast#tag_key CostCustomForecast#tag_key}.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/cost_custom_forecast#tag_value CostCustomForecast#tag_value}.
	TagValue *string `field:"required" json:"tagValue" yaml:"tagValue"`
}

