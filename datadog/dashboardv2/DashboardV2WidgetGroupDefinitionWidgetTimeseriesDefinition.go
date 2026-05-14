// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetTimeseriesDefinition struct {
	// custom_link block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#custom_link DashboardV2#custom_link}
	CustomLink interface{} `field:"optional" json:"customLink" yaml:"customLink"`
	// The description of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#description DashboardV2#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// event block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#event DashboardV2#event}
	Event interface{} `field:"optional" json:"event" yaml:"event"`
	// Hide any portion of the widget's timeframe that is incomplete due to cost data not being available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#hide_incomplete_cost_data DashboardV2#hide_incomplete_cost_data}
	HideIncompleteCostData interface{} `field:"optional" json:"hideIncompleteCostData" yaml:"hideIncompleteCostData"`
	// A list of columns to display in the legend. Valid values are `value`, `avg`, `sum`, `min`, `max`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#legend_columns DashboardV2#legend_columns}
	LegendColumns *[]*string `field:"optional" json:"legendColumns" yaml:"legendColumns"`
	// The layout of the legend displayed in the widget. Valid values are `auto`, `horizontal`, `vertical`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#legend_layout DashboardV2#legend_layout}
	LegendLayout *string `field:"optional" json:"legendLayout" yaml:"legendLayout"`
	// The size of the legend displayed in the widget. Valid values are `0`, `2`, `4`, `8`, `16`, `auto`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#legend_size DashboardV2#legend_size}
	LegendSize *string `field:"optional" json:"legendSize" yaml:"legendSize"`
	// The timeframe to use when displaying the widget.
	//
	// Valid values are `1m`, `5m`, `10m`, `15m`, `30m`, `1h`, `4h`, `1d`, `2d`, `1w`, `1mo`, `3mo`, `6mo`, `week_to_date`, `month_to_date`, `1y`, `alert`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#live_span DashboardV2#live_span}
	LiveSpan *string `field:"optional" json:"liveSpan" yaml:"liveSpan"`
	// marker block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#marker DashboardV2#marker}
	Marker interface{} `field:"optional" json:"marker" yaml:"marker"`
	// request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#request DashboardV2#request}
	Request interface{} `field:"optional" json:"request" yaml:"request"`
	// right_yaxis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#right_yaxis DashboardV2#right_yaxis}
	RightYaxis *DashboardV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRightYaxis `field:"optional" json:"rightYaxis" yaml:"rightYaxis"`
	// Whether or not to show the legend on this widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#show_legend DashboardV2#show_legend}
	ShowLegend interface{} `field:"optional" json:"showLegend" yaml:"showLegend"`
	// time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#time DashboardV2#time}
	Time *DashboardV2WidgetGroupDefinitionWidgetTimeseriesDefinitionTime `field:"optional" json:"time" yaml:"time"`
	// The title of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#title DashboardV2#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The alignment of the widget's title. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#title_align DashboardV2#title_align}
	TitleAlign *string `field:"optional" json:"titleAlign" yaml:"titleAlign"`
	// The size of the widget's title (defaults to 16).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#title_size DashboardV2#title_size}
	TitleSize *string `field:"optional" json:"titleSize" yaml:"titleSize"`
	// yaxis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#yaxis DashboardV2#yaxis}
	Yaxis *DashboardV2WidgetGroupDefinitionWidgetTimeseriesDefinitionYaxis `field:"optional" json:"yaxis" yaml:"yaxis"`
}

