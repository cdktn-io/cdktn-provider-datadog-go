// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference interface {
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
	Expression() *string
	SetExpression(val *string)
	ExpressionInput() *string
	Filters() PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference
	FiltersInput() *PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchFilters
	// Experimental.
	Fqn() *string
	InternalValue() *PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearch
	SetInternalValue(val *PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearch)
	JoinKeys() PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchJoinKeysOutputReference
	JoinKeysInput() *PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchJoinKeys
	NodeObjects() *string
	SetNodeObjects(val *string)
	NodeObjectsInput() *string
	StepAliases() *string
	SetStepAliases(val *string)
	StepAliasesInput() *string
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
	PutFilters(value *PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchFilters)
	PutJoinKeys(value *PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchJoinKeys)
	ResetFilters()
	ResetJoinKeys()
	ResetStepAliases()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference
type jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) Expression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) Filters() PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference {
	var returns PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) FiltersInput() *PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchFilters {
	var returns *PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchFilters
	_jsii_.Get(
		j,
		"filtersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) InternalValue() *PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearch {
	var returns *PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) JoinKeys() PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchJoinKeysOutputReference {
	var returns PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchJoinKeysOutputReference
	_jsii_.Get(
		j,
		"joinKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) JoinKeysInput() *PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchJoinKeys {
	var returns *PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchJoinKeys
	_jsii_.Get(
		j,
		"joinKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) NodeObjects() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) NodeObjectsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) StepAliases() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stepAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) StepAliasesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stepAliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference_Override(p PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetExpression(val *string) {
	if err := j.validateSetExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expression",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetInternalValue(val *PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetNodeObjects(val *string) {
	if err := j.validateSetNodeObjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeObjects",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetStepAliases(val *string) {
	if err := j.validateSetStepAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stepAliases",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) PutFilters(value *PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchFilters) {
	if err := p.validatePutFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFilters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) PutJoinKeys(value *PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchJoinKeys) {
	if err := p.validatePutJoinKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putJoinKeys",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ResetFilters() {
	_jsii_.InvokeVoid(
		p,
		"resetFilters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ResetJoinKeys() {
	_jsii_.InvokeVoid(
		p,
		"resetJoinKeys",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ResetStepAliases() {
	_jsii_.InvokeVoid(
		p,
		"resetStepAliases",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

