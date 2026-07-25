// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/observabilitypipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LimitExceededAction() *string
	SetLimitExceededAction(val *string)
	LimitExceededActionInput() *string
	MetricName() *string
	SetMetricName(val *string)
	MetricNameInput() *string
	OverrideType() *string
	SetOverrideType(val *string)
	OverrideTypeInput() *string
	PerTagLimit() ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitPerTagLimitList
	PerTagLimitInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ValueLimit() *float64
	SetValueLimit(val *float64)
	ValueLimitInput() *float64
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
	PutPerTagLimit(value interface{})
	ResetLimitExceededAction()
	ResetPerTagLimit()
	ResetValueLimit()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference
type jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) LimitExceededAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitExceededAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) LimitExceededActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitExceededActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) MetricNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) OverrideType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overrideType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) OverrideTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overrideTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) PerTagLimit() ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitPerTagLimitList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitPerTagLimitList
	_jsii_.Get(
		j,
		"perTagLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) PerTagLimitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"perTagLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) ValueLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) ValueLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueLimitInput",
		&returns,
	)
	return returns
}


func NewObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference_Override(o ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference)SetLimitExceededAction(val *string) {
	if err := j.validateSetLimitExceededActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limitExceededAction",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference)SetMetricName(val *string) {
	if err := j.validateSetMetricNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference)SetOverrideType(val *string) {
	if err := j.validateSetOverrideTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overrideType",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference)SetValueLimit(val *float64) {
	if err := j.validateSetValueLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueLimit",
		val,
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) PutPerTagLimit(value interface{}) {
	if err := o.validatePutPerTagLimitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putPerTagLimit",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) ResetLimitExceededAction() {
	_jsii_.InvokeVoid(
		o,
		"resetLimitExceededAction",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) ResetPerTagLimit() {
	_jsii_.InvokeVoid(
		o,
		"resetPerTagLimit",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) ResetValueLimit() {
	_jsii_.InvokeVoid(
		o,
		"resetValueLimit",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

