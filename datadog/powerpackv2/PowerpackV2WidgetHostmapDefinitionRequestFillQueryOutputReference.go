// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference interface {
	cdktn.ComplexObject
	ApmDependencyStatsQuery() PowerpackV2WidgetHostmapDefinitionRequestFillQueryApmDependencyStatsQueryOutputReference
	ApmDependencyStatsQueryInput() *PowerpackV2WidgetHostmapDefinitionRequestFillQueryApmDependencyStatsQuery
	ApmResourceStatsQuery() PowerpackV2WidgetHostmapDefinitionRequestFillQueryApmResourceStatsQueryOutputReference
	ApmResourceStatsQueryInput() *PowerpackV2WidgetHostmapDefinitionRequestFillQueryApmResourceStatsQuery
	CloudCostQuery() PowerpackV2WidgetHostmapDefinitionRequestFillQueryCloudCostQueryOutputReference
	CloudCostQueryInput() *PowerpackV2WidgetHostmapDefinitionRequestFillQueryCloudCostQuery
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
	EventQuery() PowerpackV2WidgetHostmapDefinitionRequestFillQueryEventQueryOutputReference
	EventQueryInput() *PowerpackV2WidgetHostmapDefinitionRequestFillQueryEventQuery
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricQuery() PowerpackV2WidgetHostmapDefinitionRequestFillQueryMetricQueryOutputReference
	MetricQueryInput() *PowerpackV2WidgetHostmapDefinitionRequestFillQueryMetricQuery
	ProcessQuery() PowerpackV2WidgetHostmapDefinitionRequestFillQueryProcessQueryOutputReference
	ProcessQueryInput() *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProcessQuery
	SloQuery() PowerpackV2WidgetHostmapDefinitionRequestFillQuerySloQueryOutputReference
	SloQueryInput() *PowerpackV2WidgetHostmapDefinitionRequestFillQuerySloQuery
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
	PutApmDependencyStatsQuery(value *PowerpackV2WidgetHostmapDefinitionRequestFillQueryApmDependencyStatsQuery)
	PutApmResourceStatsQuery(value *PowerpackV2WidgetHostmapDefinitionRequestFillQueryApmResourceStatsQuery)
	PutCloudCostQuery(value *PowerpackV2WidgetHostmapDefinitionRequestFillQueryCloudCostQuery)
	PutEventQuery(value *PowerpackV2WidgetHostmapDefinitionRequestFillQueryEventQuery)
	PutMetricQuery(value *PowerpackV2WidgetHostmapDefinitionRequestFillQueryMetricQuery)
	PutProcessQuery(value *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProcessQuery)
	PutSloQuery(value *PowerpackV2WidgetHostmapDefinitionRequestFillQuerySloQuery)
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

// The jsii proxy struct for PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference
type jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ApmDependencyStatsQuery() PowerpackV2WidgetHostmapDefinitionRequestFillQueryApmDependencyStatsQueryOutputReference {
	var returns PowerpackV2WidgetHostmapDefinitionRequestFillQueryApmDependencyStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmDependencyStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ApmDependencyStatsQueryInput() *PowerpackV2WidgetHostmapDefinitionRequestFillQueryApmDependencyStatsQuery {
	var returns *PowerpackV2WidgetHostmapDefinitionRequestFillQueryApmDependencyStatsQuery
	_jsii_.Get(
		j,
		"apmDependencyStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ApmResourceStatsQuery() PowerpackV2WidgetHostmapDefinitionRequestFillQueryApmResourceStatsQueryOutputReference {
	var returns PowerpackV2WidgetHostmapDefinitionRequestFillQueryApmResourceStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmResourceStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ApmResourceStatsQueryInput() *PowerpackV2WidgetHostmapDefinitionRequestFillQueryApmResourceStatsQuery {
	var returns *PowerpackV2WidgetHostmapDefinitionRequestFillQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"apmResourceStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) CloudCostQuery() PowerpackV2WidgetHostmapDefinitionRequestFillQueryCloudCostQueryOutputReference {
	var returns PowerpackV2WidgetHostmapDefinitionRequestFillQueryCloudCostQueryOutputReference
	_jsii_.Get(
		j,
		"cloudCostQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) CloudCostQueryInput() *PowerpackV2WidgetHostmapDefinitionRequestFillQueryCloudCostQuery {
	var returns *PowerpackV2WidgetHostmapDefinitionRequestFillQueryCloudCostQuery
	_jsii_.Get(
		j,
		"cloudCostQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) EventQuery() PowerpackV2WidgetHostmapDefinitionRequestFillQueryEventQueryOutputReference {
	var returns PowerpackV2WidgetHostmapDefinitionRequestFillQueryEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) EventQueryInput() *PowerpackV2WidgetHostmapDefinitionRequestFillQueryEventQuery {
	var returns *PowerpackV2WidgetHostmapDefinitionRequestFillQueryEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) MetricQuery() PowerpackV2WidgetHostmapDefinitionRequestFillQueryMetricQueryOutputReference {
	var returns PowerpackV2WidgetHostmapDefinitionRequestFillQueryMetricQueryOutputReference
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) MetricQueryInput() *PowerpackV2WidgetHostmapDefinitionRequestFillQueryMetricQuery {
	var returns *PowerpackV2WidgetHostmapDefinitionRequestFillQueryMetricQuery
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ProcessQuery() PowerpackV2WidgetHostmapDefinitionRequestFillQueryProcessQueryOutputReference {
	var returns PowerpackV2WidgetHostmapDefinitionRequestFillQueryProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ProcessQueryInput() *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProcessQuery {
	var returns *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) SloQuery() PowerpackV2WidgetHostmapDefinitionRequestFillQuerySloQueryOutputReference {
	var returns PowerpackV2WidgetHostmapDefinitionRequestFillQuerySloQueryOutputReference
	_jsii_.Get(
		j,
		"sloQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) SloQueryInput() *PowerpackV2WidgetHostmapDefinitionRequestFillQuerySloQuery {
	var returns *PowerpackV2WidgetHostmapDefinitionRequestFillQuerySloQuery
	_jsii_.Get(
		j,
		"sloQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference_Override(p PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) PutApmDependencyStatsQuery(value *PowerpackV2WidgetHostmapDefinitionRequestFillQueryApmDependencyStatsQuery) {
	if err := p.validatePutApmDependencyStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmDependencyStatsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) PutApmResourceStatsQuery(value *PowerpackV2WidgetHostmapDefinitionRequestFillQueryApmResourceStatsQuery) {
	if err := p.validatePutApmResourceStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmResourceStatsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) PutCloudCostQuery(value *PowerpackV2WidgetHostmapDefinitionRequestFillQueryCloudCostQuery) {
	if err := p.validatePutCloudCostQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCloudCostQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) PutEventQuery(value *PowerpackV2WidgetHostmapDefinitionRequestFillQueryEventQuery) {
	if err := p.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) PutMetricQuery(value *PowerpackV2WidgetHostmapDefinitionRequestFillQueryMetricQuery) {
	if err := p.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) PutProcessQuery(value *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProcessQuery) {
	if err := p.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) PutSloQuery(value *PowerpackV2WidgetHostmapDefinitionRequestFillQuerySloQuery) {
	if err := p.validatePutSloQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSloQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ResetApmDependencyStatsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmDependencyStatsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ResetApmResourceStatsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmResourceStatsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ResetCloudCostQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetCloudCostQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ResetSloQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetSloQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

