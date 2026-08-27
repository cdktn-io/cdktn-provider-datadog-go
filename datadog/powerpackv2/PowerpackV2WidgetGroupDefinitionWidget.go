// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidget struct {
	// alert_graph_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#alert_graph_definition PowerpackV2#alert_graph_definition}
	AlertGraphDefinition *PowerpackV2WidgetGroupDefinitionWidgetAlertGraphDefinition `field:"optional" json:"alertGraphDefinition" yaml:"alertGraphDefinition"`
	// alert_value_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#alert_value_definition PowerpackV2#alert_value_definition}
	AlertValueDefinition *PowerpackV2WidgetGroupDefinitionWidgetAlertValueDefinition `field:"optional" json:"alertValueDefinition" yaml:"alertValueDefinition"`
	// bar_chart_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#bar_chart_definition PowerpackV2#bar_chart_definition}
	BarChartDefinition *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinition `field:"optional" json:"barChartDefinition" yaml:"barChartDefinition"`
	// change_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#change_definition PowerpackV2#change_definition}
	ChangeDefinition *PowerpackV2WidgetGroupDefinitionWidgetChangeDefinition `field:"optional" json:"changeDefinition" yaml:"changeDefinition"`
	// check_status_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#check_status_definition PowerpackV2#check_status_definition}
	CheckStatusDefinition *PowerpackV2WidgetGroupDefinitionWidgetCheckStatusDefinition `field:"optional" json:"checkStatusDefinition" yaml:"checkStatusDefinition"`
	// cohort_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#cohort_definition PowerpackV2#cohort_definition}
	CohortDefinition *PowerpackV2WidgetGroupDefinitionWidgetCohortDefinition `field:"optional" json:"cohortDefinition" yaml:"cohortDefinition"`
	// distribution_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#distribution_definition PowerpackV2#distribution_definition}
	DistributionDefinition *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinition `field:"optional" json:"distributionDefinition" yaml:"distributionDefinition"`
	// event_stream_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#event_stream_definition PowerpackV2#event_stream_definition}
	EventStreamDefinition *PowerpackV2WidgetGroupDefinitionWidgetEventStreamDefinition `field:"optional" json:"eventStreamDefinition" yaml:"eventStreamDefinition"`
	// event_timeline_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#event_timeline_definition PowerpackV2#event_timeline_definition}
	EventTimelineDefinition *PowerpackV2WidgetGroupDefinitionWidgetEventTimelineDefinition `field:"optional" json:"eventTimelineDefinition" yaml:"eventTimelineDefinition"`
	// free_text_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#free_text_definition PowerpackV2#free_text_definition}
	FreeTextDefinition *PowerpackV2WidgetGroupDefinitionWidgetFreeTextDefinition `field:"optional" json:"freeTextDefinition" yaml:"freeTextDefinition"`
	// funnel_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#funnel_definition PowerpackV2#funnel_definition}
	FunnelDefinition *PowerpackV2WidgetGroupDefinitionWidgetFunnelDefinition `field:"optional" json:"funnelDefinition" yaml:"funnelDefinition"`
	// geomap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#geomap_definition PowerpackV2#geomap_definition}
	GeomapDefinition *PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinition `field:"optional" json:"geomapDefinition" yaml:"geomapDefinition"`
	// heatmap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#heatmap_definition PowerpackV2#heatmap_definition}
	HeatmapDefinition *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinition `field:"optional" json:"heatmapDefinition" yaml:"heatmapDefinition"`
	// hostmap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#hostmap_definition PowerpackV2#hostmap_definition}
	HostmapDefinition *PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinition `field:"optional" json:"hostmapDefinition" yaml:"hostmapDefinition"`
	// The ID of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#id PowerpackV2#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *float64 `field:"optional" json:"id" yaml:"id"`
	// iframe_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#iframe_definition PowerpackV2#iframe_definition}
	IframeDefinition *PowerpackV2WidgetGroupDefinitionWidgetIframeDefinition `field:"optional" json:"iframeDefinition" yaml:"iframeDefinition"`
	// image_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#image_definition PowerpackV2#image_definition}
	ImageDefinition *PowerpackV2WidgetGroupDefinitionWidgetImageDefinition `field:"optional" json:"imageDefinition" yaml:"imageDefinition"`
	// list_stream_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#list_stream_definition PowerpackV2#list_stream_definition}
	ListStreamDefinition *PowerpackV2WidgetGroupDefinitionWidgetListStreamDefinition `field:"optional" json:"listStreamDefinition" yaml:"listStreamDefinition"`
	// log_stream_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#log_stream_definition PowerpackV2#log_stream_definition}
	LogStreamDefinition *PowerpackV2WidgetGroupDefinitionWidgetLogStreamDefinition `field:"optional" json:"logStreamDefinition" yaml:"logStreamDefinition"`
	// manage_status_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#manage_status_definition PowerpackV2#manage_status_definition}
	ManageStatusDefinition *PowerpackV2WidgetGroupDefinitionWidgetManageStatusDefinition `field:"optional" json:"manageStatusDefinition" yaml:"manageStatusDefinition"`
	// note_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#note_definition PowerpackV2#note_definition}
	NoteDefinition *PowerpackV2WidgetGroupDefinitionWidgetNoteDefinition `field:"optional" json:"noteDefinition" yaml:"noteDefinition"`
	// point_plot_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#point_plot_definition PowerpackV2#point_plot_definition}
	PointPlotDefinition *PowerpackV2WidgetGroupDefinitionWidgetPointPlotDefinition `field:"optional" json:"pointPlotDefinition" yaml:"pointPlotDefinition"`
	// product_analytics_funnel_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#product_analytics_funnel_definition PowerpackV2#product_analytics_funnel_definition}
	ProductAnalyticsFunnelDefinition *PowerpackV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinition `field:"optional" json:"productAnalyticsFunnelDefinition" yaml:"productAnalyticsFunnelDefinition"`
	// query_table_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#query_table_definition PowerpackV2#query_table_definition}
	QueryTableDefinition *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinition `field:"optional" json:"queryTableDefinition" yaml:"queryTableDefinition"`
	// query_value_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#query_value_definition PowerpackV2#query_value_definition}
	QueryValueDefinition *PowerpackV2WidgetGroupDefinitionWidgetQueryValueDefinition `field:"optional" json:"queryValueDefinition" yaml:"queryValueDefinition"`
	// retention_curve_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#retention_curve_definition PowerpackV2#retention_curve_definition}
	RetentionCurveDefinition *PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinition `field:"optional" json:"retentionCurveDefinition" yaml:"retentionCurveDefinition"`
	// run_workflow_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#run_workflow_definition PowerpackV2#run_workflow_definition}
	RunWorkflowDefinition *PowerpackV2WidgetGroupDefinitionWidgetRunWorkflowDefinition `field:"optional" json:"runWorkflowDefinition" yaml:"runWorkflowDefinition"`
	// sankey_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#sankey_definition PowerpackV2#sankey_definition}
	SankeyDefinition *PowerpackV2WidgetGroupDefinitionWidgetSankeyDefinition `field:"optional" json:"sankeyDefinition" yaml:"sankeyDefinition"`
	// scatterplot_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#scatterplot_definition PowerpackV2#scatterplot_definition}
	ScatterplotDefinition *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinition `field:"optional" json:"scatterplotDefinition" yaml:"scatterplotDefinition"`
	// service_level_objective_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#service_level_objective_definition PowerpackV2#service_level_objective_definition}
	ServiceLevelObjectiveDefinition *PowerpackV2WidgetGroupDefinitionWidgetServiceLevelObjectiveDefinition `field:"optional" json:"serviceLevelObjectiveDefinition" yaml:"serviceLevelObjectiveDefinition"`
	// servicemap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#servicemap_definition PowerpackV2#servicemap_definition}
	ServicemapDefinition *PowerpackV2WidgetGroupDefinitionWidgetServicemapDefinition `field:"optional" json:"servicemapDefinition" yaml:"servicemapDefinition"`
	// slo_list_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#slo_list_definition PowerpackV2#slo_list_definition}
	SloListDefinition *PowerpackV2WidgetGroupDefinitionWidgetSloListDefinition `field:"optional" json:"sloListDefinition" yaml:"sloListDefinition"`
	// sunburst_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#sunburst_definition PowerpackV2#sunburst_definition}
	SunburstDefinition *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinition `field:"optional" json:"sunburstDefinition" yaml:"sunburstDefinition"`
	// timeseries_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#timeseries_definition PowerpackV2#timeseries_definition}
	TimeseriesDefinition *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinition `field:"optional" json:"timeseriesDefinition" yaml:"timeseriesDefinition"`
	// toplist_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#toplist_definition PowerpackV2#toplist_definition}
	ToplistDefinition *PowerpackV2WidgetGroupDefinitionWidgetToplistDefinition `field:"optional" json:"toplistDefinition" yaml:"toplistDefinition"`
	// topology_map_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#topology_map_definition PowerpackV2#topology_map_definition}
	TopologyMapDefinition *PowerpackV2WidgetGroupDefinitionWidgetTopologyMapDefinition `field:"optional" json:"topologyMapDefinition" yaml:"topologyMapDefinition"`
	// trace_service_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#trace_service_definition PowerpackV2#trace_service_definition}
	TraceServiceDefinition *PowerpackV2WidgetGroupDefinitionWidgetTraceServiceDefinition `field:"optional" json:"traceServiceDefinition" yaml:"traceServiceDefinition"`
	// treemap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#treemap_definition PowerpackV2#treemap_definition}
	TreemapDefinition *PowerpackV2WidgetGroupDefinitionWidgetTreemapDefinition `field:"optional" json:"treemapDefinition" yaml:"treemapDefinition"`
	// widget_layout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#widget_layout PowerpackV2#widget_layout}
	WidgetLayout *PowerpackV2WidgetGroupDefinitionWidgetWidgetLayout `field:"optional" json:"widgetLayout" yaml:"widgetLayout"`
	// wildcard_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#wildcard_definition PowerpackV2#wildcard_definition}
	WildcardDefinition *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinition `field:"optional" json:"wildcardDefinition" yaml:"wildcardDefinition"`
}

