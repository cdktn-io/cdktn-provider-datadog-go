// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference interface {
	cdktn.ComplexObject
	Canonical() PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference
	CanonicalInput() *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitCanonical
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
	Custom() PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitCustomOutputReference
	CustomInput() *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitCustom
	// Experimental.
	Fqn() *string
	InternalValue() *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnit
	SetInternalValue(val *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnit)
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
	PutCanonical(value *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitCanonical)
	PutCustom(value *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitCustom)
	ResetCanonical()
	ResetCustom()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference
type jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) Canonical() PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference {
	var returns PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference
	_jsii_.Get(
		j,
		"canonical",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) CanonicalInput() *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitCanonical {
	var returns *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitCanonical
	_jsii_.Get(
		j,
		"canonicalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) Custom() PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitCustomOutputReference {
	var returns PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitCustomOutputReference
	_jsii_.Get(
		j,
		"custom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) CustomInput() *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitCustom {
	var returns *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitCustom
	_jsii_.Get(
		j,
		"customInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) InternalValue() *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnit {
	var returns *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnit
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference_Override(p PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference)SetInternalValue(val *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnit) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) PutCanonical(value *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitCanonical) {
	if err := p.validatePutCanonicalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCanonical",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) PutCustom(value *PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitCustom) {
	if err := p.validatePutCustomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCustom",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) ResetCanonical() {
	_jsii_.InvokeVoid(
		p,
		"resetCanonical",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) ResetCustom() {
	_jsii_.InvokeVoid(
		p,
		"resetCustom",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestFormulaNumberFormatUnitOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

