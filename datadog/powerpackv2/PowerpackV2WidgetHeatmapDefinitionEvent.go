// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetHeatmapDefinitionEvent struct {
	// The event query to use in the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/powerpack_v2#q PowerpackV2#q}
	Q *string `field:"required" json:"q" yaml:"q"`
	// The execution method for multi-value filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/powerpack_v2#tags_execution PowerpackV2#tags_execution}
	TagsExecution *string `field:"optional" json:"tagsExecution" yaml:"tagsExecution"`
}

