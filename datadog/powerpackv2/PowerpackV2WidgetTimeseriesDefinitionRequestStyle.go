// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetTimeseriesDefinitionRequestStyle struct {
	// Whether to display value labels on the timeseries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#has_value_labels PowerpackV2#has_value_labels}
	HasValueLabels interface{} `field:"optional" json:"hasValueLabels" yaml:"hasValueLabels"`
	// The type of lines displayed. Valid values are `dashed`, `dotted`, `solid`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#line_type PowerpackV2#line_type}
	LineType *string `field:"optional" json:"lineType" yaml:"lineType"`
	// The width of line displayed. Valid values are `normal`, `thick`, `thin`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#line_width PowerpackV2#line_width}
	LineWidth *string `field:"optional" json:"lineWidth" yaml:"lineWidth"`
	// How to order series in timeseries visualizations. Valid values are `tags`, `values`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#order_by PowerpackV2#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// A color palette to apply to the widget. The available options are available at: https://docs.datadoghq.com/dashboards/widgets/timeseries/#appearance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#palette PowerpackV2#palette}
	Palette *string `field:"optional" json:"palette" yaml:"palette"`
}

