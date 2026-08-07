// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference interface {
	cdktn.ComplexObject
	ApmDependencyStatsQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQueryOutputReference
	ApmDependencyStatsQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQuery
	ApmResourceStatsQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference
	ApmResourceStatsQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery
	CloudCostQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryCloudCostQueryOutputReference
	CloudCostQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryCloudCostQuery
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
	EventQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryEventQueryOutputReference
	EventQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryEventQuery
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryMetricQueryOutputReference
	MetricQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryMetricQuery
	ProcessQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryProcessQueryOutputReference
	ProcessQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryProcessQuery
	SloQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQuerySloQueryOutputReference
	SloQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQuerySloQuery
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
	PutApmDependencyStatsQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQuery)
	PutApmResourceStatsQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery)
	PutCloudCostQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryCloudCostQuery)
	PutEventQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryEventQuery)
	PutMetricQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryMetricQuery)
	PutProcessQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryProcessQuery)
	PutSloQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQuerySloQuery)
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

// The jsii proxy struct for DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference
type jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ApmDependencyStatsQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmDependencyStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ApmDependencyStatsQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQuery
	_jsii_.Get(
		j,
		"apmDependencyStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ApmResourceStatsQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmResourceStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ApmResourceStatsQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"apmResourceStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) CloudCostQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryCloudCostQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryCloudCostQueryOutputReference
	_jsii_.Get(
		j,
		"cloudCostQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) CloudCostQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryCloudCostQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryCloudCostQuery
	_jsii_.Get(
		j,
		"cloudCostQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) EventQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryEventQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) EventQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryEventQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) MetricQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryMetricQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryMetricQueryOutputReference
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) MetricQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryMetricQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryMetricQuery
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ProcessQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryProcessQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ProcessQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryProcessQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) SloQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQuerySloQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQuerySloQueryOutputReference
	_jsii_.Get(
		j,
		"sloQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) SloQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQuerySloQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQuerySloQuery
	_jsii_.Get(
		j,
		"sloQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference_Override(d DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutApmDependencyStatsQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQuery) {
	if err := d.validatePutApmDependencyStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmDependencyStatsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutApmResourceStatsQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery) {
	if err := d.validatePutApmResourceStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmResourceStatsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutCloudCostQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryCloudCostQuery) {
	if err := d.validatePutCloudCostQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCloudCostQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutEventQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryEventQuery) {
	if err := d.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutMetricQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryMetricQuery) {
	if err := d.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutProcessQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryProcessQuery) {
	if err := d.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutSloQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQuerySloQuery) {
	if err := d.validatePutSloQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSloQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetApmDependencyStatsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmDependencyStatsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetApmResourceStatsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmResourceStatsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetCloudCostQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudCostQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetSloQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetSloQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

