// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinition struct {
	// The layout type of the group. Valid values are `ordered`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#layout_type PowerpackV2#layout_type}
	LayoutType *string `field:"required" json:"layoutType" yaml:"layoutType"`
	// The background color of the group title, options: `vivid_blue`, `vivid_purple`, `vivid_pink`, `vivid_orange`, `vivid_yellow`, `vivid_green`, `blue`, `purple`, `pink`, `orange`, `yellow`, `green`, `gray` or `white`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#background_color PowerpackV2#background_color}
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// The image URL to display as a banner for the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#banner_img PowerpackV2#banner_img}
	BannerImg *string `field:"optional" json:"bannerImg" yaml:"bannerImg"`
	// The description of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#description PowerpackV2#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
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
	// Whether to show the title or not. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#show_title PowerpackV2#show_title}
	ShowTitle interface{} `field:"optional" json:"showTitle" yaml:"showTitle"`
	// time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#time PowerpackV2#time}
	Time *PowerpackV2WidgetGroupDefinitionTime `field:"optional" json:"time" yaml:"time"`
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
	// widget block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#widget PowerpackV2#widget}
	Widget interface{} `field:"optional" json:"widget" yaml:"widget"`
}

