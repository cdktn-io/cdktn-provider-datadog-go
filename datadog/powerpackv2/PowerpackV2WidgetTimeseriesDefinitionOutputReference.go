// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetTimeseriesDefinitionOutputReference interface {
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
	CustomLink() PowerpackV2WidgetTimeseriesDefinitionCustomLinkList
	CustomLinkInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Event() PowerpackV2WidgetTimeseriesDefinitionEventList
	EventInput() interface{}
	// Experimental.
	Fqn() *string
	HideIncompleteCostData() interface{}
	SetHideIncompleteCostData(val interface{})
	HideIncompleteCostDataInput() interface{}
	InternalValue() *PowerpackV2WidgetTimeseriesDefinition
	SetInternalValue(val *PowerpackV2WidgetTimeseriesDefinition)
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
	Marker() PowerpackV2WidgetTimeseriesDefinitionMarkerList
	MarkerInput() interface{}
	Request() PowerpackV2WidgetTimeseriesDefinitionRequestList
	RequestInput() interface{}
	RightYaxis() PowerpackV2WidgetTimeseriesDefinitionRightYaxisOutputReference
	RightYaxisInput() *PowerpackV2WidgetTimeseriesDefinitionRightYaxis
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
	Time() PowerpackV2WidgetTimeseriesDefinitionTimeOutputReference
	TimeInput() *PowerpackV2WidgetTimeseriesDefinitionTime
	Title() *string
	SetTitle(val *string)
	TitleAlign() *string
	SetTitleAlign(val *string)
	TitleAlignInput() *string
	TitleInput() *string
	TitleSize() *string
	SetTitleSize(val *string)
	TitleSizeInput() *string
	Yaxis() PowerpackV2WidgetTimeseriesDefinitionYaxisOutputReference
	YaxisInput() *PowerpackV2WidgetTimeseriesDefinitionYaxis
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
	PutRightYaxis(value *PowerpackV2WidgetTimeseriesDefinitionRightYaxis)
	PutTime(value *PowerpackV2WidgetTimeseriesDefinitionTime)
	PutYaxis(value *PowerpackV2WidgetTimeseriesDefinitionYaxis)
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

// The jsii proxy struct for PowerpackV2WidgetTimeseriesDefinitionOutputReference
type jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) CustomLink() PowerpackV2WidgetTimeseriesDefinitionCustomLinkList {
	var returns PowerpackV2WidgetTimeseriesDefinitionCustomLinkList
	_jsii_.Get(
		j,
		"customLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) CustomLinkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) Event() PowerpackV2WidgetTimeseriesDefinitionEventList {
	var returns PowerpackV2WidgetTimeseriesDefinitionEventList
	_jsii_.Get(
		j,
		"event",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) EventInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) HideIncompleteCostData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIncompleteCostData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) HideIncompleteCostDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIncompleteCostDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) InternalValue() *PowerpackV2WidgetTimeseriesDefinition {
	var returns *PowerpackV2WidgetTimeseriesDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) LegendColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"legendColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) LegendColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"legendColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) LegendLayout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"legendLayout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) LegendLayoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"legendLayoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) LegendSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"legendSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) LegendSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"legendSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) LiveSpan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) LiveSpanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) Marker() PowerpackV2WidgetTimeseriesDefinitionMarkerList {
	var returns PowerpackV2WidgetTimeseriesDefinitionMarkerList
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) MarkerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"markerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) Request() PowerpackV2WidgetTimeseriesDefinitionRequestList {
	var returns PowerpackV2WidgetTimeseriesDefinitionRequestList
	_jsii_.Get(
		j,
		"request",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) RequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) RightYaxis() PowerpackV2WidgetTimeseriesDefinitionRightYaxisOutputReference {
	var returns PowerpackV2WidgetTimeseriesDefinitionRightYaxisOutputReference
	_jsii_.Get(
		j,
		"rightYaxis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) RightYaxisInput() *PowerpackV2WidgetTimeseriesDefinitionRightYaxis {
	var returns *PowerpackV2WidgetTimeseriesDefinitionRightYaxis
	_jsii_.Get(
		j,
		"rightYaxisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ShowLegend() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLegend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ShowLegendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLegendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) Time() PowerpackV2WidgetTimeseriesDefinitionTimeOutputReference {
	var returns PowerpackV2WidgetTimeseriesDefinitionTimeOutputReference
	_jsii_.Get(
		j,
		"time",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) TimeInput() *PowerpackV2WidgetTimeseriesDefinitionTime {
	var returns *PowerpackV2WidgetTimeseriesDefinitionTime
	_jsii_.Get(
		j,
		"timeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) TitleAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) TitleAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) TitleSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) TitleSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) Yaxis() PowerpackV2WidgetTimeseriesDefinitionYaxisOutputReference {
	var returns PowerpackV2WidgetTimeseriesDefinitionYaxisOutputReference
	_jsii_.Get(
		j,
		"yaxis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) YaxisInput() *PowerpackV2WidgetTimeseriesDefinitionYaxis {
	var returns *PowerpackV2WidgetTimeseriesDefinitionYaxis
	_jsii_.Get(
		j,
		"yaxisInput",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetTimeseriesDefinitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetTimeseriesDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetTimeseriesDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetTimeseriesDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetTimeseriesDefinitionOutputReference_Override(p PowerpackV2WidgetTimeseriesDefinitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetTimeseriesDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference)SetHideIncompleteCostData(val interface{}) {
	if err := j.validateSetHideIncompleteCostDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideIncompleteCostData",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference)SetInternalValue(val *PowerpackV2WidgetTimeseriesDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference)SetLegendColumns(val *[]*string) {
	if err := j.validateSetLegendColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"legendColumns",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference)SetLegendLayout(val *string) {
	if err := j.validateSetLegendLayoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"legendLayout",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference)SetLegendSize(val *string) {
	if err := j.validateSetLegendSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"legendSize",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference)SetLiveSpan(val *string) {
	if err := j.validateSetLiveSpanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveSpan",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference)SetShowLegend(val interface{}) {
	if err := j.validateSetShowLegendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showLegend",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference)SetTitleAlign(val *string) {
	if err := j.validateSetTitleAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleAlign",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference)SetTitleSize(val *string) {
	if err := j.validateSetTitleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleSize",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) PutCustomLink(value interface{}) {
	if err := p.validatePutCustomLinkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCustomLink",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) PutEvent(value interface{}) {
	if err := p.validatePutEventParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEvent",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) PutMarker(value interface{}) {
	if err := p.validatePutMarkerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMarker",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) PutRequest(value interface{}) {
	if err := p.validatePutRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRequest",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) PutRightYaxis(value *PowerpackV2WidgetTimeseriesDefinitionRightYaxis) {
	if err := p.validatePutRightYaxisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRightYaxis",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) PutTime(value *PowerpackV2WidgetTimeseriesDefinitionTime) {
	if err := p.validatePutTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTime",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) PutYaxis(value *PowerpackV2WidgetTimeseriesDefinitionYaxis) {
	if err := p.validatePutYaxisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putYaxis",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ResetCustomLink() {
	_jsii_.InvokeVoid(
		p,
		"resetCustomLink",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		p,
		"resetDescription",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ResetEvent() {
	_jsii_.InvokeVoid(
		p,
		"resetEvent",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ResetHideIncompleteCostData() {
	_jsii_.InvokeVoid(
		p,
		"resetHideIncompleteCostData",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ResetLegendColumns() {
	_jsii_.InvokeVoid(
		p,
		"resetLegendColumns",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ResetLegendLayout() {
	_jsii_.InvokeVoid(
		p,
		"resetLegendLayout",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ResetLegendSize() {
	_jsii_.InvokeVoid(
		p,
		"resetLegendSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ResetLiveSpan() {
	_jsii_.InvokeVoid(
		p,
		"resetLiveSpan",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ResetMarker() {
	_jsii_.InvokeVoid(
		p,
		"resetMarker",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ResetRequest() {
	_jsii_.InvokeVoid(
		p,
		"resetRequest",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ResetRightYaxis() {
	_jsii_.InvokeVoid(
		p,
		"resetRightYaxis",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ResetShowLegend() {
	_jsii_.InvokeVoid(
		p,
		"resetShowLegend",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ResetTime() {
	_jsii_.InvokeVoid(
		p,
		"resetTime",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		p,
		"resetTitle",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ResetTitleAlign() {
	_jsii_.InvokeVoid(
		p,
		"resetTitleAlign",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ResetTitleSize() {
	_jsii_.InvokeVoid(
		p,
		"resetTitleSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ResetYaxis() {
	_jsii_.InvokeVoid(
		p,
		"resetYaxis",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

