// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetManageStatusDefinition struct {
	// Query to filter the monitors with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#query DashboardV2#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// Whether to colorize text or background. Valid values are `background`, `text`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#color_preference DashboardV2#color_preference}
	ColorPreference *string `field:"optional" json:"colorPreference" yaml:"colorPreference"`
	// The description of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#description DashboardV2#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display setting to use. Valid values are `counts`, `countsAndList`, `list`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#display_format DashboardV2#display_format}
	DisplayFormat *string `field:"optional" json:"displayFormat" yaml:"displayFormat"`
	// Hide any portion of the widget's timeframe that is incomplete due to cost data not being available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#hide_incomplete_cost_data DashboardV2#hide_incomplete_cost_data}
	HideIncompleteCostData interface{} `field:"optional" json:"hideIncompleteCostData" yaml:"hideIncompleteCostData"`
	// A Boolean indicating whether to hide empty categories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#hide_zero_counts DashboardV2#hide_zero_counts}
	HideZeroCounts interface{} `field:"optional" json:"hideZeroCounts" yaml:"hideZeroCounts"`
	// The timeframe to use when displaying the widget.
	//
	// Valid values are `1m`, `5m`, `10m`, `15m`, `30m`, `1h`, `4h`, `1d`, `2d`, `1w`, `1mo`, `3mo`, `6mo`, `week_to_date`, `month_to_date`, `1y`, `alert`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#live_span DashboardV2#live_span}
	LiveSpan *string `field:"optional" json:"liveSpan" yaml:"liveSpan"`
	// Whether to show the time that has elapsed since the monitor/group triggered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#show_last_triggered DashboardV2#show_last_triggered}
	ShowLastTriggered interface{} `field:"optional" json:"showLastTriggered" yaml:"showLastTriggered"`
	// Whether to show the priorities column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#show_priority DashboardV2#show_priority}
	ShowPriority interface{} `field:"optional" json:"showPriority" yaml:"showPriority"`
	// The method to sort the monitors.
	//
	// Valid values are `name`, `group`, `status`, `tags`, `triggered`, `group,asc`, `group,desc`, `name,asc`, `name,desc`, `status,asc`, `status,desc`, `tags,asc`, `tags,desc`, `triggered,asc`, `triggered,desc`, `priority,asc`, `priority,desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#sort DashboardV2#sort}
	Sort *string `field:"optional" json:"sort" yaml:"sort"`
	// The summary type to use. Valid values are `monitors`, `groups`, `combined`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#summary_type DashboardV2#summary_type}
	SummaryType *string `field:"optional" json:"summaryType" yaml:"summaryType"`
	// time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#time DashboardV2#time}
	Time *DashboardV2WidgetManageStatusDefinitionTime `field:"optional" json:"time" yaml:"time"`
	// The title of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#title DashboardV2#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The alignment of the widget's title. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#title_align DashboardV2#title_align}
	TitleAlign *string `field:"optional" json:"titleAlign" yaml:"titleAlign"`
	// The size of the widget's title (defaults to 16).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#title_size DashboardV2#title_size}
	TitleSize *string `field:"optional" json:"titleSize" yaml:"titleSize"`
}

