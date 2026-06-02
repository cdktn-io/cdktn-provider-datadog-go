// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetServiceLevelObjectiveDefinition struct {
	// The ID of the service level objective used by the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#slo_id DashboardV2#slo_id}
	SloId *string `field:"required" json:"sloId" yaml:"sloId"`
	// A list of time windows to display in the widget.
	//
	// Valid values are `7d`, `30d`, `90d`, `week_to_date`, `previous_week`, `month_to_date`, `previous_month`, `global_time`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#time_windows DashboardV2#time_windows}
	TimeWindows *[]*string `field:"required" json:"timeWindows" yaml:"timeWindows"`
	// The view mode for the widget. Valid values are `overall`, `component`, `both`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#view_mode DashboardV2#view_mode}
	ViewMode *string `field:"required" json:"viewMode" yaml:"viewMode"`
	// The type of view to use when displaying the widget. Only `detail` is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#view_type DashboardV2#view_type}
	ViewType *string `field:"required" json:"viewType" yaml:"viewType"`
	// Additional filters applied to the SLO query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#additional_query_filters DashboardV2#additional_query_filters}
	AdditionalQueryFilters *string `field:"optional" json:"additionalQueryFilters" yaml:"additionalQueryFilters"`
	// The description of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#description DashboardV2#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The global time target of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#global_time_target DashboardV2#global_time_target}
	GlobalTimeTarget *string `field:"optional" json:"globalTimeTarget" yaml:"globalTimeTarget"`
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
	// Whether to show the error budget or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#show_error_budget DashboardV2#show_error_budget}
	ShowErrorBudget interface{} `field:"optional" json:"showErrorBudget" yaml:"showErrorBudget"`
	// time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/dashboard_v2#time DashboardV2#time}
	Time *DashboardV2WidgetServiceLevelObjectiveDefinitionTime `field:"optional" json:"time" yaml:"time"`
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

