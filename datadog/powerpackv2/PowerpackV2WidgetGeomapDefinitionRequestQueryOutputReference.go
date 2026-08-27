// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference interface {
	cdktn.ComplexObject
	ApmDependencyStatsQuery() PowerpackV2WidgetGeomapDefinitionRequestQueryApmDependencyStatsQueryOutputReference
	ApmDependencyStatsQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQueryApmDependencyStatsQuery
	ApmMetricsQuery() PowerpackV2WidgetGeomapDefinitionRequestQueryApmMetricsQueryOutputReference
	ApmMetricsQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQueryApmMetricsQuery
	ApmResourceStatsQuery() PowerpackV2WidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference
	ApmResourceStatsQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQueryApmResourceStatsQuery
	CloudCostQuery() PowerpackV2WidgetGeomapDefinitionRequestQueryCloudCostQueryOutputReference
	CloudCostQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQueryCloudCostQuery
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
	EventQuery() PowerpackV2WidgetGeomapDefinitionRequestQueryEventQueryOutputReference
	EventQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQueryEventQuery
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricQuery() PowerpackV2WidgetGeomapDefinitionRequestQueryMetricQueryOutputReference
	MetricQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQueryMetricQuery
	ProcessQuery() PowerpackV2WidgetGeomapDefinitionRequestQueryProcessQueryOutputReference
	ProcessQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQueryProcessQuery
	ProductAnalyticsExtendedQuery() PowerpackV2WidgetGeomapDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference
	ProductAnalyticsExtendedQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQueryProductAnalyticsExtendedQuery
	RetentionQuery() PowerpackV2WidgetGeomapDefinitionRequestQueryRetentionQueryOutputReference
	RetentionQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQueryRetentionQuery
	SloQuery() PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference
	SloQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQuerySloQuery
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserJourneyQuery() PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQueryOutputReference
	UserJourneyQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuery
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
	PutApmDependencyStatsQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQueryApmDependencyStatsQuery)
	PutApmMetricsQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQueryApmMetricsQuery)
	PutApmResourceStatsQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQueryApmResourceStatsQuery)
	PutCloudCostQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQueryCloudCostQuery)
	PutEventQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQueryEventQuery)
	PutMetricQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQueryMetricQuery)
	PutProcessQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQueryProcessQuery)
	PutProductAnalyticsExtendedQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQueryProductAnalyticsExtendedQuery)
	PutRetentionQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQueryRetentionQuery)
	PutSloQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQuerySloQuery)
	PutUserJourneyQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuery)
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

// The jsii proxy struct for PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference
type jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ApmDependencyStatsQuery() PowerpackV2WidgetGeomapDefinitionRequestQueryApmDependencyStatsQueryOutputReference {
	var returns PowerpackV2WidgetGeomapDefinitionRequestQueryApmDependencyStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmDependencyStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ApmDependencyStatsQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQueryApmDependencyStatsQuery {
	var returns *PowerpackV2WidgetGeomapDefinitionRequestQueryApmDependencyStatsQuery
	_jsii_.Get(
		j,
		"apmDependencyStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ApmMetricsQuery() PowerpackV2WidgetGeomapDefinitionRequestQueryApmMetricsQueryOutputReference {
	var returns PowerpackV2WidgetGeomapDefinitionRequestQueryApmMetricsQueryOutputReference
	_jsii_.Get(
		j,
		"apmMetricsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ApmMetricsQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQueryApmMetricsQuery {
	var returns *PowerpackV2WidgetGeomapDefinitionRequestQueryApmMetricsQuery
	_jsii_.Get(
		j,
		"apmMetricsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ApmResourceStatsQuery() PowerpackV2WidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference {
	var returns PowerpackV2WidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmResourceStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ApmResourceStatsQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQueryApmResourceStatsQuery {
	var returns *PowerpackV2WidgetGeomapDefinitionRequestQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"apmResourceStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) CloudCostQuery() PowerpackV2WidgetGeomapDefinitionRequestQueryCloudCostQueryOutputReference {
	var returns PowerpackV2WidgetGeomapDefinitionRequestQueryCloudCostQueryOutputReference
	_jsii_.Get(
		j,
		"cloudCostQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) CloudCostQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQueryCloudCostQuery {
	var returns *PowerpackV2WidgetGeomapDefinitionRequestQueryCloudCostQuery
	_jsii_.Get(
		j,
		"cloudCostQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) EventQuery() PowerpackV2WidgetGeomapDefinitionRequestQueryEventQueryOutputReference {
	var returns PowerpackV2WidgetGeomapDefinitionRequestQueryEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) EventQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQueryEventQuery {
	var returns *PowerpackV2WidgetGeomapDefinitionRequestQueryEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) MetricQuery() PowerpackV2WidgetGeomapDefinitionRequestQueryMetricQueryOutputReference {
	var returns PowerpackV2WidgetGeomapDefinitionRequestQueryMetricQueryOutputReference
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) MetricQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQueryMetricQuery {
	var returns *PowerpackV2WidgetGeomapDefinitionRequestQueryMetricQuery
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ProcessQuery() PowerpackV2WidgetGeomapDefinitionRequestQueryProcessQueryOutputReference {
	var returns PowerpackV2WidgetGeomapDefinitionRequestQueryProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ProcessQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQueryProcessQuery {
	var returns *PowerpackV2WidgetGeomapDefinitionRequestQueryProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ProductAnalyticsExtendedQuery() PowerpackV2WidgetGeomapDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference {
	var returns PowerpackV2WidgetGeomapDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference
	_jsii_.Get(
		j,
		"productAnalyticsExtendedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ProductAnalyticsExtendedQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQueryProductAnalyticsExtendedQuery {
	var returns *PowerpackV2WidgetGeomapDefinitionRequestQueryProductAnalyticsExtendedQuery
	_jsii_.Get(
		j,
		"productAnalyticsExtendedQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) RetentionQuery() PowerpackV2WidgetGeomapDefinitionRequestQueryRetentionQueryOutputReference {
	var returns PowerpackV2WidgetGeomapDefinitionRequestQueryRetentionQueryOutputReference
	_jsii_.Get(
		j,
		"retentionQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) RetentionQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQueryRetentionQuery {
	var returns *PowerpackV2WidgetGeomapDefinitionRequestQueryRetentionQuery
	_jsii_.Get(
		j,
		"retentionQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) SloQuery() PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference {
	var returns PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference
	_jsii_.Get(
		j,
		"sloQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) SloQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQuerySloQuery {
	var returns *PowerpackV2WidgetGeomapDefinitionRequestQuerySloQuery
	_jsii_.Get(
		j,
		"sloQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) UserJourneyQuery() PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQueryOutputReference {
	var returns PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQueryOutputReference
	_jsii_.Get(
		j,
		"userJourneyQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) UserJourneyQueryInput() *PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuery {
	var returns *PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuery
	_jsii_.Get(
		j,
		"userJourneyQueryInput",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGeomapDefinitionRequestQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference_Override(p PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) PutApmDependencyStatsQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQueryApmDependencyStatsQuery) {
	if err := p.validatePutApmDependencyStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmDependencyStatsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) PutApmMetricsQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQueryApmMetricsQuery) {
	if err := p.validatePutApmMetricsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmMetricsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) PutApmResourceStatsQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQueryApmResourceStatsQuery) {
	if err := p.validatePutApmResourceStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmResourceStatsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) PutCloudCostQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQueryCloudCostQuery) {
	if err := p.validatePutCloudCostQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCloudCostQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) PutEventQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQueryEventQuery) {
	if err := p.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) PutMetricQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQueryMetricQuery) {
	if err := p.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) PutProcessQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQueryProcessQuery) {
	if err := p.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) PutProductAnalyticsExtendedQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQueryProductAnalyticsExtendedQuery) {
	if err := p.validatePutProductAnalyticsExtendedQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProductAnalyticsExtendedQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) PutRetentionQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQueryRetentionQuery) {
	if err := p.validatePutRetentionQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRetentionQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) PutSloQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQuerySloQuery) {
	if err := p.validatePutSloQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSloQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) PutUserJourneyQuery(value *PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuery) {
	if err := p.validatePutUserJourneyQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putUserJourneyQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ResetApmDependencyStatsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmDependencyStatsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ResetApmMetricsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmMetricsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ResetApmResourceStatsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmResourceStatsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ResetCloudCostQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetCloudCostQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ResetProductAnalyticsExtendedQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProductAnalyticsExtendedQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ResetRetentionQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetRetentionQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ResetSloQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetSloQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ResetUserJourneyQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetUserJourneyQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

