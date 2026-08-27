// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference interface {
	cdktn.ComplexObject
	ApmDependencyStatsQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQueryOutputReference
	ApmDependencyStatsQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQuery
	ApmMetricsQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQueryOutputReference
	ApmMetricsQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQuery
	ApmResourceStatsQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQueryOutputReference
	ApmResourceStatsQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQuery
	CloudCostQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQueryOutputReference
	CloudCostQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQuery
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
	EventQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQueryOutputReference
	EventQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQuery
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQueryOutputReference
	MetricQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQuery
	ProcessQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQueryOutputReference
	ProcessQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQuery
	ProductAnalyticsExtendedQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQueryOutputReference
	ProductAnalyticsExtendedQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQuery
	RetentionQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQueryOutputReference
	RetentionQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQuery
	SloQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQueryOutputReference
	SloQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQuery
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserJourneyQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQueryOutputReference
	UserJourneyQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuery
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
	PutApmDependencyStatsQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQuery)
	PutApmMetricsQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQuery)
	PutApmResourceStatsQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQuery)
	PutCloudCostQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQuery)
	PutEventQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQuery)
	PutMetricQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQuery)
	PutProcessQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQuery)
	PutProductAnalyticsExtendedQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQuery)
	PutRetentionQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQuery)
	PutSloQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQuery)
	PutUserJourneyQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuery)
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

// The jsii proxy struct for PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference
type jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ApmDependencyStatsQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQueryOutputReference {
	var returns PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmDependencyStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ApmDependencyStatsQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQuery {
	var returns *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQuery
	_jsii_.Get(
		j,
		"apmDependencyStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ApmMetricsQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQueryOutputReference {
	var returns PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQueryOutputReference
	_jsii_.Get(
		j,
		"apmMetricsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ApmMetricsQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQuery {
	var returns *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQuery
	_jsii_.Get(
		j,
		"apmMetricsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ApmResourceStatsQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQueryOutputReference {
	var returns PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmResourceStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ApmResourceStatsQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQuery {
	var returns *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"apmResourceStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) CloudCostQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQueryOutputReference {
	var returns PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQueryOutputReference
	_jsii_.Get(
		j,
		"cloudCostQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) CloudCostQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQuery {
	var returns *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQuery
	_jsii_.Get(
		j,
		"cloudCostQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) EventQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQueryOutputReference {
	var returns PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) EventQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQuery {
	var returns *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) MetricQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQueryOutputReference {
	var returns PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQueryOutputReference
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) MetricQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQuery {
	var returns *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQuery
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ProcessQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQueryOutputReference {
	var returns PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ProcessQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQuery {
	var returns *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ProductAnalyticsExtendedQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQueryOutputReference {
	var returns PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQueryOutputReference
	_jsii_.Get(
		j,
		"productAnalyticsExtendedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ProductAnalyticsExtendedQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQuery {
	var returns *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQuery
	_jsii_.Get(
		j,
		"productAnalyticsExtendedQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) RetentionQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQueryOutputReference {
	var returns PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQueryOutputReference
	_jsii_.Get(
		j,
		"retentionQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) RetentionQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQuery {
	var returns *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQuery
	_jsii_.Get(
		j,
		"retentionQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) SloQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQueryOutputReference {
	var returns PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQueryOutputReference
	_jsii_.Get(
		j,
		"sloQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) SloQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQuery {
	var returns *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQuery
	_jsii_.Get(
		j,
		"sloQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) UserJourneyQuery() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQueryOutputReference {
	var returns PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQueryOutputReference
	_jsii_.Get(
		j,
		"userJourneyQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) UserJourneyQueryInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuery {
	var returns *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuery
	_jsii_.Get(
		j,
		"userJourneyQueryInput",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference_Override(p PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutApmDependencyStatsQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmDependencyStatsQuery) {
	if err := p.validatePutApmDependencyStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmDependencyStatsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutApmMetricsQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmMetricsQuery) {
	if err := p.validatePutApmMetricsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmMetricsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutApmResourceStatsQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryApmResourceStatsQuery) {
	if err := p.validatePutApmResourceStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmResourceStatsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutCloudCostQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryCloudCostQuery) {
	if err := p.validatePutCloudCostQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCloudCostQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutEventQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQuery) {
	if err := p.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutMetricQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryMetricQuery) {
	if err := p.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutProcessQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryProcessQuery) {
	if err := p.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutProductAnalyticsExtendedQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryProductAnalyticsExtendedQuery) {
	if err := p.validatePutProductAnalyticsExtendedQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProductAnalyticsExtendedQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutRetentionQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryRetentionQuery) {
	if err := p.validatePutRetentionQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRetentionQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutSloQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQuerySloQuery) {
	if err := p.validatePutSloQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSloQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) PutUserJourneyQuery(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuery) {
	if err := p.validatePutUserJourneyQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putUserJourneyQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetApmDependencyStatsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmDependencyStatsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetApmMetricsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmMetricsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetApmResourceStatsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmResourceStatsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetCloudCostQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetCloudCostQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetProductAnalyticsExtendedQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProductAnalyticsExtendedQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetRetentionQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetRetentionQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetSloQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetSloQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ResetUserJourneyQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetUserJourneyQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

