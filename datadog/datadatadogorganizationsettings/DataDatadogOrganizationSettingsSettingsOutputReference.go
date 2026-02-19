// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatadogorganizationsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/datadatadogorganizationsettings/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatadogOrganizationSettingsSettingsOutputReference interface {
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
	PrivateWidgetShare() cdktn.IResolvable
	Saml() DataDatadogOrganizationSettingsSettingsSamlList
	SamlAutocreateAccessRole() *string
	SamlAutocreateUsersDomains() DataDatadogOrganizationSettingsSettingsSamlAutocreateUsersDomainsList
	SamlAutocreateUsersDomainsInput() interface{}
	SamlCanBeEnabled() cdktn.IResolvable
	SamlIdpEndpoint() *string
	SamlIdpInitiatedLogin() DataDatadogOrganizationSettingsSettingsSamlIdpInitiatedLoginList
	SamlIdpInitiatedLoginInput() interface{}
	SamlIdpMetadataUploaded() cdktn.IResolvable
	SamlInput() interface{}
	SamlLoginUrl() *string
	SamlStrictMode() DataDatadogOrganizationSettingsSettingsSamlStrictModeList
	SamlStrictModeInput() interface{}
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
	PutSaml(value interface{})
	PutSamlAutocreateUsersDomains(value interface{})
	PutSamlIdpInitiatedLogin(value interface{})
	PutSamlStrictMode(value interface{})
	ResetSaml()
	ResetSamlAutocreateUsersDomains()
	ResetSamlIdpInitiatedLogin()
	ResetSamlStrictMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatadogOrganizationSettingsSettingsOutputReference
type jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) PrivateWidgetShare() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"privateWidgetShare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) Saml() DataDatadogOrganizationSettingsSettingsSamlList {
	var returns DataDatadogOrganizationSettingsSettingsSamlList
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) SamlAutocreateAccessRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlAutocreateAccessRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) SamlAutocreateUsersDomains() DataDatadogOrganizationSettingsSettingsSamlAutocreateUsersDomainsList {
	var returns DataDatadogOrganizationSettingsSettingsSamlAutocreateUsersDomainsList
	_jsii_.Get(
		j,
		"samlAutocreateUsersDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) SamlAutocreateUsersDomainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samlAutocreateUsersDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) SamlCanBeEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"samlCanBeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) SamlIdpEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlIdpEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) SamlIdpInitiatedLogin() DataDatadogOrganizationSettingsSettingsSamlIdpInitiatedLoginList {
	var returns DataDatadogOrganizationSettingsSettingsSamlIdpInitiatedLoginList
	_jsii_.Get(
		j,
		"samlIdpInitiatedLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) SamlIdpInitiatedLoginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samlIdpInitiatedLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) SamlIdpMetadataUploaded() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"samlIdpMetadataUploaded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) SamlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) SamlLoginUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlLoginUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) SamlStrictMode() DataDatadogOrganizationSettingsSettingsSamlStrictModeList {
	var returns DataDatadogOrganizationSettingsSettingsSamlStrictModeList
	_jsii_.Get(
		j,
		"samlStrictMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) SamlStrictModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samlStrictModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatadogOrganizationSettingsSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatadogOrganizationSettingsSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatadogOrganizationSettingsSettingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dataDatadogOrganizationSettings.DataDatadogOrganizationSettingsSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatadogOrganizationSettingsSettingsOutputReference_Override(d DataDatadogOrganizationSettingsSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dataDatadogOrganizationSettings.DataDatadogOrganizationSettingsSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) PutSaml(value interface{}) {
	if err := d.validatePutSamlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSaml",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) PutSamlAutocreateUsersDomains(value interface{}) {
	if err := d.validatePutSamlAutocreateUsersDomainsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSamlAutocreateUsersDomains",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) PutSamlIdpInitiatedLogin(value interface{}) {
	if err := d.validatePutSamlIdpInitiatedLoginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSamlIdpInitiatedLogin",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) PutSamlStrictMode(value interface{}) {
	if err := d.validatePutSamlStrictModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSamlStrictMode",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) ResetSaml() {
	_jsii_.InvokeVoid(
		d,
		"resetSaml",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) ResetSamlAutocreateUsersDomains() {
	_jsii_.InvokeVoid(
		d,
		"resetSamlAutocreateUsersDomains",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) ResetSamlIdpInitiatedLogin() {
	_jsii_.InvokeVoid(
		d,
		"resetSamlIdpInitiatedLogin",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) ResetSamlStrictMode() {
	_jsii_.InvokeVoid(
		d,
		"resetSamlStrictMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatadogOrganizationSettingsSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

