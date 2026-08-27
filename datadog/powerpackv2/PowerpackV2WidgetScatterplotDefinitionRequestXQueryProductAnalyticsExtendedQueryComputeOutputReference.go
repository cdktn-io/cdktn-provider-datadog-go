// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference interface {
	cdktn.ComplexObject
	Aggregation() *string
	SetAggregation(val *string)
	AggregationInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() *PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryCompute
	SetInternalValue(val *PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryCompute)
	Interval() *float64
	SetInterval(val *float64)
	IntervalInput() *float64
	Metric() *string
	SetMetric(val *string)
	MetricInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Rollup() PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeRollupOutputReference
	RollupInput() *PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeRollup
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
	PutRollup(value *PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeRollup)
	ResetInterval()
	ResetMetric()
	ResetName()
	ResetRollup()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference
type jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) Aggregation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) AggregationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) InternalValue() *PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryCompute {
	var returns *PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryCompute
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) IntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) Metric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) MetricInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) Rollup() PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeRollupOutputReference {
	var returns PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeRollupOutputReference
	_jsii_.Get(
		j,
		"rollup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) RollupInput() *PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeRollup {
	var returns *PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeRollup
	_jsii_.Get(
		j,
		"rollupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference_Override(p PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference)SetAggregation(val *string) {
	if err := j.validateSetAggregationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregation",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference)SetInternalValue(val *PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryCompute) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference)SetInterval(val *float64) {
	if err := j.validateSetIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference)SetMetric(val *string) {
	if err := j.validateSetMetricParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metric",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) PutRollup(value *PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeRollup) {
	if err := p.validatePutRollupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRollup",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) ResetInterval() {
	_jsii_.InvokeVoid(
		p,
		"resetInterval",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) ResetMetric() {
	_jsii_.InvokeVoid(
		p,
		"resetMetric",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		p,
		"resetName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) ResetRollup() {
	_jsii_.InvokeVoid(
		p,
		"resetRollup",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

