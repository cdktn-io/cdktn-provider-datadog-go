// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference interface {
	cdktn.ComplexObject
	CohortCriteria() PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchCohortCriteriaOutputReference
	CohortCriteriaInput() *PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchCohortCriteria
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
	Filters() PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchFiltersOutputReference
	FiltersInput() *PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchFilters
	// Experimental.
	Fqn() *string
	InternalValue() *PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearch
	SetInternalValue(val *PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearch)
	RetentionEntity() *string
	SetRetentionEntity(val *string)
	RetentionEntityInput() *string
	ReturnCondition() *string
	SetReturnCondition(val *string)
	ReturnConditionInput() *string
	ReturnCriteria() PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference
	ReturnCriteriaInput() *PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteria
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
	PutCohortCriteria(value *PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchCohortCriteria)
	PutFilters(value *PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchFilters)
	PutReturnCriteria(value *PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteria)
	ResetFilters()
	ResetReturnCriteria()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) CohortCriteria() PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchCohortCriteriaOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchCohortCriteriaOutputReference
	_jsii_.Get(
		j,
		"cohortCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) CohortCriteriaInput() *PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchCohortCriteria {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchCohortCriteria
	_jsii_.Get(
		j,
		"cohortCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) Filters() PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchFiltersOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchFiltersOutputReference
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) FiltersInput() *PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchFilters {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchFilters
	_jsii_.Get(
		j,
		"filtersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) InternalValue() *PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearch {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) RetentionEntity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionEntity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) RetentionEntityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionEntityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) ReturnCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) ReturnConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) ReturnCriteria() PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference
	_jsii_.Get(
		j,
		"returnCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) ReturnCriteriaInput() *PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteria {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteria
	_jsii_.Get(
		j,
		"returnCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference)SetInternalValue(val *PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference)SetRetentionEntity(val *string) {
	if err := j.validateSetRetentionEntityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionEntity",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference)SetReturnCondition(val *string) {
	if err := j.validateSetReturnConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnCondition",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) PutCohortCriteria(value *PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchCohortCriteria) {
	if err := p.validatePutCohortCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCohortCriteria",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) PutFilters(value *PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchFilters) {
	if err := p.validatePutFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFilters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) PutReturnCriteria(value *PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteria) {
	if err := p.validatePutReturnCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putReturnCriteria",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) ResetFilters() {
	_jsii_.InvokeVoid(
		p,
		"resetFilters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) ResetReturnCriteria() {
	_jsii_.InvokeVoid(
		p,
		"resetReturnCriteria",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetRetentionCurveDefinitionRequestQuerySearchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

