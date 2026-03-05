// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetTimeseriesDefinitionRequestStyle struct {
	// The type of lines displayed. Valid values are `dashed`, `dotted`, `solid`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/powerpack#line_type Powerpack#line_type}
	LineType *string `field:"optional" json:"lineType" yaml:"lineType"`
	// The width of line displayed. Valid values are `normal`, `thick`, `thin`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/powerpack#line_width Powerpack#line_width}
	LineWidth *string `field:"optional" json:"lineWidth" yaml:"lineWidth"`
	// How to order series in timeseries visualizations. Valid values are `tags`, `values`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/powerpack#order_by Powerpack#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// A color palette to apply to the widget. The available options are available at: https://docs.datadoghq.com/dashboards/widgets/timeseries/#appearance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/powerpack#palette Powerpack#palette}
	Palette *string `field:"optional" json:"palette" yaml:"palette"`
}

