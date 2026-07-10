// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetOutputReference interface {
	cdktn.ComplexObject
	AlertGraphDefinition() PowerpackV2WidgetAlertGraphDefinitionOutputReference
	AlertGraphDefinitionInput() *PowerpackV2WidgetAlertGraphDefinition
	AlertValueDefinition() PowerpackV2WidgetAlertValueDefinitionOutputReference
	AlertValueDefinitionInput() *PowerpackV2WidgetAlertValueDefinition
	BarChartDefinition() PowerpackV2WidgetBarChartDefinitionOutputReference
	BarChartDefinitionInput() *PowerpackV2WidgetBarChartDefinition
	ChangeDefinition() PowerpackV2WidgetChangeDefinitionOutputReference
	ChangeDefinitionInput() *PowerpackV2WidgetChangeDefinition
	CheckStatusDefinition() PowerpackV2WidgetCheckStatusDefinitionOutputReference
	CheckStatusDefinitionInput() *PowerpackV2WidgetCheckStatusDefinition
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
	DistributionDefinition() PowerpackV2WidgetDistributionDefinitionOutputReference
	DistributionDefinitionInput() *PowerpackV2WidgetDistributionDefinition
	EventStreamDefinition() PowerpackV2WidgetEventStreamDefinitionOutputReference
	EventStreamDefinitionInput() *PowerpackV2WidgetEventStreamDefinition
	EventTimelineDefinition() PowerpackV2WidgetEventTimelineDefinitionOutputReference
	EventTimelineDefinitionInput() *PowerpackV2WidgetEventTimelineDefinition
	// Experimental.
	Fqn() *string
	FreeTextDefinition() PowerpackV2WidgetFreeTextDefinitionOutputReference
	FreeTextDefinitionInput() *PowerpackV2WidgetFreeTextDefinition
	FunnelDefinition() PowerpackV2WidgetFunnelDefinitionOutputReference
	FunnelDefinitionInput() *PowerpackV2WidgetFunnelDefinition
	GeomapDefinition() PowerpackV2WidgetGeomapDefinitionOutputReference
	GeomapDefinitionInput() *PowerpackV2WidgetGeomapDefinition
	GroupDefinition() PowerpackV2WidgetGroupDefinitionOutputReference
	GroupDefinitionInput() *PowerpackV2WidgetGroupDefinition
	HeatmapDefinition() PowerpackV2WidgetHeatmapDefinitionOutputReference
	HeatmapDefinitionInput() *PowerpackV2WidgetHeatmapDefinition
	HostmapDefinition() PowerpackV2WidgetHostmapDefinitionOutputReference
	HostmapDefinitionInput() *PowerpackV2WidgetHostmapDefinition
	Id() *float64
	SetId(val *float64)
	IdInput() *float64
	IframeDefinition() PowerpackV2WidgetIframeDefinitionOutputReference
	IframeDefinitionInput() *PowerpackV2WidgetIframeDefinition
	ImageDefinition() PowerpackV2WidgetImageDefinitionOutputReference
	ImageDefinitionInput() *PowerpackV2WidgetImageDefinition
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ListStreamDefinition() PowerpackV2WidgetListStreamDefinitionOutputReference
	ListStreamDefinitionInput() *PowerpackV2WidgetListStreamDefinition
	LogStreamDefinition() PowerpackV2WidgetLogStreamDefinitionOutputReference
	LogStreamDefinitionInput() *PowerpackV2WidgetLogStreamDefinition
	ManageStatusDefinition() PowerpackV2WidgetManageStatusDefinitionOutputReference
	ManageStatusDefinitionInput() *PowerpackV2WidgetManageStatusDefinition
	NoteDefinition() PowerpackV2WidgetNoteDefinitionOutputReference
	NoteDefinitionInput() *PowerpackV2WidgetNoteDefinition
	PointPlotDefinition() PowerpackV2WidgetPointPlotDefinitionOutputReference
	PointPlotDefinitionInput() *PowerpackV2WidgetPointPlotDefinition
	QueryTableDefinition() PowerpackV2WidgetQueryTableDefinitionOutputReference
	QueryTableDefinitionInput() *PowerpackV2WidgetQueryTableDefinition
	QueryValueDefinition() PowerpackV2WidgetQueryValueDefinitionOutputReference
	QueryValueDefinitionInput() *PowerpackV2WidgetQueryValueDefinition
	RunWorkflowDefinition() PowerpackV2WidgetRunWorkflowDefinitionOutputReference
	RunWorkflowDefinitionInput() *PowerpackV2WidgetRunWorkflowDefinition
	SankeyDefinition() PowerpackV2WidgetSankeyDefinitionOutputReference
	SankeyDefinitionInput() *PowerpackV2WidgetSankeyDefinition
	ScatterplotDefinition() PowerpackV2WidgetScatterplotDefinitionOutputReference
	ScatterplotDefinitionInput() *PowerpackV2WidgetScatterplotDefinition
	ServiceLevelObjectiveDefinition() PowerpackV2WidgetServiceLevelObjectiveDefinitionOutputReference
	ServiceLevelObjectiveDefinitionInput() *PowerpackV2WidgetServiceLevelObjectiveDefinition
	ServicemapDefinition() PowerpackV2WidgetServicemapDefinitionOutputReference
	ServicemapDefinitionInput() *PowerpackV2WidgetServicemapDefinition
	SloListDefinition() PowerpackV2WidgetSloListDefinitionOutputReference
	SloListDefinitionInput() *PowerpackV2WidgetSloListDefinition
	SunburstDefinition() PowerpackV2WidgetSunburstDefinitionOutputReference
	SunburstDefinitionInput() *PowerpackV2WidgetSunburstDefinition
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeseriesDefinition() PowerpackV2WidgetTimeseriesDefinitionOutputReference
	TimeseriesDefinitionInput() *PowerpackV2WidgetTimeseriesDefinition
	ToplistDefinition() PowerpackV2WidgetToplistDefinitionOutputReference
	ToplistDefinitionInput() *PowerpackV2WidgetToplistDefinition
	TopologyMapDefinition() PowerpackV2WidgetTopologyMapDefinitionOutputReference
	TopologyMapDefinitionInput() *PowerpackV2WidgetTopologyMapDefinition
	TraceServiceDefinition() PowerpackV2WidgetTraceServiceDefinitionOutputReference
	TraceServiceDefinitionInput() *PowerpackV2WidgetTraceServiceDefinition
	TreemapDefinition() PowerpackV2WidgetTreemapDefinitionOutputReference
	TreemapDefinitionInput() *PowerpackV2WidgetTreemapDefinition
	WidgetLayout() PowerpackV2WidgetWidgetLayoutOutputReference
	WidgetLayoutInput() *PowerpackV2WidgetWidgetLayout
	WildcardDefinition() PowerpackV2WidgetWildcardDefinitionOutputReference
	WildcardDefinitionInput() *PowerpackV2WidgetWildcardDefinition
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
	PutAlertGraphDefinition(value *PowerpackV2WidgetAlertGraphDefinition)
	PutAlertValueDefinition(value *PowerpackV2WidgetAlertValueDefinition)
	PutBarChartDefinition(value *PowerpackV2WidgetBarChartDefinition)
	PutChangeDefinition(value *PowerpackV2WidgetChangeDefinition)
	PutCheckStatusDefinition(value *PowerpackV2WidgetCheckStatusDefinition)
	PutDistributionDefinition(value *PowerpackV2WidgetDistributionDefinition)
	PutEventStreamDefinition(value *PowerpackV2WidgetEventStreamDefinition)
	PutEventTimelineDefinition(value *PowerpackV2WidgetEventTimelineDefinition)
	PutFreeTextDefinition(value *PowerpackV2WidgetFreeTextDefinition)
	PutFunnelDefinition(value *PowerpackV2WidgetFunnelDefinition)
	PutGeomapDefinition(value *PowerpackV2WidgetGeomapDefinition)
	PutGroupDefinition(value *PowerpackV2WidgetGroupDefinition)
	PutHeatmapDefinition(value *PowerpackV2WidgetHeatmapDefinition)
	PutHostmapDefinition(value *PowerpackV2WidgetHostmapDefinition)
	PutIframeDefinition(value *PowerpackV2WidgetIframeDefinition)
	PutImageDefinition(value *PowerpackV2WidgetImageDefinition)
	PutListStreamDefinition(value *PowerpackV2WidgetListStreamDefinition)
	PutLogStreamDefinition(value *PowerpackV2WidgetLogStreamDefinition)
	PutManageStatusDefinition(value *PowerpackV2WidgetManageStatusDefinition)
	PutNoteDefinition(value *PowerpackV2WidgetNoteDefinition)
	PutPointPlotDefinition(value *PowerpackV2WidgetPointPlotDefinition)
	PutQueryTableDefinition(value *PowerpackV2WidgetQueryTableDefinition)
	PutQueryValueDefinition(value *PowerpackV2WidgetQueryValueDefinition)
	PutRunWorkflowDefinition(value *PowerpackV2WidgetRunWorkflowDefinition)
	PutSankeyDefinition(value *PowerpackV2WidgetSankeyDefinition)
	PutScatterplotDefinition(value *PowerpackV2WidgetScatterplotDefinition)
	PutServiceLevelObjectiveDefinition(value *PowerpackV2WidgetServiceLevelObjectiveDefinition)
	PutServicemapDefinition(value *PowerpackV2WidgetServicemapDefinition)
	PutSloListDefinition(value *PowerpackV2WidgetSloListDefinition)
	PutSunburstDefinition(value *PowerpackV2WidgetSunburstDefinition)
	PutTimeseriesDefinition(value *PowerpackV2WidgetTimeseriesDefinition)
	PutToplistDefinition(value *PowerpackV2WidgetToplistDefinition)
	PutTopologyMapDefinition(value *PowerpackV2WidgetTopologyMapDefinition)
	PutTraceServiceDefinition(value *PowerpackV2WidgetTraceServiceDefinition)
	PutTreemapDefinition(value *PowerpackV2WidgetTreemapDefinition)
	PutWidgetLayout(value *PowerpackV2WidgetWidgetLayout)
	PutWildcardDefinition(value *PowerpackV2WidgetWildcardDefinition)
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
	ResetQueryTableDefinition()
	ResetQueryValueDefinition()
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

// The jsii proxy struct for PowerpackV2WidgetOutputReference
type jsiiProxy_PowerpackV2WidgetOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) AlertGraphDefinition() PowerpackV2WidgetAlertGraphDefinitionOutputReference {
	var returns PowerpackV2WidgetAlertGraphDefinitionOutputReference
	_jsii_.Get(
		j,
		"alertGraphDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) AlertGraphDefinitionInput() *PowerpackV2WidgetAlertGraphDefinition {
	var returns *PowerpackV2WidgetAlertGraphDefinition
	_jsii_.Get(
		j,
		"alertGraphDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) AlertValueDefinition() PowerpackV2WidgetAlertValueDefinitionOutputReference {
	var returns PowerpackV2WidgetAlertValueDefinitionOutputReference
	_jsii_.Get(
		j,
		"alertValueDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) AlertValueDefinitionInput() *PowerpackV2WidgetAlertValueDefinition {
	var returns *PowerpackV2WidgetAlertValueDefinition
	_jsii_.Get(
		j,
		"alertValueDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) BarChartDefinition() PowerpackV2WidgetBarChartDefinitionOutputReference {
	var returns PowerpackV2WidgetBarChartDefinitionOutputReference
	_jsii_.Get(
		j,
		"barChartDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) BarChartDefinitionInput() *PowerpackV2WidgetBarChartDefinition {
	var returns *PowerpackV2WidgetBarChartDefinition
	_jsii_.Get(
		j,
		"barChartDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) ChangeDefinition() PowerpackV2WidgetChangeDefinitionOutputReference {
	var returns PowerpackV2WidgetChangeDefinitionOutputReference
	_jsii_.Get(
		j,
		"changeDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) ChangeDefinitionInput() *PowerpackV2WidgetChangeDefinition {
	var returns *PowerpackV2WidgetChangeDefinition
	_jsii_.Get(
		j,
		"changeDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) CheckStatusDefinition() PowerpackV2WidgetCheckStatusDefinitionOutputReference {
	var returns PowerpackV2WidgetCheckStatusDefinitionOutputReference
	_jsii_.Get(
		j,
		"checkStatusDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) CheckStatusDefinitionInput() *PowerpackV2WidgetCheckStatusDefinition {
	var returns *PowerpackV2WidgetCheckStatusDefinition
	_jsii_.Get(
		j,
		"checkStatusDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) DistributionDefinition() PowerpackV2WidgetDistributionDefinitionOutputReference {
	var returns PowerpackV2WidgetDistributionDefinitionOutputReference
	_jsii_.Get(
		j,
		"distributionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) DistributionDefinitionInput() *PowerpackV2WidgetDistributionDefinition {
	var returns *PowerpackV2WidgetDistributionDefinition
	_jsii_.Get(
		j,
		"distributionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) EventStreamDefinition() PowerpackV2WidgetEventStreamDefinitionOutputReference {
	var returns PowerpackV2WidgetEventStreamDefinitionOutputReference
	_jsii_.Get(
		j,
		"eventStreamDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) EventStreamDefinitionInput() *PowerpackV2WidgetEventStreamDefinition {
	var returns *PowerpackV2WidgetEventStreamDefinition
	_jsii_.Get(
		j,
		"eventStreamDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) EventTimelineDefinition() PowerpackV2WidgetEventTimelineDefinitionOutputReference {
	var returns PowerpackV2WidgetEventTimelineDefinitionOutputReference
	_jsii_.Get(
		j,
		"eventTimelineDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) EventTimelineDefinitionInput() *PowerpackV2WidgetEventTimelineDefinition {
	var returns *PowerpackV2WidgetEventTimelineDefinition
	_jsii_.Get(
		j,
		"eventTimelineDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) FreeTextDefinition() PowerpackV2WidgetFreeTextDefinitionOutputReference {
	var returns PowerpackV2WidgetFreeTextDefinitionOutputReference
	_jsii_.Get(
		j,
		"freeTextDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) FreeTextDefinitionInput() *PowerpackV2WidgetFreeTextDefinition {
	var returns *PowerpackV2WidgetFreeTextDefinition
	_jsii_.Get(
		j,
		"freeTextDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) FunnelDefinition() PowerpackV2WidgetFunnelDefinitionOutputReference {
	var returns PowerpackV2WidgetFunnelDefinitionOutputReference
	_jsii_.Get(
		j,
		"funnelDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) FunnelDefinitionInput() *PowerpackV2WidgetFunnelDefinition {
	var returns *PowerpackV2WidgetFunnelDefinition
	_jsii_.Get(
		j,
		"funnelDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) GeomapDefinition() PowerpackV2WidgetGeomapDefinitionOutputReference {
	var returns PowerpackV2WidgetGeomapDefinitionOutputReference
	_jsii_.Get(
		j,
		"geomapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) GeomapDefinitionInput() *PowerpackV2WidgetGeomapDefinition {
	var returns *PowerpackV2WidgetGeomapDefinition
	_jsii_.Get(
		j,
		"geomapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) GroupDefinition() PowerpackV2WidgetGroupDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionOutputReference
	_jsii_.Get(
		j,
		"groupDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) GroupDefinitionInput() *PowerpackV2WidgetGroupDefinition {
	var returns *PowerpackV2WidgetGroupDefinition
	_jsii_.Get(
		j,
		"groupDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) HeatmapDefinition() PowerpackV2WidgetHeatmapDefinitionOutputReference {
	var returns PowerpackV2WidgetHeatmapDefinitionOutputReference
	_jsii_.Get(
		j,
		"heatmapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) HeatmapDefinitionInput() *PowerpackV2WidgetHeatmapDefinition {
	var returns *PowerpackV2WidgetHeatmapDefinition
	_jsii_.Get(
		j,
		"heatmapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) HostmapDefinition() PowerpackV2WidgetHostmapDefinitionOutputReference {
	var returns PowerpackV2WidgetHostmapDefinitionOutputReference
	_jsii_.Get(
		j,
		"hostmapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) HostmapDefinitionInput() *PowerpackV2WidgetHostmapDefinition {
	var returns *PowerpackV2WidgetHostmapDefinition
	_jsii_.Get(
		j,
		"hostmapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) IdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) IframeDefinition() PowerpackV2WidgetIframeDefinitionOutputReference {
	var returns PowerpackV2WidgetIframeDefinitionOutputReference
	_jsii_.Get(
		j,
		"iframeDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) IframeDefinitionInput() *PowerpackV2WidgetIframeDefinition {
	var returns *PowerpackV2WidgetIframeDefinition
	_jsii_.Get(
		j,
		"iframeDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) ImageDefinition() PowerpackV2WidgetImageDefinitionOutputReference {
	var returns PowerpackV2WidgetImageDefinitionOutputReference
	_jsii_.Get(
		j,
		"imageDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) ImageDefinitionInput() *PowerpackV2WidgetImageDefinition {
	var returns *PowerpackV2WidgetImageDefinition
	_jsii_.Get(
		j,
		"imageDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) ListStreamDefinition() PowerpackV2WidgetListStreamDefinitionOutputReference {
	var returns PowerpackV2WidgetListStreamDefinitionOutputReference
	_jsii_.Get(
		j,
		"listStreamDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) ListStreamDefinitionInput() *PowerpackV2WidgetListStreamDefinition {
	var returns *PowerpackV2WidgetListStreamDefinition
	_jsii_.Get(
		j,
		"listStreamDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) LogStreamDefinition() PowerpackV2WidgetLogStreamDefinitionOutputReference {
	var returns PowerpackV2WidgetLogStreamDefinitionOutputReference
	_jsii_.Get(
		j,
		"logStreamDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) LogStreamDefinitionInput() *PowerpackV2WidgetLogStreamDefinition {
	var returns *PowerpackV2WidgetLogStreamDefinition
	_jsii_.Get(
		j,
		"logStreamDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) ManageStatusDefinition() PowerpackV2WidgetManageStatusDefinitionOutputReference {
	var returns PowerpackV2WidgetManageStatusDefinitionOutputReference
	_jsii_.Get(
		j,
		"manageStatusDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) ManageStatusDefinitionInput() *PowerpackV2WidgetManageStatusDefinition {
	var returns *PowerpackV2WidgetManageStatusDefinition
	_jsii_.Get(
		j,
		"manageStatusDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) NoteDefinition() PowerpackV2WidgetNoteDefinitionOutputReference {
	var returns PowerpackV2WidgetNoteDefinitionOutputReference
	_jsii_.Get(
		j,
		"noteDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) NoteDefinitionInput() *PowerpackV2WidgetNoteDefinition {
	var returns *PowerpackV2WidgetNoteDefinition
	_jsii_.Get(
		j,
		"noteDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) PointPlotDefinition() PowerpackV2WidgetPointPlotDefinitionOutputReference {
	var returns PowerpackV2WidgetPointPlotDefinitionOutputReference
	_jsii_.Get(
		j,
		"pointPlotDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) PointPlotDefinitionInput() *PowerpackV2WidgetPointPlotDefinition {
	var returns *PowerpackV2WidgetPointPlotDefinition
	_jsii_.Get(
		j,
		"pointPlotDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) QueryTableDefinition() PowerpackV2WidgetQueryTableDefinitionOutputReference {
	var returns PowerpackV2WidgetQueryTableDefinitionOutputReference
	_jsii_.Get(
		j,
		"queryTableDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) QueryTableDefinitionInput() *PowerpackV2WidgetQueryTableDefinition {
	var returns *PowerpackV2WidgetQueryTableDefinition
	_jsii_.Get(
		j,
		"queryTableDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) QueryValueDefinition() PowerpackV2WidgetQueryValueDefinitionOutputReference {
	var returns PowerpackV2WidgetQueryValueDefinitionOutputReference
	_jsii_.Get(
		j,
		"queryValueDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) QueryValueDefinitionInput() *PowerpackV2WidgetQueryValueDefinition {
	var returns *PowerpackV2WidgetQueryValueDefinition
	_jsii_.Get(
		j,
		"queryValueDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) RunWorkflowDefinition() PowerpackV2WidgetRunWorkflowDefinitionOutputReference {
	var returns PowerpackV2WidgetRunWorkflowDefinitionOutputReference
	_jsii_.Get(
		j,
		"runWorkflowDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) RunWorkflowDefinitionInput() *PowerpackV2WidgetRunWorkflowDefinition {
	var returns *PowerpackV2WidgetRunWorkflowDefinition
	_jsii_.Get(
		j,
		"runWorkflowDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) SankeyDefinition() PowerpackV2WidgetSankeyDefinitionOutputReference {
	var returns PowerpackV2WidgetSankeyDefinitionOutputReference
	_jsii_.Get(
		j,
		"sankeyDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) SankeyDefinitionInput() *PowerpackV2WidgetSankeyDefinition {
	var returns *PowerpackV2WidgetSankeyDefinition
	_jsii_.Get(
		j,
		"sankeyDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) ScatterplotDefinition() PowerpackV2WidgetScatterplotDefinitionOutputReference {
	var returns PowerpackV2WidgetScatterplotDefinitionOutputReference
	_jsii_.Get(
		j,
		"scatterplotDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) ScatterplotDefinitionInput() *PowerpackV2WidgetScatterplotDefinition {
	var returns *PowerpackV2WidgetScatterplotDefinition
	_jsii_.Get(
		j,
		"scatterplotDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) ServiceLevelObjectiveDefinition() PowerpackV2WidgetServiceLevelObjectiveDefinitionOutputReference {
	var returns PowerpackV2WidgetServiceLevelObjectiveDefinitionOutputReference
	_jsii_.Get(
		j,
		"serviceLevelObjectiveDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) ServiceLevelObjectiveDefinitionInput() *PowerpackV2WidgetServiceLevelObjectiveDefinition {
	var returns *PowerpackV2WidgetServiceLevelObjectiveDefinition
	_jsii_.Get(
		j,
		"serviceLevelObjectiveDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) ServicemapDefinition() PowerpackV2WidgetServicemapDefinitionOutputReference {
	var returns PowerpackV2WidgetServicemapDefinitionOutputReference
	_jsii_.Get(
		j,
		"servicemapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) ServicemapDefinitionInput() *PowerpackV2WidgetServicemapDefinition {
	var returns *PowerpackV2WidgetServicemapDefinition
	_jsii_.Get(
		j,
		"servicemapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) SloListDefinition() PowerpackV2WidgetSloListDefinitionOutputReference {
	var returns PowerpackV2WidgetSloListDefinitionOutputReference
	_jsii_.Get(
		j,
		"sloListDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) SloListDefinitionInput() *PowerpackV2WidgetSloListDefinition {
	var returns *PowerpackV2WidgetSloListDefinition
	_jsii_.Get(
		j,
		"sloListDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) SunburstDefinition() PowerpackV2WidgetSunburstDefinitionOutputReference {
	var returns PowerpackV2WidgetSunburstDefinitionOutputReference
	_jsii_.Get(
		j,
		"sunburstDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) SunburstDefinitionInput() *PowerpackV2WidgetSunburstDefinition {
	var returns *PowerpackV2WidgetSunburstDefinition
	_jsii_.Get(
		j,
		"sunburstDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) TimeseriesDefinition() PowerpackV2WidgetTimeseriesDefinitionOutputReference {
	var returns PowerpackV2WidgetTimeseriesDefinitionOutputReference
	_jsii_.Get(
		j,
		"timeseriesDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) TimeseriesDefinitionInput() *PowerpackV2WidgetTimeseriesDefinition {
	var returns *PowerpackV2WidgetTimeseriesDefinition
	_jsii_.Get(
		j,
		"timeseriesDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) ToplistDefinition() PowerpackV2WidgetToplistDefinitionOutputReference {
	var returns PowerpackV2WidgetToplistDefinitionOutputReference
	_jsii_.Get(
		j,
		"toplistDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) ToplistDefinitionInput() *PowerpackV2WidgetToplistDefinition {
	var returns *PowerpackV2WidgetToplistDefinition
	_jsii_.Get(
		j,
		"toplistDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) TopologyMapDefinition() PowerpackV2WidgetTopologyMapDefinitionOutputReference {
	var returns PowerpackV2WidgetTopologyMapDefinitionOutputReference
	_jsii_.Get(
		j,
		"topologyMapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) TopologyMapDefinitionInput() *PowerpackV2WidgetTopologyMapDefinition {
	var returns *PowerpackV2WidgetTopologyMapDefinition
	_jsii_.Get(
		j,
		"topologyMapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) TraceServiceDefinition() PowerpackV2WidgetTraceServiceDefinitionOutputReference {
	var returns PowerpackV2WidgetTraceServiceDefinitionOutputReference
	_jsii_.Get(
		j,
		"traceServiceDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) TraceServiceDefinitionInput() *PowerpackV2WidgetTraceServiceDefinition {
	var returns *PowerpackV2WidgetTraceServiceDefinition
	_jsii_.Get(
		j,
		"traceServiceDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) TreemapDefinition() PowerpackV2WidgetTreemapDefinitionOutputReference {
	var returns PowerpackV2WidgetTreemapDefinitionOutputReference
	_jsii_.Get(
		j,
		"treemapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) TreemapDefinitionInput() *PowerpackV2WidgetTreemapDefinition {
	var returns *PowerpackV2WidgetTreemapDefinition
	_jsii_.Get(
		j,
		"treemapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) WidgetLayout() PowerpackV2WidgetWidgetLayoutOutputReference {
	var returns PowerpackV2WidgetWidgetLayoutOutputReference
	_jsii_.Get(
		j,
		"widgetLayout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) WidgetLayoutInput() *PowerpackV2WidgetWidgetLayout {
	var returns *PowerpackV2WidgetWidgetLayout
	_jsii_.Get(
		j,
		"widgetLayoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) WildcardDefinition() PowerpackV2WidgetWildcardDefinitionOutputReference {
	var returns PowerpackV2WidgetWildcardDefinitionOutputReference
	_jsii_.Get(
		j,
		"wildcardDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference) WildcardDefinitionInput() *PowerpackV2WidgetWildcardDefinition {
	var returns *PowerpackV2WidgetWildcardDefinition
	_jsii_.Get(
		j,
		"wildcardDefinitionInput",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetOutputReference_Override(p PowerpackV2WidgetOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference)SetId(val *float64) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutAlertGraphDefinition(value *PowerpackV2WidgetAlertGraphDefinition) {
	if err := p.validatePutAlertGraphDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAlertGraphDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutAlertValueDefinition(value *PowerpackV2WidgetAlertValueDefinition) {
	if err := p.validatePutAlertValueDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAlertValueDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutBarChartDefinition(value *PowerpackV2WidgetBarChartDefinition) {
	if err := p.validatePutBarChartDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putBarChartDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutChangeDefinition(value *PowerpackV2WidgetChangeDefinition) {
	if err := p.validatePutChangeDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putChangeDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutCheckStatusDefinition(value *PowerpackV2WidgetCheckStatusDefinition) {
	if err := p.validatePutCheckStatusDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCheckStatusDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutDistributionDefinition(value *PowerpackV2WidgetDistributionDefinition) {
	if err := p.validatePutDistributionDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDistributionDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutEventStreamDefinition(value *PowerpackV2WidgetEventStreamDefinition) {
	if err := p.validatePutEventStreamDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEventStreamDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutEventTimelineDefinition(value *PowerpackV2WidgetEventTimelineDefinition) {
	if err := p.validatePutEventTimelineDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEventTimelineDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutFreeTextDefinition(value *PowerpackV2WidgetFreeTextDefinition) {
	if err := p.validatePutFreeTextDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFreeTextDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutFunnelDefinition(value *PowerpackV2WidgetFunnelDefinition) {
	if err := p.validatePutFunnelDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFunnelDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutGeomapDefinition(value *PowerpackV2WidgetGeomapDefinition) {
	if err := p.validatePutGeomapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGeomapDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutGroupDefinition(value *PowerpackV2WidgetGroupDefinition) {
	if err := p.validatePutGroupDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGroupDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutHeatmapDefinition(value *PowerpackV2WidgetHeatmapDefinition) {
	if err := p.validatePutHeatmapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHeatmapDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutHostmapDefinition(value *PowerpackV2WidgetHostmapDefinition) {
	if err := p.validatePutHostmapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHostmapDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutIframeDefinition(value *PowerpackV2WidgetIframeDefinition) {
	if err := p.validatePutIframeDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putIframeDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutImageDefinition(value *PowerpackV2WidgetImageDefinition) {
	if err := p.validatePutImageDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putImageDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutListStreamDefinition(value *PowerpackV2WidgetListStreamDefinition) {
	if err := p.validatePutListStreamDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putListStreamDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutLogStreamDefinition(value *PowerpackV2WidgetLogStreamDefinition) {
	if err := p.validatePutLogStreamDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLogStreamDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutManageStatusDefinition(value *PowerpackV2WidgetManageStatusDefinition) {
	if err := p.validatePutManageStatusDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putManageStatusDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutNoteDefinition(value *PowerpackV2WidgetNoteDefinition) {
	if err := p.validatePutNoteDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNoteDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutPointPlotDefinition(value *PowerpackV2WidgetPointPlotDefinition) {
	if err := p.validatePutPointPlotDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPointPlotDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutQueryTableDefinition(value *PowerpackV2WidgetQueryTableDefinition) {
	if err := p.validatePutQueryTableDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQueryTableDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutQueryValueDefinition(value *PowerpackV2WidgetQueryValueDefinition) {
	if err := p.validatePutQueryValueDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQueryValueDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutRunWorkflowDefinition(value *PowerpackV2WidgetRunWorkflowDefinition) {
	if err := p.validatePutRunWorkflowDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRunWorkflowDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutSankeyDefinition(value *PowerpackV2WidgetSankeyDefinition) {
	if err := p.validatePutSankeyDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSankeyDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutScatterplotDefinition(value *PowerpackV2WidgetScatterplotDefinition) {
	if err := p.validatePutScatterplotDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putScatterplotDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutServiceLevelObjectiveDefinition(value *PowerpackV2WidgetServiceLevelObjectiveDefinition) {
	if err := p.validatePutServiceLevelObjectiveDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putServiceLevelObjectiveDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutServicemapDefinition(value *PowerpackV2WidgetServicemapDefinition) {
	if err := p.validatePutServicemapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putServicemapDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutSloListDefinition(value *PowerpackV2WidgetSloListDefinition) {
	if err := p.validatePutSloListDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSloListDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutSunburstDefinition(value *PowerpackV2WidgetSunburstDefinition) {
	if err := p.validatePutSunburstDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSunburstDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutTimeseriesDefinition(value *PowerpackV2WidgetTimeseriesDefinition) {
	if err := p.validatePutTimeseriesDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTimeseriesDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutToplistDefinition(value *PowerpackV2WidgetToplistDefinition) {
	if err := p.validatePutToplistDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putToplistDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutTopologyMapDefinition(value *PowerpackV2WidgetTopologyMapDefinition) {
	if err := p.validatePutTopologyMapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTopologyMapDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutTraceServiceDefinition(value *PowerpackV2WidgetTraceServiceDefinition) {
	if err := p.validatePutTraceServiceDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTraceServiceDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutTreemapDefinition(value *PowerpackV2WidgetTreemapDefinition) {
	if err := p.validatePutTreemapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTreemapDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutWidgetLayout(value *PowerpackV2WidgetWidgetLayout) {
	if err := p.validatePutWidgetLayoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putWidgetLayout",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) PutWildcardDefinition(value *PowerpackV2WidgetWildcardDefinition) {
	if err := p.validatePutWildcardDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putWildcardDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetAlertGraphDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetAlertGraphDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetAlertValueDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetAlertValueDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetBarChartDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetBarChartDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetChangeDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetChangeDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetCheckStatusDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetCheckStatusDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetDistributionDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetDistributionDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetEventStreamDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetEventStreamDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetEventTimelineDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetEventTimelineDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetFreeTextDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetFreeTextDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetFunnelDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetFunnelDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetGeomapDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetGeomapDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetGroupDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetGroupDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetHeatmapDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetHeatmapDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetHostmapDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetHostmapDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetIframeDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetIframeDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetImageDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetImageDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetListStreamDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetListStreamDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetLogStreamDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetLogStreamDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetManageStatusDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetManageStatusDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetNoteDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetNoteDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetPointPlotDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetPointPlotDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetQueryTableDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetQueryTableDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetQueryValueDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetQueryValueDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetRunWorkflowDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetRunWorkflowDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetSankeyDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetSankeyDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetScatterplotDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetScatterplotDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetServiceLevelObjectiveDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetServiceLevelObjectiveDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetServicemapDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetServicemapDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetSloListDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetSloListDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetSunburstDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetSunburstDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetTimeseriesDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeseriesDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetToplistDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetToplistDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetTopologyMapDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetTopologyMapDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetTraceServiceDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetTraceServiceDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetTreemapDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetTreemapDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetWidgetLayout() {
	_jsii_.InvokeVoid(
		p,
		"resetWidgetLayout",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ResetWildcardDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetWildcardDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

