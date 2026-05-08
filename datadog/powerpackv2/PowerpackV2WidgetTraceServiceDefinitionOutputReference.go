// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetTraceServiceDefinitionOutputReference interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayFormat() *string
	SetDisplayFormat(val *string)
	DisplayFormatInput() *string
	Env() *string
	SetEnv(val *string)
	EnvInput() *string
	// Experimental.
	Fqn() *string
	HideIncompleteCostData() interface{}
	SetHideIncompleteCostData(val interface{})
	HideIncompleteCostDataInput() interface{}
	InternalValue() *PowerpackV2WidgetTraceServiceDefinition
	SetInternalValue(val *PowerpackV2WidgetTraceServiceDefinition)
	LiveSpan() *string
	SetLiveSpan(val *string)
	LiveSpanInput() *string
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	ShowBreakdown() interface{}
	SetShowBreakdown(val interface{})
	ShowBreakdownInput() interface{}
	ShowDistribution() interface{}
	SetShowDistribution(val interface{})
	ShowDistributionInput() interface{}
	ShowErrors() interface{}
	SetShowErrors(val interface{})
	ShowErrorsInput() interface{}
	ShowHits() interface{}
	SetShowHits(val interface{})
	ShowHitsInput() interface{}
	ShowLatency() interface{}
	SetShowLatency(val interface{})
	ShowLatencyInput() interface{}
	ShowResourceList() interface{}
	SetShowResourceList(val interface{})
	ShowResourceListInput() interface{}
	SizeFormat() *string
	SetSizeFormat(val *string)
	SizeFormatInput() *string
	SpanName() *string
	SetSpanName(val *string)
	SpanNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Time() PowerpackV2WidgetTraceServiceDefinitionTimeOutputReference
	TimeInput() *PowerpackV2WidgetTraceServiceDefinitionTime
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
	PutTime(value *PowerpackV2WidgetTraceServiceDefinitionTime)
	ResetDescription()
	ResetDisplayFormat()
	ResetHideIncompleteCostData()
	ResetLiveSpan()
	ResetShowBreakdown()
	ResetShowDistribution()
	ResetShowErrors()
	ResetShowHits()
	ResetShowLatency()
	ResetShowResourceList()
	ResetSizeFormat()
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

// The jsii proxy struct for PowerpackV2WidgetTraceServiceDefinitionOutputReference
type jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) DisplayFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) DisplayFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) Env() *string {
	var returns *string
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) EnvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) HideIncompleteCostData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIncompleteCostData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) HideIncompleteCostDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIncompleteCostDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) InternalValue() *PowerpackV2WidgetTraceServiceDefinition {
	var returns *PowerpackV2WidgetTraceServiceDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) LiveSpan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) LiveSpanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ShowBreakdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showBreakdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ShowBreakdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showBreakdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ShowDistribution() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showDistribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ShowDistributionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showDistributionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ShowErrors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ShowErrorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ShowHits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showHits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ShowHitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showHitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ShowLatency() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ShowLatencyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ShowResourceList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showResourceList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ShowResourceListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showResourceListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) SizeFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) SizeFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) SpanName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spanName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) SpanNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spanNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) Time() PowerpackV2WidgetTraceServiceDefinitionTimeOutputReference {
	var returns PowerpackV2WidgetTraceServiceDefinitionTimeOutputReference
	_jsii_.Get(
		j,
		"time",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) TimeInput() *PowerpackV2WidgetTraceServiceDefinitionTime {
	var returns *PowerpackV2WidgetTraceServiceDefinitionTime
	_jsii_.Get(
		j,
		"timeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) TitleAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) TitleAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) TitleSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) TitleSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSizeInput",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetTraceServiceDefinitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetTraceServiceDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetTraceServiceDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetTraceServiceDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetTraceServiceDefinitionOutputReference_Override(p PowerpackV2WidgetTraceServiceDefinitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetTraceServiceDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetDisplayFormat(val *string) {
	if err := j.validateSetDisplayFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayFormat",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetEnv(val *string) {
	if err := j.validateSetEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"env",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetHideIncompleteCostData(val interface{}) {
	if err := j.validateSetHideIncompleteCostDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideIncompleteCostData",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetInternalValue(val *PowerpackV2WidgetTraceServiceDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetLiveSpan(val *string) {
	if err := j.validateSetLiveSpanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveSpan",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetShowBreakdown(val interface{}) {
	if err := j.validateSetShowBreakdownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showBreakdown",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetShowDistribution(val interface{}) {
	if err := j.validateSetShowDistributionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showDistribution",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetShowErrors(val interface{}) {
	if err := j.validateSetShowErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showErrors",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetShowHits(val interface{}) {
	if err := j.validateSetShowHitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showHits",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetShowLatency(val interface{}) {
	if err := j.validateSetShowLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showLatency",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetShowResourceList(val interface{}) {
	if err := j.validateSetShowResourceListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showResourceList",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetSizeFormat(val *string) {
	if err := j.validateSetSizeFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeFormat",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetSpanName(val *string) {
	if err := j.validateSetSpanNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spanName",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetTitleAlign(val *string) {
	if err := j.validateSetTitleAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleAlign",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference)SetTitleSize(val *string) {
	if err := j.validateSetTitleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleSize",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) PutTime(value *PowerpackV2WidgetTraceServiceDefinitionTime) {
	if err := p.validatePutTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTime",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		p,
		"resetDescription",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ResetDisplayFormat() {
	_jsii_.InvokeVoid(
		p,
		"resetDisplayFormat",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ResetHideIncompleteCostData() {
	_jsii_.InvokeVoid(
		p,
		"resetHideIncompleteCostData",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ResetLiveSpan() {
	_jsii_.InvokeVoid(
		p,
		"resetLiveSpan",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ResetShowBreakdown() {
	_jsii_.InvokeVoid(
		p,
		"resetShowBreakdown",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ResetShowDistribution() {
	_jsii_.InvokeVoid(
		p,
		"resetShowDistribution",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ResetShowErrors() {
	_jsii_.InvokeVoid(
		p,
		"resetShowErrors",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ResetShowHits() {
	_jsii_.InvokeVoid(
		p,
		"resetShowHits",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ResetShowLatency() {
	_jsii_.InvokeVoid(
		p,
		"resetShowLatency",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ResetShowResourceList() {
	_jsii_.InvokeVoid(
		p,
		"resetShowResourceList",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ResetSizeFormat() {
	_jsii_.InvokeVoid(
		p,
		"resetSizeFormat",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ResetTime() {
	_jsii_.InvokeVoid(
		p,
		"resetTime",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		p,
		"resetTitle",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ResetTitleAlign() {
	_jsii_.InvokeVoid(
		p,
		"resetTitleAlign",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ResetTitleSize() {
	_jsii_.InvokeVoid(
		p,
		"resetTitleSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetTraceServiceDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

