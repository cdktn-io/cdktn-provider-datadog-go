// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package referencetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/referencetable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference interface {
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
	FilePath() *string
	SetFilePath(val *string)
	FilePathInput() *string
	// Experimental.
	Fqn() *string
	GcpBucketName() *string
	SetGcpBucketName(val *string)
	GcpBucketNameInput() *string
	GcpProjectId() *string
	SetGcpProjectId(val *string)
	GcpProjectIdInput() *string
	GcpServiceAccountEmail() *string
	SetGcpServiceAccountEmail(val *string)
	GcpServiceAccountEmailInput() *string
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
	ResetFilePath()
	ResetGcpBucketName()
	ResetGcpProjectId()
	ResetGcpServiceAccountEmail()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference
type jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) FilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) FilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) GcpBucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpBucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) GcpBucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpBucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) GcpProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) GcpProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpProjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) GcpServiceAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpServiceAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) GcpServiceAccountEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpServiceAccountEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference {
	_init_.Initialize()

	if err := validateNewReferenceTableFileMetadataAccessDetailsGcpDetailOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.referenceTable.ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference_Override(r ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.referenceTable.ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference)SetFilePath(val *string) {
	if err := j.validateSetFilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filePath",
		val,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference)SetGcpBucketName(val *string) {
	if err := j.validateSetGcpBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcpBucketName",
		val,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference)SetGcpProjectId(val *string) {
	if err := j.validateSetGcpProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcpProjectId",
		val,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference)SetGcpServiceAccountEmail(val *string) {
	if err := j.validateSetGcpServiceAccountEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcpServiceAccountEmail",
		val,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) ResetFilePath() {
	_jsii_.InvokeVoid(
		r,
		"resetFilePath",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) ResetGcpBucketName() {
	_jsii_.InvokeVoid(
		r,
		"resetGcpBucketName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) ResetGcpProjectId() {
	_jsii_.InvokeVoid(
		r,
		"resetGcpProjectId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) ResetGcpServiceAccountEmail() {
	_jsii_.InvokeVoid(
		r,
		"resetGcpServiceAccountEmail",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

