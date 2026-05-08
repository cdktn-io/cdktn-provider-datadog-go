// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetOutputReference interface {
	cdktn.ComplexObject
	AlertGraphDefinition() PowerpackV2WidgetGroupDefinitionWidgetAlertGraphDefinitionOutputReference
	AlertGraphDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetAlertGraphDefinition
	AlertValueDefinition() PowerpackV2WidgetGroupDefinitionWidgetAlertValueDefinitionOutputReference
	AlertValueDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetAlertValueDefinition
	ChangeDefinition() PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionOutputReference
	ChangeDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetChangeDefinition
	CheckStatusDefinition() PowerpackV2WidgetGroupDefinitionWidgetCheckStatusDefinitionOutputReference
	CheckStatusDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetCheckStatusDefinition
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
	DistributionDefinition() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionOutputReference
	DistributionDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinition
	EventStreamDefinition() PowerpackV2WidgetGroupDefinitionWidgetEventStreamDefinitionOutputReference
	EventStreamDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetEventStreamDefinition
	EventTimelineDefinition() PowerpackV2WidgetGroupDefinitionWidgetEventTimelineDefinitionOutputReference
	EventTimelineDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetEventTimelineDefinition
	// Experimental.
	Fqn() *string
	FreeTextDefinition() PowerpackV2WidgetGroupDefinitionWidgetFreeTextDefinitionOutputReference
	FreeTextDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetFreeTextDefinition
	FunnelDefinition() PowerpackV2WidgetGroupDefinitionWidgetFunnelDefinitionOutputReference
	FunnelDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetFunnelDefinition
	GeomapDefinition() PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionOutputReference
	GeomapDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinition
	HeatmapDefinition() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference
	HeatmapDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinition
	HostmapDefinition() PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionOutputReference
	HostmapDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinition
	Id() *float64
	SetId(val *float64)
	IdInput() *float64
	IframeDefinition() PowerpackV2WidgetGroupDefinitionWidgetIframeDefinitionOutputReference
	IframeDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetIframeDefinition
	ImageDefinition() PowerpackV2WidgetGroupDefinitionWidgetImageDefinitionOutputReference
	ImageDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetImageDefinition
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ListStreamDefinition() PowerpackV2WidgetGroupDefinitionWidgetListStreamDefinitionOutputReference
	ListStreamDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetListStreamDefinition
	LogStreamDefinition() PowerpackV2WidgetGroupDefinitionWidgetLogStreamDefinitionOutputReference
	LogStreamDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetLogStreamDefinition
	ManageStatusDefinition() PowerpackV2WidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference
	ManageStatusDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetManageStatusDefinition
	NoteDefinition() PowerpackV2WidgetGroupDefinitionWidgetNoteDefinitionOutputReference
	NoteDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetNoteDefinition
	QueryTableDefinition() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionOutputReference
	QueryTableDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinition
	QueryValueDefinition() PowerpackV2WidgetGroupDefinitionWidgetQueryValueDefinitionOutputReference
	QueryValueDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryValueDefinition
	RunWorkflowDefinition() PowerpackV2WidgetGroupDefinitionWidgetRunWorkflowDefinitionOutputReference
	RunWorkflowDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetRunWorkflowDefinition
	ScatterplotDefinition() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionOutputReference
	ScatterplotDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinition
	ServiceLevelObjectiveDefinition() PowerpackV2WidgetGroupDefinitionWidgetServiceLevelObjectiveDefinitionOutputReference
	ServiceLevelObjectiveDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetServiceLevelObjectiveDefinition
	ServicemapDefinition() PowerpackV2WidgetGroupDefinitionWidgetServicemapDefinitionOutputReference
	ServicemapDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetServicemapDefinition
	SloListDefinition() PowerpackV2WidgetGroupDefinitionWidgetSloListDefinitionOutputReference
	SloListDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetSloListDefinition
	SunburstDefinition() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionOutputReference
	SunburstDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinition
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeseriesDefinition() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionOutputReference
	TimeseriesDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinition
	ToplistDefinition() PowerpackV2WidgetGroupDefinitionWidgetToplistDefinitionOutputReference
	ToplistDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetToplistDefinition
	TopologyMapDefinition() PowerpackV2WidgetGroupDefinitionWidgetTopologyMapDefinitionOutputReference
	TopologyMapDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetTopologyMapDefinition
	TraceServiceDefinition() PowerpackV2WidgetGroupDefinitionWidgetTraceServiceDefinitionOutputReference
	TraceServiceDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetTraceServiceDefinition
	TreemapDefinition() PowerpackV2WidgetGroupDefinitionWidgetTreemapDefinitionOutputReference
	TreemapDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetTreemapDefinition
	WidgetLayout() PowerpackV2WidgetGroupDefinitionWidgetWidgetLayoutOutputReference
	WidgetLayoutInput() *PowerpackV2WidgetGroupDefinitionWidgetWidgetLayout
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
	PutAlertGraphDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetAlertGraphDefinition)
	PutAlertValueDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetAlertValueDefinition)
	PutChangeDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetChangeDefinition)
	PutCheckStatusDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetCheckStatusDefinition)
	PutDistributionDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinition)
	PutEventStreamDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetEventStreamDefinition)
	PutEventTimelineDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetEventTimelineDefinition)
	PutFreeTextDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetFreeTextDefinition)
	PutFunnelDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetFunnelDefinition)
	PutGeomapDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinition)
	PutHeatmapDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinition)
	PutHostmapDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinition)
	PutIframeDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetIframeDefinition)
	PutImageDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetImageDefinition)
	PutListStreamDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetListStreamDefinition)
	PutLogStreamDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetLogStreamDefinition)
	PutManageStatusDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetManageStatusDefinition)
	PutNoteDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetNoteDefinition)
	PutQueryTableDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinition)
	PutQueryValueDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetQueryValueDefinition)
	PutRunWorkflowDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetRunWorkflowDefinition)
	PutScatterplotDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinition)
	PutServiceLevelObjectiveDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetServiceLevelObjectiveDefinition)
	PutServicemapDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetServicemapDefinition)
	PutSloListDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetSloListDefinition)
	PutSunburstDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinition)
	PutTimeseriesDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinition)
	PutToplistDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetToplistDefinition)
	PutTopologyMapDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetTopologyMapDefinition)
	PutTraceServiceDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetTraceServiceDefinition)
	PutTreemapDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetTreemapDefinition)
	PutWidgetLayout(value *PowerpackV2WidgetGroupDefinitionWidgetWidgetLayout)
	ResetAlertGraphDefinition()
	ResetAlertValueDefinition()
	ResetChangeDefinition()
	ResetCheckStatusDefinition()
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
	ResetQueryTableDefinition()
	ResetQueryValueDefinition()
	ResetRunWorkflowDefinition()
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) AlertGraphDefinition() PowerpackV2WidgetGroupDefinitionWidgetAlertGraphDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetAlertGraphDefinitionOutputReference
	_jsii_.Get(
		j,
		"alertGraphDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) AlertGraphDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetAlertGraphDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetAlertGraphDefinition
	_jsii_.Get(
		j,
		"alertGraphDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) AlertValueDefinition() PowerpackV2WidgetGroupDefinitionWidgetAlertValueDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetAlertValueDefinitionOutputReference
	_jsii_.Get(
		j,
		"alertValueDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) AlertValueDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetAlertValueDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetAlertValueDefinition
	_jsii_.Get(
		j,
		"alertValueDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ChangeDefinition() PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionOutputReference
	_jsii_.Get(
		j,
		"changeDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ChangeDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetChangeDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetChangeDefinition
	_jsii_.Get(
		j,
		"changeDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) CheckStatusDefinition() PowerpackV2WidgetGroupDefinitionWidgetCheckStatusDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetCheckStatusDefinitionOutputReference
	_jsii_.Get(
		j,
		"checkStatusDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) CheckStatusDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetCheckStatusDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetCheckStatusDefinition
	_jsii_.Get(
		j,
		"checkStatusDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) DistributionDefinition() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionOutputReference
	_jsii_.Get(
		j,
		"distributionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) DistributionDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinition
	_jsii_.Get(
		j,
		"distributionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) EventStreamDefinition() PowerpackV2WidgetGroupDefinitionWidgetEventStreamDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetEventStreamDefinitionOutputReference
	_jsii_.Get(
		j,
		"eventStreamDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) EventStreamDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetEventStreamDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetEventStreamDefinition
	_jsii_.Get(
		j,
		"eventStreamDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) EventTimelineDefinition() PowerpackV2WidgetGroupDefinitionWidgetEventTimelineDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetEventTimelineDefinitionOutputReference
	_jsii_.Get(
		j,
		"eventTimelineDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) EventTimelineDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetEventTimelineDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetEventTimelineDefinition
	_jsii_.Get(
		j,
		"eventTimelineDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) FreeTextDefinition() PowerpackV2WidgetGroupDefinitionWidgetFreeTextDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetFreeTextDefinitionOutputReference
	_jsii_.Get(
		j,
		"freeTextDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) FreeTextDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetFreeTextDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetFreeTextDefinition
	_jsii_.Get(
		j,
		"freeTextDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) FunnelDefinition() PowerpackV2WidgetGroupDefinitionWidgetFunnelDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetFunnelDefinitionOutputReference
	_jsii_.Get(
		j,
		"funnelDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) FunnelDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetFunnelDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetFunnelDefinition
	_jsii_.Get(
		j,
		"funnelDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) GeomapDefinition() PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionOutputReference
	_jsii_.Get(
		j,
		"geomapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) GeomapDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinition
	_jsii_.Get(
		j,
		"geomapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) HeatmapDefinition() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference
	_jsii_.Get(
		j,
		"heatmapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) HeatmapDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinition
	_jsii_.Get(
		j,
		"heatmapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) HostmapDefinition() PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionOutputReference
	_jsii_.Get(
		j,
		"hostmapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) HostmapDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinition
	_jsii_.Get(
		j,
		"hostmapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) IdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) IframeDefinition() PowerpackV2WidgetGroupDefinitionWidgetIframeDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetIframeDefinitionOutputReference
	_jsii_.Get(
		j,
		"iframeDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) IframeDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetIframeDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetIframeDefinition
	_jsii_.Get(
		j,
		"iframeDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ImageDefinition() PowerpackV2WidgetGroupDefinitionWidgetImageDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetImageDefinitionOutputReference
	_jsii_.Get(
		j,
		"imageDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ImageDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetImageDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetImageDefinition
	_jsii_.Get(
		j,
		"imageDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ListStreamDefinition() PowerpackV2WidgetGroupDefinitionWidgetListStreamDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetListStreamDefinitionOutputReference
	_jsii_.Get(
		j,
		"listStreamDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ListStreamDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetListStreamDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetListStreamDefinition
	_jsii_.Get(
		j,
		"listStreamDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) LogStreamDefinition() PowerpackV2WidgetGroupDefinitionWidgetLogStreamDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetLogStreamDefinitionOutputReference
	_jsii_.Get(
		j,
		"logStreamDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) LogStreamDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetLogStreamDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetLogStreamDefinition
	_jsii_.Get(
		j,
		"logStreamDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ManageStatusDefinition() PowerpackV2WidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference
	_jsii_.Get(
		j,
		"manageStatusDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ManageStatusDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetManageStatusDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetManageStatusDefinition
	_jsii_.Get(
		j,
		"manageStatusDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) NoteDefinition() PowerpackV2WidgetGroupDefinitionWidgetNoteDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetNoteDefinitionOutputReference
	_jsii_.Get(
		j,
		"noteDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) NoteDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetNoteDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetNoteDefinition
	_jsii_.Get(
		j,
		"noteDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) QueryTableDefinition() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionOutputReference
	_jsii_.Get(
		j,
		"queryTableDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) QueryTableDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinition
	_jsii_.Get(
		j,
		"queryTableDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) QueryValueDefinition() PowerpackV2WidgetGroupDefinitionWidgetQueryValueDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetQueryValueDefinitionOutputReference
	_jsii_.Get(
		j,
		"queryValueDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) QueryValueDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryValueDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetQueryValueDefinition
	_jsii_.Get(
		j,
		"queryValueDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) RunWorkflowDefinition() PowerpackV2WidgetGroupDefinitionWidgetRunWorkflowDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetRunWorkflowDefinitionOutputReference
	_jsii_.Get(
		j,
		"runWorkflowDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) RunWorkflowDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetRunWorkflowDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetRunWorkflowDefinition
	_jsii_.Get(
		j,
		"runWorkflowDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ScatterplotDefinition() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionOutputReference
	_jsii_.Get(
		j,
		"scatterplotDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ScatterplotDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinition
	_jsii_.Get(
		j,
		"scatterplotDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ServiceLevelObjectiveDefinition() PowerpackV2WidgetGroupDefinitionWidgetServiceLevelObjectiveDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetServiceLevelObjectiveDefinitionOutputReference
	_jsii_.Get(
		j,
		"serviceLevelObjectiveDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ServiceLevelObjectiveDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetServiceLevelObjectiveDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetServiceLevelObjectiveDefinition
	_jsii_.Get(
		j,
		"serviceLevelObjectiveDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ServicemapDefinition() PowerpackV2WidgetGroupDefinitionWidgetServicemapDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetServicemapDefinitionOutputReference
	_jsii_.Get(
		j,
		"servicemapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ServicemapDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetServicemapDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetServicemapDefinition
	_jsii_.Get(
		j,
		"servicemapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) SloListDefinition() PowerpackV2WidgetGroupDefinitionWidgetSloListDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetSloListDefinitionOutputReference
	_jsii_.Get(
		j,
		"sloListDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) SloListDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetSloListDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetSloListDefinition
	_jsii_.Get(
		j,
		"sloListDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) SunburstDefinition() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionOutputReference
	_jsii_.Get(
		j,
		"sunburstDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) SunburstDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinition
	_jsii_.Get(
		j,
		"sunburstDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) TimeseriesDefinition() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionOutputReference
	_jsii_.Get(
		j,
		"timeseriesDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) TimeseriesDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinition
	_jsii_.Get(
		j,
		"timeseriesDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ToplistDefinition() PowerpackV2WidgetGroupDefinitionWidgetToplistDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetToplistDefinitionOutputReference
	_jsii_.Get(
		j,
		"toplistDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ToplistDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetToplistDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetToplistDefinition
	_jsii_.Get(
		j,
		"toplistDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) TopologyMapDefinition() PowerpackV2WidgetGroupDefinitionWidgetTopologyMapDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetTopologyMapDefinitionOutputReference
	_jsii_.Get(
		j,
		"topologyMapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) TopologyMapDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetTopologyMapDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetTopologyMapDefinition
	_jsii_.Get(
		j,
		"topologyMapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) TraceServiceDefinition() PowerpackV2WidgetGroupDefinitionWidgetTraceServiceDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetTraceServiceDefinitionOutputReference
	_jsii_.Get(
		j,
		"traceServiceDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) TraceServiceDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetTraceServiceDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetTraceServiceDefinition
	_jsii_.Get(
		j,
		"traceServiceDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) TreemapDefinition() PowerpackV2WidgetGroupDefinitionWidgetTreemapDefinitionOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetTreemapDefinitionOutputReference
	_jsii_.Get(
		j,
		"treemapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) TreemapDefinitionInput() *PowerpackV2WidgetGroupDefinitionWidgetTreemapDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetTreemapDefinition
	_jsii_.Get(
		j,
		"treemapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) WidgetLayout() PowerpackV2WidgetGroupDefinitionWidgetWidgetLayoutOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetWidgetLayoutOutputReference
	_jsii_.Get(
		j,
		"widgetLayout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) WidgetLayoutInput() *PowerpackV2WidgetGroupDefinitionWidgetWidgetLayout {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetWidgetLayout
	_jsii_.Get(
		j,
		"widgetLayoutInput",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetGroupDefinitionWidgetOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference)SetId(val *float64) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutAlertGraphDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetAlertGraphDefinition) {
	if err := p.validatePutAlertGraphDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAlertGraphDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutAlertValueDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetAlertValueDefinition) {
	if err := p.validatePutAlertValueDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAlertValueDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutChangeDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetChangeDefinition) {
	if err := p.validatePutChangeDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putChangeDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutCheckStatusDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetCheckStatusDefinition) {
	if err := p.validatePutCheckStatusDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCheckStatusDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutDistributionDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinition) {
	if err := p.validatePutDistributionDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDistributionDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutEventStreamDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetEventStreamDefinition) {
	if err := p.validatePutEventStreamDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEventStreamDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutEventTimelineDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetEventTimelineDefinition) {
	if err := p.validatePutEventTimelineDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEventTimelineDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutFreeTextDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetFreeTextDefinition) {
	if err := p.validatePutFreeTextDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFreeTextDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutFunnelDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetFunnelDefinition) {
	if err := p.validatePutFunnelDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFunnelDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutGeomapDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinition) {
	if err := p.validatePutGeomapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGeomapDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutHeatmapDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinition) {
	if err := p.validatePutHeatmapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHeatmapDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutHostmapDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinition) {
	if err := p.validatePutHostmapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHostmapDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutIframeDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetIframeDefinition) {
	if err := p.validatePutIframeDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putIframeDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutImageDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetImageDefinition) {
	if err := p.validatePutImageDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putImageDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutListStreamDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetListStreamDefinition) {
	if err := p.validatePutListStreamDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putListStreamDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutLogStreamDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetLogStreamDefinition) {
	if err := p.validatePutLogStreamDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLogStreamDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutManageStatusDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetManageStatusDefinition) {
	if err := p.validatePutManageStatusDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putManageStatusDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutNoteDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetNoteDefinition) {
	if err := p.validatePutNoteDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNoteDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutQueryTableDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinition) {
	if err := p.validatePutQueryTableDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQueryTableDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutQueryValueDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetQueryValueDefinition) {
	if err := p.validatePutQueryValueDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQueryValueDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutRunWorkflowDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetRunWorkflowDefinition) {
	if err := p.validatePutRunWorkflowDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRunWorkflowDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutScatterplotDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinition) {
	if err := p.validatePutScatterplotDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putScatterplotDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutServiceLevelObjectiveDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetServiceLevelObjectiveDefinition) {
	if err := p.validatePutServiceLevelObjectiveDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putServiceLevelObjectiveDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutServicemapDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetServicemapDefinition) {
	if err := p.validatePutServicemapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putServicemapDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutSloListDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetSloListDefinition) {
	if err := p.validatePutSloListDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSloListDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutSunburstDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinition) {
	if err := p.validatePutSunburstDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSunburstDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutTimeseriesDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinition) {
	if err := p.validatePutTimeseriesDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTimeseriesDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutToplistDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetToplistDefinition) {
	if err := p.validatePutToplistDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putToplistDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutTopologyMapDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetTopologyMapDefinition) {
	if err := p.validatePutTopologyMapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTopologyMapDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutTraceServiceDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetTraceServiceDefinition) {
	if err := p.validatePutTraceServiceDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTraceServiceDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutTreemapDefinition(value *PowerpackV2WidgetGroupDefinitionWidgetTreemapDefinition) {
	if err := p.validatePutTreemapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTreemapDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) PutWidgetLayout(value *PowerpackV2WidgetGroupDefinitionWidgetWidgetLayout) {
	if err := p.validatePutWidgetLayoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putWidgetLayout",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetAlertGraphDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetAlertGraphDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetAlertValueDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetAlertValueDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetChangeDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetChangeDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetCheckStatusDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetCheckStatusDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetDistributionDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetDistributionDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetEventStreamDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetEventStreamDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetEventTimelineDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetEventTimelineDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetFreeTextDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetFreeTextDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetFunnelDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetFunnelDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetGeomapDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetGeomapDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetHeatmapDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetHeatmapDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetHostmapDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetHostmapDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetIframeDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetIframeDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetImageDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetImageDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetListStreamDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetListStreamDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetLogStreamDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetLogStreamDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetManageStatusDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetManageStatusDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetNoteDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetNoteDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetQueryTableDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetQueryTableDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetQueryValueDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetQueryValueDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetRunWorkflowDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetRunWorkflowDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetScatterplotDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetScatterplotDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetServiceLevelObjectiveDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetServiceLevelObjectiveDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetServicemapDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetServicemapDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetSloListDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetSloListDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetSunburstDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetSunburstDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetTimeseriesDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeseriesDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetToplistDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetToplistDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetTopologyMapDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetTopologyMapDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetTraceServiceDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetTraceServiceDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetTreemapDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetTreemapDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ResetWidgetLayout() {
	_jsii_.InvokeVoid(
		p,
		"resetWidgetLayout",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

