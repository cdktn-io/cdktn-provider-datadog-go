// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference interface {
	cdktn.ComplexObject
	ApmDependencyStatsQuery() PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQueryOutputReference
	ApmDependencyStatsQueryInput() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQuery
	ApmResourceStatsQuery() PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference
	ApmResourceStatsQueryInput() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery
	CloudCostQuery() PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQueryOutputReference
	CloudCostQueryInput() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQuery
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
	EventQuery() PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryEventQueryOutputReference
	EventQueryInput() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryEventQuery
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricQuery() PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQueryOutputReference
	MetricQueryInput() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQuery
	ProcessQuery() PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQueryOutputReference
	ProcessQueryInput() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQuery
	SloQuery() PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQuerySloQueryOutputReference
	SloQueryInput() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQuerySloQuery
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
	PutApmDependencyStatsQuery(value *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQuery)
	PutApmResourceStatsQuery(value *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery)
	PutCloudCostQuery(value *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQuery)
	PutEventQuery(value *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryEventQuery)
	PutMetricQuery(value *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQuery)
	PutProcessQuery(value *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQuery)
	PutSloQuery(value *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQuerySloQuery)
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

// The jsii proxy struct for PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference
type jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ApmDependencyStatsQuery() PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQueryOutputReference {
	var returns PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmDependencyStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ApmDependencyStatsQueryInput() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQuery {
	var returns *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQuery
	_jsii_.Get(
		j,
		"apmDependencyStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ApmResourceStatsQuery() PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference {
	var returns PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmResourceStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ApmResourceStatsQueryInput() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery {
	var returns *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"apmResourceStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) CloudCostQuery() PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQueryOutputReference {
	var returns PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQueryOutputReference
	_jsii_.Get(
		j,
		"cloudCostQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) CloudCostQueryInput() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQuery {
	var returns *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQuery
	_jsii_.Get(
		j,
		"cloudCostQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) EventQuery() PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryEventQueryOutputReference {
	var returns PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) EventQueryInput() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryEventQuery {
	var returns *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) MetricQuery() PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQueryOutputReference {
	var returns PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQueryOutputReference
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) MetricQueryInput() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQuery {
	var returns *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQuery
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ProcessQuery() PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQueryOutputReference {
	var returns PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ProcessQueryInput() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQuery {
	var returns *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) SloQuery() PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQuerySloQueryOutputReference {
	var returns PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQuerySloQueryOutputReference
	_jsii_.Get(
		j,
		"sloQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) SloQueryInput() *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQuerySloQuery {
	var returns *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQuerySloQuery
	_jsii_.Get(
		j,
		"sloQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference_Override(p PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutApmDependencyStatsQuery(value *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQuery) {
	if err := p.validatePutApmDependencyStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmDependencyStatsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutApmResourceStatsQuery(value *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery) {
	if err := p.validatePutApmResourceStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmResourceStatsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutCloudCostQuery(value *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryCloudCostQuery) {
	if err := p.validatePutCloudCostQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCloudCostQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutEventQuery(value *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryEventQuery) {
	if err := p.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutMetricQuery(value *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQuery) {
	if err := p.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutProcessQuery(value *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQuery) {
	if err := p.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) PutSloQuery(value *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQuerySloQuery) {
	if err := p.validatePutSloQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSloQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetApmDependencyStatsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmDependencyStatsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetApmResourceStatsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmResourceStatsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetCloudCostQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetCloudCostQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ResetSloQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetSloQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

