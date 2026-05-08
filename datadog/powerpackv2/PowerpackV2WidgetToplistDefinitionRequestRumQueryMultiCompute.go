// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetToplistDefinitionRequestRumQueryMultiCompute struct {
	// The aggregation method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/powerpack_v2#aggregation PowerpackV2#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// The facet name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/powerpack_v2#facet PowerpackV2#facet}
	Facet *string `field:"optional" json:"facet" yaml:"facet"`
	// Define the time interval in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/powerpack_v2#interval PowerpackV2#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
}

