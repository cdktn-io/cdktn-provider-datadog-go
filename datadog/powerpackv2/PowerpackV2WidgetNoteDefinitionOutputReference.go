// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetNoteDefinitionOutputReference interface {
	cdktn.ComplexObject
	BackgroundColor() *string
	SetBackgroundColor(val *string)
	BackgroundColorInput() *string
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
	Content() *string
	SetContent(val *string)
	ContentInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	FontSize() *string
	SetFontSize(val *string)
	FontSizeInput() *string
	// Experimental.
	Fqn() *string
	HasPadding() interface{}
	SetHasPadding(val interface{})
	HasPaddingInput() interface{}
	HideIncompleteCostData() interface{}
	SetHideIncompleteCostData(val interface{})
	HideIncompleteCostDataInput() interface{}
	InternalValue() *PowerpackV2WidgetNoteDefinition
	SetInternalValue(val *PowerpackV2WidgetNoteDefinition)
	LiveSpan() *string
	SetLiveSpan(val *string)
	LiveSpanInput() *string
	ShowTick() interface{}
	SetShowTick(val interface{})
	ShowTickInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TextAlign() *string
	SetTextAlign(val *string)
	TextAlignInput() *string
	TickEdge() *string
	SetTickEdge(val *string)
	TickEdgeInput() *string
	TickPos() *string
	SetTickPos(val *string)
	TickPosInput() *string
	Time() PowerpackV2WidgetNoteDefinitionTimeOutputReference
	TimeInput() *PowerpackV2WidgetNoteDefinitionTime
	Title() *string
	SetTitle(val *string)
	TitleAlign() *string
	SetTitleAlign(val *string)
	TitleAlignInput() *string
	TitleInput() *string
	TitleSize() *string
	SetTitleSize(val *string)
	TitleSizeInput() *string
	VerticalAlign() *string
	SetVerticalAlign(val *string)
	VerticalAlignInput() *string
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
	PutTime(value *PowerpackV2WidgetNoteDefinitionTime)
	ResetBackgroundColor()
	ResetDescription()
	ResetFontSize()
	ResetHasPadding()
	ResetHideIncompleteCostData()
	ResetLiveSpan()
	ResetShowTick()
	ResetTextAlign()
	ResetTickEdge()
	ResetTickPos()
	ResetTime()
	ResetTitle()
	ResetTitleAlign()
	ResetTitleSize()
	ResetVerticalAlign()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetNoteDefinitionOutputReference
type jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) BackgroundColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) BackgroundColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) FontSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) FontSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) HasPadding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasPadding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) HasPaddingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasPaddingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) HideIncompleteCostData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIncompleteCostData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) HideIncompleteCostDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIncompleteCostDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) InternalValue() *PowerpackV2WidgetNoteDefinition {
	var returns *PowerpackV2WidgetNoteDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) LiveSpan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) LiveSpanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ShowTick() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showTick",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ShowTickInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showTickInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) TextAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) TextAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) TickEdge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tickEdge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) TickEdgeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tickEdgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) TickPos() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tickPos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) TickPosInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tickPosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) Time() PowerpackV2WidgetNoteDefinitionTimeOutputReference {
	var returns PowerpackV2WidgetNoteDefinitionTimeOutputReference
	_jsii_.Get(
		j,
		"time",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) TimeInput() *PowerpackV2WidgetNoteDefinitionTime {
	var returns *PowerpackV2WidgetNoteDefinitionTime
	_jsii_.Get(
		j,
		"timeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) TitleAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) TitleAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) TitleSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) TitleSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) VerticalAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verticalAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) VerticalAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verticalAlignInput",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetNoteDefinitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetNoteDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetNoteDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetNoteDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetNoteDefinitionOutputReference_Override(p PowerpackV2WidgetNoteDefinitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetNoteDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference)SetBackgroundColor(val *string) {
	if err := j.validateSetBackgroundColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backgroundColor",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference)SetContent(val *string) {
	if err := j.validateSetContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference)SetFontSize(val *string) {
	if err := j.validateSetFontSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontSize",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference)SetHasPadding(val interface{}) {
	if err := j.validateSetHasPaddingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasPadding",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference)SetHideIncompleteCostData(val interface{}) {
	if err := j.validateSetHideIncompleteCostDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideIncompleteCostData",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference)SetInternalValue(val *PowerpackV2WidgetNoteDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference)SetLiveSpan(val *string) {
	if err := j.validateSetLiveSpanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveSpan",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference)SetShowTick(val interface{}) {
	if err := j.validateSetShowTickParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showTick",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference)SetTextAlign(val *string) {
	if err := j.validateSetTextAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textAlign",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference)SetTickEdge(val *string) {
	if err := j.validateSetTickEdgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tickEdge",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference)SetTickPos(val *string) {
	if err := j.validateSetTickPosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tickPos",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference)SetTitleAlign(val *string) {
	if err := j.validateSetTitleAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleAlign",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference)SetTitleSize(val *string) {
	if err := j.validateSetTitleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleSize",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference)SetVerticalAlign(val *string) {
	if err := j.validateSetVerticalAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verticalAlign",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) PutTime(value *PowerpackV2WidgetNoteDefinitionTime) {
	if err := p.validatePutTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTime",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ResetBackgroundColor() {
	_jsii_.InvokeVoid(
		p,
		"resetBackgroundColor",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		p,
		"resetDescription",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ResetFontSize() {
	_jsii_.InvokeVoid(
		p,
		"resetFontSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ResetHasPadding() {
	_jsii_.InvokeVoid(
		p,
		"resetHasPadding",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ResetHideIncompleteCostData() {
	_jsii_.InvokeVoid(
		p,
		"resetHideIncompleteCostData",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ResetLiveSpan() {
	_jsii_.InvokeVoid(
		p,
		"resetLiveSpan",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ResetShowTick() {
	_jsii_.InvokeVoid(
		p,
		"resetShowTick",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ResetTextAlign() {
	_jsii_.InvokeVoid(
		p,
		"resetTextAlign",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ResetTickEdge() {
	_jsii_.InvokeVoid(
		p,
		"resetTickEdge",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ResetTickPos() {
	_jsii_.InvokeVoid(
		p,
		"resetTickPos",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ResetTime() {
	_jsii_.InvokeVoid(
		p,
		"resetTime",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		p,
		"resetTitle",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ResetTitleAlign() {
	_jsii_.InvokeVoid(
		p,
		"resetTitleAlign",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ResetTitleSize() {
	_jsii_.InvokeVoid(
		p,
		"resetTitleSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ResetVerticalAlign() {
	_jsii_.InvokeVoid(
		p,
		"resetVerticalAlign",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetNoteDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

