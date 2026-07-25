// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetImageDefinition struct {
	// URL of the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#url DashboardV2#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// The description of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#description DashboardV2#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether to display a background or not. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#has_background DashboardV2#has_background}
	HasBackground interface{} `field:"optional" json:"hasBackground" yaml:"hasBackground"`
	// Whether to display a border or not. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#has_border DashboardV2#has_border}
	HasBorder interface{} `field:"optional" json:"hasBorder" yaml:"hasBorder"`
	// Hide any portion of the widget's timeframe that is incomplete due to cost data not being available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#hide_incomplete_cost_data DashboardV2#hide_incomplete_cost_data}
	HideIncompleteCostData interface{} `field:"optional" json:"hideIncompleteCostData" yaml:"hideIncompleteCostData"`
	// The horizontal alignment for the widget. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#horizontal_align DashboardV2#horizontal_align}
	HorizontalAlign *string `field:"optional" json:"horizontalAlign" yaml:"horizontalAlign"`
	// The timeframe to use when displaying the widget.
	//
	// Valid values are `1m`, `5m`, `10m`, `15m`, `30m`, `1h`, `4h`, `1d`, `2d`, `1w`, `1mo`, `3mo`, `6mo`, `week_to_date`, `month_to_date`, `1y`, `alert`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#live_span DashboardV2#live_span}
	LiveSpan *string `field:"optional" json:"liveSpan" yaml:"liveSpan"`
	// The margins to use around the image.
	//
	// Note: `small` and `large` values are deprecated. Valid values are `sm`, `md`, `lg`, `small`, `large`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#margin DashboardV2#margin}
	Margin *string `field:"optional" json:"margin" yaml:"margin"`
	// The preferred method to adapt the dimensions of the image.
	//
	// The values are based on the image `object-fit` CSS properties. Note: `zoom`, `fit` and `center` values are deprecated. Valid values are `fill`, `contain`, `cover`, `none`, `scale-down`, `zoom`, `fit`, `center`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#sizing DashboardV2#sizing}
	Sizing *string `field:"optional" json:"sizing" yaml:"sizing"`
	// time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#time DashboardV2#time}
	Time *DashboardV2WidgetImageDefinitionTime `field:"optional" json:"time" yaml:"time"`
	// The title of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#title DashboardV2#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The alignment of the widget's title. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#title_align DashboardV2#title_align}
	TitleAlign *string `field:"optional" json:"titleAlign" yaml:"titleAlign"`
	// The size of the widget's title (defaults to 16).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#title_size DashboardV2#title_size}
	TitleSize *string `field:"optional" json:"titleSize" yaml:"titleSize"`
	// URL of the image in dark mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#url_dark_theme DashboardV2#url_dark_theme}
	UrlDarkTheme *string `field:"optional" json:"urlDarkTheme" yaml:"urlDarkTheme"`
	// The vertical alignment for the widget. Valid values are `center`, `top`, `bottom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2#vertical_align DashboardV2#vertical_align}
	VerticalAlign *string `field:"optional" json:"verticalAlign" yaml:"verticalAlign"`
}

