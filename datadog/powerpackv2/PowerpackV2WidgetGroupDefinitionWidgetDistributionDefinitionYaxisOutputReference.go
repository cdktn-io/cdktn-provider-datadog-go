// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference interface {
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
	IncludeZero() interface{}
	SetIncludeZero(val interface{})
	IncludeZeroInput() interface{}
	InternalValue() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxis
	SetInternalValue(val *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxis)
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	Max() *string
	SetMax(val *string)
	MaxInput() *string
	Min() *string
	SetMin(val *string)
	MinInput() *string
	Scale() *string
	SetScale(val *string)
	ScaleInput() *string
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
	ResetIncludeZero()
	ResetLabel()
	ResetMax()
	ResetMin()
	ResetScale()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) IncludeZero() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeZero",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) IncludeZeroInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeZeroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) InternalValue() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxis {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxis
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) Max() *string {
	var returns *string
	_jsii_.Get(
		j,
		"max",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) MaxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) Min() *string {
	var returns *string
	_jsii_.Get(
		j,
		"min",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) MinInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) Scale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) ScaleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference)SetIncludeZero(val interface{}) {
	if err := j.validateSetIncludeZeroParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeZero",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference)SetInternalValue(val *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxis) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference)SetMax(val *string) {
	if err := j.validateSetMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"max",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference)SetMin(val *string) {
	if err := j.validateSetMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"min",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference)SetScale(val *string) {
	if err := j.validateSetScaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scale",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) ResetIncludeZero() {
	_jsii_.InvokeVoid(
		p,
		"resetIncludeZero",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) ResetLabel() {
	_jsii_.InvokeVoid(
		p,
		"resetLabel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) ResetMax() {
	_jsii_.InvokeVoid(
		p,
		"resetMax",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) ResetMin() {
	_jsii_.InvokeVoid(
		p,
		"resetMin",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) ResetScale() {
	_jsii_.InvokeVoid(
		p,
		"resetScale",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionYaxisOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

