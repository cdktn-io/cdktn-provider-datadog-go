// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequest
	SetInternalValue(val *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequest)
	ScatterplotTable() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableList
	ScatterplotTableInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	X() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXOutputReference
	XInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestX
	Y() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYOutputReference
	YInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestY
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
	PutScatterplotTable(value interface{})
	PutX(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestX)
	PutY(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestY)
	ResetScatterplotTable()
	ResetX()
	ResetY()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) InternalValue() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequest {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequest
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) ScatterplotTable() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableList
	_jsii_.Get(
		j,
		"scatterplotTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) ScatterplotTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scatterplotTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) X() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXOutputReference
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) XInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestX {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestX
	_jsii_.Get(
		j,
		"xInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) Y() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYOutputReference
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) YInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestY {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestY
	_jsii_.Get(
		j,
		"yInput",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference)SetInternalValue(val *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequest) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) PutScatterplotTable(value interface{}) {
	if err := p.validatePutScatterplotTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putScatterplotTable",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) PutX(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestX) {
	if err := p.validatePutXParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putX",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) PutY(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestY) {
	if err := p.validatePutYParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putY",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) ResetScatterplotTable() {
	_jsii_.InvokeVoid(
		p,
		"resetScatterplotTable",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) ResetX() {
	_jsii_.InvokeVoid(
		p,
		"resetX",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) ResetY() {
	_jsii_.InvokeVoid(
		p,
		"resetY",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

