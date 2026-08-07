// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationawsaccountccmconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/integrationawsaccountccmconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference interface {
	cdktn.ComplexObject
	BucketName() *string
	SetBucketName(val *string)
	BucketNameInput() *string
	BucketRegion() *string
	SetBucketRegion(val *string)
	BucketRegionInput() *string
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
	ReportName() *string
	SetReportName(val *string)
	ReportNameInput() *string
	ReportPrefix() *string
	SetReportPrefix(val *string)
	ReportPrefixInput() *string
	ReportType() *string
	SetReportType(val *string)
	ReportTypeInput() *string
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
	ResetBucketName()
	ResetBucketRegion()
	ResetReportName()
	ResetReportPrefix()
	ResetReportType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference
type jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) BucketRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) BucketRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) ReportName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) ReportNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) ReportPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) ReportPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) ReportType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) ReportTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewIntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.integrationAwsAccountCcmConfig.IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference_Override(i IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.integrationAwsAccountCcmConfig.IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference)SetBucketRegion(val *string) {
	if err := j.validateSetBucketRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketRegion",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference)SetReportName(val *string) {
	if err := j.validateSetReportNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportName",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference)SetReportPrefix(val *string) {
	if err := j.validateSetReportPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportPrefix",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference)SetReportType(val *string) {
	if err := j.validateSetReportTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportType",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) ResetBucketName() {
	_jsii_.InvokeVoid(
		i,
		"resetBucketName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) ResetBucketRegion() {
	_jsii_.InvokeVoid(
		i,
		"resetBucketRegion",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) ResetReportName() {
	_jsii_.InvokeVoid(
		i,
		"resetReportName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) ResetReportPrefix() {
	_jsii_.InvokeVoid(
		i,
		"resetReportPrefix",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) ResetReportType() {
	_jsii_.InvokeVoid(
		i,
		"resetReportType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

