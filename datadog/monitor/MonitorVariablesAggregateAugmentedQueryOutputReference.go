// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/monitor/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MonitorVariablesAggregateAugmentedQueryOutputReference interface {
	cdktn.ComplexObject
	AugmentEventQuery() MonitorVariablesAggregateAugmentedQueryAugmentEventQueryOutputReference
	AugmentEventQueryInput() *MonitorVariablesAggregateAugmentedQueryAugmentEventQuery
	AugmentReferenceTable() MonitorVariablesAggregateAugmentedQueryAugmentReferenceTableOutputReference
	AugmentReferenceTableInput() *MonitorVariablesAggregateAugmentedQueryAugmentReferenceTable
	BaseEventQuery() MonitorVariablesAggregateAugmentedQueryBaseEventQueryOutputReference
	BaseEventQueryInput() *MonitorVariablesAggregateAugmentedQueryBaseEventQuery
	BaseMetricsQuery() MonitorVariablesAggregateAugmentedQueryBaseMetricsQueryOutputReference
	BaseMetricsQueryInput() *MonitorVariablesAggregateAugmentedQueryBaseMetricsQuery
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
	Compute() MonitorVariablesAggregateAugmentedQueryComputeList
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
	GroupBy() MonitorVariablesAggregateAugmentedQueryGroupByList
	GroupByInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JoinCondition() MonitorVariablesAggregateAugmentedQueryJoinConditionOutputReference
	JoinConditionInput() *MonitorVariablesAggregateAugmentedQueryJoinCondition
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
	PutAugmentEventQuery(value *MonitorVariablesAggregateAugmentedQueryAugmentEventQuery)
	PutAugmentReferenceTable(value *MonitorVariablesAggregateAugmentedQueryAugmentReferenceTable)
	PutBaseEventQuery(value *MonitorVariablesAggregateAugmentedQueryBaseEventQuery)
	PutBaseMetricsQuery(value *MonitorVariablesAggregateAugmentedQueryBaseMetricsQuery)
	PutCompute(value interface{})
	PutGroupBy(value interface{})
	PutJoinCondition(value *MonitorVariablesAggregateAugmentedQueryJoinCondition)
	ResetAugmentEventQuery()
	ResetAugmentReferenceTable()
	ResetBaseEventQuery()
	ResetBaseMetricsQuery()
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

// The jsii proxy struct for MonitorVariablesAggregateAugmentedQueryOutputReference
type jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) AugmentEventQuery() MonitorVariablesAggregateAugmentedQueryAugmentEventQueryOutputReference {
	var returns MonitorVariablesAggregateAugmentedQueryAugmentEventQueryOutputReference
	_jsii_.Get(
		j,
		"augmentEventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) AugmentEventQueryInput() *MonitorVariablesAggregateAugmentedQueryAugmentEventQuery {
	var returns *MonitorVariablesAggregateAugmentedQueryAugmentEventQuery
	_jsii_.Get(
		j,
		"augmentEventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) AugmentReferenceTable() MonitorVariablesAggregateAugmentedQueryAugmentReferenceTableOutputReference {
	var returns MonitorVariablesAggregateAugmentedQueryAugmentReferenceTableOutputReference
	_jsii_.Get(
		j,
		"augmentReferenceTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) AugmentReferenceTableInput() *MonitorVariablesAggregateAugmentedQueryAugmentReferenceTable {
	var returns *MonitorVariablesAggregateAugmentedQueryAugmentReferenceTable
	_jsii_.Get(
		j,
		"augmentReferenceTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) BaseEventQuery() MonitorVariablesAggregateAugmentedQueryBaseEventQueryOutputReference {
	var returns MonitorVariablesAggregateAugmentedQueryBaseEventQueryOutputReference
	_jsii_.Get(
		j,
		"baseEventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) BaseEventQueryInput() *MonitorVariablesAggregateAugmentedQueryBaseEventQuery {
	var returns *MonitorVariablesAggregateAugmentedQueryBaseEventQuery
	_jsii_.Get(
		j,
		"baseEventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) BaseMetricsQuery() MonitorVariablesAggregateAugmentedQueryBaseMetricsQueryOutputReference {
	var returns MonitorVariablesAggregateAugmentedQueryBaseMetricsQueryOutputReference
	_jsii_.Get(
		j,
		"baseMetricsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) BaseMetricsQueryInput() *MonitorVariablesAggregateAugmentedQueryBaseMetricsQuery {
	var returns *MonitorVariablesAggregateAugmentedQueryBaseMetricsQuery
	_jsii_.Get(
		j,
		"baseMetricsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) Compute() MonitorVariablesAggregateAugmentedQueryComputeList {
	var returns MonitorVariablesAggregateAugmentedQueryComputeList
	_jsii_.Get(
		j,
		"compute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) ComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) GroupBy() MonitorVariablesAggregateAugmentedQueryGroupByList {
	var returns MonitorVariablesAggregateAugmentedQueryGroupByList
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) GroupByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) JoinCondition() MonitorVariablesAggregateAugmentedQueryJoinConditionOutputReference {
	var returns MonitorVariablesAggregateAugmentedQueryJoinConditionOutputReference
	_jsii_.Get(
		j,
		"joinCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) JoinConditionInput() *MonitorVariablesAggregateAugmentedQueryJoinCondition {
	var returns *MonitorVariablesAggregateAugmentedQueryJoinCondition
	_jsii_.Get(
		j,
		"joinConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorVariablesAggregateAugmentedQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MonitorVariablesAggregateAugmentedQueryOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorVariablesAggregateAugmentedQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.monitor.MonitorVariablesAggregateAugmentedQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMonitorVariablesAggregateAugmentedQueryOutputReference_Override(m MonitorVariablesAggregateAugmentedQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.monitor.MonitorVariablesAggregateAugmentedQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) PutAugmentEventQuery(value *MonitorVariablesAggregateAugmentedQueryAugmentEventQuery) {
	if err := m.validatePutAugmentEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAugmentEventQuery",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) PutAugmentReferenceTable(value *MonitorVariablesAggregateAugmentedQueryAugmentReferenceTable) {
	if err := m.validatePutAugmentReferenceTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAugmentReferenceTable",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) PutBaseEventQuery(value *MonitorVariablesAggregateAugmentedQueryBaseEventQuery) {
	if err := m.validatePutBaseEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putBaseEventQuery",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) PutBaseMetricsQuery(value *MonitorVariablesAggregateAugmentedQueryBaseMetricsQuery) {
	if err := m.validatePutBaseMetricsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putBaseMetricsQuery",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) PutCompute(value interface{}) {
	if err := m.validatePutComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCompute",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) PutGroupBy(value interface{}) {
	if err := m.validatePutGroupByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putGroupBy",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) PutJoinCondition(value *MonitorVariablesAggregateAugmentedQueryJoinCondition) {
	if err := m.validatePutJoinConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putJoinCondition",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) ResetAugmentEventQuery() {
	_jsii_.InvokeVoid(
		m,
		"resetAugmentEventQuery",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) ResetAugmentReferenceTable() {
	_jsii_.InvokeVoid(
		m,
		"resetAugmentReferenceTable",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) ResetBaseEventQuery() {
	_jsii_.InvokeVoid(
		m,
		"resetBaseEventQuery",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) ResetBaseMetricsQuery() {
	_jsii_.InvokeVoid(
		m,
		"resetBaseMetricsQuery",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		m,
		"resetName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MonitorVariablesAggregateAugmentedQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

