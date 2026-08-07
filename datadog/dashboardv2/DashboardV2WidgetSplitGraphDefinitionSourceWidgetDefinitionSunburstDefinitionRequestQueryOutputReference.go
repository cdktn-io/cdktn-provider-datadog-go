// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference interface {
	cdktn.ComplexObject
	ApmDependencyStatsQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmDependencyStatsQueryOutputReference
	ApmDependencyStatsQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmDependencyStatsQuery
	ApmResourceStatsQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmResourceStatsQueryOutputReference
	ApmResourceStatsQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmResourceStatsQuery
	CloudCostQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryCloudCostQueryOutputReference
	CloudCostQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryCloudCostQuery
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
	EventQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryEventQueryOutputReference
	EventQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryEventQuery
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryMetricQueryOutputReference
	MetricQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryMetricQuery
	ProcessQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryProcessQueryOutputReference
	ProcessQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryProcessQuery
	SloQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQuerySloQueryOutputReference
	SloQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQuerySloQuery
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
	PutApmDependencyStatsQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmDependencyStatsQuery)
	PutApmResourceStatsQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmResourceStatsQuery)
	PutCloudCostQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryCloudCostQuery)
	PutEventQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryEventQuery)
	PutMetricQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryMetricQuery)
	PutProcessQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryProcessQuery)
	PutSloQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQuerySloQuery)
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

// The jsii proxy struct for DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference
type jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ApmDependencyStatsQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmDependencyStatsQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmDependencyStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmDependencyStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ApmDependencyStatsQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmDependencyStatsQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmDependencyStatsQuery
	_jsii_.Get(
		j,
		"apmDependencyStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ApmResourceStatsQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmResourceStatsQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmResourceStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmResourceStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ApmResourceStatsQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmResourceStatsQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"apmResourceStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) CloudCostQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryCloudCostQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryCloudCostQueryOutputReference
	_jsii_.Get(
		j,
		"cloudCostQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) CloudCostQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryCloudCostQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryCloudCostQuery
	_jsii_.Get(
		j,
		"cloudCostQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) EventQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryEventQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) EventQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryEventQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) MetricQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryMetricQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryMetricQueryOutputReference
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) MetricQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryMetricQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryMetricQuery
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ProcessQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryProcessQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ProcessQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryProcessQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) SloQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQuerySloQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQuerySloQueryOutputReference
	_jsii_.Get(
		j,
		"sloQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) SloQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQuerySloQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQuerySloQuery
	_jsii_.Get(
		j,
		"sloQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference_Override(d DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) PutApmDependencyStatsQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmDependencyStatsQuery) {
	if err := d.validatePutApmDependencyStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmDependencyStatsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) PutApmResourceStatsQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmResourceStatsQuery) {
	if err := d.validatePutApmResourceStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmResourceStatsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) PutCloudCostQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryCloudCostQuery) {
	if err := d.validatePutCloudCostQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCloudCostQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) PutEventQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryEventQuery) {
	if err := d.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) PutMetricQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryMetricQuery) {
	if err := d.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) PutProcessQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryProcessQuery) {
	if err := d.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) PutSloQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQuerySloQuery) {
	if err := d.validatePutSloQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSloQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ResetApmDependencyStatsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmDependencyStatsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ResetApmResourceStatsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmResourceStatsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ResetCloudCostQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudCostQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ResetSloQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetSloQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

