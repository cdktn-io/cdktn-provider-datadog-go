// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatadogsyntheticstest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/datadatadogsyntheticstest/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatadogSyntheticsTestOptionsListStructOutputReference interface {
	cdktn.ComplexObject
	AcceptSelfSigned() cdktn.IResolvable
	AllowInsecure() cdktn.IResolvable
	BlockedRequestPatterns() *[]*string
	CaptureNetworkPayloads() cdktn.IResolvable
	CheckCertificateRevocation() cdktn.IResolvable
	Ci() DataDatadogSyntheticsTestOptionsListCiList
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
	DisableAiaIntermediateFetching() cdktn.IResolvable
	DisableCors() cdktn.IResolvable
	DisableCsp() cdktn.IResolvable
	FollowRedirects() cdktn.IResolvable
	// Experimental.
	Fqn() *string
	HttpVersion() *string
	IgnoreCertificateValidation() cdktn.IResolvable
	IgnoreServerCertificateError() cdktn.IResolvable
	InitialNavigationTimeout() *float64
	InternalValue() *DataDatadogSyntheticsTestOptionsListStruct
	SetInternalValue(val *DataDatadogSyntheticsTestOptionsListStruct)
	MinFailureDuration() *float64
	MinLocationFailed() *float64
	MonitorName() *string
	MonitorOptions() DataDatadogSyntheticsTestOptionsListMonitorOptionsList
	MonitorPriority() *float64
	NoScreenshot() cdktn.IResolvable
	RestrictedRoles() *[]*string
	Retry() DataDatadogSyntheticsTestOptionsListRetryList
	RumSettings() DataDatadogSyntheticsTestOptionsListRumSettingsList
	Scheduling() DataDatadogSyntheticsTestOptionsListSchedulingList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TickEvery() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatadogSyntheticsTestOptionsListStructOutputReference
type jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) AcceptSelfSigned() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"acceptSelfSigned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) AllowInsecure() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"allowInsecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) BlockedRequestPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedRequestPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) CaptureNetworkPayloads() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"captureNetworkPayloads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) CheckCertificateRevocation() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"checkCertificateRevocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) Ci() DataDatadogSyntheticsTestOptionsListCiList {
	var returns DataDatadogSyntheticsTestOptionsListCiList
	_jsii_.Get(
		j,
		"ci",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) DisableAiaIntermediateFetching() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"disableAiaIntermediateFetching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) DisableCors() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"disableCors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) DisableCsp() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"disableCsp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) FollowRedirects() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"followRedirects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) HttpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) IgnoreCertificateValidation() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"ignoreCertificateValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) IgnoreServerCertificateError() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"ignoreServerCertificateError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) InitialNavigationTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialNavigationTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) InternalValue() *DataDatadogSyntheticsTestOptionsListStruct {
	var returns *DataDatadogSyntheticsTestOptionsListStruct
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) MinFailureDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minFailureDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) MinLocationFailed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLocationFailed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) MonitorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) MonitorOptions() DataDatadogSyntheticsTestOptionsListMonitorOptionsList {
	var returns DataDatadogSyntheticsTestOptionsListMonitorOptionsList
	_jsii_.Get(
		j,
		"monitorOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) MonitorPriority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitorPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) NoScreenshot() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"noScreenshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) RestrictedRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) Retry() DataDatadogSyntheticsTestOptionsListRetryList {
	var returns DataDatadogSyntheticsTestOptionsListRetryList
	_jsii_.Get(
		j,
		"retry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) RumSettings() DataDatadogSyntheticsTestOptionsListRumSettingsList {
	var returns DataDatadogSyntheticsTestOptionsListRumSettingsList
	_jsii_.Get(
		j,
		"rumSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) Scheduling() DataDatadogSyntheticsTestOptionsListSchedulingList {
	var returns DataDatadogSyntheticsTestOptionsListSchedulingList
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) TickEvery() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tickEvery",
		&returns,
	)
	return returns
}


func NewDataDatadogSyntheticsTestOptionsListStructOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatadogSyntheticsTestOptionsListStructOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatadogSyntheticsTestOptionsListStructOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dataDatadogSyntheticsTest.DataDatadogSyntheticsTestOptionsListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatadogSyntheticsTestOptionsListStructOutputReference_Override(d DataDatadogSyntheticsTestOptionsListStructOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dataDatadogSyntheticsTest.DataDatadogSyntheticsTestOptionsListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference)SetInternalValue(val *DataDatadogSyntheticsTestOptionsListStruct) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSyntheticsTestOptionsListStructOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

