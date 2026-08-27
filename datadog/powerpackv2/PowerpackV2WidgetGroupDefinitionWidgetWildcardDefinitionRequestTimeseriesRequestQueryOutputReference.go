// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference interface {
	cdktn.ComplexObject
	ApmDependencyStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQueryOutputReference
	ApmDependencyStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQuery
	ApmMetricsQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQueryOutputReference
	ApmMetricsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQuery
	ApmResourceStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQueryOutputReference
	ApmResourceStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQuery
	CloudCostQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQueryOutputReference
	CloudCostQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQuery
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
	EventQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQueryOutputReference
	EventQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQuery
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQueryOutputReference
	MetricQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQuery
	ProcessQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQueryOutputReference
	ProcessQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQuery
	ProductAnalyticsExtendedQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQueryOutputReference
	ProductAnalyticsExtendedQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQuery
	RetentionQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQueryOutputReference
	RetentionQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQuery
	SloQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQueryOutputReference
	SloQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQuery
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserJourneyQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQueryOutputReference
	UserJourneyQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuery
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
	PutApmDependencyStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQuery)
	PutApmMetricsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQuery)
	PutApmResourceStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQuery)
	PutCloudCostQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQuery)
	PutEventQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQuery)
	PutMetricQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQuery)
	PutProcessQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQuery)
	PutProductAnalyticsExtendedQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQuery)
	PutRetentionQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQuery)
	PutSloQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQuery)
	PutUserJourneyQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuery)
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

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ApmDependencyStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmDependencyStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ApmDependencyStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQuery
	_jsii_.Get(
		j,
		"apmDependencyStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ApmMetricsQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQueryOutputReference
	_jsii_.Get(
		j,
		"apmMetricsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ApmMetricsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQuery
	_jsii_.Get(
		j,
		"apmMetricsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ApmResourceStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmResourceStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ApmResourceStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"apmResourceStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) CloudCostQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQueryOutputReference
	_jsii_.Get(
		j,
		"cloudCostQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) CloudCostQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQuery
	_jsii_.Get(
		j,
		"cloudCostQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) EventQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) EventQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) MetricQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQueryOutputReference
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) MetricQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQuery
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ProcessQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ProcessQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ProductAnalyticsExtendedQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQueryOutputReference
	_jsii_.Get(
		j,
		"productAnalyticsExtendedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ProductAnalyticsExtendedQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQuery
	_jsii_.Get(
		j,
		"productAnalyticsExtendedQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) RetentionQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQueryOutputReference
	_jsii_.Get(
		j,
		"retentionQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) RetentionQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQuery
	_jsii_.Get(
		j,
		"retentionQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) SloQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQueryOutputReference
	_jsii_.Get(
		j,
		"sloQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) SloQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQuery
	_jsii_.Get(
		j,
		"sloQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) UserJourneyQuery() PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQueryOutputReference
	_jsii_.Get(
		j,
		"userJourneyQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) UserJourneyQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuery
	_jsii_.Get(
		j,
		"userJourneyQueryInput",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutApmDependencyStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQuery) {
	if err := p.validatePutApmDependencyStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmDependencyStatsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutApmMetricsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQuery) {
	if err := p.validatePutApmMetricsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmMetricsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutApmResourceStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQuery) {
	if err := p.validatePutApmResourceStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmResourceStatsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutCloudCostQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQuery) {
	if err := p.validatePutCloudCostQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCloudCostQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutEventQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQuery) {
	if err := p.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutMetricQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQuery) {
	if err := p.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutProcessQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQuery) {
	if err := p.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutProductAnalyticsExtendedQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQuery) {
	if err := p.validatePutProductAnalyticsExtendedQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProductAnalyticsExtendedQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutRetentionQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQuery) {
	if err := p.validatePutRetentionQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRetentionQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutSloQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQuery) {
	if err := p.validatePutSloQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSloQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutUserJourneyQuery(value *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuery) {
	if err := p.validatePutUserJourneyQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putUserJourneyQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetApmDependencyStatsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmDependencyStatsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetApmMetricsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmMetricsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetApmResourceStatsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmResourceStatsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetCloudCostQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetCloudCostQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetProductAnalyticsExtendedQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProductAnalyticsExtendedQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetRetentionQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetRetentionQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetSloQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetSloQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetUserJourneyQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetUserJourneyQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

