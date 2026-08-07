// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpack/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference interface {
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
	Fields() *[]*string
	SetFields(val *[]*string)
	FieldsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFields
	SetInternalValue(val *PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFields)
	Limit() *float64
	SetLimit(val *float64)
	LimitInput() *float64
	Sort() PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference
	SortInput() *PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSort
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
	PutSort(value *PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSort)
	ResetLimit()
	ResetSort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference
type jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) Fields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) FieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) InternalValue() *PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFields {
	var returns *PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFields
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) Limit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) LimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) Sort() PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference {
	var returns PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference
	_jsii_.Get(
		j,
		"sort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) SortInput() *PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSort {
	var returns *PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSort
	_jsii_.Get(
		j,
		"sortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpack.PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference_Override(p PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpack.PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference)SetFields(val *[]*string) {
	if err := j.validateSetFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fields",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference)SetInternalValue(val *PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFields) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference)SetLimit(val *float64) {
	if err := j.validateSetLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limit",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) PutSort(value *PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSort) {
	if err := p.validatePutSortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSort",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) ResetLimit() {
	_jsii_.InvokeVoid(
		p,
		"resetLimit",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) ResetSort() {
	_jsii_.InvokeVoid(
		p,
		"resetSort",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

