// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetBarChartDefinitionStyleOutputReference interface {
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
	Display() PowerpackV2WidgetBarChartDefinitionStyleDisplayOutputReference
	DisplayInput() *PowerpackV2WidgetBarChartDefinitionStyleDisplay
	// Experimental.
	Fqn() *string
	InternalValue() *PowerpackV2WidgetBarChartDefinitionStyle
	SetInternalValue(val *PowerpackV2WidgetBarChartDefinitionStyle)
	Palette() *string
	SetPalette(val *string)
	PaletteInput() *string
	Scaling() *string
	SetScaling(val *string)
	ScalingInput() *string
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
	PutDisplay(value *PowerpackV2WidgetBarChartDefinitionStyleDisplay)
	ResetDisplay()
	ResetPalette()
	ResetScaling()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetBarChartDefinitionStyleOutputReference
type jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) Display() PowerpackV2WidgetBarChartDefinitionStyleDisplayOutputReference {
	var returns PowerpackV2WidgetBarChartDefinitionStyleDisplayOutputReference
	_jsii_.Get(
		j,
		"display",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) DisplayInput() *PowerpackV2WidgetBarChartDefinitionStyleDisplay {
	var returns *PowerpackV2WidgetBarChartDefinitionStyleDisplay
	_jsii_.Get(
		j,
		"displayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) InternalValue() *PowerpackV2WidgetBarChartDefinitionStyle {
	var returns *PowerpackV2WidgetBarChartDefinitionStyle
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) Palette() *string {
	var returns *string
	_jsii_.Get(
		j,
		"palette",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) PaletteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paletteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) Scaling() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) ScalingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetBarChartDefinitionStyleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetBarChartDefinitionStyleOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetBarChartDefinitionStyleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetBarChartDefinitionStyleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetBarChartDefinitionStyleOutputReference_Override(p PowerpackV2WidgetBarChartDefinitionStyleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetBarChartDefinitionStyleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference)SetInternalValue(val *PowerpackV2WidgetBarChartDefinitionStyle) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference)SetPalette(val *string) {
	if err := j.validateSetPaletteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"palette",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference)SetScaling(val *string) {
	if err := j.validateSetScalingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaling",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) PutDisplay(value *PowerpackV2WidgetBarChartDefinitionStyleDisplay) {
	if err := p.validatePutDisplayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDisplay",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) ResetDisplay() {
	_jsii_.InvokeVoid(
		p,
		"resetDisplay",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) ResetPalette() {
	_jsii_.InvokeVoid(
		p,
		"resetPalette",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) ResetScaling() {
	_jsii_.InvokeVoid(
		p,
		"resetScaling",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionStyleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

