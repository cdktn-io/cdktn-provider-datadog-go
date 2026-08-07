// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference interface {
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
	InternalValue() *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormat
	SetInternalValue(val *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormat)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Unit() PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference
	UnitInput() *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnit
	UnitScale() PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitScaleOutputReference
	UnitScaleInput() *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitScale
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
	PutUnit(value *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnit)
	PutUnitScale(value *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitScale)
	ResetUnitScale()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference
type jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) InternalValue() *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormat {
	var returns *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormat
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) Unit() PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference {
	var returns PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) UnitInput() *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnit {
	var returns *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnit
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) UnitScale() PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitScaleOutputReference {
	var returns PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitScaleOutputReference
	_jsii_.Get(
		j,
		"unitScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) UnitScaleInput() *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitScale {
	var returns *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitScale
	_jsii_.Get(
		j,
		"unitScaleInput",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference_Override(p PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference)SetInternalValue(val *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormat) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) PutUnit(value *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnit) {
	if err := p.validatePutUnitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putUnit",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) PutUnitScale(value *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitScale) {
	if err := p.validatePutUnitScaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putUnitScale",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) ResetUnitScale() {
	_jsii_.InvokeVoid(
		p,
		"resetUnitScale",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

