// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetFreeTextDefinition struct {
	// The text to display in the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#text DashboardV2#text}
	Text *string `field:"required" json:"text" yaml:"text"`
	// The background color of the text widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#background_color DashboardV2#background_color}
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// The color of the text in the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#color DashboardV2#color}
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The description of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#description DashboardV2#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The size of the text in the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#font_size DashboardV2#font_size}
	FontSize *string `field:"optional" json:"fontSize" yaml:"fontSize"`
	// Hide any portion of the widget's timeframe that is incomplete due to cost data not being available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#hide_incomplete_cost_data DashboardV2#hide_incomplete_cost_data}
	HideIncompleteCostData interface{} `field:"optional" json:"hideIncompleteCostData" yaml:"hideIncompleteCostData"`
	// The timeframe to use when displaying the widget.
	//
	// Valid values are `1m`, `5m`, `10m`, `15m`, `30m`, `1h`, `4h`, `1d`, `2d`, `1w`, `1mo`, `3mo`, `6mo`, `week_to_date`, `month_to_date`, `1y`, `alert`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#live_span DashboardV2#live_span}
	LiveSpan *string `field:"optional" json:"liveSpan" yaml:"liveSpan"`
	// The alignment of the text in the widget. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#text_align DashboardV2#text_align}
	TextAlign *string `field:"optional" json:"textAlign" yaml:"textAlign"`
	// time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#time DashboardV2#time}
	Time *DashboardV2WidgetFreeTextDefinitionTime `field:"optional" json:"time" yaml:"time"`
	// The title of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#title DashboardV2#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The alignment of the widget's title. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#title_align DashboardV2#title_align}
	TitleAlign *string `field:"optional" json:"titleAlign" yaml:"titleAlign"`
	// The size of the widget's title (defaults to 16).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#title_size DashboardV2#title_size}
	TitleSize *string `field:"optional" json:"titleSize" yaml:"titleSize"`
}

