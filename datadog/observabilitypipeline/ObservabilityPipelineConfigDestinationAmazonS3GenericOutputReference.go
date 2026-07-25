// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/observabilitypipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference interface {
	cdktn.ComplexObject
	Auth() ObservabilityPipelineConfigDestinationAmazonS3GenericAuthList
	AuthInput() interface{}
	BatchSettings() ObservabilityPipelineConfigDestinationAmazonS3GenericBatchSettingsList
	BatchSettingsInput() interface{}
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
	Buffer() ObservabilityPipelineConfigDestinationAmazonS3GenericBufferList
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
	Compression() ObservabilityPipelineConfigDestinationAmazonS3GenericCompressionList
	CompressionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Encoding() ObservabilityPipelineConfigDestinationAmazonS3GenericEncodingList
	EncodingInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyPrefix() *string
	SetKeyPrefix(val *string)
	KeyPrefixInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ServerSideEncryption() *string
	SetServerSideEncryption(val *string)
	ServerSideEncryptionInput() *string
	SsekmsKeyId() *string
	SetSsekmsKeyId(val *string)
	SsekmsKeyIdInput() *string
	StorageClass() *string
	SetStorageClass(val *string)
	StorageClassInput() *string
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
	PutBatchSettings(value interface{})
	PutBuffer(value interface{})
	PutCompression(value interface{})
	PutEncoding(value interface{})
	ResetAuth()
	ResetBatchSettings()
	ResetBuffer()
	ResetCompression()
	ResetEncoding()
	ResetKeyPrefix()
	ResetServerSideEncryption()
	ResetSsekmsKeyId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference
type jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) Auth() ObservabilityPipelineConfigDestinationAmazonS3GenericAuthList {
	var returns ObservabilityPipelineConfigDestinationAmazonS3GenericAuthList
	_jsii_.Get(
		j,
		"auth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) AuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) BatchSettings() ObservabilityPipelineConfigDestinationAmazonS3GenericBatchSettingsList {
	var returns ObservabilityPipelineConfigDestinationAmazonS3GenericBatchSettingsList
	_jsii_.Get(
		j,
		"batchSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) BatchSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) Buffer() ObservabilityPipelineConfigDestinationAmazonS3GenericBufferList {
	var returns ObservabilityPipelineConfigDestinationAmazonS3GenericBufferList
	_jsii_.Get(
		j,
		"buffer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) BufferInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bufferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) Compression() ObservabilityPipelineConfigDestinationAmazonS3GenericCompressionList {
	var returns ObservabilityPipelineConfigDestinationAmazonS3GenericCompressionList
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) CompressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) Encoding() ObservabilityPipelineConfigDestinationAmazonS3GenericEncodingList {
	var returns ObservabilityPipelineConfigDestinationAmazonS3GenericEncodingList
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) EncodingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) KeyPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) ServerSideEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) ServerSideEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) SsekmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssekmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) SsekmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssekmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) StorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityPipelineConfigDestinationAmazonS3GenericOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference_Override(o ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference)SetKeyPrefix(val *string) {
	if err := j.validateSetKeyPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyPrefix",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference)SetServerSideEncryption(val *string) {
	if err := j.validateSetServerSideEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverSideEncryption",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference)SetSsekmsKeyId(val *string) {
	if err := j.validateSetSsekmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssekmsKeyId",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference)SetStorageClass(val *string) {
	if err := j.validateSetStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageClass",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) PutAuth(value interface{}) {
	if err := o.validatePutAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAuth",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) PutBatchSettings(value interface{}) {
	if err := o.validatePutBatchSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putBatchSettings",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) PutBuffer(value interface{}) {
	if err := o.validatePutBufferParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putBuffer",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) PutCompression(value interface{}) {
	if err := o.validatePutCompressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCompression",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) PutEncoding(value interface{}) {
	if err := o.validatePutEncodingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putEncoding",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) ResetAuth() {
	_jsii_.InvokeVoid(
		o,
		"resetAuth",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) ResetBatchSettings() {
	_jsii_.InvokeVoid(
		o,
		"resetBatchSettings",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) ResetBuffer() {
	_jsii_.InvokeVoid(
		o,
		"resetBuffer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) ResetCompression() {
	_jsii_.InvokeVoid(
		o,
		"resetCompression",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) ResetEncoding() {
	_jsii_.InvokeVoid(
		o,
		"resetEncoding",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) ResetKeyPrefix() {
	_jsii_.InvokeVoid(
		o,
		"resetKeyPrefix",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) ResetServerSideEncryption() {
	_jsii_.InvokeVoid(
		o,
		"resetServerSideEncryption",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) ResetSsekmsKeyId() {
	_jsii_.InvokeVoid(
		o,
		"resetSsekmsKeyId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationAmazonS3GenericOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

