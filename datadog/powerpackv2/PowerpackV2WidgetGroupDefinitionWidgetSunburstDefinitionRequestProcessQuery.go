// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestProcessQuery struct {
	// Your chosen metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#metric PowerpackV2#metric}
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// A list of processes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#filter_by PowerpackV2#filter_by}
	FilterBy *[]*string `field:"optional" json:"filterBy" yaml:"filterBy"`
	// The max number of items in the filter list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#limit PowerpackV2#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Your chosen search term.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#search_by PowerpackV2#search_by}
	SearchBy *string `field:"optional" json:"searchBy" yaml:"searchBy"`
}

