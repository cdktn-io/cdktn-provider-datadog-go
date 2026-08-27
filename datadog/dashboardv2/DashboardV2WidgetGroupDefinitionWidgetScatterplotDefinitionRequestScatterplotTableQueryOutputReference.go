// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference interface {
	cdktn.ComplexObject
	ApmDependencyStatsQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQueryOutputReference
	ApmDependencyStatsQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQuery
	ApmMetricsQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmMetricsQueryOutputReference
	ApmMetricsQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmMetricsQuery
	ApmResourceStatsQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference
	ApmResourceStatsQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery
	CloudCostQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQueryOutputReference
	CloudCostQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQuery
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
	EventQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryEventQueryOutputReference
	EventQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryEventQuery
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQueryOutputReference
	MetricQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQuery
	ProcessQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQueryOutputReference
	ProcessQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQuery
	ProductAnalyticsExtendedQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryOutputReference
	ProductAnalyticsExtendedQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQuery
	RetentionQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQueryOutputReference
	RetentionQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuery
	SloQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQuerySloQueryOutputReference
	SloQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQuerySloQuery
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserJourneyQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQueryOutputReference
	UserJourneyQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuery
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
	PutApmDependencyStatsQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQuery)
	PutApmMetricsQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmMetricsQuery)
	PutApmResourceStatsQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery)
	PutCloudCostQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQuery)
	PutEventQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryEventQuery)
	PutMetricQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQuery)
	PutProcessQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQuery)
	PutProductAnalyticsExtendedQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQuery)
	PutRetentionQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuery)
	PutSloQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQuerySloQuery)
	PutUserJourneyQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuery)
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

// The jsii proxy struct for DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference
type jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ApmDependencyStatsQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmDependencyStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ApmDependencyStatsQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQuery
	_jsii_.Get(
		j,
		"apmDependencyStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ApmMetricsQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmMetricsQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmMetricsQueryOutputReference
	_jsii_.Get(
		j,
		"apmMetricsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ApmMetricsQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmMetricsQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmMetricsQuery
	_jsii_.Get(
		j,
		"apmMetricsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ApmResourceStatsQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmResourceStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ApmResourceStatsQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"apmResourceStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) CloudCostQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQueryOutputReference
	_jsii_.Get(
		j,
		"cloudCostQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) CloudCostQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQuery
	_jsii_.Get(
		j,
		"cloudCostQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) EventQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryEventQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) EventQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryEventQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) MetricQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQueryOutputReference
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) MetricQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQuery
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ProcessQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ProcessQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ProductAnalyticsExtendedQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryOutputReference
	_jsii_.Get(
		j,
		"productAnalyticsExtendedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ProductAnalyticsExtendedQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQuery
	_jsii_.Get(
		j,
		"productAnalyticsExtendedQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) RetentionQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQueryOutputReference
	_jsii_.Get(
		j,
		"retentionQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) RetentionQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuery
	_jsii_.Get(
		j,
		"retentionQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) SloQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQuerySloQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQuerySloQueryOutputReference
	_jsii_.Get(
		j,
		"sloQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) SloQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQuerySloQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQuerySloQuery
	_jsii_.Get(
		j,
		"sloQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) UserJourneyQuery() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQueryOutputReference
	_jsii_.Get(
		j,
		"userJourneyQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) UserJourneyQueryInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuery
	_jsii_.Get(
		j,
		"userJourneyQueryInput",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference_Override(d DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutApmDependencyStatsQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQuery) {
	if err := d.validatePutApmDependencyStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmDependencyStatsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutApmMetricsQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmMetricsQuery) {
	if err := d.validatePutApmMetricsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmMetricsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutApmResourceStatsQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery) {
	if err := d.validatePutApmResourceStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmResourceStatsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutCloudCostQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQuery) {
	if err := d.validatePutCloudCostQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCloudCostQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutEventQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryEventQuery) {
	if err := d.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutMetricQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQuery) {
	if err := d.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutProcessQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQuery) {
	if err := d.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutProductAnalyticsExtendedQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQuery) {
	if err := d.validatePutProductAnalyticsExtendedQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProductAnalyticsExtendedQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutRetentionQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuery) {
	if err := d.validatePutRetentionQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRetentionQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutSloQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQuerySloQuery) {
	if err := d.validatePutSloQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSloQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutUserJourneyQuery(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuery) {
	if err := d.validatePutUserJourneyQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUserJourneyQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetApmDependencyStatsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmDependencyStatsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetApmMetricsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmMetricsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetApmResourceStatsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmResourceStatsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetCloudCostQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudCostQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetProductAnalyticsExtendedQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProductAnalyticsExtendedQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetRetentionQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetRetentionQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetSloQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetSloQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetUserJourneyQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetUserJourneyQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

