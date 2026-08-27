// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference interface {
	cdktn.ComplexObject
	ApmDependencyStatsQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmDependencyStatsQueryOutputReference
	ApmDependencyStatsQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmDependencyStatsQuery
	ApmMetricsQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmMetricsQueryOutputReference
	ApmMetricsQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmMetricsQuery
	ApmResourceStatsQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference
	ApmResourceStatsQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmResourceStatsQuery
	CloudCostQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryCloudCostQueryOutputReference
	CloudCostQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryCloudCostQuery
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
	EventQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryEventQueryOutputReference
	EventQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryEventQuery
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryMetricQueryOutputReference
	MetricQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryMetricQuery
	ProcessQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProcessQueryOutputReference
	ProcessQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProcessQuery
	ProductAnalyticsExtendedQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference
	ProductAnalyticsExtendedQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQuery
	RetentionQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryRetentionQueryOutputReference
	RetentionQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryRetentionQuery
	SloQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference
	SloQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQuerySloQuery
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserJourneyQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryUserJourneyQueryOutputReference
	UserJourneyQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryUserJourneyQuery
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
	PutApmDependencyStatsQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmDependencyStatsQuery)
	PutApmMetricsQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmMetricsQuery)
	PutApmResourceStatsQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmResourceStatsQuery)
	PutCloudCostQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryCloudCostQuery)
	PutEventQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryEventQuery)
	PutMetricQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryMetricQuery)
	PutProcessQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProcessQuery)
	PutProductAnalyticsExtendedQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQuery)
	PutRetentionQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryRetentionQuery)
	PutSloQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQuerySloQuery)
	PutUserJourneyQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryUserJourneyQuery)
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

// The jsii proxy struct for DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference
type jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ApmDependencyStatsQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmDependencyStatsQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmDependencyStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmDependencyStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ApmDependencyStatsQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmDependencyStatsQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmDependencyStatsQuery
	_jsii_.Get(
		j,
		"apmDependencyStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ApmMetricsQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmMetricsQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmMetricsQueryOutputReference
	_jsii_.Get(
		j,
		"apmMetricsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ApmMetricsQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmMetricsQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmMetricsQuery
	_jsii_.Get(
		j,
		"apmMetricsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ApmResourceStatsQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmResourceStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ApmResourceStatsQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmResourceStatsQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"apmResourceStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) CloudCostQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryCloudCostQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryCloudCostQueryOutputReference
	_jsii_.Get(
		j,
		"cloudCostQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) CloudCostQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryCloudCostQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryCloudCostQuery
	_jsii_.Get(
		j,
		"cloudCostQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) EventQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryEventQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) EventQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryEventQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) MetricQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryMetricQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryMetricQueryOutputReference
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) MetricQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryMetricQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryMetricQuery
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ProcessQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProcessQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ProcessQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProcessQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ProductAnalyticsExtendedQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference
	_jsii_.Get(
		j,
		"productAnalyticsExtendedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ProductAnalyticsExtendedQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQuery
	_jsii_.Get(
		j,
		"productAnalyticsExtendedQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) RetentionQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryRetentionQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryRetentionQueryOutputReference
	_jsii_.Get(
		j,
		"retentionQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) RetentionQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryRetentionQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryRetentionQuery
	_jsii_.Get(
		j,
		"retentionQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) SloQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference
	_jsii_.Get(
		j,
		"sloQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) SloQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQuerySloQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQuerySloQuery
	_jsii_.Get(
		j,
		"sloQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) UserJourneyQuery() DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryUserJourneyQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryUserJourneyQueryOutputReference
	_jsii_.Get(
		j,
		"userJourneyQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) UserJourneyQueryInput() *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryUserJourneyQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryUserJourneyQuery
	_jsii_.Get(
		j,
		"userJourneyQueryInput",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference_Override(d DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) PutApmDependencyStatsQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmDependencyStatsQuery) {
	if err := d.validatePutApmDependencyStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmDependencyStatsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) PutApmMetricsQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmMetricsQuery) {
	if err := d.validatePutApmMetricsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmMetricsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) PutApmResourceStatsQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryApmResourceStatsQuery) {
	if err := d.validatePutApmResourceStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmResourceStatsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) PutCloudCostQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryCloudCostQuery) {
	if err := d.validatePutCloudCostQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCloudCostQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) PutEventQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryEventQuery) {
	if err := d.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) PutMetricQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryMetricQuery) {
	if err := d.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) PutProcessQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProcessQuery) {
	if err := d.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) PutProductAnalyticsExtendedQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQuery) {
	if err := d.validatePutProductAnalyticsExtendedQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProductAnalyticsExtendedQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) PutRetentionQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryRetentionQuery) {
	if err := d.validatePutRetentionQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRetentionQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) PutSloQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQuerySloQuery) {
	if err := d.validatePutSloQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSloQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) PutUserJourneyQuery(value *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryUserJourneyQuery) {
	if err := d.validatePutUserJourneyQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUserJourneyQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ResetApmDependencyStatsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmDependencyStatsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ResetApmMetricsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmMetricsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ResetApmResourceStatsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmResourceStatsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ResetCloudCostQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudCostQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ResetProductAnalyticsExtendedQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProductAnalyticsExtendedQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ResetRetentionQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetRetentionQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ResetSloQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetSloQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ResetUserJourneyQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetUserJourneyQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

