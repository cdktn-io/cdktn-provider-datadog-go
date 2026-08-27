// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference interface {
	cdktn.ComplexObject
	CohortCriteria() PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaOutputReference
	CohortCriteriaInput() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteria
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
	Filters() PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchFiltersOutputReference
	FiltersInput() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchFilters
	// Experimental.
	Fqn() *string
	InternalValue() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearch
	SetInternalValue(val *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearch)
	RetentionEntity() *string
	SetRetentionEntity(val *string)
	RetentionEntityInput() *string
	ReturnCondition() *string
	SetReturnCondition(val *string)
	ReturnConditionInput() *string
	ReturnCriteria() PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchReturnCriteriaOutputReference
	ReturnCriteriaInput() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchReturnCriteria
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
	PutCohortCriteria(value *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteria)
	PutFilters(value *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchFilters)
	PutReturnCriteria(value *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchReturnCriteria)
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

// The jsii proxy struct for PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference
type jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) CohortCriteria() PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaOutputReference {
	var returns PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaOutputReference
	_jsii_.Get(
		j,
		"cohortCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) CohortCriteriaInput() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteria {
	var returns *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteria
	_jsii_.Get(
		j,
		"cohortCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) Filters() PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchFiltersOutputReference {
	var returns PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchFiltersOutputReference
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) FiltersInput() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchFilters {
	var returns *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchFilters
	_jsii_.Get(
		j,
		"filtersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) InternalValue() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearch {
	var returns *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) RetentionEntity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionEntity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) RetentionEntityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionEntityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) ReturnCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) ReturnConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) ReturnCriteria() PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchReturnCriteriaOutputReference {
	var returns PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchReturnCriteriaOutputReference
	_jsii_.Get(
		j,
		"returnCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) ReturnCriteriaInput() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchReturnCriteria {
	var returns *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchReturnCriteria
	_jsii_.Get(
		j,
		"returnCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference_Override(p PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference)SetInternalValue(val *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference)SetRetentionEntity(val *string) {
	if err := j.validateSetRetentionEntityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionEntity",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference)SetReturnCondition(val *string) {
	if err := j.validateSetReturnConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnCondition",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) PutCohortCriteria(value *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteria) {
	if err := p.validatePutCohortCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCohortCriteria",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) PutFilters(value *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchFilters) {
	if err := p.validatePutFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFilters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) PutReturnCriteria(value *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchReturnCriteria) {
	if err := p.validatePutReturnCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putReturnCriteria",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) ResetFilters() {
	_jsii_.InvokeVoid(
		p,
		"resetFilters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) ResetReturnCriteria() {
	_jsii_.InvokeVoid(
		p,
		"resetReturnCriteria",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

