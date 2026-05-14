// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionXaxis struct {
	// True includes zero.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/powerpack_v2#include_zero PowerpackV2#include_zero}
	IncludeZero interface{} `field:"optional" json:"includeZero" yaml:"includeZero"`
	// Specifies maximum value to show on the x-axis.
	//
	// It takes a number, percentile (p90 === 90th percentile), or auto for default behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/powerpack_v2#max PowerpackV2#max}
	Max *string `field:"optional" json:"max" yaml:"max"`
	// Specifies minimum value to show on the x-axis.
	//
	// It takes a number, percentile (p90 === 90th percentile), or auto for default behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/powerpack_v2#min PowerpackV2#min}
	Min *string `field:"optional" json:"min" yaml:"min"`
	// Number of value buckets to target, also known as the resolution of the value bins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/powerpack_v2#num_buckets PowerpackV2#num_buckets}
	NumBuckets *float64 `field:"optional" json:"numBuckets" yaml:"numBuckets"`
	// Specifies the scale type. Possible values are `linear`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/powerpack_v2#scale PowerpackV2#scale}
	Scale *string `field:"optional" json:"scale" yaml:"scale"`
}

