// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/monitor/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference interface {
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
	Compute() MonitorVariablesAggregateFilteredQueryBaseEventQueryComputeList
	ComputeInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataSource() *string
	SetDataSource(val *string)
	DataSourceInput() *string
	// Experimental.
	Fqn() *string
	GroupBy() MonitorVariablesAggregateFilteredQueryBaseEventQueryGroupByList
	GroupByInput() interface{}
	Indexes() *[]*string
	SetIndexes(val *[]*string)
	IndexesInput() *[]*string
	InternalValue() *MonitorVariablesAggregateFilteredQueryBaseEventQuery
	SetInternalValue(val *MonitorVariablesAggregateFilteredQueryBaseEventQuery)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Search() MonitorVariablesAggregateFilteredQueryBaseEventQuerySearchOutputReference
	SearchInput() *MonitorVariablesAggregateFilteredQueryBaseEventQuerySearch
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
	PutCompute(value interface{})
	PutGroupBy(value interface{})
	PutSearch(value *MonitorVariablesAggregateFilteredQueryBaseEventQuerySearch)
	ResetGroupBy()
	ResetIndexes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference
type jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) Compute() MonitorVariablesAggregateFilteredQueryBaseEventQueryComputeList {
	var returns MonitorVariablesAggregateFilteredQueryBaseEventQueryComputeList
	_jsii_.Get(
		j,
		"compute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) ComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) GroupBy() MonitorVariablesAggregateFilteredQueryBaseEventQueryGroupByList {
	var returns MonitorVariablesAggregateFilteredQueryBaseEventQueryGroupByList
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) GroupByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) Indexes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"indexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) IndexesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"indexesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) InternalValue() *MonitorVariablesAggregateFilteredQueryBaseEventQuery {
	var returns *MonitorVariablesAggregateFilteredQueryBaseEventQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) Search() MonitorVariablesAggregateFilteredQueryBaseEventQuerySearchOutputReference {
	var returns MonitorVariablesAggregateFilteredQueryBaseEventQuerySearchOutputReference
	_jsii_.Get(
		j,
		"search",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) SearchInput() *MonitorVariablesAggregateFilteredQueryBaseEventQuerySearch {
	var returns *MonitorVariablesAggregateFilteredQueryBaseEventQuerySearch
	_jsii_.Get(
		j,
		"searchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.monitor.MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference_Override(m MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.monitor.MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference)SetIndexes(val *[]*string) {
	if err := j.validateSetIndexesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexes",
		val,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference)SetInternalValue(val *MonitorVariablesAggregateFilteredQueryBaseEventQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) PutCompute(value interface{}) {
	if err := m.validatePutComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCompute",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) PutGroupBy(value interface{}) {
	if err := m.validatePutGroupByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putGroupBy",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) PutSearch(value *MonitorVariablesAggregateFilteredQueryBaseEventQuerySearch) {
	if err := m.validatePutSearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSearch",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) ResetGroupBy() {
	_jsii_.InvokeVoid(
		m,
		"resetGroupBy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) ResetIndexes() {
	_jsii_.InvokeVoid(
		m,
		"resetIndexes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

