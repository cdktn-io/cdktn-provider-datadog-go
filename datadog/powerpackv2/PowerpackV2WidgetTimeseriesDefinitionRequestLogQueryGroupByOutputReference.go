// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference interface {
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
	Facet() *string
	SetFacet(val *string)
	FacetInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Limit() *float64
	SetLimit(val *float64)
	LimitInput() *float64
	SortQuery() PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupBySortQueryOutputReference
	SortQueryInput() *PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupBySortQuery
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
	PutSortQuery(value *PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupBySortQuery)
	ResetFacet()
	ResetLimit()
	ResetSortQuery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference
type jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) Facet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"facet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) FacetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"facetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) Limit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) LimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) SortQuery() PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupBySortQueryOutputReference {
	var returns PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupBySortQueryOutputReference
	_jsii_.Get(
		j,
		"sortQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) SortQueryInput() *PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupBySortQuery {
	var returns *PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupBySortQuery
	_jsii_.Get(
		j,
		"sortQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference_Override(p PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference)SetFacet(val *string) {
	if err := j.validateSetFacetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"facet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference)SetLimit(val *float64) {
	if err := j.validateSetLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limit",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) PutSortQuery(value *PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupBySortQuery) {
	if err := p.validatePutSortQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSortQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) ResetFacet() {
	_jsii_.InvokeVoid(
		p,
		"resetFacet",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) ResetLimit() {
	_jsii_.InvokeVoid(
		p,
		"resetLimit",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) ResetSortQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetSortQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetTimeseriesDefinitionRequestLogQueryGroupByOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

