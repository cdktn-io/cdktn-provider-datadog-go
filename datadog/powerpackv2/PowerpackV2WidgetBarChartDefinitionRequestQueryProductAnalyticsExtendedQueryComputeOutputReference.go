// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference interface {
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
	InternalValue() *PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryCompute
	SetInternalValue(val *PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryCompute)
	Interval() *float64
	SetInterval(val *float64)
	IntervalInput() *float64
	Metric() *string
	SetMetric(val *string)
	MetricInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Rollup() PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeRollupOutputReference
	RollupInput() *PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeRollup
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
	PutRollup(value *PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeRollup)
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

// The jsii proxy struct for PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference
type jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) Aggregation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) AggregationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) InternalValue() *PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryCompute {
	var returns *PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryCompute
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) IntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) Metric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) MetricInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) Rollup() PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeRollupOutputReference {
	var returns PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeRollupOutputReference
	_jsii_.Get(
		j,
		"rollup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) RollupInput() *PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeRollup {
	var returns *PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeRollup
	_jsii_.Get(
		j,
		"rollupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference_Override(p PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference)SetAggregation(val *string) {
	if err := j.validateSetAggregationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregation",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference)SetInternalValue(val *PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryCompute) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference)SetInterval(val *float64) {
	if err := j.validateSetIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference)SetMetric(val *string) {
	if err := j.validateSetMetricParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metric",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) PutRollup(value *PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeRollup) {
	if err := p.validatePutRollupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRollup",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) ResetInterval() {
	_jsii_.InvokeVoid(
		p,
		"resetInterval",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) ResetMetric() {
	_jsii_.InvokeVoid(
		p,
		"resetMetric",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		p,
		"resetName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) ResetRollup() {
	_jsii_.InvokeVoid(
		p,
		"resetRollup",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetBarChartDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

