// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference interface {
	cdktn.ComplexObject
	ApmDependencyStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryApmDependencyStatsQueryOutputReference
	ApmDependencyStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryApmDependencyStatsQuery
	ApmResourceStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryApmResourceStatsQueryOutputReference
	ApmResourceStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryApmResourceStatsQuery
	CloudCostQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryCloudCostQueryOutputReference
	CloudCostQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryCloudCostQuery
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
	EventQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryEventQueryOutputReference
	EventQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryEventQuery
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryMetricQueryOutputReference
	MetricQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryMetricQuery
	ProcessQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryProcessQueryOutputReference
	ProcessQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryProcessQuery
	SloQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQuerySloQueryOutputReference
	SloQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQuerySloQuery
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
	PutApmDependencyStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryApmDependencyStatsQuery)
	PutApmResourceStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryApmResourceStatsQuery)
	PutCloudCostQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryCloudCostQuery)
	PutEventQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryEventQuery)
	PutMetricQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryMetricQuery)
	PutProcessQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryProcessQuery)
	PutSloQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQuerySloQuery)
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

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) ApmDependencyStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryApmDependencyStatsQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryApmDependencyStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmDependencyStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) ApmDependencyStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryApmDependencyStatsQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryApmDependencyStatsQuery
	_jsii_.Get(
		j,
		"apmDependencyStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) ApmResourceStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryApmResourceStatsQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryApmResourceStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmResourceStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) ApmResourceStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryApmResourceStatsQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"apmResourceStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) CloudCostQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryCloudCostQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryCloudCostQueryOutputReference
	_jsii_.Get(
		j,
		"cloudCostQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) CloudCostQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryCloudCostQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryCloudCostQuery
	_jsii_.Get(
		j,
		"cloudCostQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) EventQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryEventQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) EventQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryEventQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) MetricQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryMetricQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryMetricQueryOutputReference
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) MetricQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryMetricQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryMetricQuery
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) ProcessQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryProcessQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) ProcessQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryProcessQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) SloQuery() PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQuerySloQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQuerySloQueryOutputReference
	_jsii_.Get(
		j,
		"sloQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) SloQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQuerySloQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQuerySloQuery
	_jsii_.Get(
		j,
		"sloQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) PutApmDependencyStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryApmDependencyStatsQuery) {
	if err := p.validatePutApmDependencyStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmDependencyStatsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) PutApmResourceStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryApmResourceStatsQuery) {
	if err := p.validatePutApmResourceStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmResourceStatsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) PutCloudCostQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryCloudCostQuery) {
	if err := p.validatePutCloudCostQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCloudCostQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) PutEventQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryEventQuery) {
	if err := p.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) PutMetricQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryMetricQuery) {
	if err := p.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) PutProcessQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryProcessQuery) {
	if err := p.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) PutSloQuery(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQuerySloQuery) {
	if err := p.validatePutSloQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSloQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) ResetApmDependencyStatsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmDependencyStatsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) ResetApmResourceStatsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmResourceStatsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) ResetCloudCostQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetCloudCostQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) ResetSloQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetSloQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestXQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

