// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package referencetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/referencetable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference interface {
	cdktn.ComplexObject
	AzureClientId() *string
	SetAzureClientId(val *string)
	AzureClientIdInput() *string
	AzureContainerName() *string
	SetAzureContainerName(val *string)
	AzureContainerNameInput() *string
	AzureStorageAccountName() *string
	SetAzureStorageAccountName(val *string)
	AzureStorageAccountNameInput() *string
	AzureTenantId() *string
	SetAzureTenantId(val *string)
	AzureTenantIdInput() *string
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
	FilePath() *string
	SetFilePath(val *string)
	FilePathInput() *string
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
	ResetAzureClientId()
	ResetAzureContainerName()
	ResetAzureStorageAccountName()
	ResetAzureTenantId()
	ResetFilePath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference
type jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) AzureClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) AzureClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) AzureContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureContainerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) AzureContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureContainerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) AzureStorageAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureStorageAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) AzureStorageAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureStorageAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) AzureTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) AzureTenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureTenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) FilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) FilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference {
	_init_.Initialize()

	if err := validateNewReferenceTableFileMetadataAccessDetailsAzureDetailOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.referenceTable.ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference_Override(r ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.referenceTable.ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference)SetAzureClientId(val *string) {
	if err := j.validateSetAzureClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureClientId",
		val,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference)SetAzureContainerName(val *string) {
	if err := j.validateSetAzureContainerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureContainerName",
		val,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference)SetAzureStorageAccountName(val *string) {
	if err := j.validateSetAzureStorageAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureStorageAccountName",
		val,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference)SetAzureTenantId(val *string) {
	if err := j.validateSetAzureTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureTenantId",
		val,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference)SetFilePath(val *string) {
	if err := j.validateSetFilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filePath",
		val,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) ResetAzureClientId() {
	_jsii_.InvokeVoid(
		r,
		"resetAzureClientId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) ResetAzureContainerName() {
	_jsii_.InvokeVoid(
		r,
		"resetAzureContainerName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) ResetAzureStorageAccountName() {
	_jsii_.InvokeVoid(
		r,
		"resetAzureStorageAccountName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) ResetAzureTenantId() {
	_jsii_.InvokeVoid(
		r,
		"resetAzureTenantId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) ResetFilePath() {
	_jsii_.InvokeVoid(
		r,
		"resetFilePath",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

