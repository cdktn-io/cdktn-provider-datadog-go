// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/observabilitypipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityPipelineConfigDestinationClickhouseOutputReference interface {
	cdktn.ComplexObject
	Auth() ObservabilityPipelineConfigDestinationClickhouseAuthList
	AuthInput() interface{}
	Batch() ObservabilityPipelineConfigDestinationClickhouseBatchList
	BatchEncoding() ObservabilityPipelineConfigDestinationClickhouseBatchEncodingList
	BatchEncodingInput() interface{}
	BatchInput() interface{}
	Buffer() ObservabilityPipelineConfigDestinationClickhouseBufferList
	BufferInput() interface{}
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
	Compression() ObservabilityPipelineConfigDestinationClickhouseCompressionList
	CompressionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	DateTimeBestEffort() interface{}
	SetDateTimeBestEffort(val interface{})
	DateTimeBestEffortInput() interface{}
	EndpointUrlKey() *string
	SetEndpointUrlKey(val *string)
	EndpointUrlKeyInput() *string
	Format() *string
	SetFormat(val *string)
	FormatInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SkipUnknownFields() interface{}
	SetSkipUnknownFields(val interface{})
	SkipUnknownFieldsInput() interface{}
	Table() *string
	SetTable(val *string)
	TableInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Tls() ObservabilityPipelineConfigDestinationClickhouseTlsList
	TlsInput() interface{}
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
	PutBatch(value interface{})
	PutBatchEncoding(value interface{})
	PutBuffer(value interface{})
	PutCompression(value interface{})
	PutTls(value interface{})
	ResetAuth()
	ResetBatch()
	ResetBatchEncoding()
	ResetBuffer()
	ResetCompression()
	ResetDatabase()
	ResetDateTimeBestEffort()
	ResetEndpointUrlKey()
	ResetFormat()
	ResetSkipUnknownFields()
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

// The jsii proxy struct for ObservabilityPipelineConfigDestinationClickhouseOutputReference
type jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) Auth() ObservabilityPipelineConfigDestinationClickhouseAuthList {
	var returns ObservabilityPipelineConfigDestinationClickhouseAuthList
	_jsii_.Get(
		j,
		"auth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) AuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) Batch() ObservabilityPipelineConfigDestinationClickhouseBatchList {
	var returns ObservabilityPipelineConfigDestinationClickhouseBatchList
	_jsii_.Get(
		j,
		"batch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) BatchEncoding() ObservabilityPipelineConfigDestinationClickhouseBatchEncodingList {
	var returns ObservabilityPipelineConfigDestinationClickhouseBatchEncodingList
	_jsii_.Get(
		j,
		"batchEncoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) BatchEncodingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchEncodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) BatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) Buffer() ObservabilityPipelineConfigDestinationClickhouseBufferList {
	var returns ObservabilityPipelineConfigDestinationClickhouseBufferList
	_jsii_.Get(
		j,
		"buffer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) BufferInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bufferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) Compression() ObservabilityPipelineConfigDestinationClickhouseCompressionList {
	var returns ObservabilityPipelineConfigDestinationClickhouseCompressionList
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) CompressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) DateTimeBestEffort() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dateTimeBestEffort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) DateTimeBestEffortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dateTimeBestEffortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) EndpointUrlKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUrlKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) EndpointUrlKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUrlKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) SkipUnknownFields() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUnknownFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) SkipUnknownFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUnknownFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) Table() *string {
	var returns *string
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) TableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) Tls() ObservabilityPipelineConfigDestinationClickhouseTlsList {
	var returns ObservabilityPipelineConfigDestinationClickhouseTlsList
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) TlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}


func NewObservabilityPipelineConfigDestinationClickhouseOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ObservabilityPipelineConfigDestinationClickhouseOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityPipelineConfigDestinationClickhouseOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigDestinationClickhouseOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewObservabilityPipelineConfigDestinationClickhouseOutputReference_Override(o ObservabilityPipelineConfigDestinationClickhouseOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigDestinationClickhouseOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference)SetDateTimeBestEffort(val interface{}) {
	if err := j.validateSetDateTimeBestEffortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateTimeBestEffort",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference)SetEndpointUrlKey(val *string) {
	if err := j.validateSetEndpointUrlKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointUrlKey",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference)SetFormat(val *string) {
	if err := j.validateSetFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference)SetSkipUnknownFields(val interface{}) {
	if err := j.validateSetSkipUnknownFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipUnknownFields",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference)SetTable(val *string) {
	if err := j.validateSetTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"table",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) PutAuth(value interface{}) {
	if err := o.validatePutAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAuth",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) PutBatch(value interface{}) {
	if err := o.validatePutBatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putBatch",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) PutBatchEncoding(value interface{}) {
	if err := o.validatePutBatchEncodingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putBatchEncoding",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) PutBuffer(value interface{}) {
	if err := o.validatePutBufferParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putBuffer",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) PutCompression(value interface{}) {
	if err := o.validatePutCompressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCompression",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) PutTls(value interface{}) {
	if err := o.validatePutTlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTls",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) ResetAuth() {
	_jsii_.InvokeVoid(
		o,
		"resetAuth",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) ResetBatch() {
	_jsii_.InvokeVoid(
		o,
		"resetBatch",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) ResetBatchEncoding() {
	_jsii_.InvokeVoid(
		o,
		"resetBatchEncoding",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) ResetBuffer() {
	_jsii_.InvokeVoid(
		o,
		"resetBuffer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) ResetCompression() {
	_jsii_.InvokeVoid(
		o,
		"resetCompression",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		o,
		"resetDatabase",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) ResetDateTimeBestEffort() {
	_jsii_.InvokeVoid(
		o,
		"resetDateTimeBestEffort",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) ResetEndpointUrlKey() {
	_jsii_.InvokeVoid(
		o,
		"resetEndpointUrlKey",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) ResetFormat() {
	_jsii_.InvokeVoid(
		o,
		"resetFormat",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) ResetSkipUnknownFields() {
	_jsii_.InvokeVoid(
		o,
		"resetSkipUnknownFields",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) ResetTls() {
	_jsii_.InvokeVoid(
		o,
		"resetTls",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationClickhouseOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

