// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference interface {
	cdktn.ComplexObject
	ApmDependencyStatsQuery() DashboardV2WidgetHostmapDefinitionRequestFillQueryApmDependencyStatsQueryOutputReference
	ApmDependencyStatsQueryInput() *DashboardV2WidgetHostmapDefinitionRequestFillQueryApmDependencyStatsQuery
	ApmResourceStatsQuery() DashboardV2WidgetHostmapDefinitionRequestFillQueryApmResourceStatsQueryOutputReference
	ApmResourceStatsQueryInput() *DashboardV2WidgetHostmapDefinitionRequestFillQueryApmResourceStatsQuery
	CloudCostQuery() DashboardV2WidgetHostmapDefinitionRequestFillQueryCloudCostQueryOutputReference
	CloudCostQueryInput() *DashboardV2WidgetHostmapDefinitionRequestFillQueryCloudCostQuery
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
	EventQuery() DashboardV2WidgetHostmapDefinitionRequestFillQueryEventQueryOutputReference
	EventQueryInput() *DashboardV2WidgetHostmapDefinitionRequestFillQueryEventQuery
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricQuery() DashboardV2WidgetHostmapDefinitionRequestFillQueryMetricQueryOutputReference
	MetricQueryInput() *DashboardV2WidgetHostmapDefinitionRequestFillQueryMetricQuery
	ProcessQuery() DashboardV2WidgetHostmapDefinitionRequestFillQueryProcessQueryOutputReference
	ProcessQueryInput() *DashboardV2WidgetHostmapDefinitionRequestFillQueryProcessQuery
	SloQuery() DashboardV2WidgetHostmapDefinitionRequestFillQuerySloQueryOutputReference
	SloQueryInput() *DashboardV2WidgetHostmapDefinitionRequestFillQuerySloQuery
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
	PutApmDependencyStatsQuery(value *DashboardV2WidgetHostmapDefinitionRequestFillQueryApmDependencyStatsQuery)
	PutApmResourceStatsQuery(value *DashboardV2WidgetHostmapDefinitionRequestFillQueryApmResourceStatsQuery)
	PutCloudCostQuery(value *DashboardV2WidgetHostmapDefinitionRequestFillQueryCloudCostQuery)
	PutEventQuery(value *DashboardV2WidgetHostmapDefinitionRequestFillQueryEventQuery)
	PutMetricQuery(value *DashboardV2WidgetHostmapDefinitionRequestFillQueryMetricQuery)
	PutProcessQuery(value *DashboardV2WidgetHostmapDefinitionRequestFillQueryProcessQuery)
	PutSloQuery(value *DashboardV2WidgetHostmapDefinitionRequestFillQuerySloQuery)
	ResetApmDependencyStatsQuery()
	ResetApmResourceStatsQuery()
	ResetCloudCostQuery()
	ResetEventQuery()
	ResetMetricQuery()
	ResetProcessQuery()
	ResetSloQuery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference
type jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ApmDependencyStatsQuery() DashboardV2WidgetHostmapDefinitionRequestFillQueryApmDependencyStatsQueryOutputReference {
	var returns DashboardV2WidgetHostmapDefinitionRequestFillQueryApmDependencyStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmDependencyStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ApmDependencyStatsQueryInput() *DashboardV2WidgetHostmapDefinitionRequestFillQueryApmDependencyStatsQuery {
	var returns *DashboardV2WidgetHostmapDefinitionRequestFillQueryApmDependencyStatsQuery
	_jsii_.Get(
		j,
		"apmDependencyStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ApmResourceStatsQuery() DashboardV2WidgetHostmapDefinitionRequestFillQueryApmResourceStatsQueryOutputReference {
	var returns DashboardV2WidgetHostmapDefinitionRequestFillQueryApmResourceStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmResourceStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ApmResourceStatsQueryInput() *DashboardV2WidgetHostmapDefinitionRequestFillQueryApmResourceStatsQuery {
	var returns *DashboardV2WidgetHostmapDefinitionRequestFillQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"apmResourceStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) CloudCostQuery() DashboardV2WidgetHostmapDefinitionRequestFillQueryCloudCostQueryOutputReference {
	var returns DashboardV2WidgetHostmapDefinitionRequestFillQueryCloudCostQueryOutputReference
	_jsii_.Get(
		j,
		"cloudCostQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) CloudCostQueryInput() *DashboardV2WidgetHostmapDefinitionRequestFillQueryCloudCostQuery {
	var returns *DashboardV2WidgetHostmapDefinitionRequestFillQueryCloudCostQuery
	_jsii_.Get(
		j,
		"cloudCostQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) EventQuery() DashboardV2WidgetHostmapDefinitionRequestFillQueryEventQueryOutputReference {
	var returns DashboardV2WidgetHostmapDefinitionRequestFillQueryEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) EventQueryInput() *DashboardV2WidgetHostmapDefinitionRequestFillQueryEventQuery {
	var returns *DashboardV2WidgetHostmapDefinitionRequestFillQueryEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) MetricQuery() DashboardV2WidgetHostmapDefinitionRequestFillQueryMetricQueryOutputReference {
	var returns DashboardV2WidgetHostmapDefinitionRequestFillQueryMetricQueryOutputReference
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) MetricQueryInput() *DashboardV2WidgetHostmapDefinitionRequestFillQueryMetricQuery {
	var returns *DashboardV2WidgetHostmapDefinitionRequestFillQueryMetricQuery
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ProcessQuery() DashboardV2WidgetHostmapDefinitionRequestFillQueryProcessQueryOutputReference {
	var returns DashboardV2WidgetHostmapDefinitionRequestFillQueryProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ProcessQueryInput() *DashboardV2WidgetHostmapDefinitionRequestFillQueryProcessQuery {
	var returns *DashboardV2WidgetHostmapDefinitionRequestFillQueryProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) SloQuery() DashboardV2WidgetHostmapDefinitionRequestFillQuerySloQueryOutputReference {
	var returns DashboardV2WidgetHostmapDefinitionRequestFillQuerySloQueryOutputReference
	_jsii_.Get(
		j,
		"sloQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) SloQueryInput() *DashboardV2WidgetHostmapDefinitionRequestFillQuerySloQuery {
	var returns *DashboardV2WidgetHostmapDefinitionRequestFillQuerySloQuery
	_jsii_.Get(
		j,
		"sloQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference_Override(d DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) PutApmDependencyStatsQuery(value *DashboardV2WidgetHostmapDefinitionRequestFillQueryApmDependencyStatsQuery) {
	if err := d.validatePutApmDependencyStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmDependencyStatsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) PutApmResourceStatsQuery(value *DashboardV2WidgetHostmapDefinitionRequestFillQueryApmResourceStatsQuery) {
	if err := d.validatePutApmResourceStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmResourceStatsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) PutCloudCostQuery(value *DashboardV2WidgetHostmapDefinitionRequestFillQueryCloudCostQuery) {
	if err := d.validatePutCloudCostQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCloudCostQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) PutEventQuery(value *DashboardV2WidgetHostmapDefinitionRequestFillQueryEventQuery) {
	if err := d.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) PutMetricQuery(value *DashboardV2WidgetHostmapDefinitionRequestFillQueryMetricQuery) {
	if err := d.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) PutProcessQuery(value *DashboardV2WidgetHostmapDefinitionRequestFillQueryProcessQuery) {
	if err := d.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) PutSloQuery(value *DashboardV2WidgetHostmapDefinitionRequestFillQuerySloQuery) {
	if err := d.validatePutSloQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSloQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ResetApmDependencyStatsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmDependencyStatsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ResetApmResourceStatsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmResourceStatsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ResetCloudCostQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudCostQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ResetSloQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetSloQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

