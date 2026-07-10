// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetOutputReference interface {
	cdktn.ComplexObject
	AlertGraphDefinition() DashboardV2WidgetAlertGraphDefinitionOutputReference
	AlertGraphDefinitionInput() *DashboardV2WidgetAlertGraphDefinition
	AlertValueDefinition() DashboardV2WidgetAlertValueDefinitionOutputReference
	AlertValueDefinitionInput() *DashboardV2WidgetAlertValueDefinition
	BarChartDefinition() DashboardV2WidgetBarChartDefinitionOutputReference
	BarChartDefinitionInput() *DashboardV2WidgetBarChartDefinition
	ChangeDefinition() DashboardV2WidgetChangeDefinitionOutputReference
	ChangeDefinitionInput() *DashboardV2WidgetChangeDefinition
	CheckStatusDefinition() DashboardV2WidgetCheckStatusDefinitionOutputReference
	CheckStatusDefinitionInput() *DashboardV2WidgetCheckStatusDefinition
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
	DistributionDefinition() DashboardV2WidgetDistributionDefinitionOutputReference
	DistributionDefinitionInput() *DashboardV2WidgetDistributionDefinition
	EventStreamDefinition() DashboardV2WidgetEventStreamDefinitionOutputReference
	EventStreamDefinitionInput() *DashboardV2WidgetEventStreamDefinition
	EventTimelineDefinition() DashboardV2WidgetEventTimelineDefinitionOutputReference
	EventTimelineDefinitionInput() *DashboardV2WidgetEventTimelineDefinition
	// Experimental.
	Fqn() *string
	FreeTextDefinition() DashboardV2WidgetFreeTextDefinitionOutputReference
	FreeTextDefinitionInput() *DashboardV2WidgetFreeTextDefinition
	FunnelDefinition() DashboardV2WidgetFunnelDefinitionOutputReference
	FunnelDefinitionInput() *DashboardV2WidgetFunnelDefinition
	GeomapDefinition() DashboardV2WidgetGeomapDefinitionOutputReference
	GeomapDefinitionInput() *DashboardV2WidgetGeomapDefinition
	GroupDefinition() DashboardV2WidgetGroupDefinitionOutputReference
	GroupDefinitionInput() *DashboardV2WidgetGroupDefinition
	HeatmapDefinition() DashboardV2WidgetHeatmapDefinitionOutputReference
	HeatmapDefinitionInput() *DashboardV2WidgetHeatmapDefinition
	HostmapDefinition() DashboardV2WidgetHostmapDefinitionOutputReference
	HostmapDefinitionInput() *DashboardV2WidgetHostmapDefinition
	Id() *float64
	SetId(val *float64)
	IdInput() *float64
	IframeDefinition() DashboardV2WidgetIframeDefinitionOutputReference
	IframeDefinitionInput() *DashboardV2WidgetIframeDefinition
	ImageDefinition() DashboardV2WidgetImageDefinitionOutputReference
	ImageDefinitionInput() *DashboardV2WidgetImageDefinition
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ListStreamDefinition() DashboardV2WidgetListStreamDefinitionOutputReference
	ListStreamDefinitionInput() *DashboardV2WidgetListStreamDefinition
	LogStreamDefinition() DashboardV2WidgetLogStreamDefinitionOutputReference
	LogStreamDefinitionInput() *DashboardV2WidgetLogStreamDefinition
	ManageStatusDefinition() DashboardV2WidgetManageStatusDefinitionOutputReference
	ManageStatusDefinitionInput() *DashboardV2WidgetManageStatusDefinition
	NoteDefinition() DashboardV2WidgetNoteDefinitionOutputReference
	NoteDefinitionInput() *DashboardV2WidgetNoteDefinition
	PointPlotDefinition() DashboardV2WidgetPointPlotDefinitionOutputReference
	PointPlotDefinitionInput() *DashboardV2WidgetPointPlotDefinition
	PowerpackDefinition() DashboardV2WidgetPowerpackDefinitionOutputReference
	PowerpackDefinitionInput() *DashboardV2WidgetPowerpackDefinition
	QueryTableDefinition() DashboardV2WidgetQueryTableDefinitionOutputReference
	QueryTableDefinitionInput() *DashboardV2WidgetQueryTableDefinition
	QueryValueDefinition() DashboardV2WidgetQueryValueDefinitionOutputReference
	QueryValueDefinitionInput() *DashboardV2WidgetQueryValueDefinition
	RunWorkflowDefinition() DashboardV2WidgetRunWorkflowDefinitionOutputReference
	RunWorkflowDefinitionInput() *DashboardV2WidgetRunWorkflowDefinition
	SankeyDefinition() DashboardV2WidgetSankeyDefinitionOutputReference
	SankeyDefinitionInput() *DashboardV2WidgetSankeyDefinition
	ScatterplotDefinition() DashboardV2WidgetScatterplotDefinitionOutputReference
	ScatterplotDefinitionInput() *DashboardV2WidgetScatterplotDefinition
	ServiceLevelObjectiveDefinition() DashboardV2WidgetServiceLevelObjectiveDefinitionOutputReference
	ServiceLevelObjectiveDefinitionInput() *DashboardV2WidgetServiceLevelObjectiveDefinition
	ServicemapDefinition() DashboardV2WidgetServicemapDefinitionOutputReference
	ServicemapDefinitionInput() *DashboardV2WidgetServicemapDefinition
	SloListDefinition() DashboardV2WidgetSloListDefinitionOutputReference
	SloListDefinitionInput() *DashboardV2WidgetSloListDefinition
	SplitGraphDefinition() DashboardV2WidgetSplitGraphDefinitionOutputReference
	SplitGraphDefinitionInput() *DashboardV2WidgetSplitGraphDefinition
	SunburstDefinition() DashboardV2WidgetSunburstDefinitionOutputReference
	SunburstDefinitionInput() *DashboardV2WidgetSunburstDefinition
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeseriesDefinition() DashboardV2WidgetTimeseriesDefinitionOutputReference
	TimeseriesDefinitionInput() *DashboardV2WidgetTimeseriesDefinition
	ToplistDefinition() DashboardV2WidgetToplistDefinitionOutputReference
	ToplistDefinitionInput() *DashboardV2WidgetToplistDefinition
	TopologyMapDefinition() DashboardV2WidgetTopologyMapDefinitionOutputReference
	TopologyMapDefinitionInput() *DashboardV2WidgetTopologyMapDefinition
	TraceServiceDefinition() DashboardV2WidgetTraceServiceDefinitionOutputReference
	TraceServiceDefinitionInput() *DashboardV2WidgetTraceServiceDefinition
	TreemapDefinition() DashboardV2WidgetTreemapDefinitionOutputReference
	TreemapDefinitionInput() *DashboardV2WidgetTreemapDefinition
	WidgetLayout() DashboardV2WidgetWidgetLayoutOutputReference
	WidgetLayoutInput() *DashboardV2WidgetWidgetLayout
	WildcardDefinition() DashboardV2WidgetWildcardDefinitionOutputReference
	WildcardDefinitionInput() *DashboardV2WidgetWildcardDefinition
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
	PutAlertGraphDefinition(value *DashboardV2WidgetAlertGraphDefinition)
	PutAlertValueDefinition(value *DashboardV2WidgetAlertValueDefinition)
	PutBarChartDefinition(value *DashboardV2WidgetBarChartDefinition)
	PutChangeDefinition(value *DashboardV2WidgetChangeDefinition)
	PutCheckStatusDefinition(value *DashboardV2WidgetCheckStatusDefinition)
	PutDistributionDefinition(value *DashboardV2WidgetDistributionDefinition)
	PutEventStreamDefinition(value *DashboardV2WidgetEventStreamDefinition)
	PutEventTimelineDefinition(value *DashboardV2WidgetEventTimelineDefinition)
	PutFreeTextDefinition(value *DashboardV2WidgetFreeTextDefinition)
	PutFunnelDefinition(value *DashboardV2WidgetFunnelDefinition)
	PutGeomapDefinition(value *DashboardV2WidgetGeomapDefinition)
	PutGroupDefinition(value *DashboardV2WidgetGroupDefinition)
	PutHeatmapDefinition(value *DashboardV2WidgetHeatmapDefinition)
	PutHostmapDefinition(value *DashboardV2WidgetHostmapDefinition)
	PutIframeDefinition(value *DashboardV2WidgetIframeDefinition)
	PutImageDefinition(value *DashboardV2WidgetImageDefinition)
	PutListStreamDefinition(value *DashboardV2WidgetListStreamDefinition)
	PutLogStreamDefinition(value *DashboardV2WidgetLogStreamDefinition)
	PutManageStatusDefinition(value *DashboardV2WidgetManageStatusDefinition)
	PutNoteDefinition(value *DashboardV2WidgetNoteDefinition)
	PutPointPlotDefinition(value *DashboardV2WidgetPointPlotDefinition)
	PutPowerpackDefinition(value *DashboardV2WidgetPowerpackDefinition)
	PutQueryTableDefinition(value *DashboardV2WidgetQueryTableDefinition)
	PutQueryValueDefinition(value *DashboardV2WidgetQueryValueDefinition)
	PutRunWorkflowDefinition(value *DashboardV2WidgetRunWorkflowDefinition)
	PutSankeyDefinition(value *DashboardV2WidgetSankeyDefinition)
	PutScatterplotDefinition(value *DashboardV2WidgetScatterplotDefinition)
	PutServiceLevelObjectiveDefinition(value *DashboardV2WidgetServiceLevelObjectiveDefinition)
	PutServicemapDefinition(value *DashboardV2WidgetServicemapDefinition)
	PutSloListDefinition(value *DashboardV2WidgetSloListDefinition)
	PutSplitGraphDefinition(value *DashboardV2WidgetSplitGraphDefinition)
	PutSunburstDefinition(value *DashboardV2WidgetSunburstDefinition)
	PutTimeseriesDefinition(value *DashboardV2WidgetTimeseriesDefinition)
	PutToplistDefinition(value *DashboardV2WidgetToplistDefinition)
	PutTopologyMapDefinition(value *DashboardV2WidgetTopologyMapDefinition)
	PutTraceServiceDefinition(value *DashboardV2WidgetTraceServiceDefinition)
	PutTreemapDefinition(value *DashboardV2WidgetTreemapDefinition)
	PutWidgetLayout(value *DashboardV2WidgetWidgetLayout)
	PutWildcardDefinition(value *DashboardV2WidgetWildcardDefinition)
	ResetAlertGraphDefinition()
	ResetAlertValueDefinition()
	ResetBarChartDefinition()
	ResetChangeDefinition()
	ResetCheckStatusDefinition()
	ResetDistributionDefinition()
	ResetEventStreamDefinition()
	ResetEventTimelineDefinition()
	ResetFreeTextDefinition()
	ResetFunnelDefinition()
	ResetGeomapDefinition()
	ResetGroupDefinition()
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
	ResetPowerpackDefinition()
	ResetQueryTableDefinition()
	ResetQueryValueDefinition()
	ResetRunWorkflowDefinition()
	ResetSankeyDefinition()
	ResetScatterplotDefinition()
	ResetServiceLevelObjectiveDefinition()
	ResetServicemapDefinition()
	ResetSloListDefinition()
	ResetSplitGraphDefinition()
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

// The jsii proxy struct for DashboardV2WidgetOutputReference
type jsiiProxy_DashboardV2WidgetOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) AlertGraphDefinition() DashboardV2WidgetAlertGraphDefinitionOutputReference {
	var returns DashboardV2WidgetAlertGraphDefinitionOutputReference
	_jsii_.Get(
		j,
		"alertGraphDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) AlertGraphDefinitionInput() *DashboardV2WidgetAlertGraphDefinition {
	var returns *DashboardV2WidgetAlertGraphDefinition
	_jsii_.Get(
		j,
		"alertGraphDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) AlertValueDefinition() DashboardV2WidgetAlertValueDefinitionOutputReference {
	var returns DashboardV2WidgetAlertValueDefinitionOutputReference
	_jsii_.Get(
		j,
		"alertValueDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) AlertValueDefinitionInput() *DashboardV2WidgetAlertValueDefinition {
	var returns *DashboardV2WidgetAlertValueDefinition
	_jsii_.Get(
		j,
		"alertValueDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) BarChartDefinition() DashboardV2WidgetBarChartDefinitionOutputReference {
	var returns DashboardV2WidgetBarChartDefinitionOutputReference
	_jsii_.Get(
		j,
		"barChartDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) BarChartDefinitionInput() *DashboardV2WidgetBarChartDefinition {
	var returns *DashboardV2WidgetBarChartDefinition
	_jsii_.Get(
		j,
		"barChartDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) ChangeDefinition() DashboardV2WidgetChangeDefinitionOutputReference {
	var returns DashboardV2WidgetChangeDefinitionOutputReference
	_jsii_.Get(
		j,
		"changeDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) ChangeDefinitionInput() *DashboardV2WidgetChangeDefinition {
	var returns *DashboardV2WidgetChangeDefinition
	_jsii_.Get(
		j,
		"changeDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) CheckStatusDefinition() DashboardV2WidgetCheckStatusDefinitionOutputReference {
	var returns DashboardV2WidgetCheckStatusDefinitionOutputReference
	_jsii_.Get(
		j,
		"checkStatusDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) CheckStatusDefinitionInput() *DashboardV2WidgetCheckStatusDefinition {
	var returns *DashboardV2WidgetCheckStatusDefinition
	_jsii_.Get(
		j,
		"checkStatusDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) DistributionDefinition() DashboardV2WidgetDistributionDefinitionOutputReference {
	var returns DashboardV2WidgetDistributionDefinitionOutputReference
	_jsii_.Get(
		j,
		"distributionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) DistributionDefinitionInput() *DashboardV2WidgetDistributionDefinition {
	var returns *DashboardV2WidgetDistributionDefinition
	_jsii_.Get(
		j,
		"distributionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) EventStreamDefinition() DashboardV2WidgetEventStreamDefinitionOutputReference {
	var returns DashboardV2WidgetEventStreamDefinitionOutputReference
	_jsii_.Get(
		j,
		"eventStreamDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) EventStreamDefinitionInput() *DashboardV2WidgetEventStreamDefinition {
	var returns *DashboardV2WidgetEventStreamDefinition
	_jsii_.Get(
		j,
		"eventStreamDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) EventTimelineDefinition() DashboardV2WidgetEventTimelineDefinitionOutputReference {
	var returns DashboardV2WidgetEventTimelineDefinitionOutputReference
	_jsii_.Get(
		j,
		"eventTimelineDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) EventTimelineDefinitionInput() *DashboardV2WidgetEventTimelineDefinition {
	var returns *DashboardV2WidgetEventTimelineDefinition
	_jsii_.Get(
		j,
		"eventTimelineDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) FreeTextDefinition() DashboardV2WidgetFreeTextDefinitionOutputReference {
	var returns DashboardV2WidgetFreeTextDefinitionOutputReference
	_jsii_.Get(
		j,
		"freeTextDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) FreeTextDefinitionInput() *DashboardV2WidgetFreeTextDefinition {
	var returns *DashboardV2WidgetFreeTextDefinition
	_jsii_.Get(
		j,
		"freeTextDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) FunnelDefinition() DashboardV2WidgetFunnelDefinitionOutputReference {
	var returns DashboardV2WidgetFunnelDefinitionOutputReference
	_jsii_.Get(
		j,
		"funnelDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) FunnelDefinitionInput() *DashboardV2WidgetFunnelDefinition {
	var returns *DashboardV2WidgetFunnelDefinition
	_jsii_.Get(
		j,
		"funnelDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) GeomapDefinition() DashboardV2WidgetGeomapDefinitionOutputReference {
	var returns DashboardV2WidgetGeomapDefinitionOutputReference
	_jsii_.Get(
		j,
		"geomapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) GeomapDefinitionInput() *DashboardV2WidgetGeomapDefinition {
	var returns *DashboardV2WidgetGeomapDefinition
	_jsii_.Get(
		j,
		"geomapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) GroupDefinition() DashboardV2WidgetGroupDefinitionOutputReference {
	var returns DashboardV2WidgetGroupDefinitionOutputReference
	_jsii_.Get(
		j,
		"groupDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) GroupDefinitionInput() *DashboardV2WidgetGroupDefinition {
	var returns *DashboardV2WidgetGroupDefinition
	_jsii_.Get(
		j,
		"groupDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) HeatmapDefinition() DashboardV2WidgetHeatmapDefinitionOutputReference {
	var returns DashboardV2WidgetHeatmapDefinitionOutputReference
	_jsii_.Get(
		j,
		"heatmapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) HeatmapDefinitionInput() *DashboardV2WidgetHeatmapDefinition {
	var returns *DashboardV2WidgetHeatmapDefinition
	_jsii_.Get(
		j,
		"heatmapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) HostmapDefinition() DashboardV2WidgetHostmapDefinitionOutputReference {
	var returns DashboardV2WidgetHostmapDefinitionOutputReference
	_jsii_.Get(
		j,
		"hostmapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) HostmapDefinitionInput() *DashboardV2WidgetHostmapDefinition {
	var returns *DashboardV2WidgetHostmapDefinition
	_jsii_.Get(
		j,
		"hostmapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) IdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) IframeDefinition() DashboardV2WidgetIframeDefinitionOutputReference {
	var returns DashboardV2WidgetIframeDefinitionOutputReference
	_jsii_.Get(
		j,
		"iframeDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) IframeDefinitionInput() *DashboardV2WidgetIframeDefinition {
	var returns *DashboardV2WidgetIframeDefinition
	_jsii_.Get(
		j,
		"iframeDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) ImageDefinition() DashboardV2WidgetImageDefinitionOutputReference {
	var returns DashboardV2WidgetImageDefinitionOutputReference
	_jsii_.Get(
		j,
		"imageDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) ImageDefinitionInput() *DashboardV2WidgetImageDefinition {
	var returns *DashboardV2WidgetImageDefinition
	_jsii_.Get(
		j,
		"imageDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) ListStreamDefinition() DashboardV2WidgetListStreamDefinitionOutputReference {
	var returns DashboardV2WidgetListStreamDefinitionOutputReference
	_jsii_.Get(
		j,
		"listStreamDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) ListStreamDefinitionInput() *DashboardV2WidgetListStreamDefinition {
	var returns *DashboardV2WidgetListStreamDefinition
	_jsii_.Get(
		j,
		"listStreamDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) LogStreamDefinition() DashboardV2WidgetLogStreamDefinitionOutputReference {
	var returns DashboardV2WidgetLogStreamDefinitionOutputReference
	_jsii_.Get(
		j,
		"logStreamDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) LogStreamDefinitionInput() *DashboardV2WidgetLogStreamDefinition {
	var returns *DashboardV2WidgetLogStreamDefinition
	_jsii_.Get(
		j,
		"logStreamDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) ManageStatusDefinition() DashboardV2WidgetManageStatusDefinitionOutputReference {
	var returns DashboardV2WidgetManageStatusDefinitionOutputReference
	_jsii_.Get(
		j,
		"manageStatusDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) ManageStatusDefinitionInput() *DashboardV2WidgetManageStatusDefinition {
	var returns *DashboardV2WidgetManageStatusDefinition
	_jsii_.Get(
		j,
		"manageStatusDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) NoteDefinition() DashboardV2WidgetNoteDefinitionOutputReference {
	var returns DashboardV2WidgetNoteDefinitionOutputReference
	_jsii_.Get(
		j,
		"noteDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) NoteDefinitionInput() *DashboardV2WidgetNoteDefinition {
	var returns *DashboardV2WidgetNoteDefinition
	_jsii_.Get(
		j,
		"noteDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) PointPlotDefinition() DashboardV2WidgetPointPlotDefinitionOutputReference {
	var returns DashboardV2WidgetPointPlotDefinitionOutputReference
	_jsii_.Get(
		j,
		"pointPlotDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) PointPlotDefinitionInput() *DashboardV2WidgetPointPlotDefinition {
	var returns *DashboardV2WidgetPointPlotDefinition
	_jsii_.Get(
		j,
		"pointPlotDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) PowerpackDefinition() DashboardV2WidgetPowerpackDefinitionOutputReference {
	var returns DashboardV2WidgetPowerpackDefinitionOutputReference
	_jsii_.Get(
		j,
		"powerpackDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) PowerpackDefinitionInput() *DashboardV2WidgetPowerpackDefinition {
	var returns *DashboardV2WidgetPowerpackDefinition
	_jsii_.Get(
		j,
		"powerpackDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) QueryTableDefinition() DashboardV2WidgetQueryTableDefinitionOutputReference {
	var returns DashboardV2WidgetQueryTableDefinitionOutputReference
	_jsii_.Get(
		j,
		"queryTableDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) QueryTableDefinitionInput() *DashboardV2WidgetQueryTableDefinition {
	var returns *DashboardV2WidgetQueryTableDefinition
	_jsii_.Get(
		j,
		"queryTableDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) QueryValueDefinition() DashboardV2WidgetQueryValueDefinitionOutputReference {
	var returns DashboardV2WidgetQueryValueDefinitionOutputReference
	_jsii_.Get(
		j,
		"queryValueDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) QueryValueDefinitionInput() *DashboardV2WidgetQueryValueDefinition {
	var returns *DashboardV2WidgetQueryValueDefinition
	_jsii_.Get(
		j,
		"queryValueDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) RunWorkflowDefinition() DashboardV2WidgetRunWorkflowDefinitionOutputReference {
	var returns DashboardV2WidgetRunWorkflowDefinitionOutputReference
	_jsii_.Get(
		j,
		"runWorkflowDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) RunWorkflowDefinitionInput() *DashboardV2WidgetRunWorkflowDefinition {
	var returns *DashboardV2WidgetRunWorkflowDefinition
	_jsii_.Get(
		j,
		"runWorkflowDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) SankeyDefinition() DashboardV2WidgetSankeyDefinitionOutputReference {
	var returns DashboardV2WidgetSankeyDefinitionOutputReference
	_jsii_.Get(
		j,
		"sankeyDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) SankeyDefinitionInput() *DashboardV2WidgetSankeyDefinition {
	var returns *DashboardV2WidgetSankeyDefinition
	_jsii_.Get(
		j,
		"sankeyDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) ScatterplotDefinition() DashboardV2WidgetScatterplotDefinitionOutputReference {
	var returns DashboardV2WidgetScatterplotDefinitionOutputReference
	_jsii_.Get(
		j,
		"scatterplotDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) ScatterplotDefinitionInput() *DashboardV2WidgetScatterplotDefinition {
	var returns *DashboardV2WidgetScatterplotDefinition
	_jsii_.Get(
		j,
		"scatterplotDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) ServiceLevelObjectiveDefinition() DashboardV2WidgetServiceLevelObjectiveDefinitionOutputReference {
	var returns DashboardV2WidgetServiceLevelObjectiveDefinitionOutputReference
	_jsii_.Get(
		j,
		"serviceLevelObjectiveDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) ServiceLevelObjectiveDefinitionInput() *DashboardV2WidgetServiceLevelObjectiveDefinition {
	var returns *DashboardV2WidgetServiceLevelObjectiveDefinition
	_jsii_.Get(
		j,
		"serviceLevelObjectiveDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) ServicemapDefinition() DashboardV2WidgetServicemapDefinitionOutputReference {
	var returns DashboardV2WidgetServicemapDefinitionOutputReference
	_jsii_.Get(
		j,
		"servicemapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) ServicemapDefinitionInput() *DashboardV2WidgetServicemapDefinition {
	var returns *DashboardV2WidgetServicemapDefinition
	_jsii_.Get(
		j,
		"servicemapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) SloListDefinition() DashboardV2WidgetSloListDefinitionOutputReference {
	var returns DashboardV2WidgetSloListDefinitionOutputReference
	_jsii_.Get(
		j,
		"sloListDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) SloListDefinitionInput() *DashboardV2WidgetSloListDefinition {
	var returns *DashboardV2WidgetSloListDefinition
	_jsii_.Get(
		j,
		"sloListDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) SplitGraphDefinition() DashboardV2WidgetSplitGraphDefinitionOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionOutputReference
	_jsii_.Get(
		j,
		"splitGraphDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) SplitGraphDefinitionInput() *DashboardV2WidgetSplitGraphDefinition {
	var returns *DashboardV2WidgetSplitGraphDefinition
	_jsii_.Get(
		j,
		"splitGraphDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) SunburstDefinition() DashboardV2WidgetSunburstDefinitionOutputReference {
	var returns DashboardV2WidgetSunburstDefinitionOutputReference
	_jsii_.Get(
		j,
		"sunburstDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) SunburstDefinitionInput() *DashboardV2WidgetSunburstDefinition {
	var returns *DashboardV2WidgetSunburstDefinition
	_jsii_.Get(
		j,
		"sunburstDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) TimeseriesDefinition() DashboardV2WidgetTimeseriesDefinitionOutputReference {
	var returns DashboardV2WidgetTimeseriesDefinitionOutputReference
	_jsii_.Get(
		j,
		"timeseriesDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) TimeseriesDefinitionInput() *DashboardV2WidgetTimeseriesDefinition {
	var returns *DashboardV2WidgetTimeseriesDefinition
	_jsii_.Get(
		j,
		"timeseriesDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) ToplistDefinition() DashboardV2WidgetToplistDefinitionOutputReference {
	var returns DashboardV2WidgetToplistDefinitionOutputReference
	_jsii_.Get(
		j,
		"toplistDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) ToplistDefinitionInput() *DashboardV2WidgetToplistDefinition {
	var returns *DashboardV2WidgetToplistDefinition
	_jsii_.Get(
		j,
		"toplistDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) TopologyMapDefinition() DashboardV2WidgetTopologyMapDefinitionOutputReference {
	var returns DashboardV2WidgetTopologyMapDefinitionOutputReference
	_jsii_.Get(
		j,
		"topologyMapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) TopologyMapDefinitionInput() *DashboardV2WidgetTopologyMapDefinition {
	var returns *DashboardV2WidgetTopologyMapDefinition
	_jsii_.Get(
		j,
		"topologyMapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) TraceServiceDefinition() DashboardV2WidgetTraceServiceDefinitionOutputReference {
	var returns DashboardV2WidgetTraceServiceDefinitionOutputReference
	_jsii_.Get(
		j,
		"traceServiceDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) TraceServiceDefinitionInput() *DashboardV2WidgetTraceServiceDefinition {
	var returns *DashboardV2WidgetTraceServiceDefinition
	_jsii_.Get(
		j,
		"traceServiceDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) TreemapDefinition() DashboardV2WidgetTreemapDefinitionOutputReference {
	var returns DashboardV2WidgetTreemapDefinitionOutputReference
	_jsii_.Get(
		j,
		"treemapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) TreemapDefinitionInput() *DashboardV2WidgetTreemapDefinition {
	var returns *DashboardV2WidgetTreemapDefinition
	_jsii_.Get(
		j,
		"treemapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) WidgetLayout() DashboardV2WidgetWidgetLayoutOutputReference {
	var returns DashboardV2WidgetWidgetLayoutOutputReference
	_jsii_.Get(
		j,
		"widgetLayout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) WidgetLayoutInput() *DashboardV2WidgetWidgetLayout {
	var returns *DashboardV2WidgetWidgetLayout
	_jsii_.Get(
		j,
		"widgetLayoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) WildcardDefinition() DashboardV2WidgetWildcardDefinitionOutputReference {
	var returns DashboardV2WidgetWildcardDefinitionOutputReference
	_jsii_.Get(
		j,
		"wildcardDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference) WildcardDefinitionInput() *DashboardV2WidgetWildcardDefinition {
	var returns *DashboardV2WidgetWildcardDefinition
	_jsii_.Get(
		j,
		"wildcardDefinitionInput",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardV2WidgetOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetOutputReference_Override(d DashboardV2WidgetOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference)SetId(val *float64) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutAlertGraphDefinition(value *DashboardV2WidgetAlertGraphDefinition) {
	if err := d.validatePutAlertGraphDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAlertGraphDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutAlertValueDefinition(value *DashboardV2WidgetAlertValueDefinition) {
	if err := d.validatePutAlertValueDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAlertValueDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutBarChartDefinition(value *DashboardV2WidgetBarChartDefinition) {
	if err := d.validatePutBarChartDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBarChartDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutChangeDefinition(value *DashboardV2WidgetChangeDefinition) {
	if err := d.validatePutChangeDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putChangeDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutCheckStatusDefinition(value *DashboardV2WidgetCheckStatusDefinition) {
	if err := d.validatePutCheckStatusDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCheckStatusDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutDistributionDefinition(value *DashboardV2WidgetDistributionDefinition) {
	if err := d.validatePutDistributionDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDistributionDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutEventStreamDefinition(value *DashboardV2WidgetEventStreamDefinition) {
	if err := d.validatePutEventStreamDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventStreamDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutEventTimelineDefinition(value *DashboardV2WidgetEventTimelineDefinition) {
	if err := d.validatePutEventTimelineDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventTimelineDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutFreeTextDefinition(value *DashboardV2WidgetFreeTextDefinition) {
	if err := d.validatePutFreeTextDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFreeTextDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutFunnelDefinition(value *DashboardV2WidgetFunnelDefinition) {
	if err := d.validatePutFunnelDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFunnelDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutGeomapDefinition(value *DashboardV2WidgetGeomapDefinition) {
	if err := d.validatePutGeomapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGeomapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutGroupDefinition(value *DashboardV2WidgetGroupDefinition) {
	if err := d.validatePutGroupDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGroupDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutHeatmapDefinition(value *DashboardV2WidgetHeatmapDefinition) {
	if err := d.validatePutHeatmapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHeatmapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutHostmapDefinition(value *DashboardV2WidgetHostmapDefinition) {
	if err := d.validatePutHostmapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHostmapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutIframeDefinition(value *DashboardV2WidgetIframeDefinition) {
	if err := d.validatePutIframeDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIframeDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutImageDefinition(value *DashboardV2WidgetImageDefinition) {
	if err := d.validatePutImageDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putImageDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutListStreamDefinition(value *DashboardV2WidgetListStreamDefinition) {
	if err := d.validatePutListStreamDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putListStreamDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutLogStreamDefinition(value *DashboardV2WidgetLogStreamDefinition) {
	if err := d.validatePutLogStreamDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLogStreamDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutManageStatusDefinition(value *DashboardV2WidgetManageStatusDefinition) {
	if err := d.validatePutManageStatusDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putManageStatusDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutNoteDefinition(value *DashboardV2WidgetNoteDefinition) {
	if err := d.validatePutNoteDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNoteDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutPointPlotDefinition(value *DashboardV2WidgetPointPlotDefinition) {
	if err := d.validatePutPointPlotDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPointPlotDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutPowerpackDefinition(value *DashboardV2WidgetPowerpackDefinition) {
	if err := d.validatePutPowerpackDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPowerpackDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutQueryTableDefinition(value *DashboardV2WidgetQueryTableDefinition) {
	if err := d.validatePutQueryTableDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQueryTableDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutQueryValueDefinition(value *DashboardV2WidgetQueryValueDefinition) {
	if err := d.validatePutQueryValueDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQueryValueDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutRunWorkflowDefinition(value *DashboardV2WidgetRunWorkflowDefinition) {
	if err := d.validatePutRunWorkflowDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRunWorkflowDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutSankeyDefinition(value *DashboardV2WidgetSankeyDefinition) {
	if err := d.validatePutSankeyDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSankeyDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutScatterplotDefinition(value *DashboardV2WidgetScatterplotDefinition) {
	if err := d.validatePutScatterplotDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putScatterplotDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutServiceLevelObjectiveDefinition(value *DashboardV2WidgetServiceLevelObjectiveDefinition) {
	if err := d.validatePutServiceLevelObjectiveDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServiceLevelObjectiveDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutServicemapDefinition(value *DashboardV2WidgetServicemapDefinition) {
	if err := d.validatePutServicemapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServicemapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutSloListDefinition(value *DashboardV2WidgetSloListDefinition) {
	if err := d.validatePutSloListDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSloListDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutSplitGraphDefinition(value *DashboardV2WidgetSplitGraphDefinition) {
	if err := d.validatePutSplitGraphDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSplitGraphDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutSunburstDefinition(value *DashboardV2WidgetSunburstDefinition) {
	if err := d.validatePutSunburstDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSunburstDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutTimeseriesDefinition(value *DashboardV2WidgetTimeseriesDefinition) {
	if err := d.validatePutTimeseriesDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeseriesDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutToplistDefinition(value *DashboardV2WidgetToplistDefinition) {
	if err := d.validatePutToplistDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putToplistDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutTopologyMapDefinition(value *DashboardV2WidgetTopologyMapDefinition) {
	if err := d.validatePutTopologyMapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTopologyMapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutTraceServiceDefinition(value *DashboardV2WidgetTraceServiceDefinition) {
	if err := d.validatePutTraceServiceDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTraceServiceDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutTreemapDefinition(value *DashboardV2WidgetTreemapDefinition) {
	if err := d.validatePutTreemapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTreemapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutWidgetLayout(value *DashboardV2WidgetWidgetLayout) {
	if err := d.validatePutWidgetLayoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWidgetLayout",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) PutWildcardDefinition(value *DashboardV2WidgetWildcardDefinition) {
	if err := d.validatePutWildcardDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWildcardDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetAlertGraphDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetAlertGraphDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetAlertValueDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetAlertValueDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetBarChartDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetBarChartDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetChangeDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetChangeDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetCheckStatusDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetCheckStatusDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetDistributionDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetDistributionDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetEventStreamDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetEventStreamDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetEventTimelineDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetEventTimelineDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetFreeTextDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetFreeTextDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetFunnelDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetFunnelDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetGeomapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetGeomapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetGroupDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetGroupDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetHeatmapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetHeatmapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetHostmapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetHostmapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetIframeDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetIframeDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetImageDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetImageDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetListStreamDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetListStreamDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetLogStreamDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetLogStreamDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetManageStatusDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetManageStatusDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetNoteDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetNoteDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetPointPlotDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetPointPlotDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetPowerpackDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetPowerpackDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetQueryTableDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryTableDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetQueryValueDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryValueDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetRunWorkflowDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetRunWorkflowDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetSankeyDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetSankeyDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetScatterplotDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetScatterplotDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetServiceLevelObjectiveDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceLevelObjectiveDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetServicemapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetServicemapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetSloListDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetSloListDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetSplitGraphDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetSplitGraphDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetSunburstDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetSunburstDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetTimeseriesDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeseriesDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetToplistDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetToplistDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetTopologyMapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetTopologyMapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetTraceServiceDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetTraceServiceDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetTreemapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetTreemapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetWidgetLayout() {
	_jsii_.InvokeVoid(
		d,
		"resetWidgetLayout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ResetWildcardDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetWildcardDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

