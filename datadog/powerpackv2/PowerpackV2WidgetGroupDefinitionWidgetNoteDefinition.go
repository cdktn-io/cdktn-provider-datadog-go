// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetNoteDefinition struct {
	// The content of the note.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#content PowerpackV2#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The background color of the note.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#background_color PowerpackV2#background_color}
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// The description of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#description PowerpackV2#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The size of the text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#font_size PowerpackV2#font_size}
	FontSize *string `field:"optional" json:"fontSize" yaml:"fontSize"`
	// Whether to add padding or not. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#has_padding PowerpackV2#has_padding}
	HasPadding interface{} `field:"optional" json:"hasPadding" yaml:"hasPadding"`
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
	// Whether to show a tick or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#show_tick PowerpackV2#show_tick}
	ShowTick interface{} `field:"optional" json:"showTick" yaml:"showTick"`
	// The alignment of the widget's text. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#text_align PowerpackV2#text_align}
	TextAlign *string `field:"optional" json:"textAlign" yaml:"textAlign"`
	// When `tick = true`, a string indicating on which side of the widget the tick should be displayed.
	//
	// Valid values are `bottom`, `left`, `right`, `top`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#tick_edge PowerpackV2#tick_edge}
	TickEdge *string `field:"optional" json:"tickEdge" yaml:"tickEdge"`
	// When `tick = true`, a string with a percent sign indicating the position of the tick, for example: `tick_pos = "50%"` is centered alignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#tick_pos PowerpackV2#tick_pos}
	TickPos *string `field:"optional" json:"tickPos" yaml:"tickPos"`
	// time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#time PowerpackV2#time}
	Time *PowerpackV2WidgetGroupDefinitionWidgetNoteDefinitionTime `field:"optional" json:"time" yaml:"time"`
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
	// The vertical alignment for the widget. Valid values are `center`, `top`, `bottom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/powerpack_v2#vertical_align PowerpackV2#vertical_align}
	VerticalAlign *string `field:"optional" json:"verticalAlign" yaml:"verticalAlign"`
}

