// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference interface {
	cdktn.ComplexObject
	ApmDependencyStatsQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmDependencyStatsQueryOutputReference
	ApmDependencyStatsQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmDependencyStatsQuery
	ApmResourceStatsQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference
	ApmResourceStatsQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQuery
	CloudCostQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryCloudCostQueryOutputReference
	CloudCostQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryCloudCostQuery
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
	EventQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryEventQueryOutputReference
	EventQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryEventQuery
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryMetricQueryOutputReference
	MetricQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryMetricQuery
	ProcessQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference
	ProcessQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQuery
	SloQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQuerySloQueryOutputReference
	SloQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQuerySloQuery
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
	PutApmDependencyStatsQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmDependencyStatsQuery)
	PutApmResourceStatsQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQuery)
	PutCloudCostQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryCloudCostQuery)
	PutEventQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryEventQuery)
	PutMetricQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryMetricQuery)
	PutProcessQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQuery)
	PutSloQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQuerySloQuery)
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

// The jsii proxy struct for DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference
type jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) ApmDependencyStatsQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmDependencyStatsQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmDependencyStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmDependencyStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) ApmDependencyStatsQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmDependencyStatsQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmDependencyStatsQuery
	_jsii_.Get(
		j,
		"apmDependencyStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) ApmResourceStatsQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmResourceStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) ApmResourceStatsQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"apmResourceStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) CloudCostQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryCloudCostQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryCloudCostQueryOutputReference
	_jsii_.Get(
		j,
		"cloudCostQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) CloudCostQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryCloudCostQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryCloudCostQuery
	_jsii_.Get(
		j,
		"cloudCostQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) EventQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryEventQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) EventQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryEventQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) MetricQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryMetricQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryMetricQueryOutputReference
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) MetricQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryMetricQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryMetricQuery
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) ProcessQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) ProcessQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) SloQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQuerySloQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQuerySloQueryOutputReference
	_jsii_.Get(
		j,
		"sloQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) SloQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQuerySloQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQuerySloQuery
	_jsii_.Get(
		j,
		"sloQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference_Override(d DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) PutApmDependencyStatsQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmDependencyStatsQuery) {
	if err := d.validatePutApmDependencyStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmDependencyStatsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) PutApmResourceStatsQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQuery) {
	if err := d.validatePutApmResourceStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmResourceStatsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) PutCloudCostQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryCloudCostQuery) {
	if err := d.validatePutCloudCostQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCloudCostQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) PutEventQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryEventQuery) {
	if err := d.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) PutMetricQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryMetricQuery) {
	if err := d.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) PutProcessQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQuery) {
	if err := d.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) PutSloQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQuerySloQuery) {
	if err := d.validatePutSloQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSloQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) ResetApmDependencyStatsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmDependencyStatsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) ResetApmResourceStatsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmResourceStatsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) ResetCloudCostQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudCostQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) ResetSloQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetSloQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

