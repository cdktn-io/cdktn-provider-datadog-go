// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference interface {
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
	ComputeQuery() PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryComputeQueryOutputReference
	ComputeQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryComputeQuery
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GroupBy() PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryGroupByList
	GroupByInput() interface{}
	Index() *string
	SetIndex(val *string)
	IndexInput() *string
	InternalValue() *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQuery
	SetInternalValue(val *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQuery)
	MultiCompute() PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryMultiComputeList
	MultiComputeInput() interface{}
	SearchQuery() *string
	SetSearchQuery(val *string)
	SearchQueryInput() *string
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
	PutComputeQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryComputeQuery)
	PutGroupBy(value interface{})
	PutMultiCompute(value interface{})
	ResetComputeQuery()
	ResetGroupBy()
	ResetMultiCompute()
	ResetSearchQuery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference
type jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) ComputeQuery() PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryComputeQueryOutputReference {
	var returns PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryComputeQueryOutputReference
	_jsii_.Get(
		j,
		"computeQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) ComputeQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryComputeQuery {
	var returns *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryComputeQuery
	_jsii_.Get(
		j,
		"computeQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) GroupBy() PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryGroupByList {
	var returns PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryGroupByList
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) GroupByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) Index() *string {
	var returns *string
	_jsii_.Get(
		j,
		"index",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) IndexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) InternalValue() *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQuery {
	var returns *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) MultiCompute() PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryMultiComputeList {
	var returns PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryMultiComputeList
	_jsii_.Get(
		j,
		"multiCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) MultiComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiComputeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) SearchQuery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) SearchQueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference_Override(p PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference)SetIndex(val *string) {
	if err := j.validateSetIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"index",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference)SetInternalValue(val *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference)SetSearchQuery(val *string) {
	if err := j.validateSetSearchQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchQuery",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) PutComputeQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryComputeQuery) {
	if err := p.validatePutComputeQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putComputeQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) PutGroupBy(value interface{}) {
	if err := p.validatePutGroupByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGroupBy",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) PutMultiCompute(value interface{}) {
	if err := p.validatePutMultiComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMultiCompute",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) ResetComputeQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetComputeQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) ResetGroupBy() {
	_jsii_.InvokeVoid(
		p,
		"resetGroupBy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) ResetMultiCompute() {
	_jsii_.InvokeVoid(
		p,
		"resetMultiCompute",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) ResetSearchQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetSearchQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestApmQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

