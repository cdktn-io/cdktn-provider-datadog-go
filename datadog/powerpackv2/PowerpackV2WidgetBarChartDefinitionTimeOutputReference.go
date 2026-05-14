// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetBarChartDefinitionTimeOutputReference interface {
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
	Fixed() PowerpackV2WidgetBarChartDefinitionTimeFixedOutputReference
	FixedInput() *PowerpackV2WidgetBarChartDefinitionTimeFixed
	// Experimental.
	Fqn() *string
	InternalValue() *PowerpackV2WidgetBarChartDefinitionTime
	SetInternalValue(val *PowerpackV2WidgetBarChartDefinitionTime)
	Live() PowerpackV2WidgetBarChartDefinitionTimeLiveOutputReference
	LiveInput() *PowerpackV2WidgetBarChartDefinitionTimeLive
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutFixed(value *PowerpackV2WidgetBarChartDefinitionTimeFixed)
	PutLive(value *PowerpackV2WidgetBarChartDefinitionTimeLive)
	ResetFixed()
	ResetLive()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetBarChartDefinitionTimeOutputReference
type jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) Fixed() PowerpackV2WidgetBarChartDefinitionTimeFixedOutputReference {
	var returns PowerpackV2WidgetBarChartDefinitionTimeFixedOutputReference
	_jsii_.Get(
		j,
		"fixed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) FixedInput() *PowerpackV2WidgetBarChartDefinitionTimeFixed {
	var returns *PowerpackV2WidgetBarChartDefinitionTimeFixed
	_jsii_.Get(
		j,
		"fixedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) InternalValue() *PowerpackV2WidgetBarChartDefinitionTime {
	var returns *PowerpackV2WidgetBarChartDefinitionTime
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) Live() PowerpackV2WidgetBarChartDefinitionTimeLiveOutputReference {
	var returns PowerpackV2WidgetBarChartDefinitionTimeLiveOutputReference
	_jsii_.Get(
		j,
		"live",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) LiveInput() *PowerpackV2WidgetBarChartDefinitionTimeLive {
	var returns *PowerpackV2WidgetBarChartDefinitionTimeLive
	_jsii_.Get(
		j,
		"liveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetBarChartDefinitionTimeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetBarChartDefinitionTimeOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetBarChartDefinitionTimeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetBarChartDefinitionTimeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetBarChartDefinitionTimeOutputReference_Override(p PowerpackV2WidgetBarChartDefinitionTimeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetBarChartDefinitionTimeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference)SetInternalValue(val *PowerpackV2WidgetBarChartDefinitionTime) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) PutFixed(value *PowerpackV2WidgetBarChartDefinitionTimeFixed) {
	if err := p.validatePutFixedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFixed",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) PutLive(value *PowerpackV2WidgetBarChartDefinitionTimeLive) {
	if err := p.validatePutLiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLive",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) ResetFixed() {
	_jsii_.InvokeVoid(
		p,
		"resetFixed",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) ResetLive() {
	_jsii_.InvokeVoid(
		p,
		"resetLive",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionTimeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

