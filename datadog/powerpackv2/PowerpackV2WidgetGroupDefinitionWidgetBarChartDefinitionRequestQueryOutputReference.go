// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference interface {
	cdktn.ComplexObject
	ApmDependencyStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmDependencyStatsQueryOutputReference
	ApmDependencyStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmDependencyStatsQuery
	ApmMetricsQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmMetricsQueryOutputReference
	ApmMetricsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmMetricsQuery
	ApmResourceStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmResourceStatsQueryOutputReference
	ApmResourceStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmResourceStatsQuery
	CloudCostQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryCloudCostQueryOutputReference
	CloudCostQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryCloudCostQuery
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
	EventQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryEventQueryOutputReference
	EventQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryEventQuery
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryMetricQueryOutputReference
	MetricQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryMetricQuery
	ProcessQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryProcessQueryOutputReference
	ProcessQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryProcessQuery
	ProductAnalyticsExtendedQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference
	ProductAnalyticsExtendedQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQuery
	RetentionQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryRetentionQueryOutputReference
	RetentionQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryRetentionQuery
	SloQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQuerySloQueryOutputReference
	SloQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQuerySloQuery
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserJourneyQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryUserJourneyQueryOutputReference
	UserJourneyQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryUserJourneyQuery
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
	PutApmDependencyStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmDependencyStatsQuery)
	PutApmMetricsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmMetricsQuery)
	PutApmResourceStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmResourceStatsQuery)
	PutCloudCostQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryCloudCostQuery)
	PutEventQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryEventQuery)
	PutMetricQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryMetricQuery)
	PutProcessQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryProcessQuery)
	PutProductAnalyticsExtendedQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQuery)
	PutRetentionQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryRetentionQuery)
	PutSloQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQuerySloQuery)
	PutUserJourneyQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryUserJourneyQuery)
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

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ApmDependencyStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmDependencyStatsQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmDependencyStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmDependencyStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ApmDependencyStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmDependencyStatsQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmDependencyStatsQuery
	_jsii_.Get(
		j,
		"apmDependencyStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ApmMetricsQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmMetricsQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmMetricsQueryOutputReference
	_jsii_.Get(
		j,
		"apmMetricsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ApmMetricsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmMetricsQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmMetricsQuery
	_jsii_.Get(
		j,
		"apmMetricsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ApmResourceStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmResourceStatsQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmResourceStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmResourceStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ApmResourceStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmResourceStatsQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"apmResourceStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) CloudCostQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryCloudCostQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryCloudCostQueryOutputReference
	_jsii_.Get(
		j,
		"cloudCostQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) CloudCostQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryCloudCostQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryCloudCostQuery
	_jsii_.Get(
		j,
		"cloudCostQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) EventQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryEventQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) EventQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryEventQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) MetricQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryMetricQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryMetricQueryOutputReference
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) MetricQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryMetricQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryMetricQuery
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ProcessQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryProcessQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ProcessQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryProcessQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ProductAnalyticsExtendedQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference
	_jsii_.Get(
		j,
		"productAnalyticsExtendedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ProductAnalyticsExtendedQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQuery
	_jsii_.Get(
		j,
		"productAnalyticsExtendedQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) RetentionQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryRetentionQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryRetentionQueryOutputReference
	_jsii_.Get(
		j,
		"retentionQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) RetentionQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryRetentionQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryRetentionQuery
	_jsii_.Get(
		j,
		"retentionQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) SloQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQuerySloQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQuerySloQueryOutputReference
	_jsii_.Get(
		j,
		"sloQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) SloQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQuerySloQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQuerySloQuery
	_jsii_.Get(
		j,
		"sloQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) UserJourneyQuery() PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryUserJourneyQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryUserJourneyQueryOutputReference
	_jsii_.Get(
		j,
		"userJourneyQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) UserJourneyQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryUserJourneyQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryUserJourneyQuery
	_jsii_.Get(
		j,
		"userJourneyQueryInput",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) PutApmDependencyStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmDependencyStatsQuery) {
	if err := p.validatePutApmDependencyStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmDependencyStatsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) PutApmMetricsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmMetricsQuery) {
	if err := p.validatePutApmMetricsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmMetricsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) PutApmResourceStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryApmResourceStatsQuery) {
	if err := p.validatePutApmResourceStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmResourceStatsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) PutCloudCostQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryCloudCostQuery) {
	if err := p.validatePutCloudCostQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCloudCostQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) PutEventQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryEventQuery) {
	if err := p.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) PutMetricQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryMetricQuery) {
	if err := p.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) PutProcessQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryProcessQuery) {
	if err := p.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) PutProductAnalyticsExtendedQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQuery) {
	if err := p.validatePutProductAnalyticsExtendedQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProductAnalyticsExtendedQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) PutRetentionQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryRetentionQuery) {
	if err := p.validatePutRetentionQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRetentionQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) PutSloQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQuerySloQuery) {
	if err := p.validatePutSloQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSloQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) PutUserJourneyQuery(value *PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryUserJourneyQuery) {
	if err := p.validatePutUserJourneyQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putUserJourneyQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ResetApmDependencyStatsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmDependencyStatsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ResetApmMetricsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmMetricsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ResetApmResourceStatsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmResourceStatsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ResetCloudCostQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetCloudCostQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ResetProductAnalyticsExtendedQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProductAnalyticsExtendedQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ResetRetentionQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetRetentionQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ResetSloQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetSloQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ResetUserJourneyQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetUserJourneyQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetBarChartDefinitionRequestQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

