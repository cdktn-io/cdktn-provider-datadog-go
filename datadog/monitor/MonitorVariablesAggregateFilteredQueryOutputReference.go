// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/monitor/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MonitorVariablesAggregateFilteredQueryOutputReference interface {
	cdktn.ComplexObject
	BaseEventQuery() MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference
	BaseEventQueryInput() *MonitorVariablesAggregateFilteredQueryBaseEventQuery
	BaseMetricsQuery() MonitorVariablesAggregateFilteredQueryBaseMetricsQueryOutputReference
	BaseMetricsQueryInput() *MonitorVariablesAggregateFilteredQueryBaseMetricsQuery
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
	Compute() MonitorVariablesAggregateFilteredQueryComputeList
	ComputeInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataSource() *string
	SetDataSource(val *string)
	DataSourceInput() *string
	FilterEventQuery() MonitorVariablesAggregateFilteredQueryFilterEventQueryOutputReference
	FilterEventQueryInput() *MonitorVariablesAggregateFilteredQueryFilterEventQuery
	FilterReferenceTable() MonitorVariablesAggregateFilteredQueryFilterReferenceTableOutputReference
	FilterReferenceTableInput() *MonitorVariablesAggregateFilteredQueryFilterReferenceTable
	Filters() MonitorVariablesAggregateFilteredQueryFiltersList
	FiltersInput() interface{}
	// Experimental.
	Fqn() *string
	GroupBy() MonitorVariablesAggregateFilteredQueryGroupByList
	GroupByInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	PutBaseEventQuery(value *MonitorVariablesAggregateFilteredQueryBaseEventQuery)
	PutBaseMetricsQuery(value *MonitorVariablesAggregateFilteredQueryBaseMetricsQuery)
	PutCompute(value interface{})
	PutFilterEventQuery(value *MonitorVariablesAggregateFilteredQueryFilterEventQuery)
	PutFilterReferenceTable(value *MonitorVariablesAggregateFilteredQueryFilterReferenceTable)
	PutFilters(value interface{})
	PutGroupBy(value interface{})
	ResetBaseEventQuery()
	ResetBaseMetricsQuery()
	ResetCompute()
	ResetFilterEventQuery()
	ResetFilterReferenceTable()
	ResetGroupBy()
	ResetName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorVariablesAggregateFilteredQueryOutputReference
type jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) BaseEventQuery() MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference {
	var returns MonitorVariablesAggregateFilteredQueryBaseEventQueryOutputReference
	_jsii_.Get(
		j,
		"baseEventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) BaseEventQueryInput() *MonitorVariablesAggregateFilteredQueryBaseEventQuery {
	var returns *MonitorVariablesAggregateFilteredQueryBaseEventQuery
	_jsii_.Get(
		j,
		"baseEventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) BaseMetricsQuery() MonitorVariablesAggregateFilteredQueryBaseMetricsQueryOutputReference {
	var returns MonitorVariablesAggregateFilteredQueryBaseMetricsQueryOutputReference
	_jsii_.Get(
		j,
		"baseMetricsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) BaseMetricsQueryInput() *MonitorVariablesAggregateFilteredQueryBaseMetricsQuery {
	var returns *MonitorVariablesAggregateFilteredQueryBaseMetricsQuery
	_jsii_.Get(
		j,
		"baseMetricsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) Compute() MonitorVariablesAggregateFilteredQueryComputeList {
	var returns MonitorVariablesAggregateFilteredQueryComputeList
	_jsii_.Get(
		j,
		"compute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) ComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) FilterEventQuery() MonitorVariablesAggregateFilteredQueryFilterEventQueryOutputReference {
	var returns MonitorVariablesAggregateFilteredQueryFilterEventQueryOutputReference
	_jsii_.Get(
		j,
		"filterEventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) FilterEventQueryInput() *MonitorVariablesAggregateFilteredQueryFilterEventQuery {
	var returns *MonitorVariablesAggregateFilteredQueryFilterEventQuery
	_jsii_.Get(
		j,
		"filterEventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) FilterReferenceTable() MonitorVariablesAggregateFilteredQueryFilterReferenceTableOutputReference {
	var returns MonitorVariablesAggregateFilteredQueryFilterReferenceTableOutputReference
	_jsii_.Get(
		j,
		"filterReferenceTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) FilterReferenceTableInput() *MonitorVariablesAggregateFilteredQueryFilterReferenceTable {
	var returns *MonitorVariablesAggregateFilteredQueryFilterReferenceTable
	_jsii_.Get(
		j,
		"filterReferenceTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) Filters() MonitorVariablesAggregateFilteredQueryFiltersList {
	var returns MonitorVariablesAggregateFilteredQueryFiltersList
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) FiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filtersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) GroupBy() MonitorVariablesAggregateFilteredQueryGroupByList {
	var returns MonitorVariablesAggregateFilteredQueryGroupByList
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) GroupByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorVariablesAggregateFilteredQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MonitorVariablesAggregateFilteredQueryOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorVariablesAggregateFilteredQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.monitor.MonitorVariablesAggregateFilteredQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMonitorVariablesAggregateFilteredQueryOutputReference_Override(m MonitorVariablesAggregateFilteredQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.monitor.MonitorVariablesAggregateFilteredQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) PutBaseEventQuery(value *MonitorVariablesAggregateFilteredQueryBaseEventQuery) {
	if err := m.validatePutBaseEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putBaseEventQuery",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) PutBaseMetricsQuery(value *MonitorVariablesAggregateFilteredQueryBaseMetricsQuery) {
	if err := m.validatePutBaseMetricsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putBaseMetricsQuery",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) PutCompute(value interface{}) {
	if err := m.validatePutComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCompute",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) PutFilterEventQuery(value *MonitorVariablesAggregateFilteredQueryFilterEventQuery) {
	if err := m.validatePutFilterEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFilterEventQuery",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) PutFilterReferenceTable(value *MonitorVariablesAggregateFilteredQueryFilterReferenceTable) {
	if err := m.validatePutFilterReferenceTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFilterReferenceTable",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) PutFilters(value interface{}) {
	if err := m.validatePutFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFilters",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) PutGroupBy(value interface{}) {
	if err := m.validatePutGroupByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putGroupBy",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) ResetBaseEventQuery() {
	_jsii_.InvokeVoid(
		m,
		"resetBaseEventQuery",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) ResetBaseMetricsQuery() {
	_jsii_.InvokeVoid(
		m,
		"resetBaseMetricsQuery",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) ResetCompute() {
	_jsii_.InvokeVoid(
		m,
		"resetCompute",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) ResetFilterEventQuery() {
	_jsii_.InvokeVoid(
		m,
		"resetFilterEventQuery",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) ResetFilterReferenceTable() {
	_jsii_.InvokeVoid(
		m,
		"resetFilterReferenceTable",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) ResetGroupBy() {
	_jsii_.InvokeVoid(
		m,
		"resetGroupBy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		m,
		"resetName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MonitorVariablesAggregateFilteredQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

