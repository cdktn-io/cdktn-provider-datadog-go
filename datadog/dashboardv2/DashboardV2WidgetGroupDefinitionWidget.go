// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetGroupDefinitionWidget struct {
	// alert_graph_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#alert_graph_definition DashboardV2#alert_graph_definition}
	AlertGraphDefinition *DashboardV2WidgetGroupDefinitionWidgetAlertGraphDefinition `field:"optional" json:"alertGraphDefinition" yaml:"alertGraphDefinition"`
	// alert_value_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#alert_value_definition DashboardV2#alert_value_definition}
	AlertValueDefinition *DashboardV2WidgetGroupDefinitionWidgetAlertValueDefinition `field:"optional" json:"alertValueDefinition" yaml:"alertValueDefinition"`
	// bar_chart_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#bar_chart_definition DashboardV2#bar_chart_definition}
	BarChartDefinition *DashboardV2WidgetGroupDefinitionWidgetBarChartDefinition `field:"optional" json:"barChartDefinition" yaml:"barChartDefinition"`
	// change_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#change_definition DashboardV2#change_definition}
	ChangeDefinition *DashboardV2WidgetGroupDefinitionWidgetChangeDefinition `field:"optional" json:"changeDefinition" yaml:"changeDefinition"`
	// check_status_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#check_status_definition DashboardV2#check_status_definition}
	CheckStatusDefinition *DashboardV2WidgetGroupDefinitionWidgetCheckStatusDefinition `field:"optional" json:"checkStatusDefinition" yaml:"checkStatusDefinition"`
	// cohort_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#cohort_definition DashboardV2#cohort_definition}
	CohortDefinition *DashboardV2WidgetGroupDefinitionWidgetCohortDefinition `field:"optional" json:"cohortDefinition" yaml:"cohortDefinition"`
	// distribution_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#distribution_definition DashboardV2#distribution_definition}
	DistributionDefinition *DashboardV2WidgetGroupDefinitionWidgetDistributionDefinition `field:"optional" json:"distributionDefinition" yaml:"distributionDefinition"`
	// event_stream_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#event_stream_definition DashboardV2#event_stream_definition}
	EventStreamDefinition *DashboardV2WidgetGroupDefinitionWidgetEventStreamDefinition `field:"optional" json:"eventStreamDefinition" yaml:"eventStreamDefinition"`
	// event_timeline_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#event_timeline_definition DashboardV2#event_timeline_definition}
	EventTimelineDefinition *DashboardV2WidgetGroupDefinitionWidgetEventTimelineDefinition `field:"optional" json:"eventTimelineDefinition" yaml:"eventTimelineDefinition"`
	// free_text_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#free_text_definition DashboardV2#free_text_definition}
	FreeTextDefinition *DashboardV2WidgetGroupDefinitionWidgetFreeTextDefinition `field:"optional" json:"freeTextDefinition" yaml:"freeTextDefinition"`
	// funnel_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#funnel_definition DashboardV2#funnel_definition}
	FunnelDefinition *DashboardV2WidgetGroupDefinitionWidgetFunnelDefinition `field:"optional" json:"funnelDefinition" yaml:"funnelDefinition"`
	// geomap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#geomap_definition DashboardV2#geomap_definition}
	GeomapDefinition *DashboardV2WidgetGroupDefinitionWidgetGeomapDefinition `field:"optional" json:"geomapDefinition" yaml:"geomapDefinition"`
	// heatmap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#heatmap_definition DashboardV2#heatmap_definition}
	HeatmapDefinition *DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinition `field:"optional" json:"heatmapDefinition" yaml:"heatmapDefinition"`
	// hostmap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#hostmap_definition DashboardV2#hostmap_definition}
	HostmapDefinition *DashboardV2WidgetGroupDefinitionWidgetHostmapDefinition `field:"optional" json:"hostmapDefinition" yaml:"hostmapDefinition"`
	// The ID of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#id DashboardV2#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *float64 `field:"optional" json:"id" yaml:"id"`
	// iframe_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#iframe_definition DashboardV2#iframe_definition}
	IframeDefinition *DashboardV2WidgetGroupDefinitionWidgetIframeDefinition `field:"optional" json:"iframeDefinition" yaml:"iframeDefinition"`
	// image_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#image_definition DashboardV2#image_definition}
	ImageDefinition *DashboardV2WidgetGroupDefinitionWidgetImageDefinition `field:"optional" json:"imageDefinition" yaml:"imageDefinition"`
	// list_stream_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#list_stream_definition DashboardV2#list_stream_definition}
	ListStreamDefinition *DashboardV2WidgetGroupDefinitionWidgetListStreamDefinition `field:"optional" json:"listStreamDefinition" yaml:"listStreamDefinition"`
	// log_stream_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#log_stream_definition DashboardV2#log_stream_definition}
	LogStreamDefinition *DashboardV2WidgetGroupDefinitionWidgetLogStreamDefinition `field:"optional" json:"logStreamDefinition" yaml:"logStreamDefinition"`
	// manage_status_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#manage_status_definition DashboardV2#manage_status_definition}
	ManageStatusDefinition *DashboardV2WidgetGroupDefinitionWidgetManageStatusDefinition `field:"optional" json:"manageStatusDefinition" yaml:"manageStatusDefinition"`
	// note_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#note_definition DashboardV2#note_definition}
	NoteDefinition *DashboardV2WidgetGroupDefinitionWidgetNoteDefinition `field:"optional" json:"noteDefinition" yaml:"noteDefinition"`
	// point_plot_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#point_plot_definition DashboardV2#point_plot_definition}
	PointPlotDefinition *DashboardV2WidgetGroupDefinitionWidgetPointPlotDefinition `field:"optional" json:"pointPlotDefinition" yaml:"pointPlotDefinition"`
	// product_analytics_funnel_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#product_analytics_funnel_definition DashboardV2#product_analytics_funnel_definition}
	ProductAnalyticsFunnelDefinition *DashboardV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinition `field:"optional" json:"productAnalyticsFunnelDefinition" yaml:"productAnalyticsFunnelDefinition"`
	// query_table_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#query_table_definition DashboardV2#query_table_definition}
	QueryTableDefinition *DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinition `field:"optional" json:"queryTableDefinition" yaml:"queryTableDefinition"`
	// query_value_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#query_value_definition DashboardV2#query_value_definition}
	QueryValueDefinition *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinition `field:"optional" json:"queryValueDefinition" yaml:"queryValueDefinition"`
	// retention_curve_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#retention_curve_definition DashboardV2#retention_curve_definition}
	RetentionCurveDefinition *DashboardV2WidgetGroupDefinitionWidgetRetentionCurveDefinition `field:"optional" json:"retentionCurveDefinition" yaml:"retentionCurveDefinition"`
	// run_workflow_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#run_workflow_definition DashboardV2#run_workflow_definition}
	RunWorkflowDefinition *DashboardV2WidgetGroupDefinitionWidgetRunWorkflowDefinition `field:"optional" json:"runWorkflowDefinition" yaml:"runWorkflowDefinition"`
	// sankey_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#sankey_definition DashboardV2#sankey_definition}
	SankeyDefinition *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinition `field:"optional" json:"sankeyDefinition" yaml:"sankeyDefinition"`
	// scatterplot_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#scatterplot_definition DashboardV2#scatterplot_definition}
	ScatterplotDefinition *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinition `field:"optional" json:"scatterplotDefinition" yaml:"scatterplotDefinition"`
	// service_level_objective_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#service_level_objective_definition DashboardV2#service_level_objective_definition}
	ServiceLevelObjectiveDefinition *DashboardV2WidgetGroupDefinitionWidgetServiceLevelObjectiveDefinition `field:"optional" json:"serviceLevelObjectiveDefinition" yaml:"serviceLevelObjectiveDefinition"`
	// servicemap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#servicemap_definition DashboardV2#servicemap_definition}
	ServicemapDefinition *DashboardV2WidgetGroupDefinitionWidgetServicemapDefinition `field:"optional" json:"servicemapDefinition" yaml:"servicemapDefinition"`
	// slo_list_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#slo_list_definition DashboardV2#slo_list_definition}
	SloListDefinition *DashboardV2WidgetGroupDefinitionWidgetSloListDefinition `field:"optional" json:"sloListDefinition" yaml:"sloListDefinition"`
	// sunburst_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#sunburst_definition DashboardV2#sunburst_definition}
	SunburstDefinition *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinition `field:"optional" json:"sunburstDefinition" yaml:"sunburstDefinition"`
	// timeseries_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#timeseries_definition DashboardV2#timeseries_definition}
	TimeseriesDefinition *DashboardV2WidgetGroupDefinitionWidgetTimeseriesDefinition `field:"optional" json:"timeseriesDefinition" yaml:"timeseriesDefinition"`
	// toplist_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#toplist_definition DashboardV2#toplist_definition}
	ToplistDefinition *DashboardV2WidgetGroupDefinitionWidgetToplistDefinition `field:"optional" json:"toplistDefinition" yaml:"toplistDefinition"`
	// topology_map_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#topology_map_definition DashboardV2#topology_map_definition}
	TopologyMapDefinition *DashboardV2WidgetGroupDefinitionWidgetTopologyMapDefinition `field:"optional" json:"topologyMapDefinition" yaml:"topologyMapDefinition"`
	// trace_service_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#trace_service_definition DashboardV2#trace_service_definition}
	TraceServiceDefinition *DashboardV2WidgetGroupDefinitionWidgetTraceServiceDefinition `field:"optional" json:"traceServiceDefinition" yaml:"traceServiceDefinition"`
	// treemap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#treemap_definition DashboardV2#treemap_definition}
	TreemapDefinition *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinition `field:"optional" json:"treemapDefinition" yaml:"treemapDefinition"`
	// widget_layout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#widget_layout DashboardV2#widget_layout}
	WidgetLayout *DashboardV2WidgetGroupDefinitionWidgetWidgetLayout `field:"optional" json:"widgetLayout" yaml:"widgetLayout"`
	// wildcard_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#wildcard_definition DashboardV2#wildcard_definition}
	WildcardDefinition *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinition `field:"optional" json:"wildcardDefinition" yaml:"wildcardDefinition"`
}

