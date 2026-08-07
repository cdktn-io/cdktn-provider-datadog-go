// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference interface {
	cdktn.ComplexObject
	ApmDependencyStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQueryOutputReference
	ApmDependencyStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQuery
	ApmResourceStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference
	ApmResourceStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery
	CloudCostQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQueryOutputReference
	CloudCostQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQuery
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
	EventQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryEventQueryOutputReference
	EventQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryEventQuery
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQueryOutputReference
	MetricQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQuery
	ProcessQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQueryOutputReference
	ProcessQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQuery
	SloQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQuerySloQueryOutputReference
	SloQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQuerySloQuery
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
	PutApmDependencyStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQuery)
	PutApmResourceStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery)
	PutCloudCostQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQuery)
	PutEventQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryEventQuery)
	PutMetricQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQuery)
	PutProcessQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQuery)
	PutSloQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQuerySloQuery)
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

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ApmDependencyStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmDependencyStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ApmDependencyStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQuery
	_jsii_.Get(
		j,
		"apmDependencyStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ApmResourceStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmResourceStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ApmResourceStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"apmResourceStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) CloudCostQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQueryOutputReference
	_jsii_.Get(
		j,
		"cloudCostQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) CloudCostQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQuery
	_jsii_.Get(
		j,
		"cloudCostQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) EventQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryEventQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) EventQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryEventQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) MetricQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQueryOutputReference
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) MetricQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQuery
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ProcessQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ProcessQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) SloQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQuerySloQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQuerySloQueryOutputReference
	_jsii_.Get(
		j,
		"sloQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) SloQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQuerySloQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQuerySloQuery
	_jsii_.Get(
		j,
		"sloQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutApmDependencyStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQuery) {
	if err := p.validatePutApmDependencyStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmDependencyStatsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutApmResourceStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery) {
	if err := p.validatePutApmResourceStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmResourceStatsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutCloudCostQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQuery) {
	if err := p.validatePutCloudCostQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCloudCostQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutEventQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryEventQuery) {
	if err := p.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutMetricQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQuery) {
	if err := p.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutProcessQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQuery) {
	if err := p.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutSloQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQuerySloQuery) {
	if err := p.validatePutSloQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSloQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetApmDependencyStatsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmDependencyStatsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetApmResourceStatsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmResourceStatsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetCloudCostQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetCloudCostQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetSloQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetSloQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

