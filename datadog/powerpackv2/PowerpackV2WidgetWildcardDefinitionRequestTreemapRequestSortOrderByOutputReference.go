// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference interface {
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
	FormulaSort() PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByFormulaSortOutputReference
	FormulaSortInput() *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByFormulaSort
	// Experimental.
	Fqn() *string
	GroupSort() PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByGroupSortOutputReference
	GroupSortInput() *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByGroupSort
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutFormulaSort(value *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByFormulaSort)
	PutGroupSort(value *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByGroupSort)
	ResetFormulaSort()
	ResetGroupSort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference
type jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) FormulaSort() PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByFormulaSortOutputReference {
	var returns PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByFormulaSortOutputReference
	_jsii_.Get(
		j,
		"formulaSort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) FormulaSortInput() *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByFormulaSort {
	var returns *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByFormulaSort
	_jsii_.Get(
		j,
		"formulaSortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) GroupSort() PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByGroupSortOutputReference {
	var returns PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByGroupSortOutputReference
	_jsii_.Get(
		j,
		"groupSort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) GroupSortInput() *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByGroupSort {
	var returns *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByGroupSort
	_jsii_.Get(
		j,
		"groupSortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference_Override(p PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) PutFormulaSort(value *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByFormulaSort) {
	if err := p.validatePutFormulaSortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFormulaSort",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) PutGroupSort(value *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByGroupSort) {
	if err := p.validatePutGroupSortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGroupSort",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) ResetFormulaSort() {
	_jsii_.InvokeVoid(
		p,
		"resetFormulaSort",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) ResetGroupSort() {
	_jsii_.InvokeVoid(
		p,
		"resetGroupSort",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestSortOrderByOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

