// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinition struct {
	// change_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#change_definition DashboardV2#change_definition}
	ChangeDefinition *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinition `field:"optional" json:"changeDefinition" yaml:"changeDefinition"`
	// geomap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#geomap_definition DashboardV2#geomap_definition}
	GeomapDefinition *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinition `field:"optional" json:"geomapDefinition" yaml:"geomapDefinition"`
	// query_table_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#query_table_definition DashboardV2#query_table_definition}
	QueryTableDefinition *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryTableDefinition `field:"optional" json:"queryTableDefinition" yaml:"queryTableDefinition"`
	// query_value_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#query_value_definition DashboardV2#query_value_definition}
	QueryValueDefinition *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinition `field:"optional" json:"queryValueDefinition" yaml:"queryValueDefinition"`
	// scatterplot_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#scatterplot_definition DashboardV2#scatterplot_definition}
	ScatterplotDefinition *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinition `field:"optional" json:"scatterplotDefinition" yaml:"scatterplotDefinition"`
	// sunburst_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#sunburst_definition DashboardV2#sunburst_definition}
	SunburstDefinition *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinition `field:"optional" json:"sunburstDefinition" yaml:"sunburstDefinition"`
	// timeseries_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#timeseries_definition DashboardV2#timeseries_definition}
	TimeseriesDefinition *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinition `field:"optional" json:"timeseriesDefinition" yaml:"timeseriesDefinition"`
	// toplist_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#toplist_definition DashboardV2#toplist_definition}
	ToplistDefinition *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinition `field:"optional" json:"toplistDefinition" yaml:"toplistDefinition"`
	// treemap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#treemap_definition DashboardV2#treemap_definition}
	TreemapDefinition *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinition `field:"optional" json:"treemapDefinition" yaml:"treemapDefinition"`
}

