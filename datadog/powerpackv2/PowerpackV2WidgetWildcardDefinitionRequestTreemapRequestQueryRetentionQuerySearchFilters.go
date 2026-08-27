// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchFilters struct {
	// audience_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#audience_filters PowerpackV2#audience_filters}
	AudienceFilters *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchFiltersAudienceFilters `field:"optional" json:"audienceFilters" yaml:"audienceFilters"`
	// String filter for the retention query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#string_filter PowerpackV2#string_filter}
	StringFilter *string `field:"optional" json:"stringFilter" yaml:"stringFilter"`
}

