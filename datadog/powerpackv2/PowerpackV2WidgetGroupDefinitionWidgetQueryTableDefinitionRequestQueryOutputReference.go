// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference interface {
	cdktn.ComplexObject
	ApmDependencyStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmDependencyStatsQueryOutputReference
	ApmDependencyStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmDependencyStatsQuery
	ApmMetricsQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmMetricsQueryOutputReference
	ApmMetricsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmMetricsQuery
	ApmResourceStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference
	ApmResourceStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQuery
	CloudCostQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryCloudCostQueryOutputReference
	CloudCostQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryCloudCostQuery
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
	EventQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryEventQueryOutputReference
	EventQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryEventQuery
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryMetricQueryOutputReference
	MetricQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryMetricQuery
	ProcessQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryProcessQueryOutputReference
	ProcessQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryProcessQuery
	ProductAnalyticsExtendedQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference
	ProductAnalyticsExtendedQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQuery
	RetentionQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryRetentionQueryOutputReference
	RetentionQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryRetentionQuery
	SloQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQuerySloQueryOutputReference
	SloQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQuerySloQuery
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserJourneyQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryUserJourneyQueryOutputReference
	UserJourneyQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryUserJourneyQuery
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
	PutApmDependencyStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmDependencyStatsQuery)
	PutApmMetricsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmMetricsQuery)
	PutApmResourceStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQuery)
	PutCloudCostQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryCloudCostQuery)
	PutEventQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryEventQuery)
	PutMetricQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryMetricQuery)
	PutProcessQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryProcessQuery)
	PutProductAnalyticsExtendedQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQuery)
	PutRetentionQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryRetentionQuery)
	PutSloQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQuerySloQuery)
	PutUserJourneyQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryUserJourneyQuery)
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

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ApmDependencyStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmDependencyStatsQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmDependencyStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmDependencyStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ApmDependencyStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmDependencyStatsQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmDependencyStatsQuery
	_jsii_.Get(
		j,
		"apmDependencyStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ApmMetricsQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmMetricsQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmMetricsQueryOutputReference
	_jsii_.Get(
		j,
		"apmMetricsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ApmMetricsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmMetricsQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmMetricsQuery
	_jsii_.Get(
		j,
		"apmMetricsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ApmResourceStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmResourceStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ApmResourceStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"apmResourceStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) CloudCostQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryCloudCostQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryCloudCostQueryOutputReference
	_jsii_.Get(
		j,
		"cloudCostQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) CloudCostQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryCloudCostQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryCloudCostQuery
	_jsii_.Get(
		j,
		"cloudCostQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) EventQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryEventQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) EventQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryEventQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) MetricQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryMetricQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryMetricQueryOutputReference
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) MetricQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryMetricQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryMetricQuery
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ProcessQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryProcessQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ProcessQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryProcessQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ProductAnalyticsExtendedQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference
	_jsii_.Get(
		j,
		"productAnalyticsExtendedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ProductAnalyticsExtendedQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQuery
	_jsii_.Get(
		j,
		"productAnalyticsExtendedQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) RetentionQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryRetentionQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryRetentionQueryOutputReference
	_jsii_.Get(
		j,
		"retentionQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) RetentionQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryRetentionQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryRetentionQuery
	_jsii_.Get(
		j,
		"retentionQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) SloQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQuerySloQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQuerySloQueryOutputReference
	_jsii_.Get(
		j,
		"sloQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) SloQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQuerySloQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQuerySloQuery
	_jsii_.Get(
		j,
		"sloQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) UserJourneyQuery() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryUserJourneyQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryUserJourneyQueryOutputReference
	_jsii_.Get(
		j,
		"userJourneyQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) UserJourneyQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryUserJourneyQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryUserJourneyQuery
	_jsii_.Get(
		j,
		"userJourneyQueryInput",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) PutApmDependencyStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmDependencyStatsQuery) {
	if err := p.validatePutApmDependencyStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmDependencyStatsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) PutApmMetricsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmMetricsQuery) {
	if err := p.validatePutApmMetricsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmMetricsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) PutApmResourceStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQuery) {
	if err := p.validatePutApmResourceStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmResourceStatsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) PutCloudCostQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryCloudCostQuery) {
	if err := p.validatePutCloudCostQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCloudCostQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) PutEventQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryEventQuery) {
	if err := p.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) PutMetricQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryMetricQuery) {
	if err := p.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) PutProcessQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryProcessQuery) {
	if err := p.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) PutProductAnalyticsExtendedQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQuery) {
	if err := p.validatePutProductAnalyticsExtendedQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProductAnalyticsExtendedQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) PutRetentionQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryRetentionQuery) {
	if err := p.validatePutRetentionQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRetentionQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) PutSloQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQuerySloQuery) {
	if err := p.validatePutSloQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSloQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) PutUserJourneyQuery(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryUserJourneyQuery) {
	if err := p.validatePutUserJourneyQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putUserJourneyQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ResetApmDependencyStatsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmDependencyStatsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ResetApmMetricsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmMetricsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ResetApmResourceStatsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmResourceStatsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ResetCloudCostQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetCloudCostQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ResetProductAnalyticsExtendedQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProductAnalyticsExtendedQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ResetRetentionQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetRetentionQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ResetSloQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetSloQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ResetUserJourneyQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetUserJourneyQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

