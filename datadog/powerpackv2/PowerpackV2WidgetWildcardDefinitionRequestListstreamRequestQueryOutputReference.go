// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference interface {
	cdktn.ComplexObject
	ClusteringPatternFieldPath() *string
	SetClusteringPatternFieldPath(val *string)
	ClusteringPatternFieldPathInput() *string
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
	DataSource() *string
	SetDataSource(val *string)
	DataSourceInput() *string
	EventSize() *string
	SetEventSize(val *string)
	EventSizeInput() *string
	// Experimental.
	Fqn() *string
	GroupBy() PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryGroupByList
	GroupByInput() interface{}
	Indexes() *[]*string
	SetIndexes(val *[]*string)
	IndexesInput() *[]*string
	InternalValue() *PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQuery
	SetInternalValue(val *PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQuery)
	QueryString() *string
	SetQueryString(val *string)
	QueryStringInput() *string
	Sort() PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQuerySortOutputReference
	SortInput() *PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQuerySort
	Storage() *string
	SetStorage(val *string)
	StorageInput() *string
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
	PutGroupBy(value interface{})
	PutSort(value *PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQuerySort)
	ResetClusteringPatternFieldPath()
	ResetEventSize()
	ResetGroupBy()
	ResetIndexes()
	ResetQueryString()
	ResetSort()
	ResetStorage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference
type jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) ClusteringPatternFieldPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusteringPatternFieldPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) ClusteringPatternFieldPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusteringPatternFieldPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) EventSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) EventSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) GroupBy() PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryGroupByList {
	var returns PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryGroupByList
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) GroupByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) Indexes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"indexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) IndexesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"indexesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) InternalValue() *PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQuery {
	var returns *PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) QueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) QueryStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) Sort() PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQuerySortOutputReference {
	var returns PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQuerySortOutputReference
	_jsii_.Get(
		j,
		"sort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) SortInput() *PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQuerySort {
	var returns *PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQuerySort
	_jsii_.Get(
		j,
		"sortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) Storage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) StorageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference_Override(p PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference)SetClusteringPatternFieldPath(val *string) {
	if err := j.validateSetClusteringPatternFieldPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusteringPatternFieldPath",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference)SetEventSize(val *string) {
	if err := j.validateSetEventSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventSize",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference)SetIndexes(val *[]*string) {
	if err := j.validateSetIndexesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexes",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference)SetInternalValue(val *PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference)SetQueryString(val *string) {
	if err := j.validateSetQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference)SetStorage(val *string) {
	if err := j.validateSetStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storage",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) PutGroupBy(value interface{}) {
	if err := p.validatePutGroupByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGroupBy",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) PutSort(value *PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQuerySort) {
	if err := p.validatePutSortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSort",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) ResetClusteringPatternFieldPath() {
	_jsii_.InvokeVoid(
		p,
		"resetClusteringPatternFieldPath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) ResetEventSize() {
	_jsii_.InvokeVoid(
		p,
		"resetEventSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) ResetGroupBy() {
	_jsii_.InvokeVoid(
		p,
		"resetGroupBy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) ResetIndexes() {
	_jsii_.InvokeVoid(
		p,
		"resetIndexes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		p,
		"resetQueryString",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) ResetSort() {
	_jsii_.InvokeVoid(
		p,
		"resetSort",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) ResetStorage() {
	_jsii_.InvokeVoid(
		p,
		"resetStorage",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

