// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/observabilitypipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityPipelineConfigDestinationKafkaOutputReference interface {
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
	Compression() *string
	SetCompression(val *string)
	CompressionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Encoding() *string
	SetEncoding(val *string)
	EncodingInput() *string
	// Experimental.
	Fqn() *string
	HeadersKey() *string
	SetHeadersKey(val *string)
	HeadersKeyInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyField() *string
	SetKeyField(val *string)
	KeyFieldInput() *string
	LibrdkafkaOption() ObservabilityPipelineConfigDestinationKafkaLibrdkafkaOptionList
	LibrdkafkaOptionInput() interface{}
	MessageTimeoutMs() *float64
	SetMessageTimeoutMs(val *float64)
	MessageTimeoutMsInput() *float64
	RateLimitDurationSecs() *float64
	SetRateLimitDurationSecs(val *float64)
	RateLimitDurationSecsInput() *float64
	RateLimitNum() *float64
	SetRateLimitNum(val *float64)
	RateLimitNumInput() *float64
	Sasl() ObservabilityPipelineConfigDestinationKafkaSaslList
	SaslInput() interface{}
	SocketTimeoutMs() *float64
	SetSocketTimeoutMs(val *float64)
	SocketTimeoutMsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Tls() ObservabilityPipelineConfigDestinationKafkaTlsList
	TlsInput() interface{}
	Topic() *string
	SetTopic(val *string)
	TopicInput() *string
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
	PutLibrdkafkaOption(value interface{})
	PutSasl(value interface{})
	PutTls(value interface{})
	ResetCompression()
	ResetHeadersKey()
	ResetKeyField()
	ResetLibrdkafkaOption()
	ResetMessageTimeoutMs()
	ResetRateLimitDurationSecs()
	ResetRateLimitNum()
	ResetSasl()
	ResetSocketTimeoutMs()
	ResetTls()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityPipelineConfigDestinationKafkaOutputReference
type jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) CompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) Encoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) EncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) HeadersKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headersKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) HeadersKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headersKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) KeyField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) KeyFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) LibrdkafkaOption() ObservabilityPipelineConfigDestinationKafkaLibrdkafkaOptionList {
	var returns ObservabilityPipelineConfigDestinationKafkaLibrdkafkaOptionList
	_jsii_.Get(
		j,
		"librdkafkaOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) LibrdkafkaOptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"librdkafkaOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) MessageTimeoutMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messageTimeoutMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) MessageTimeoutMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messageTimeoutMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) RateLimitDurationSecs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rateLimitDurationSecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) RateLimitDurationSecsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rateLimitDurationSecsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) RateLimitNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rateLimitNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) RateLimitNumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rateLimitNumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) Sasl() ObservabilityPipelineConfigDestinationKafkaSaslList {
	var returns ObservabilityPipelineConfigDestinationKafkaSaslList
	_jsii_.Get(
		j,
		"sasl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) SaslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"saslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) SocketTimeoutMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"socketTimeoutMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) SocketTimeoutMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"socketTimeoutMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) Tls() ObservabilityPipelineConfigDestinationKafkaTlsList {
	var returns ObservabilityPipelineConfigDestinationKafkaTlsList
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) TlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) Topic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) TopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicInput",
		&returns,
	)
	return returns
}


func NewObservabilityPipelineConfigDestinationKafkaOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ObservabilityPipelineConfigDestinationKafkaOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityPipelineConfigDestinationKafkaOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigDestinationKafkaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewObservabilityPipelineConfigDestinationKafkaOutputReference_Override(o ObservabilityPipelineConfigDestinationKafkaOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigDestinationKafkaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference)SetCompression(val *string) {
	if err := j.validateSetCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compression",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference)SetEncoding(val *string) {
	if err := j.validateSetEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encoding",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference)SetHeadersKey(val *string) {
	if err := j.validateSetHeadersKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headersKey",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference)SetKeyField(val *string) {
	if err := j.validateSetKeyFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyField",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference)SetMessageTimeoutMs(val *float64) {
	if err := j.validateSetMessageTimeoutMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageTimeoutMs",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference)SetRateLimitDurationSecs(val *float64) {
	if err := j.validateSetRateLimitDurationSecsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rateLimitDurationSecs",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference)SetRateLimitNum(val *float64) {
	if err := j.validateSetRateLimitNumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rateLimitNum",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference)SetSocketTimeoutMs(val *float64) {
	if err := j.validateSetSocketTimeoutMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"socketTimeoutMs",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference)SetTopic(val *string) {
	if err := j.validateSetTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topic",
		val,
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) PutLibrdkafkaOption(value interface{}) {
	if err := o.validatePutLibrdkafkaOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLibrdkafkaOption",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) PutSasl(value interface{}) {
	if err := o.validatePutSaslParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSasl",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) PutTls(value interface{}) {
	if err := o.validatePutTlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTls",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) ResetCompression() {
	_jsii_.InvokeVoid(
		o,
		"resetCompression",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) ResetHeadersKey() {
	_jsii_.InvokeVoid(
		o,
		"resetHeadersKey",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) ResetKeyField() {
	_jsii_.InvokeVoid(
		o,
		"resetKeyField",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) ResetLibrdkafkaOption() {
	_jsii_.InvokeVoid(
		o,
		"resetLibrdkafkaOption",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) ResetMessageTimeoutMs() {
	_jsii_.InvokeVoid(
		o,
		"resetMessageTimeoutMs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) ResetRateLimitDurationSecs() {
	_jsii_.InvokeVoid(
		o,
		"resetRateLimitDurationSecs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) ResetRateLimitNum() {
	_jsii_.InvokeVoid(
		o,
		"resetRateLimitNum",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) ResetSasl() {
	_jsii_.InvokeVoid(
		o,
		"resetSasl",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) ResetSocketTimeoutMs() {
	_jsii_.InvokeVoid(
		o,
		"resetSocketTimeoutMs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) ResetTls() {
	_jsii_.InvokeVoid(
		o,
		"resetTls",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationKafkaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

