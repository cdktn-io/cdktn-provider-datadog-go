// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/observabilitypipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference interface {
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
	MaxSize() *float64
	SetMaxSize(val *float64)
	MaxSizeInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WhenFull() *string
	SetWhenFull(val *string)
	WhenFullInput() *string
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
	ResetMaxSize()
	ResetWhenFull()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference
type jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) MaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) WhenFull() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whenFull",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) WhenFullInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whenFullInput",
		&returns,
	)
	return returns
}


func NewObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference_Override(o ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference)SetMaxSize(val *float64) {
	if err := j.validateSetMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference)SetWhenFull(val *string) {
	if err := j.validateSetWhenFullParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whenFull",
		val,
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) ResetMaxSize() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxSize",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) ResetWhenFull() {
	_jsii_.InvokeVoid(
		o,
		"resetWhenFull",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationHttpClientBufferDiskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

