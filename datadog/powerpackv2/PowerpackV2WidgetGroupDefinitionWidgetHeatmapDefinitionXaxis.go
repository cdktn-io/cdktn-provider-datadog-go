// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionXaxis struct {
	// Number of time buckets to target, also known as the resolution of the time bins.
	//
	// This is only applicable for distribution of points (group distributions use the roll-up modifier).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#num_buckets PowerpackV2#num_buckets}
	NumBuckets *float64 `field:"optional" json:"numBuckets" yaml:"numBuckets"`
}

