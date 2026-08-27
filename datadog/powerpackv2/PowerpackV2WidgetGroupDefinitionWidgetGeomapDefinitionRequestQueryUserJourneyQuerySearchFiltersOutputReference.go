// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference interface {
	cdktn.ComplexObject
	AudienceFilters() PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference
	AudienceFiltersInput() *PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters
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
	GraphFilter() PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterList
	GraphFilterInput() interface{}
	InternalValue() *PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFilters
	SetInternalValue(val *PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFilters)
	StringFilter() *string
	SetStringFilter(val *string)
	StringFilterInput() *string
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
	PutAudienceFilters(value *PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters)
	PutGraphFilter(value interface{})
	ResetAudienceFilters()
	ResetGraphFilter()
	ResetStringFilter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) AudienceFilters() PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference
	_jsii_.Get(
		j,
		"audienceFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) AudienceFiltersInput() *PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters
	_jsii_.Get(
		j,
		"audienceFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) GraphFilter() PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterList
	_jsii_.Get(
		j,
		"graphFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) GraphFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"graphFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) InternalValue() *PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFilters {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFilters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) StringFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) StringFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference)SetInternalValue(val *PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFilters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference)SetStringFilter(val *string) {
	if err := j.validateSetStringFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringFilter",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) PutAudienceFilters(value *PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters) {
	if err := p.validatePutAudienceFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAudienceFilters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) PutGraphFilter(value interface{}) {
	if err := p.validatePutGraphFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGraphFilter",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) ResetAudienceFilters() {
	_jsii_.InvokeVoid(
		p,
		"resetAudienceFilters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) ResetGraphFilter() {
	_jsii_.InvokeVoid(
		p,
		"resetGraphFilter",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) ResetStringFilter() {
	_jsii_.InvokeVoid(
		p,
		"resetStringFilter",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

