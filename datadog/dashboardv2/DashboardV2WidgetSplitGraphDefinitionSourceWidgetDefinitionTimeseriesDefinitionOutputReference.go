// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference interface {
	cdktn.ComplexObject
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
	CustomLink() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionCustomLinkList
	CustomLinkInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Event() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionEventList
	EventInput() interface{}
	// Experimental.
	Fqn() *string
	HideIncompleteCostData() interface{}
	SetHideIncompleteCostData(val interface{})
	HideIncompleteCostDataInput() interface{}
	InternalValue() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinition
	SetInternalValue(val *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinition)
	LegendColumns() *[]*string
	SetLegendColumns(val *[]*string)
	LegendColumnsInput() *[]*string
	LegendLayout() *string
	SetLegendLayout(val *string)
	LegendLayoutInput() *string
	LegendSize() *string
	SetLegendSize(val *string)
	LegendSizeInput() *string
	LiveSpan() *string
	SetLiveSpan(val *string)
	LiveSpanInput() *string
	Marker() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionMarkerList
	MarkerInput() interface{}
	Request() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestList
	RequestInput() interface{}
	RightYaxis() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRightYaxisOutputReference
	RightYaxisInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRightYaxis
	ShowLegend() interface{}
	SetShowLegend(val interface{})
	ShowLegendInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Time() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionTimeOutputReference
	TimeInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionTime
	Title() *string
	SetTitle(val *string)
	TitleAlign() *string
	SetTitleAlign(val *string)
	TitleAlignInput() *string
	TitleInput() *string
	TitleSize() *string
	SetTitleSize(val *string)
	TitleSizeInput() *string
	Yaxis() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionYaxisOutputReference
	YaxisInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionYaxis
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
	PutCustomLink(value interface{})
	PutEvent(value interface{})
	PutMarker(value interface{})
	PutRequest(value interface{})
	PutRightYaxis(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRightYaxis)
	PutTime(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionTime)
	PutYaxis(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionYaxis)
	ResetCustomLink()
	ResetDescription()
	ResetEvent()
	ResetHideIncompleteCostData()
	ResetLegendColumns()
	ResetLegendLayout()
	ResetLegendSize()
	ResetLiveSpan()
	ResetMarker()
	ResetRequest()
	ResetRightYaxis()
	ResetShowLegend()
	ResetTime()
	ResetTitle()
	ResetTitleAlign()
	ResetTitleSize()
	ResetYaxis()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference
type jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) CustomLink() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionCustomLinkList {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionCustomLinkList
	_jsii_.Get(
		j,
		"customLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) CustomLinkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) Event() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionEventList {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionEventList
	_jsii_.Get(
		j,
		"event",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) EventInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) HideIncompleteCostData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIncompleteCostData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) HideIncompleteCostDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIncompleteCostDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) InternalValue() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinition {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) LegendColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"legendColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) LegendColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"legendColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) LegendLayout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"legendLayout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) LegendLayoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"legendLayoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) LegendSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"legendSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) LegendSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"legendSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) LiveSpan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) LiveSpanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) Marker() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionMarkerList {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionMarkerList
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) MarkerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"markerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) Request() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestList {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestList
	_jsii_.Get(
		j,
		"request",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) RequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) RightYaxis() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRightYaxisOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRightYaxisOutputReference
	_jsii_.Get(
		j,
		"rightYaxis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) RightYaxisInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRightYaxis {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRightYaxis
	_jsii_.Get(
		j,
		"rightYaxisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ShowLegend() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLegend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ShowLegendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLegendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) Time() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionTimeOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionTimeOutputReference
	_jsii_.Get(
		j,
		"time",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) TimeInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionTime {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionTime
	_jsii_.Get(
		j,
		"timeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) TitleAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) TitleAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) TitleSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) TitleSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) Yaxis() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionYaxisOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionYaxisOutputReference
	_jsii_.Get(
		j,
		"yaxis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) YaxisInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionYaxis {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionYaxis
	_jsii_.Get(
		j,
		"yaxisInput",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference_Override(d DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetHideIncompleteCostData(val interface{}) {
	if err := j.validateSetHideIncompleteCostDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideIncompleteCostData",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetInternalValue(val *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetLegendColumns(val *[]*string) {
	if err := j.validateSetLegendColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"legendColumns",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetLegendLayout(val *string) {
	if err := j.validateSetLegendLayoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"legendLayout",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetLegendSize(val *string) {
	if err := j.validateSetLegendSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"legendSize",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetLiveSpan(val *string) {
	if err := j.validateSetLiveSpanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveSpan",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetShowLegend(val interface{}) {
	if err := j.validateSetShowLegendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showLegend",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetTitleAlign(val *string) {
	if err := j.validateSetTitleAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleAlign",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetTitleSize(val *string) {
	if err := j.validateSetTitleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleSize",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) PutCustomLink(value interface{}) {
	if err := d.validatePutCustomLinkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustomLink",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) PutEvent(value interface{}) {
	if err := d.validatePutEventParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEvent",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) PutMarker(value interface{}) {
	if err := d.validatePutMarkerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMarker",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) PutRequest(value interface{}) {
	if err := d.validatePutRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRequest",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) PutRightYaxis(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRightYaxis) {
	if err := d.validatePutRightYaxisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRightYaxis",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) PutTime(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionTime) {
	if err := d.validatePutTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTime",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) PutYaxis(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionYaxis) {
	if err := d.validatePutYaxisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putYaxis",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetCustomLink() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomLink",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetEvent() {
	_jsii_.InvokeVoid(
		d,
		"resetEvent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetHideIncompleteCostData() {
	_jsii_.InvokeVoid(
		d,
		"resetHideIncompleteCostData",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetLegendColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetLegendColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetLegendLayout() {
	_jsii_.InvokeVoid(
		d,
		"resetLegendLayout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetLegendSize() {
	_jsii_.InvokeVoid(
		d,
		"resetLegendSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetLiveSpan() {
	_jsii_.InvokeVoid(
		d,
		"resetLiveSpan",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetMarker() {
	_jsii_.InvokeVoid(
		d,
		"resetMarker",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetRequest() {
	_jsii_.InvokeVoid(
		d,
		"resetRequest",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetRightYaxis() {
	_jsii_.InvokeVoid(
		d,
		"resetRightYaxis",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetShowLegend() {
	_jsii_.InvokeVoid(
		d,
		"resetShowLegend",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetTime() {
	_jsii_.InvokeVoid(
		d,
		"resetTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		d,
		"resetTitle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetTitleAlign() {
	_jsii_.InvokeVoid(
		d,
		"resetTitleAlign",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetTitleSize() {
	_jsii_.InvokeVoid(
		d,
		"resetTitleSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetYaxis() {
	_jsii_.InvokeVoid(
		d,
		"resetYaxis",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

