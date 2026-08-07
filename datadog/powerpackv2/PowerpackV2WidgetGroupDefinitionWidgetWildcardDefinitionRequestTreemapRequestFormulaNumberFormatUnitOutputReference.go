// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference interface {
	cdktn.ComplexObject
	Canonical() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitCanonicalOutputReference
	CanonicalInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitCanonical
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
	Custom() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitCustomOutputReference
	CustomInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitCustom
	// Experimental.
	Fqn() *string
	InternalValue() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnit
	SetInternalValue(val *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnit)
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
	PutCanonical(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitCanonical)
	PutCustom(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitCustom)
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

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) Canonical() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitCanonicalOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitCanonicalOutputReference
	_jsii_.Get(
		j,
		"canonical",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) CanonicalInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitCanonical {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitCanonical
	_jsii_.Get(
		j,
		"canonicalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) Custom() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitCustomOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitCustomOutputReference
	_jsii_.Get(
		j,
		"custom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) CustomInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitCustom {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitCustom
	_jsii_.Get(
		j,
		"customInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) InternalValue() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnit {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnit
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference)SetInternalValue(val *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnit) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) PutCanonical(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitCanonical) {
	if err := p.validatePutCanonicalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCanonical",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) PutCustom(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitCustom) {
	if err := p.validatePutCustomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCustom",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) ResetCanonical() {
	_jsii_.InvokeVoid(
		p,
		"resetCanonical",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) ResetCustom() {
	_jsii_.InvokeVoid(
		p,
		"resetCustom",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestFormulaNumberFormatUnitOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

