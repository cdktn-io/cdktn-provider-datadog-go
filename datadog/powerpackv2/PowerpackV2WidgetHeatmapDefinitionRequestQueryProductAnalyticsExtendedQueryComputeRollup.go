// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetHeatmapDefinitionRequestQueryProductAnalyticsExtendedQueryComputeRollup struct {
	// Type of calendar interval. Valid values are `day`, `week`, `month`, `year`, `quarter`, `minute`, `hour`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#type PowerpackV2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Alignment of the calendar interval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#alignment PowerpackV2#alignment}
	Alignment *string `field:"optional" json:"alignment" yaml:"alignment"`
	// Quantity of the calendar interval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#quantity PowerpackV2#quantity}
	Quantity *float64 `field:"optional" json:"quantity" yaml:"quantity"`
	// Timezone for the calendar interval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#timezone PowerpackV2#timezone}
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

