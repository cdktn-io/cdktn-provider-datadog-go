// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference interface {
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
	InternalValue() *PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonical
	SetInternalValue(val *PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonical)
	PerUnitName() *string
	SetPerUnitName(val *string)
	PerUnitNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UnitName() *string
	SetUnitName(val *string)
	UnitNameInput() *string
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
	ResetPerUnitName()
	ResetUnitName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference
type jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) InternalValue() *PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonical {
	var returns *PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonical
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) PerUnitName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"perUnitName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) PerUnitNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"perUnitNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) UnitName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) UnitNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitNameInput",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference_Override(p PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference)SetInternalValue(val *PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonical) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference)SetPerUnitName(val *string) {
	if err := j.validateSetPerUnitNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"perUnitName",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference)SetUnitName(val *string) {
	if err := j.validateSetUnitNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unitName",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) ResetPerUnitName() {
	_jsii_.InvokeVoid(
		p,
		"resetPerUnitName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) ResetUnitName() {
	_jsii_.InvokeVoid(
		p,
		"resetUnitName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXFormulaNumberFormatUnitCanonicalOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

