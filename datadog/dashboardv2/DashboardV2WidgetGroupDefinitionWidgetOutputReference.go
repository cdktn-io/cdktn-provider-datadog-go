// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetGroupDefinitionWidgetOutputReference interface {
	cdktn.ComplexObject
	AlertGraphDefinition() DashboardV2WidgetGroupDefinitionWidgetAlertGraphDefinitionOutputReference
	AlertGraphDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetAlertGraphDefinition
	AlertValueDefinition() DashboardV2WidgetGroupDefinitionWidgetAlertValueDefinitionOutputReference
	AlertValueDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetAlertValueDefinition
	BarChartDefinition() DashboardV2WidgetGroupDefinitionWidgetBarChartDefinitionOutputReference
	BarChartDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetBarChartDefinition
	ChangeDefinition() DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionOutputReference
	ChangeDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetChangeDefinition
	CheckStatusDefinition() DashboardV2WidgetGroupDefinitionWidgetCheckStatusDefinitionOutputReference
	CheckStatusDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetCheckStatusDefinition
	CohortDefinition() DashboardV2WidgetGroupDefinitionWidgetCohortDefinitionOutputReference
	CohortDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetCohortDefinition
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DistributionDefinition() DashboardV2WidgetGroupDefinitionWidgetDistributionDefinitionOutputReference
	DistributionDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetDistributionDefinition
	EventStreamDefinition() DashboardV2WidgetGroupDefinitionWidgetEventStreamDefinitionOutputReference
	EventStreamDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetEventStreamDefinition
	EventTimelineDefinition() DashboardV2WidgetGroupDefinitionWidgetEventTimelineDefinitionOutputReference
	EventTimelineDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetEventTimelineDefinition
	// Experimental.
	Fqn() *string
	FreeTextDefinition() DashboardV2WidgetGroupDefinitionWidgetFreeTextDefinitionOutputReference
	FreeTextDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetFreeTextDefinition
	FunnelDefinition() DashboardV2WidgetGroupDefinitionWidgetFunnelDefinitionOutputReference
	FunnelDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetFunnelDefinition
	GeomapDefinition() DashboardV2WidgetGroupDefinitionWidgetGeomapDefinitionOutputReference
	GeomapDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetGeomapDefinition
	HeatmapDefinition() DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference
	HeatmapDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinition
	HostmapDefinition() DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionOutputReference
	HostmapDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetHostmapDefinition
	Id() *float64
	SetId(val *float64)
	IdInput() *float64
	IframeDefinition() DashboardV2WidgetGroupDefinitionWidgetIframeDefinitionOutputReference
	IframeDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetIframeDefinition
	ImageDefinition() DashboardV2WidgetGroupDefinitionWidgetImageDefinitionOutputReference
	ImageDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetImageDefinition
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ListStreamDefinition() DashboardV2WidgetGroupDefinitionWidgetListStreamDefinitionOutputReference
	ListStreamDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetListStreamDefinition
	LogStreamDefinition() DashboardV2WidgetGroupDefinitionWidgetLogStreamDefinitionOutputReference
	LogStreamDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetLogStreamDefinition
	ManageStatusDefinition() DashboardV2WidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference
	ManageStatusDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetManageStatusDefinition
	NoteDefinition() DashboardV2WidgetGroupDefinitionWidgetNoteDefinitionOutputReference
	NoteDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetNoteDefinition
	PointPlotDefinition() DashboardV2WidgetGroupDefinitionWidgetPointPlotDefinitionOutputReference
	PointPlotDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetPointPlotDefinition
	ProductAnalyticsFunnelDefinition() DashboardV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinitionOutputReference
	ProductAnalyticsFunnelDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinition
	QueryTableDefinition() DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionOutputReference
	QueryTableDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinition
	QueryValueDefinition() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionOutputReference
	QueryValueDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinition
	RetentionCurveDefinition() DashboardV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionOutputReference
	RetentionCurveDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetRetentionCurveDefinition
	RunWorkflowDefinition() DashboardV2WidgetGroupDefinitionWidgetRunWorkflowDefinitionOutputReference
	RunWorkflowDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetRunWorkflowDefinition
	SankeyDefinition() DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference
	SankeyDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinition
	ScatterplotDefinition() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionOutputReference
	ScatterplotDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinition
	ServiceLevelObjectiveDefinition() DashboardV2WidgetGroupDefinitionWidgetServiceLevelObjectiveDefinitionOutputReference
	ServiceLevelObjectiveDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetServiceLevelObjectiveDefinition
	ServicemapDefinition() DashboardV2WidgetGroupDefinitionWidgetServicemapDefinitionOutputReference
	ServicemapDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetServicemapDefinition
	SloListDefinition() DashboardV2WidgetGroupDefinitionWidgetSloListDefinitionOutputReference
	SloListDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetSloListDefinition
	SunburstDefinition() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionOutputReference
	SunburstDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinition
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeseriesDefinition() DashboardV2WidgetGroupDefinitionWidgetTimeseriesDefinitionOutputReference
	TimeseriesDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetTimeseriesDefinition
	ToplistDefinition() DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionOutputReference
	ToplistDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetToplistDefinition
	TopologyMapDefinition() DashboardV2WidgetGroupDefinitionWidgetTopologyMapDefinitionOutputReference
	TopologyMapDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetTopologyMapDefinition
	TraceServiceDefinition() DashboardV2WidgetGroupDefinitionWidgetTraceServiceDefinitionOutputReference
	TraceServiceDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetTraceServiceDefinition
	TreemapDefinition() DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionOutputReference
	TreemapDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinition
	WidgetLayout() DashboardV2WidgetGroupDefinitionWidgetWidgetLayoutOutputReference
	WidgetLayoutInput() *DashboardV2WidgetGroupDefinitionWidgetWidgetLayout
	WildcardDefinition() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionOutputReference
	WildcardDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinition
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutAlertGraphDefinition(value *DashboardV2WidgetGroupDefinitionWidgetAlertGraphDefinition)
	PutAlertValueDefinition(value *DashboardV2WidgetGroupDefinitionWidgetAlertValueDefinition)
	PutBarChartDefinition(value *DashboardV2WidgetGroupDefinitionWidgetBarChartDefinition)
	PutChangeDefinition(value *DashboardV2WidgetGroupDefinitionWidgetChangeDefinition)
	PutCheckStatusDefinition(value *DashboardV2WidgetGroupDefinitionWidgetCheckStatusDefinition)
	PutCohortDefinition(value *DashboardV2WidgetGroupDefinitionWidgetCohortDefinition)
	PutDistributionDefinition(value *DashboardV2WidgetGroupDefinitionWidgetDistributionDefinition)
	PutEventStreamDefinition(value *DashboardV2WidgetGroupDefinitionWidgetEventStreamDefinition)
	PutEventTimelineDefinition(value *DashboardV2WidgetGroupDefinitionWidgetEventTimelineDefinition)
	PutFreeTextDefinition(value *DashboardV2WidgetGroupDefinitionWidgetFreeTextDefinition)
	PutFunnelDefinition(value *DashboardV2WidgetGroupDefinitionWidgetFunnelDefinition)
	PutGeomapDefinition(value *DashboardV2WidgetGroupDefinitionWidgetGeomapDefinition)
	PutHeatmapDefinition(value *DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinition)
	PutHostmapDefinition(value *DashboardV2WidgetGroupDefinitionWidgetHostmapDefinition)
	PutIframeDefinition(value *DashboardV2WidgetGroupDefinitionWidgetIframeDefinition)
	PutImageDefinition(value *DashboardV2WidgetGroupDefinitionWidgetImageDefinition)
	PutListStreamDefinition(value *DashboardV2WidgetGroupDefinitionWidgetListStreamDefinition)
	PutLogStreamDefinition(value *DashboardV2WidgetGroupDefinitionWidgetLogStreamDefinition)
	PutManageStatusDefinition(value *DashboardV2WidgetGroupDefinitionWidgetManageStatusDefinition)
	PutNoteDefinition(value *DashboardV2WidgetGroupDefinitionWidgetNoteDefinition)
	PutPointPlotDefinition(value *DashboardV2WidgetGroupDefinitionWidgetPointPlotDefinition)
	PutProductAnalyticsFunnelDefinition(value *DashboardV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinition)
	PutQueryTableDefinition(value *DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinition)
	PutQueryValueDefinition(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinition)
	PutRetentionCurveDefinition(value *DashboardV2WidgetGroupDefinitionWidgetRetentionCurveDefinition)
	PutRunWorkflowDefinition(value *DashboardV2WidgetGroupDefinitionWidgetRunWorkflowDefinition)
	PutSankeyDefinition(value *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinition)
	PutScatterplotDefinition(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinition)
	PutServiceLevelObjectiveDefinition(value *DashboardV2WidgetGroupDefinitionWidgetServiceLevelObjectiveDefinition)
	PutServicemapDefinition(value *DashboardV2WidgetGroupDefinitionWidgetServicemapDefinition)
	PutSloListDefinition(value *DashboardV2WidgetGroupDefinitionWidgetSloListDefinition)
	PutSunburstDefinition(value *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinition)
	PutTimeseriesDefinition(value *DashboardV2WidgetGroupDefinitionWidgetTimeseriesDefinition)
	PutToplistDefinition(value *DashboardV2WidgetGroupDefinitionWidgetToplistDefinition)
	PutTopologyMapDefinition(value *DashboardV2WidgetGroupDefinitionWidgetTopologyMapDefinition)
	PutTraceServiceDefinition(value *DashboardV2WidgetGroupDefinitionWidgetTraceServiceDefinition)
	PutTreemapDefinition(value *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinition)
	PutWidgetLayout(value *DashboardV2WidgetGroupDefinitionWidgetWidgetLayout)
	PutWildcardDefinition(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinition)
	ResetAlertGraphDefinition()
	ResetAlertValueDefinition()
	ResetBarChartDefinition()
	ResetChangeDefinition()
	ResetCheckStatusDefinition()
	ResetCohortDefinition()
	ResetDistributionDefinition()
	ResetEventStreamDefinition()
	ResetEventTimelineDefinition()
	ResetFreeTextDefinition()
	ResetFunnelDefinition()
	ResetGeomapDefinition()
	ResetHeatmapDefinition()
	ResetHostmapDefinition()
	ResetId()
	ResetIframeDefinition()
	ResetImageDefinition()
	ResetListStreamDefinition()
	ResetLogStreamDefinition()
	ResetManageStatusDefinition()
	ResetNoteDefinition()
	ResetPointPlotDefinition()
	ResetProductAnalyticsFunnelDefinition()
	ResetQueryTableDefinition()
	ResetQueryValueDefinition()
	ResetRetentionCurveDefinition()
	ResetRunWorkflowDefinition()
	ResetSankeyDefinition()
	ResetScatterplotDefinition()
	ResetServiceLevelObjectiveDefinition()
	ResetServicemapDefinition()
	ResetSloListDefinition()
	ResetSunburstDefinition()
	ResetTimeseriesDefinition()
	ResetToplistDefinition()
	ResetTopologyMapDefinition()
	ResetTraceServiceDefinition()
	ResetTreemapDefinition()
	ResetWidgetLayout()
	ResetWildcardDefinition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetGroupDefinitionWidgetOutputReference
type jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) AlertGraphDefinition() DashboardV2WidgetGroupDefinitionWidgetAlertGraphDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetAlertGraphDefinitionOutputReference
	_jsii_.Get(
		j,
		"alertGraphDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) AlertGraphDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetAlertGraphDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetAlertGraphDefinition
	_jsii_.Get(
		j,
		"alertGraphDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) AlertValueDefinition() DashboardV2WidgetGroupDefinitionWidgetAlertValueDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetAlertValueDefinitionOutputReference
	_jsii_.Get(
		j,
		"alertValueDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) AlertValueDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetAlertValueDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetAlertValueDefinition
	_jsii_.Get(
		j,
		"alertValueDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) BarChartDefinition() DashboardV2WidgetGroupDefinitionWidgetBarChartDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetBarChartDefinitionOutputReference
	_jsii_.Get(
		j,
		"barChartDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) BarChartDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetBarChartDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetBarChartDefinition
	_jsii_.Get(
		j,
		"barChartDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ChangeDefinition() DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionOutputReference
	_jsii_.Get(
		j,
		"changeDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ChangeDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetChangeDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetChangeDefinition
	_jsii_.Get(
		j,
		"changeDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) CheckStatusDefinition() DashboardV2WidgetGroupDefinitionWidgetCheckStatusDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetCheckStatusDefinitionOutputReference
	_jsii_.Get(
		j,
		"checkStatusDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) CheckStatusDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetCheckStatusDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetCheckStatusDefinition
	_jsii_.Get(
		j,
		"checkStatusDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) CohortDefinition() DashboardV2WidgetGroupDefinitionWidgetCohortDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetCohortDefinitionOutputReference
	_jsii_.Get(
		j,
		"cohortDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) CohortDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetCohortDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetCohortDefinition
	_jsii_.Get(
		j,
		"cohortDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) DistributionDefinition() DashboardV2WidgetGroupDefinitionWidgetDistributionDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetDistributionDefinitionOutputReference
	_jsii_.Get(
		j,
		"distributionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) DistributionDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetDistributionDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetDistributionDefinition
	_jsii_.Get(
		j,
		"distributionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) EventStreamDefinition() DashboardV2WidgetGroupDefinitionWidgetEventStreamDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetEventStreamDefinitionOutputReference
	_jsii_.Get(
		j,
		"eventStreamDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) EventStreamDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetEventStreamDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetEventStreamDefinition
	_jsii_.Get(
		j,
		"eventStreamDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) EventTimelineDefinition() DashboardV2WidgetGroupDefinitionWidgetEventTimelineDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetEventTimelineDefinitionOutputReference
	_jsii_.Get(
		j,
		"eventTimelineDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) EventTimelineDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetEventTimelineDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetEventTimelineDefinition
	_jsii_.Get(
		j,
		"eventTimelineDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) FreeTextDefinition() DashboardV2WidgetGroupDefinitionWidgetFreeTextDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetFreeTextDefinitionOutputReference
	_jsii_.Get(
		j,
		"freeTextDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) FreeTextDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetFreeTextDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetFreeTextDefinition
	_jsii_.Get(
		j,
		"freeTextDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) FunnelDefinition() DashboardV2WidgetGroupDefinitionWidgetFunnelDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetFunnelDefinitionOutputReference
	_jsii_.Get(
		j,
		"funnelDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) FunnelDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetFunnelDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetFunnelDefinition
	_jsii_.Get(
		j,
		"funnelDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) GeomapDefinition() DashboardV2WidgetGroupDefinitionWidgetGeomapDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetGeomapDefinitionOutputReference
	_jsii_.Get(
		j,
		"geomapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) GeomapDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetGeomapDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetGeomapDefinition
	_jsii_.Get(
		j,
		"geomapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) HeatmapDefinition() DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference
	_jsii_.Get(
		j,
		"heatmapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) HeatmapDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinition
	_jsii_.Get(
		j,
		"heatmapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) HostmapDefinition() DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionOutputReference
	_jsii_.Get(
		j,
		"hostmapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) HostmapDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetHostmapDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetHostmapDefinition
	_jsii_.Get(
		j,
		"hostmapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) IdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) IframeDefinition() DashboardV2WidgetGroupDefinitionWidgetIframeDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetIframeDefinitionOutputReference
	_jsii_.Get(
		j,
		"iframeDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) IframeDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetIframeDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetIframeDefinition
	_jsii_.Get(
		j,
		"iframeDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ImageDefinition() DashboardV2WidgetGroupDefinitionWidgetImageDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetImageDefinitionOutputReference
	_jsii_.Get(
		j,
		"imageDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ImageDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetImageDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetImageDefinition
	_jsii_.Get(
		j,
		"imageDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ListStreamDefinition() DashboardV2WidgetGroupDefinitionWidgetListStreamDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetListStreamDefinitionOutputReference
	_jsii_.Get(
		j,
		"listStreamDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ListStreamDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetListStreamDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetListStreamDefinition
	_jsii_.Get(
		j,
		"listStreamDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) LogStreamDefinition() DashboardV2WidgetGroupDefinitionWidgetLogStreamDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetLogStreamDefinitionOutputReference
	_jsii_.Get(
		j,
		"logStreamDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) LogStreamDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetLogStreamDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetLogStreamDefinition
	_jsii_.Get(
		j,
		"logStreamDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ManageStatusDefinition() DashboardV2WidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference
	_jsii_.Get(
		j,
		"manageStatusDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ManageStatusDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetManageStatusDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetManageStatusDefinition
	_jsii_.Get(
		j,
		"manageStatusDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) NoteDefinition() DashboardV2WidgetGroupDefinitionWidgetNoteDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetNoteDefinitionOutputReference
	_jsii_.Get(
		j,
		"noteDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) NoteDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetNoteDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetNoteDefinition
	_jsii_.Get(
		j,
		"noteDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PointPlotDefinition() DashboardV2WidgetGroupDefinitionWidgetPointPlotDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetPointPlotDefinitionOutputReference
	_jsii_.Get(
		j,
		"pointPlotDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PointPlotDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetPointPlotDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetPointPlotDefinition
	_jsii_.Get(
		j,
		"pointPlotDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ProductAnalyticsFunnelDefinition() DashboardV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinitionOutputReference
	_jsii_.Get(
		j,
		"productAnalyticsFunnelDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ProductAnalyticsFunnelDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinition
	_jsii_.Get(
		j,
		"productAnalyticsFunnelDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) QueryTableDefinition() DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionOutputReference
	_jsii_.Get(
		j,
		"queryTableDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) QueryTableDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinition
	_jsii_.Get(
		j,
		"queryTableDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) QueryValueDefinition() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionOutputReference
	_jsii_.Get(
		j,
		"queryValueDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) QueryValueDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinition
	_jsii_.Get(
		j,
		"queryValueDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) RetentionCurveDefinition() DashboardV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionOutputReference
	_jsii_.Get(
		j,
		"retentionCurveDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) RetentionCurveDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetRetentionCurveDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetRetentionCurveDefinition
	_jsii_.Get(
		j,
		"retentionCurveDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) RunWorkflowDefinition() DashboardV2WidgetGroupDefinitionWidgetRunWorkflowDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetRunWorkflowDefinitionOutputReference
	_jsii_.Get(
		j,
		"runWorkflowDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) RunWorkflowDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetRunWorkflowDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetRunWorkflowDefinition
	_jsii_.Get(
		j,
		"runWorkflowDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) SankeyDefinition() DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference
	_jsii_.Get(
		j,
		"sankeyDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) SankeyDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinition
	_jsii_.Get(
		j,
		"sankeyDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ScatterplotDefinition() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionOutputReference
	_jsii_.Get(
		j,
		"scatterplotDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ScatterplotDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinition
	_jsii_.Get(
		j,
		"scatterplotDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ServiceLevelObjectiveDefinition() DashboardV2WidgetGroupDefinitionWidgetServiceLevelObjectiveDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetServiceLevelObjectiveDefinitionOutputReference
	_jsii_.Get(
		j,
		"serviceLevelObjectiveDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ServiceLevelObjectiveDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetServiceLevelObjectiveDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetServiceLevelObjectiveDefinition
	_jsii_.Get(
		j,
		"serviceLevelObjectiveDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ServicemapDefinition() DashboardV2WidgetGroupDefinitionWidgetServicemapDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetServicemapDefinitionOutputReference
	_jsii_.Get(
		j,
		"servicemapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ServicemapDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetServicemapDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetServicemapDefinition
	_jsii_.Get(
		j,
		"servicemapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) SloListDefinition() DashboardV2WidgetGroupDefinitionWidgetSloListDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetSloListDefinitionOutputReference
	_jsii_.Get(
		j,
		"sloListDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) SloListDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetSloListDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetSloListDefinition
	_jsii_.Get(
		j,
		"sloListDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) SunburstDefinition() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionOutputReference
	_jsii_.Get(
		j,
		"sunburstDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) SunburstDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinition
	_jsii_.Get(
		j,
		"sunburstDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) TimeseriesDefinition() DashboardV2WidgetGroupDefinitionWidgetTimeseriesDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetTimeseriesDefinitionOutputReference
	_jsii_.Get(
		j,
		"timeseriesDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) TimeseriesDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetTimeseriesDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetTimeseriesDefinition
	_jsii_.Get(
		j,
		"timeseriesDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ToplistDefinition() DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionOutputReference
	_jsii_.Get(
		j,
		"toplistDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ToplistDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetToplistDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetToplistDefinition
	_jsii_.Get(
		j,
		"toplistDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) TopologyMapDefinition() DashboardV2WidgetGroupDefinitionWidgetTopologyMapDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetTopologyMapDefinitionOutputReference
	_jsii_.Get(
		j,
		"topologyMapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) TopologyMapDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetTopologyMapDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetTopologyMapDefinition
	_jsii_.Get(
		j,
		"topologyMapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) TraceServiceDefinition() DashboardV2WidgetGroupDefinitionWidgetTraceServiceDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetTraceServiceDefinitionOutputReference
	_jsii_.Get(
		j,
		"traceServiceDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) TraceServiceDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetTraceServiceDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetTraceServiceDefinition
	_jsii_.Get(
		j,
		"traceServiceDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) TreemapDefinition() DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionOutputReference
	_jsii_.Get(
		j,
		"treemapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) TreemapDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinition
	_jsii_.Get(
		j,
		"treemapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) WidgetLayout() DashboardV2WidgetGroupDefinitionWidgetWidgetLayoutOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetWidgetLayoutOutputReference
	_jsii_.Get(
		j,
		"widgetLayout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) WidgetLayoutInput() *DashboardV2WidgetGroupDefinitionWidgetWidgetLayout {
	var returns *DashboardV2WidgetGroupDefinitionWidgetWidgetLayout
	_jsii_.Get(
		j,
		"widgetLayoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) WildcardDefinition() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionOutputReference
	_jsii_.Get(
		j,
		"wildcardDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) WildcardDefinitionInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinition
	_jsii_.Get(
		j,
		"wildcardDefinitionInput",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetGroupDefinitionWidgetOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardV2WidgetGroupDefinitionWidgetOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetGroupDefinitionWidgetOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetGroupDefinitionWidgetOutputReference_Override(d DashboardV2WidgetGroupDefinitionWidgetOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference)SetId(val *float64) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutAlertGraphDefinition(value *DashboardV2WidgetGroupDefinitionWidgetAlertGraphDefinition) {
	if err := d.validatePutAlertGraphDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAlertGraphDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutAlertValueDefinition(value *DashboardV2WidgetGroupDefinitionWidgetAlertValueDefinition) {
	if err := d.validatePutAlertValueDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAlertValueDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutBarChartDefinition(value *DashboardV2WidgetGroupDefinitionWidgetBarChartDefinition) {
	if err := d.validatePutBarChartDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBarChartDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutChangeDefinition(value *DashboardV2WidgetGroupDefinitionWidgetChangeDefinition) {
	if err := d.validatePutChangeDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putChangeDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutCheckStatusDefinition(value *DashboardV2WidgetGroupDefinitionWidgetCheckStatusDefinition) {
	if err := d.validatePutCheckStatusDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCheckStatusDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutCohortDefinition(value *DashboardV2WidgetGroupDefinitionWidgetCohortDefinition) {
	if err := d.validatePutCohortDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCohortDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutDistributionDefinition(value *DashboardV2WidgetGroupDefinitionWidgetDistributionDefinition) {
	if err := d.validatePutDistributionDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDistributionDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutEventStreamDefinition(value *DashboardV2WidgetGroupDefinitionWidgetEventStreamDefinition) {
	if err := d.validatePutEventStreamDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventStreamDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutEventTimelineDefinition(value *DashboardV2WidgetGroupDefinitionWidgetEventTimelineDefinition) {
	if err := d.validatePutEventTimelineDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventTimelineDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutFreeTextDefinition(value *DashboardV2WidgetGroupDefinitionWidgetFreeTextDefinition) {
	if err := d.validatePutFreeTextDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFreeTextDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutFunnelDefinition(value *DashboardV2WidgetGroupDefinitionWidgetFunnelDefinition) {
	if err := d.validatePutFunnelDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFunnelDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutGeomapDefinition(value *DashboardV2WidgetGroupDefinitionWidgetGeomapDefinition) {
	if err := d.validatePutGeomapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGeomapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutHeatmapDefinition(value *DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinition) {
	if err := d.validatePutHeatmapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHeatmapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutHostmapDefinition(value *DashboardV2WidgetGroupDefinitionWidgetHostmapDefinition) {
	if err := d.validatePutHostmapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHostmapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutIframeDefinition(value *DashboardV2WidgetGroupDefinitionWidgetIframeDefinition) {
	if err := d.validatePutIframeDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIframeDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutImageDefinition(value *DashboardV2WidgetGroupDefinitionWidgetImageDefinition) {
	if err := d.validatePutImageDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putImageDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutListStreamDefinition(value *DashboardV2WidgetGroupDefinitionWidgetListStreamDefinition) {
	if err := d.validatePutListStreamDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putListStreamDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutLogStreamDefinition(value *DashboardV2WidgetGroupDefinitionWidgetLogStreamDefinition) {
	if err := d.validatePutLogStreamDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLogStreamDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutManageStatusDefinition(value *DashboardV2WidgetGroupDefinitionWidgetManageStatusDefinition) {
	if err := d.validatePutManageStatusDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putManageStatusDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutNoteDefinition(value *DashboardV2WidgetGroupDefinitionWidgetNoteDefinition) {
	if err := d.validatePutNoteDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNoteDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutPointPlotDefinition(value *DashboardV2WidgetGroupDefinitionWidgetPointPlotDefinition) {
	if err := d.validatePutPointPlotDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPointPlotDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutProductAnalyticsFunnelDefinition(value *DashboardV2WidgetGroupDefinitionWidgetProductAnalyticsFunnelDefinition) {
	if err := d.validatePutProductAnalyticsFunnelDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProductAnalyticsFunnelDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutQueryTableDefinition(value *DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinition) {
	if err := d.validatePutQueryTableDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQueryTableDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutQueryValueDefinition(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinition) {
	if err := d.validatePutQueryValueDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQueryValueDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutRetentionCurveDefinition(value *DashboardV2WidgetGroupDefinitionWidgetRetentionCurveDefinition) {
	if err := d.validatePutRetentionCurveDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRetentionCurveDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutRunWorkflowDefinition(value *DashboardV2WidgetGroupDefinitionWidgetRunWorkflowDefinition) {
	if err := d.validatePutRunWorkflowDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRunWorkflowDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutSankeyDefinition(value *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinition) {
	if err := d.validatePutSankeyDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSankeyDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutScatterplotDefinition(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinition) {
	if err := d.validatePutScatterplotDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putScatterplotDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutServiceLevelObjectiveDefinition(value *DashboardV2WidgetGroupDefinitionWidgetServiceLevelObjectiveDefinition) {
	if err := d.validatePutServiceLevelObjectiveDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServiceLevelObjectiveDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutServicemapDefinition(value *DashboardV2WidgetGroupDefinitionWidgetServicemapDefinition) {
	if err := d.validatePutServicemapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServicemapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutSloListDefinition(value *DashboardV2WidgetGroupDefinitionWidgetSloListDefinition) {
	if err := d.validatePutSloListDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSloListDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutSunburstDefinition(value *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinition) {
	if err := d.validatePutSunburstDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSunburstDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutTimeseriesDefinition(value *DashboardV2WidgetGroupDefinitionWidgetTimeseriesDefinition) {
	if err := d.validatePutTimeseriesDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeseriesDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutToplistDefinition(value *DashboardV2WidgetGroupDefinitionWidgetToplistDefinition) {
	if err := d.validatePutToplistDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putToplistDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutTopologyMapDefinition(value *DashboardV2WidgetGroupDefinitionWidgetTopologyMapDefinition) {
	if err := d.validatePutTopologyMapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTopologyMapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutTraceServiceDefinition(value *DashboardV2WidgetGroupDefinitionWidgetTraceServiceDefinition) {
	if err := d.validatePutTraceServiceDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTraceServiceDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutTreemapDefinition(value *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinition) {
	if err := d.validatePutTreemapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTreemapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutWidgetLayout(value *DashboardV2WidgetGroupDefinitionWidgetWidgetLayout) {
	if err := d.validatePutWidgetLayoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWidgetLayout",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) PutWildcardDefinition(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinition) {
	if err := d.validatePutWildcardDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWildcardDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetAlertGraphDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetAlertGraphDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetAlertValueDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetAlertValueDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetBarChartDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetBarChartDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetChangeDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetChangeDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetCheckStatusDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetCheckStatusDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetCohortDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetCohortDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetDistributionDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetDistributionDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetEventStreamDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetEventStreamDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetEventTimelineDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetEventTimelineDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetFreeTextDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetFreeTextDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetFunnelDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetFunnelDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetGeomapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetGeomapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetHeatmapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetHeatmapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetHostmapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetHostmapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetIframeDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetIframeDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetImageDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetImageDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetListStreamDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetListStreamDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetLogStreamDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetLogStreamDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetManageStatusDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetManageStatusDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetNoteDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetNoteDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetPointPlotDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetPointPlotDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetProductAnalyticsFunnelDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetProductAnalyticsFunnelDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetQueryTableDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryTableDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetQueryValueDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryValueDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetRetentionCurveDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetRetentionCurveDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetRunWorkflowDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetRunWorkflowDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetSankeyDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetSankeyDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetScatterplotDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetScatterplotDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetServiceLevelObjectiveDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceLevelObjectiveDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetServicemapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetServicemapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetSloListDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetSloListDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetSunburstDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetSunburstDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetTimeseriesDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeseriesDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetToplistDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetToplistDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetTopologyMapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetTopologyMapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetTraceServiceDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetTraceServiceDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetTreemapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetTreemapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetWidgetLayout() {
	_jsii_.InvokeVoid(
		d,
		"resetWidgetLayout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ResetWildcardDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetWildcardDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

