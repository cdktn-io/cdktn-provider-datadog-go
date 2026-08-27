// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference interface {
	cdktn.ComplexObject
	ApmDependencyStatsQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQueryOutputReference
	ApmDependencyStatsQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQuery
	ApmMetricsQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQueryOutputReference
	ApmMetricsQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQuery
	ApmResourceStatsQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQueryOutputReference
	ApmResourceStatsQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQuery
	CloudCostQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQueryOutputReference
	CloudCostQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQuery
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
	EventQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQueryOutputReference
	EventQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQuery
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQueryOutputReference
	MetricQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQuery
	ProcessQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQueryOutputReference
	ProcessQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQuery
	ProductAnalyticsExtendedQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQueryOutputReference
	ProductAnalyticsExtendedQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQuery
	RetentionQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQueryOutputReference
	RetentionQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQuery
	SloQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQueryOutputReference
	SloQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQuery
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserJourneyQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQueryOutputReference
	UserJourneyQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuery
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
	PutApmDependencyStatsQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQuery)
	PutApmMetricsQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQuery)
	PutApmResourceStatsQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQuery)
	PutCloudCostQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQuery)
	PutEventQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQuery)
	PutMetricQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQuery)
	PutProcessQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQuery)
	PutProductAnalyticsExtendedQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQuery)
	PutRetentionQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQuery)
	PutSloQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQuery)
	PutUserJourneyQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuery)
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

// The jsii proxy struct for DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference
type jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ApmDependencyStatsQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmDependencyStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ApmDependencyStatsQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQuery
	_jsii_.Get(
		j,
		"apmDependencyStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ApmMetricsQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQueryOutputReference
	_jsii_.Get(
		j,
		"apmMetricsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ApmMetricsQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQuery
	_jsii_.Get(
		j,
		"apmMetricsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ApmResourceStatsQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmResourceStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ApmResourceStatsQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"apmResourceStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) CloudCostQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQueryOutputReference
	_jsii_.Get(
		j,
		"cloudCostQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) CloudCostQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQuery
	_jsii_.Get(
		j,
		"cloudCostQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) EventQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) EventQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) MetricQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQueryOutputReference
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) MetricQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQuery
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ProcessQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ProcessQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ProductAnalyticsExtendedQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQueryOutputReference
	_jsii_.Get(
		j,
		"productAnalyticsExtendedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ProductAnalyticsExtendedQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQuery
	_jsii_.Get(
		j,
		"productAnalyticsExtendedQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) RetentionQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQueryOutputReference
	_jsii_.Get(
		j,
		"retentionQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) RetentionQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQuery
	_jsii_.Get(
		j,
		"retentionQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) SloQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQueryOutputReference
	_jsii_.Get(
		j,
		"sloQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) SloQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQuery
	_jsii_.Get(
		j,
		"sloQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) UserJourneyQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQueryOutputReference
	_jsii_.Get(
		j,
		"userJourneyQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) UserJourneyQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuery
	_jsii_.Get(
		j,
		"userJourneyQueryInput",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference_Override(d DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutApmDependencyStatsQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQuery) {
	if err := d.validatePutApmDependencyStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmDependencyStatsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutApmMetricsQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQuery) {
	if err := d.validatePutApmMetricsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmMetricsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutApmResourceStatsQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQuery) {
	if err := d.validatePutApmResourceStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmResourceStatsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutCloudCostQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQuery) {
	if err := d.validatePutCloudCostQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCloudCostQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutEventQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQuery) {
	if err := d.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutMetricQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQuery) {
	if err := d.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutProcessQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQuery) {
	if err := d.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutProductAnalyticsExtendedQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQuery) {
	if err := d.validatePutProductAnalyticsExtendedQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProductAnalyticsExtendedQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutRetentionQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQuery) {
	if err := d.validatePutRetentionQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRetentionQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutSloQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQuery) {
	if err := d.validatePutSloQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSloQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutUserJourneyQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuery) {
	if err := d.validatePutUserJourneyQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUserJourneyQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetApmDependencyStatsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmDependencyStatsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetApmMetricsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmMetricsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetApmResourceStatsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmResourceStatsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetCloudCostQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudCostQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetProductAnalyticsExtendedQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProductAnalyticsExtendedQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetRetentionQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetRetentionQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetSloQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetSloQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetUserJourneyQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetUserJourneyQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

