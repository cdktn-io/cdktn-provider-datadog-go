// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package referencetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/referencetable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ReferenceTableFileMetadataAccessDetailsOutputReference interface {
	cdktn.ComplexObject
	AwsDetail() ReferenceTableFileMetadataAccessDetailsAwsDetailOutputReference
	AwsDetailInput() interface{}
	AzureDetail() ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference
	AzureDetailInput() interface{}
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
	GcpDetail() ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference
	GcpDetailInput() interface{}
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
	PutAwsDetail(value *ReferenceTableFileMetadataAccessDetailsAwsDetail)
	PutAzureDetail(value *ReferenceTableFileMetadataAccessDetailsAzureDetail)
	PutGcpDetail(value *ReferenceTableFileMetadataAccessDetailsGcpDetail)
	ResetAwsDetail()
	ResetAzureDetail()
	ResetGcpDetail()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ReferenceTableFileMetadataAccessDetailsOutputReference
type jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) AwsDetail() ReferenceTableFileMetadataAccessDetailsAwsDetailOutputReference {
	var returns ReferenceTableFileMetadataAccessDetailsAwsDetailOutputReference
	_jsii_.Get(
		j,
		"awsDetail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) AwsDetailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsDetailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) AzureDetail() ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference {
	var returns ReferenceTableFileMetadataAccessDetailsAzureDetailOutputReference
	_jsii_.Get(
		j,
		"azureDetail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) AzureDetailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureDetailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) GcpDetail() ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference {
	var returns ReferenceTableFileMetadataAccessDetailsGcpDetailOutputReference
	_jsii_.Get(
		j,
		"gcpDetail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) GcpDetailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gcpDetailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewReferenceTableFileMetadataAccessDetailsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ReferenceTableFileMetadataAccessDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewReferenceTableFileMetadataAccessDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.referenceTable.ReferenceTableFileMetadataAccessDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewReferenceTableFileMetadataAccessDetailsOutputReference_Override(r ReferenceTableFileMetadataAccessDetailsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.referenceTable.ReferenceTableFileMetadataAccessDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) PutAwsDetail(value *ReferenceTableFileMetadataAccessDetailsAwsDetail) {
	if err := r.validatePutAwsDetailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAwsDetail",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) PutAzureDetail(value *ReferenceTableFileMetadataAccessDetailsAzureDetail) {
	if err := r.validatePutAzureDetailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAzureDetail",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) PutGcpDetail(value *ReferenceTableFileMetadataAccessDetailsGcpDetail) {
	if err := r.validatePutGcpDetailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putGcpDetail",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) ResetAwsDetail() {
	_jsii_.InvokeVoid(
		r,
		"resetAwsDetail",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) ResetAzureDetail() {
	_jsii_.InvokeVoid(
		r,
		"resetAzureDetail",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) ResetGcpDetail() {
	_jsii_.InvokeVoid(
		r,
		"resetGcpDetail",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (r *jsiiProxy_ReferenceTableFileMetadataAccessDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

