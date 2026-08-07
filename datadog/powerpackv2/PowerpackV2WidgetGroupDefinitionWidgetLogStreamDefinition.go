// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetLogStreamDefinition struct {
	// Stringified list of columns to use, for example: `["column1","column2","column3"]`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#columns PowerpackV2#columns}
	Columns *[]*string `field:"optional" json:"columns" yaml:"columns"`
	// The description of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#description PowerpackV2#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Hide any portion of the widget's timeframe that is incomplete due to cost data not being available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#hide_incomplete_cost_data PowerpackV2#hide_incomplete_cost_data}
	HideIncompleteCostData interface{} `field:"optional" json:"hideIncompleteCostData" yaml:"hideIncompleteCostData"`
	// An array of index names to query in the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#indexes PowerpackV2#indexes}
	Indexes *[]*string `field:"optional" json:"indexes" yaml:"indexes"`
	// The timeframe to use when displaying the widget.
	//
	// Valid values are `1m`, `5m`, `10m`, `15m`, `30m`, `1h`, `4h`, `1d`, `2d`, `1w`, `1mo`, `3mo`, `6mo`, `week_to_date`, `month_to_date`, `1y`, `alert`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#live_span PowerpackV2#live_span}
	LiveSpan *string `field:"optional" json:"liveSpan" yaml:"liveSpan"`
	// The number of log lines to display. Valid values are `inline`, `expanded-md`, `expanded-lg`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#message_display PowerpackV2#message_display}
	MessageDisplay *string `field:"optional" json:"messageDisplay" yaml:"messageDisplay"`
	// Query to filter the log stream with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#query PowerpackV2#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
	// If the date column should be displayed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#show_date_column PowerpackV2#show_date_column}
	ShowDateColumn interface{} `field:"optional" json:"showDateColumn" yaml:"showDateColumn"`
	// If the message column should be displayed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#show_message_column PowerpackV2#show_message_column}
	ShowMessageColumn interface{} `field:"optional" json:"showMessageColumn" yaml:"showMessageColumn"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#sort PowerpackV2#sort}
	Sort *PowerpackV2WidgetGroupDefinitionWidgetLogStreamDefinitionSort `field:"optional" json:"sort" yaml:"sort"`
	// time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#time PowerpackV2#time}
	Time *PowerpackV2WidgetGroupDefinitionWidgetLogStreamDefinitionTime `field:"optional" json:"time" yaml:"time"`
	// The title of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#title PowerpackV2#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The alignment of the widget's title. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#title_align PowerpackV2#title_align}
	TitleAlign *string `field:"optional" json:"titleAlign" yaml:"titleAlign"`
	// The size of the widget's title (defaults to 16).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#title_size PowerpackV2#title_size}
	TitleSize *string `field:"optional" json:"titleSize" yaml:"titleSize"`
}

