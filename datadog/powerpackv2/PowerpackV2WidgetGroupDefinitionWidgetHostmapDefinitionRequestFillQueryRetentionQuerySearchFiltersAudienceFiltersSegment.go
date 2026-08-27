// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryRetentionQuerySearchFiltersAudienceFiltersSegment struct {
	// The name of the segment subquery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#name PowerpackV2#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The unique identifier of the segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#segment_id PowerpackV2#segment_id}
	SegmentId *string `field:"optional" json:"segmentId" yaml:"segmentId"`
}

