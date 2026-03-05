// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/observabilitypipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityPipelineConfigDestinationElasticsearchOutputReference interface {
	cdktn.ComplexObject
	ApiVersion() *string
	SetApiVersion(val *string)
	ApiVersionInput() *string
	Auth() ObservabilityPipelineConfigDestinationElasticsearchAuthList
	AuthInput() interface{}
	Buffer() ObservabilityPipelineConfigDestinationElasticsearchBufferList
	BufferInput() interface{}
	BulkIndex() *string
	SetBulkIndex(val *string)
	BulkIndexInput() *string
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
	DataStream() ObservabilityPipelineConfigDestinationElasticsearchDataStreamList
	DataStreamInput() interface{}
	EndpointUrlKey() *string
	SetEndpointUrlKey(val *string)
	EndpointUrlKeyInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutAuth(value interface{})
	PutBuffer(value interface{})
	PutDataStream(value interface{})
	ResetApiVersion()
	ResetAuth()
	ResetBuffer()
	ResetBulkIndex()
	ResetDataStream()
	ResetEndpointUrlKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityPipelineConfigDestinationElasticsearchOutputReference
type jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) ApiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) Auth() ObservabilityPipelineConfigDestinationElasticsearchAuthList {
	var returns ObservabilityPipelineConfigDestinationElasticsearchAuthList
	_jsii_.Get(
		j,
		"auth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) AuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) Buffer() ObservabilityPipelineConfigDestinationElasticsearchBufferList {
	var returns ObservabilityPipelineConfigDestinationElasticsearchBufferList
	_jsii_.Get(
		j,
		"buffer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) BufferInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bufferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) BulkIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bulkIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) BulkIndexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bulkIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) DataStream() ObservabilityPipelineConfigDestinationElasticsearchDataStreamList {
	var returns ObservabilityPipelineConfigDestinationElasticsearchDataStreamList
	_jsii_.Get(
		j,
		"dataStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) DataStreamInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataStreamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) EndpointUrlKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUrlKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) EndpointUrlKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUrlKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewObservabilityPipelineConfigDestinationElasticsearchOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ObservabilityPipelineConfigDestinationElasticsearchOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityPipelineConfigDestinationElasticsearchOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigDestinationElasticsearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewObservabilityPipelineConfigDestinationElasticsearchOutputReference_Override(o ObservabilityPipelineConfigDestinationElasticsearchOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigDestinationElasticsearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference)SetApiVersion(val *string) {
	if err := j.validateSetApiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiVersion",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference)SetBulkIndex(val *string) {
	if err := j.validateSetBulkIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bulkIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference)SetEndpointUrlKey(val *string) {
	if err := j.validateSetEndpointUrlKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointUrlKey",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) PutAuth(value interface{}) {
	if err := o.validatePutAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAuth",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) PutBuffer(value interface{}) {
	if err := o.validatePutBufferParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putBuffer",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) PutDataStream(value interface{}) {
	if err := o.validatePutDataStreamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDataStream",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) ResetApiVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetApiVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) ResetAuth() {
	_jsii_.InvokeVoid(
		o,
		"resetAuth",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) ResetBuffer() {
	_jsii_.InvokeVoid(
		o,
		"resetBuffer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) ResetBulkIndex() {
	_jsii_.InvokeVoid(
		o,
		"resetBulkIndex",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) ResetDataStream() {
	_jsii_.InvokeVoid(
		o,
		"resetDataStream",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) ResetEndpointUrlKey() {
	_jsii_.InvokeVoid(
		o,
		"resetEndpointUrlKey",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationElasticsearchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

