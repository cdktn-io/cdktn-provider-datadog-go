// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionMarker struct {
	// Value to apply.
	//
	// Can be a single value `y = 15` or a range of values `0 < y < 10`. For Distribution widgets with `display_type` set to `percentile`, this should be a numeric percentile value (for example, `90` for P90).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#value DashboardV2#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Combination of a severity (`error`, `warning`, `ok`, or `info`) and a line type (`dashed`, `solid`, or `bold`).
	//
	// For Distribution widgets, this can be set to `percentile`. Example: `error dashed`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#display_type DashboardV2#display_type}
	DisplayType *string `field:"optional" json:"displayType" yaml:"displayType"`
	// Label to display over the marker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#label DashboardV2#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Timestamp for the marker position.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#time DashboardV2#time}
	Time *string `field:"optional" json:"time" yaml:"time"`
}

