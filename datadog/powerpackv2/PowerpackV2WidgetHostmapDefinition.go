// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetHostmapDefinition struct {
	// custom_link block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#custom_link PowerpackV2#custom_link}
	CustomLink interface{} `field:"optional" json:"customLink" yaml:"customLink"`
	// The description of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#description PowerpackV2#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The list of tag prefixes to group by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#group PowerpackV2#group}
	Group *[]*string `field:"optional" json:"group" yaml:"group"`
	// Hide any portion of the widget's timeframe that is incomplete due to cost data not being available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#hide_incomplete_cost_data PowerpackV2#hide_incomplete_cost_data}
	HideIncompleteCostData interface{} `field:"optional" json:"hideIncompleteCostData" yaml:"hideIncompleteCostData"`
	// The timeframe to use when displaying the widget.
	//
	// Valid values are `1m`, `5m`, `10m`, `15m`, `30m`, `1h`, `4h`, `1d`, `2d`, `1w`, `1mo`, `3mo`, `6mo`, `week_to_date`, `month_to_date`, `1y`, `alert`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#live_span PowerpackV2#live_span}
	LiveSpan *string `field:"optional" json:"liveSpan" yaml:"liveSpan"`
	// The type of node used. Valid values are `host`, `container`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#node_type PowerpackV2#node_type}
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
	// A Boolean indicating whether to show the hosts that don't fit in a group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#no_group_hosts PowerpackV2#no_group_hosts}
	NoGroupHosts interface{} `field:"optional" json:"noGroupHosts" yaml:"noGroupHosts"`
	// A Boolean indicating whether to show nodes with no metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#no_metric_hosts PowerpackV2#no_metric_hosts}
	NoMetricHosts interface{} `field:"optional" json:"noMetricHosts" yaml:"noMetricHosts"`
	// Notes/description text for the host map widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#notes PowerpackV2#notes}
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#request PowerpackV2#request}
	Request *PowerpackV2WidgetHostmapDefinitionRequest `field:"optional" json:"request" yaml:"request"`
	// The list of tags used to filter the map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#scope PowerpackV2#scope}
	Scope *[]*string `field:"optional" json:"scope" yaml:"scope"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#style PowerpackV2#style}
	Style *PowerpackV2WidgetHostmapDefinitionStyle `field:"optional" json:"style" yaml:"style"`
	// time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#time PowerpackV2#time}
	Time *PowerpackV2WidgetHostmapDefinitionTime `field:"optional" json:"time" yaml:"time"`
	// The title of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#title PowerpackV2#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The alignment of the widget's title. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#title_align PowerpackV2#title_align}
	TitleAlign *string `field:"optional" json:"titleAlign" yaml:"titleAlign"`
	// The size of the widget's title (defaults to 16).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#title_size PowerpackV2#title_size}
	TitleSize *string `field:"optional" json:"titleSize" yaml:"titleSize"`
}

