// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/observabilitypipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference interface {
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
	PerMetricLimit() ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitList
	PerMetricLimitInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrackingMode() ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitTrackingModeList
	TrackingModeInput() interface{}
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
	PutPerMetricLimit(value interface{})
	PutTrackingMode(value interface{})
	ResetPerMetricLimit()
	ResetTrackingMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference
type jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) LimitExceededAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitExceededAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) LimitExceededActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitExceededActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) PerMetricLimit() ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitPerMetricLimitList
	_jsii_.Get(
		j,
		"perMetricLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) PerMetricLimitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"perMetricLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) TrackingMode() ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitTrackingModeList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitTrackingModeList
	_jsii_.Get(
		j,
		"trackingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) TrackingModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trackingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) ValueLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) ValueLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueLimitInput",
		&returns,
	)
	return returns
}


func NewObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference_Override(o ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference)SetLimitExceededAction(val *string) {
	if err := j.validateSetLimitExceededActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limitExceededAction",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference)SetValueLimit(val *float64) {
	if err := j.validateSetValueLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueLimit",
		val,
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) PutPerMetricLimit(value interface{}) {
	if err := o.validatePutPerMetricLimitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putPerMetricLimit",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) PutTrackingMode(value interface{}) {
	if err := o.validatePutTrackingModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTrackingMode",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) ResetPerMetricLimit() {
	_jsii_.InvokeVoid(
		o,
		"resetPerMetricLimit",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) ResetTrackingMode() {
	_jsii_.InvokeVoid(
		o,
		"resetTrackingMode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorTagCardinalityLimitOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

