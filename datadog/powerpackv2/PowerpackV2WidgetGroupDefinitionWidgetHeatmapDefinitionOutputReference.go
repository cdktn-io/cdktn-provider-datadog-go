// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference interface {
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
	CustomLink() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionCustomLinkList
	CustomLinkInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Event() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionEventList
	EventInput() interface{}
	// Experimental.
	Fqn() *string
	HideIncompleteCostData() interface{}
	SetHideIncompleteCostData(val interface{})
	HideIncompleteCostDataInput() interface{}
	InternalValue() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinition
	SetInternalValue(val *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinition)
	LegendSize() *string
	SetLegendSize(val *string)
	LegendSizeInput() *string
	LiveSpan() *string
	SetLiveSpan(val *string)
	LiveSpanInput() *string
	Marker() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionMarkerList
	MarkerInput() interface{}
	Request() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestList
	RequestInput() interface{}
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
	Time() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionTimeOutputReference
	TimeInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionTime
	Title() *string
	SetTitle(val *string)
	TitleAlign() *string
	SetTitleAlign(val *string)
	TitleAlignInput() *string
	TitleInput() *string
	TitleSize() *string
	SetTitleSize(val *string)
	TitleSizeInput() *string
	Xaxis() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionXaxisOutputReference
	XaxisInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionXaxis
	Yaxis() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionYaxisOutputReference
	YaxisInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionYaxis
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
	PutTime(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionTime)
	PutXaxis(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionXaxis)
	PutYaxis(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionYaxis)
	ResetCustomLink()
	ResetDescription()
	ResetEvent()
	ResetHideIncompleteCostData()
	ResetLegendSize()
	ResetLiveSpan()
	ResetMarker()
	ResetRequest()
	ResetShowLegend()
	ResetTime()
	ResetTitle()
	ResetTitleAlign()
	ResetTitleSize()
	ResetXaxis()
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

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) CustomLink() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionCustomLinkList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionCustomLinkList
	_jsii_.Get(
		j,
		"customLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) CustomLinkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) Event() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionEventList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionEventList
	_jsii_.Get(
		j,
		"event",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) EventInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) HideIncompleteCostData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIncompleteCostData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) HideIncompleteCostDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIncompleteCostDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) InternalValue() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinition {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) LegendSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"legendSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) LegendSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"legendSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) LiveSpan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) LiveSpanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) Marker() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionMarkerList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionMarkerList
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) MarkerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"markerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) Request() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestList
	_jsii_.Get(
		j,
		"request",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) RequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) ShowLegend() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLegend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) ShowLegendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLegendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) Time() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionTimeOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionTimeOutputReference
	_jsii_.Get(
		j,
		"time",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) TimeInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionTime {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionTime
	_jsii_.Get(
		j,
		"timeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) TitleAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) TitleAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) TitleSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) TitleSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) Xaxis() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionXaxisOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionXaxisOutputReference
	_jsii_.Get(
		j,
		"xaxis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) XaxisInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionXaxis {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionXaxis
	_jsii_.Get(
		j,
		"xaxisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) Yaxis() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionYaxisOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionYaxisOutputReference
	_jsii_.Get(
		j,
		"yaxis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) YaxisInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionYaxis {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionYaxis
	_jsii_.Get(
		j,
		"yaxisInput",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference)SetHideIncompleteCostData(val interface{}) {
	if err := j.validateSetHideIncompleteCostDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideIncompleteCostData",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference)SetInternalValue(val *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference)SetLegendSize(val *string) {
	if err := j.validateSetLegendSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"legendSize",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference)SetLiveSpan(val *string) {
	if err := j.validateSetLiveSpanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveSpan",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference)SetShowLegend(val interface{}) {
	if err := j.validateSetShowLegendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showLegend",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference)SetTitleAlign(val *string) {
	if err := j.validateSetTitleAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleAlign",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference)SetTitleSize(val *string) {
	if err := j.validateSetTitleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleSize",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) PutCustomLink(value interface{}) {
	if err := p.validatePutCustomLinkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCustomLink",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) PutEvent(value interface{}) {
	if err := p.validatePutEventParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEvent",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) PutMarker(value interface{}) {
	if err := p.validatePutMarkerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMarker",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) PutRequest(value interface{}) {
	if err := p.validatePutRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRequest",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) PutTime(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionTime) {
	if err := p.validatePutTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTime",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) PutXaxis(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionXaxis) {
	if err := p.validatePutXaxisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putXaxis",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) PutYaxis(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionYaxis) {
	if err := p.validatePutYaxisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putYaxis",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) ResetCustomLink() {
	_jsii_.InvokeVoid(
		p,
		"resetCustomLink",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		p,
		"resetDescription",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) ResetEvent() {
	_jsii_.InvokeVoid(
		p,
		"resetEvent",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) ResetHideIncompleteCostData() {
	_jsii_.InvokeVoid(
		p,
		"resetHideIncompleteCostData",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) ResetLegendSize() {
	_jsii_.InvokeVoid(
		p,
		"resetLegendSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) ResetLiveSpan() {
	_jsii_.InvokeVoid(
		p,
		"resetLiveSpan",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) ResetMarker() {
	_jsii_.InvokeVoid(
		p,
		"resetMarker",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) ResetRequest() {
	_jsii_.InvokeVoid(
		p,
		"resetRequest",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) ResetShowLegend() {
	_jsii_.InvokeVoid(
		p,
		"resetShowLegend",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) ResetTime() {
	_jsii_.InvokeVoid(
		p,
		"resetTime",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		p,
		"resetTitle",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) ResetTitleAlign() {
	_jsii_.InvokeVoid(
		p,
		"resetTitleAlign",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) ResetTitleSize() {
	_jsii_.InvokeVoid(
		p,
		"resetTitleSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) ResetXaxis() {
	_jsii_.InvokeVoid(
		p,
		"resetXaxis",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) ResetYaxis() {
	_jsii_.InvokeVoid(
		p,
		"resetYaxis",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

