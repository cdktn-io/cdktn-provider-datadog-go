// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference interface {
	cdktn.ComplexObject
	ApmDependencyStatsQuery() DashboardV2WidgetScatterplotDefinitionRequestYQueryApmDependencyStatsQueryOutputReference
	ApmDependencyStatsQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQueryApmDependencyStatsQuery
	ApmMetricsQuery() DashboardV2WidgetScatterplotDefinitionRequestYQueryApmMetricsQueryOutputReference
	ApmMetricsQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQueryApmMetricsQuery
	ApmResourceStatsQuery() DashboardV2WidgetScatterplotDefinitionRequestYQueryApmResourceStatsQueryOutputReference
	ApmResourceStatsQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQueryApmResourceStatsQuery
	CloudCostQuery() DashboardV2WidgetScatterplotDefinitionRequestYQueryCloudCostQueryOutputReference
	CloudCostQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQueryCloudCostQuery
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
	EventQuery() DashboardV2WidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference
	EventQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQueryEventQuery
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricQuery() DashboardV2WidgetScatterplotDefinitionRequestYQueryMetricQueryOutputReference
	MetricQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQueryMetricQuery
	ProcessQuery() DashboardV2WidgetScatterplotDefinitionRequestYQueryProcessQueryOutputReference
	ProcessQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQueryProcessQuery
	ProductAnalyticsExtendedQuery() DashboardV2WidgetScatterplotDefinitionRequestYQueryProductAnalyticsExtendedQueryOutputReference
	ProductAnalyticsExtendedQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQueryProductAnalyticsExtendedQuery
	RetentionQuery() DashboardV2WidgetScatterplotDefinitionRequestYQueryRetentionQueryOutputReference
	RetentionQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQueryRetentionQuery
	SloQuery() DashboardV2WidgetScatterplotDefinitionRequestYQuerySloQueryOutputReference
	SloQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQuerySloQuery
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserJourneyQuery() DashboardV2WidgetScatterplotDefinitionRequestYQueryUserJourneyQueryOutputReference
	UserJourneyQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQueryUserJourneyQuery
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
	PutApmDependencyStatsQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQueryApmDependencyStatsQuery)
	PutApmMetricsQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQueryApmMetricsQuery)
	PutApmResourceStatsQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQueryApmResourceStatsQuery)
	PutCloudCostQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQueryCloudCostQuery)
	PutEventQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQueryEventQuery)
	PutMetricQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQueryMetricQuery)
	PutProcessQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQueryProcessQuery)
	PutProductAnalyticsExtendedQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQueryProductAnalyticsExtendedQuery)
	PutRetentionQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQueryRetentionQuery)
	PutSloQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQuerySloQuery)
	PutUserJourneyQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQueryUserJourneyQuery)
	ResetApmDependencyStatsQuery()
	ResetApmMetricsQuery()
	ResetApmResourceStatsQuery()
	ResetCloudCostQuery()
	ResetEventQuery()
	ResetMetricQuery()
	ResetProcessQuery()
	ResetProductAnalyticsExtendedQuery()
	ResetRetentionQuery()
	ResetSloQuery()
	ResetUserJourneyQuery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference
type jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ApmDependencyStatsQuery() DashboardV2WidgetScatterplotDefinitionRequestYQueryApmDependencyStatsQueryOutputReference {
	var returns DashboardV2WidgetScatterplotDefinitionRequestYQueryApmDependencyStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmDependencyStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ApmDependencyStatsQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQueryApmDependencyStatsQuery {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestYQueryApmDependencyStatsQuery
	_jsii_.Get(
		j,
		"apmDependencyStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ApmMetricsQuery() DashboardV2WidgetScatterplotDefinitionRequestYQueryApmMetricsQueryOutputReference {
	var returns DashboardV2WidgetScatterplotDefinitionRequestYQueryApmMetricsQueryOutputReference
	_jsii_.Get(
		j,
		"apmMetricsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ApmMetricsQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQueryApmMetricsQuery {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestYQueryApmMetricsQuery
	_jsii_.Get(
		j,
		"apmMetricsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ApmResourceStatsQuery() DashboardV2WidgetScatterplotDefinitionRequestYQueryApmResourceStatsQueryOutputReference {
	var returns DashboardV2WidgetScatterplotDefinitionRequestYQueryApmResourceStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmResourceStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ApmResourceStatsQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQueryApmResourceStatsQuery {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestYQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"apmResourceStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) CloudCostQuery() DashboardV2WidgetScatterplotDefinitionRequestYQueryCloudCostQueryOutputReference {
	var returns DashboardV2WidgetScatterplotDefinitionRequestYQueryCloudCostQueryOutputReference
	_jsii_.Get(
		j,
		"cloudCostQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) CloudCostQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQueryCloudCostQuery {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestYQueryCloudCostQuery
	_jsii_.Get(
		j,
		"cloudCostQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) EventQuery() DashboardV2WidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference {
	var returns DashboardV2WidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) EventQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQueryEventQuery {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestYQueryEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) MetricQuery() DashboardV2WidgetScatterplotDefinitionRequestYQueryMetricQueryOutputReference {
	var returns DashboardV2WidgetScatterplotDefinitionRequestYQueryMetricQueryOutputReference
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) MetricQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQueryMetricQuery {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestYQueryMetricQuery
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ProcessQuery() DashboardV2WidgetScatterplotDefinitionRequestYQueryProcessQueryOutputReference {
	var returns DashboardV2WidgetScatterplotDefinitionRequestYQueryProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ProcessQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQueryProcessQuery {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestYQueryProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ProductAnalyticsExtendedQuery() DashboardV2WidgetScatterplotDefinitionRequestYQueryProductAnalyticsExtendedQueryOutputReference {
	var returns DashboardV2WidgetScatterplotDefinitionRequestYQueryProductAnalyticsExtendedQueryOutputReference
	_jsii_.Get(
		j,
		"productAnalyticsExtendedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ProductAnalyticsExtendedQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQueryProductAnalyticsExtendedQuery {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestYQueryProductAnalyticsExtendedQuery
	_jsii_.Get(
		j,
		"productAnalyticsExtendedQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) RetentionQuery() DashboardV2WidgetScatterplotDefinitionRequestYQueryRetentionQueryOutputReference {
	var returns DashboardV2WidgetScatterplotDefinitionRequestYQueryRetentionQueryOutputReference
	_jsii_.Get(
		j,
		"retentionQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) RetentionQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQueryRetentionQuery {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestYQueryRetentionQuery
	_jsii_.Get(
		j,
		"retentionQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) SloQuery() DashboardV2WidgetScatterplotDefinitionRequestYQuerySloQueryOutputReference {
	var returns DashboardV2WidgetScatterplotDefinitionRequestYQuerySloQueryOutputReference
	_jsii_.Get(
		j,
		"sloQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) SloQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQuerySloQuery {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestYQuerySloQuery
	_jsii_.Get(
		j,
		"sloQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) UserJourneyQuery() DashboardV2WidgetScatterplotDefinitionRequestYQueryUserJourneyQueryOutputReference {
	var returns DashboardV2WidgetScatterplotDefinitionRequestYQueryUserJourneyQueryOutputReference
	_jsii_.Get(
		j,
		"userJourneyQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) UserJourneyQueryInput() *DashboardV2WidgetScatterplotDefinitionRequestYQueryUserJourneyQuery {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestYQueryUserJourneyQuery
	_jsii_.Get(
		j,
		"userJourneyQueryInput",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference_Override(d DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) PutApmDependencyStatsQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQueryApmDependencyStatsQuery) {
	if err := d.validatePutApmDependencyStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmDependencyStatsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) PutApmMetricsQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQueryApmMetricsQuery) {
	if err := d.validatePutApmMetricsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmMetricsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) PutApmResourceStatsQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQueryApmResourceStatsQuery) {
	if err := d.validatePutApmResourceStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmResourceStatsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) PutCloudCostQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQueryCloudCostQuery) {
	if err := d.validatePutCloudCostQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCloudCostQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) PutEventQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQueryEventQuery) {
	if err := d.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) PutMetricQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQueryMetricQuery) {
	if err := d.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) PutProcessQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQueryProcessQuery) {
	if err := d.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) PutProductAnalyticsExtendedQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQueryProductAnalyticsExtendedQuery) {
	if err := d.validatePutProductAnalyticsExtendedQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProductAnalyticsExtendedQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) PutRetentionQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQueryRetentionQuery) {
	if err := d.validatePutRetentionQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRetentionQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) PutSloQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQuerySloQuery) {
	if err := d.validatePutSloQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSloQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) PutUserJourneyQuery(value *DashboardV2WidgetScatterplotDefinitionRequestYQueryUserJourneyQuery) {
	if err := d.validatePutUserJourneyQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUserJourneyQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ResetApmDependencyStatsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmDependencyStatsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ResetApmMetricsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmMetricsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ResetApmResourceStatsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmResourceStatsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ResetCloudCostQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudCostQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ResetProductAnalyticsExtendedQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProductAnalyticsExtendedQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ResetRetentionQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetRetentionQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ResetSloQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetSloQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ResetUserJourneyQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetUserJourneyQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

