// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetManageStatusDefinitionOutputReference interface {
	cdktn.ComplexObject
	ColorPreference() *string
	SetColorPreference(val *string)
	ColorPreferenceInput() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayFormat() *string
	SetDisplayFormat(val *string)
	DisplayFormatInput() *string
	// Experimental.
	Fqn() *string
	HideIncompleteCostData() interface{}
	SetHideIncompleteCostData(val interface{})
	HideIncompleteCostDataInput() interface{}
	HideZeroCounts() interface{}
	SetHideZeroCounts(val interface{})
	HideZeroCountsInput() interface{}
	InternalValue() *PowerpackV2WidgetManageStatusDefinition
	SetInternalValue(val *PowerpackV2WidgetManageStatusDefinition)
	LiveSpan() *string
	SetLiveSpan(val *string)
	LiveSpanInput() *string
	Query() *string
	SetQuery(val *string)
	QueryInput() *string
	ShowLastTriggered() interface{}
	SetShowLastTriggered(val interface{})
	ShowLastTriggeredInput() interface{}
	ShowPriority() interface{}
	SetShowPriority(val interface{})
	ShowPriorityInput() interface{}
	Sort() *string
	SetSort(val *string)
	SortInput() *string
	SummaryType() *string
	SetSummaryType(val *string)
	SummaryTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Time() PowerpackV2WidgetManageStatusDefinitionTimeOutputReference
	TimeInput() *PowerpackV2WidgetManageStatusDefinitionTime
	Title() *string
	SetTitle(val *string)
	TitleAlign() *string
	SetTitleAlign(val *string)
	TitleAlignInput() *string
	TitleInput() *string
	TitleSize() *string
	SetTitleSize(val *string)
	TitleSizeInput() *string
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
	PutTime(value *PowerpackV2WidgetManageStatusDefinitionTime)
	ResetColorPreference()
	ResetDescription()
	ResetDisplayFormat()
	ResetHideIncompleteCostData()
	ResetHideZeroCounts()
	ResetLiveSpan()
	ResetShowLastTriggered()
	ResetShowPriority()
	ResetSort()
	ResetSummaryType()
	ResetTime()
	ResetTitle()
	ResetTitleAlign()
	ResetTitleSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetManageStatusDefinitionOutputReference
type jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ColorPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ColorPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) DisplayFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) DisplayFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) HideIncompleteCostData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIncompleteCostData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) HideIncompleteCostDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIncompleteCostDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) HideZeroCounts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideZeroCounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) HideZeroCountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideZeroCountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) InternalValue() *PowerpackV2WidgetManageStatusDefinition {
	var returns *PowerpackV2WidgetManageStatusDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) LiveSpan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) LiveSpanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ShowLastTriggered() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLastTriggered",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ShowLastTriggeredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLastTriggeredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ShowPriority() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ShowPriorityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showPriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) Sort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) SortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) SummaryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summaryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) SummaryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summaryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) Time() PowerpackV2WidgetManageStatusDefinitionTimeOutputReference {
	var returns PowerpackV2WidgetManageStatusDefinitionTimeOutputReference
	_jsii_.Get(
		j,
		"time",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) TimeInput() *PowerpackV2WidgetManageStatusDefinitionTime {
	var returns *PowerpackV2WidgetManageStatusDefinitionTime
	_jsii_.Get(
		j,
		"timeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) TitleAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) TitleAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) TitleSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) TitleSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSizeInput",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetManageStatusDefinitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetManageStatusDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetManageStatusDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetManageStatusDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetManageStatusDefinitionOutputReference_Override(p PowerpackV2WidgetManageStatusDefinitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetManageStatusDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference)SetColorPreference(val *string) {
	if err := j.validateSetColorPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"colorPreference",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference)SetDisplayFormat(val *string) {
	if err := j.validateSetDisplayFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayFormat",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference)SetHideIncompleteCostData(val interface{}) {
	if err := j.validateSetHideIncompleteCostDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideIncompleteCostData",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference)SetHideZeroCounts(val interface{}) {
	if err := j.validateSetHideZeroCountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideZeroCounts",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference)SetInternalValue(val *PowerpackV2WidgetManageStatusDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference)SetLiveSpan(val *string) {
	if err := j.validateSetLiveSpanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveSpan",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference)SetShowLastTriggered(val interface{}) {
	if err := j.validateSetShowLastTriggeredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showLastTriggered",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference)SetShowPriority(val interface{}) {
	if err := j.validateSetShowPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showPriority",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference)SetSort(val *string) {
	if err := j.validateSetSortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sort",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference)SetSummaryType(val *string) {
	if err := j.validateSetSummaryTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"summaryType",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference)SetTitleAlign(val *string) {
	if err := j.validateSetTitleAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleAlign",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference)SetTitleSize(val *string) {
	if err := j.validateSetTitleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleSize",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) PutTime(value *PowerpackV2WidgetManageStatusDefinitionTime) {
	if err := p.validatePutTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTime",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ResetColorPreference() {
	_jsii_.InvokeVoid(
		p,
		"resetColorPreference",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		p,
		"resetDescription",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ResetDisplayFormat() {
	_jsii_.InvokeVoid(
		p,
		"resetDisplayFormat",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ResetHideIncompleteCostData() {
	_jsii_.InvokeVoid(
		p,
		"resetHideIncompleteCostData",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ResetHideZeroCounts() {
	_jsii_.InvokeVoid(
		p,
		"resetHideZeroCounts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ResetLiveSpan() {
	_jsii_.InvokeVoid(
		p,
		"resetLiveSpan",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ResetShowLastTriggered() {
	_jsii_.InvokeVoid(
		p,
		"resetShowLastTriggered",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ResetShowPriority() {
	_jsii_.InvokeVoid(
		p,
		"resetShowPriority",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ResetSort() {
	_jsii_.InvokeVoid(
		p,
		"resetSort",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ResetSummaryType() {
	_jsii_.InvokeVoid(
		p,
		"resetSummaryType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ResetTime() {
	_jsii_.InvokeVoid(
		p,
		"resetTime",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		p,
		"resetTitle",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ResetTitleAlign() {
	_jsii_.InvokeVoid(
		p,
		"resetTitleAlign",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ResetTitleSize() {
	_jsii_.InvokeVoid(
		p,
		"resetTitleSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetManageStatusDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

