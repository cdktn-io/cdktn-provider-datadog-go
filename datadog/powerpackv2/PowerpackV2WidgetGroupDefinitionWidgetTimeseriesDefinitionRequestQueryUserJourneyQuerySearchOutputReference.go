// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference interface {
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
	Filters() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference
	FiltersInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchFilters
	// Experimental.
	Fqn() *string
	InternalValue() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearch
	SetInternalValue(val *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearch)
	JoinKeys() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchJoinKeysOutputReference
	JoinKeysInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchJoinKeys
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
	PutFilters(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchFilters)
	PutJoinKeys(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchJoinKeys)
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

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) Expression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) Filters() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) FiltersInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchFilters {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchFilters
	_jsii_.Get(
		j,
		"filtersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) InternalValue() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearch {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) JoinKeys() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchJoinKeysOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchJoinKeysOutputReference
	_jsii_.Get(
		j,
		"joinKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) JoinKeysInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchJoinKeys {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchJoinKeys
	_jsii_.Get(
		j,
		"joinKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) NodeObjects() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) NodeObjectsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) StepAliases() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stepAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) StepAliasesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stepAliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetExpression(val *string) {
	if err := j.validateSetExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expression",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetInternalValue(val *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetNodeObjects(val *string) {
	if err := j.validateSetNodeObjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeObjects",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetStepAliases(val *string) {
	if err := j.validateSetStepAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stepAliases",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) PutFilters(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchFilters) {
	if err := p.validatePutFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFilters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) PutJoinKeys(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchJoinKeys) {
	if err := p.validatePutJoinKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putJoinKeys",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ResetFilters() {
	_jsii_.InvokeVoid(
		p,
		"resetFilters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ResetJoinKeys() {
	_jsii_.InvokeVoid(
		p,
		"resetJoinKeys",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ResetStepAliases() {
	_jsii_.InvokeVoid(
		p,
		"resetStepAliases",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

