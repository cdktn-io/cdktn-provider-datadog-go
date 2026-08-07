// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticstest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/syntheticstest/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SyntheticsTestApiStepRequestBasicauthOutputReference interface {
	cdktn.ComplexObject
	AccessKey() *string
	SetAccessKey(val *string)
	AccessKeyInput() *string
	AccessTokenUrl() *string
	SetAccessTokenUrl(val *string)
	AccessTokenUrlInput() *string
	AddClaimsExp() interface{}
	SetAddClaimsExp(val interface{})
	AddClaimsExpInput() interface{}
	AddClaimsIat() interface{}
	SetAddClaimsIat(val interface{})
	AddClaimsIatInput() interface{}
	Algorithm() *string
	SetAlgorithm(val *string)
	AlgorithmInput() *string
	Audience() *string
	SetAudience(val *string)
	AudienceInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
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
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	ExpiresIn() *float64
	SetExpiresIn(val *float64)
	ExpiresInInput() *float64
	// Experimental.
	Fqn() *string
	Header() *string
	SetHeader(val *string)
	HeaderInput() *string
	InternalValue() *SyntheticsTestApiStepRequestBasicauth
	SetInternalValue(val *SyntheticsTestApiStepRequestBasicauth)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Payload() *string
	SetPayload(val *string)
	PayloadInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	Resource() *string
	SetResource(val *string)
	ResourceInput() *string
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	Secret() *string
	SetSecret(val *string)
	SecretInput() *string
	SecretKey() *string
	SetSecretKey(val *string)
	SecretKeyInput() *string
	ServiceName() *string
	SetServiceName(val *string)
	ServiceNameInput() *string
	SessionToken() *string
	SetSessionToken(val *string)
	SessionTokenInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TokenApiAuthentication() *string
	SetTokenApiAuthentication(val *string)
	TokenApiAuthenticationInput() *string
	TokenPrefix() *string
	SetTokenPrefix(val *string)
	TokenPrefixInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	Workstation() *string
	SetWorkstation(val *string)
	WorkstationInput() *string
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
	ResetAccessKey()
	ResetAccessTokenUrl()
	ResetAddClaimsExp()
	ResetAddClaimsIat()
	ResetAlgorithm()
	ResetAudience()
	ResetClientId()
	ResetClientSecret()
	ResetDomain()
	ResetExpiresIn()
	ResetHeader()
	ResetPassword()
	ResetPayload()
	ResetRegion()
	ResetResource()
	ResetScope()
	ResetSecret()
	ResetSecretKey()
	ResetServiceName()
	ResetSessionToken()
	ResetTokenApiAuthentication()
	ResetTokenPrefix()
	ResetType()
	ResetUsername()
	ResetWorkstation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SyntheticsTestApiStepRequestBasicauthOutputReference
type jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) AccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) AccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) AccessTokenUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) AccessTokenUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) AddClaimsExp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addClaimsExp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) AddClaimsExpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addClaimsExpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) AddClaimsIat() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addClaimsIat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) AddClaimsIatInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addClaimsIatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) Algorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) AlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) Audience() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) AudienceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ExpiresIn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expiresIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ExpiresInInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expiresInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) Header() *string {
	var returns *string
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) HeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) InternalValue() *SyntheticsTestApiStepRequestBasicauth {
	var returns *SyntheticsTestApiStepRequestBasicauth
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) Payload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) PayloadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) Resource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) Secret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) SecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) SecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) SecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) SessionToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) SessionTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) TokenApiAuthentication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenApiAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) TokenApiAuthenticationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenApiAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) TokenPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) TokenPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) Workstation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workstation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) WorkstationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workstationInput",
		&returns,
	)
	return returns
}


func NewSyntheticsTestApiStepRequestBasicauthOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SyntheticsTestApiStepRequestBasicauthOutputReference {
	_init_.Initialize()

	if err := validateNewSyntheticsTestApiStepRequestBasicauthOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.syntheticsTest.SyntheticsTestApiStepRequestBasicauthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSyntheticsTestApiStepRequestBasicauthOutputReference_Override(s SyntheticsTestApiStepRequestBasicauthOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.syntheticsTest.SyntheticsTestApiStepRequestBasicauthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetAccessKey(val *string) {
	if err := j.validateSetAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessKey",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetAccessTokenUrl(val *string) {
	if err := j.validateSetAccessTokenUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessTokenUrl",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetAddClaimsExp(val interface{}) {
	if err := j.validateSetAddClaimsExpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addClaimsExp",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetAddClaimsIat(val interface{}) {
	if err := j.validateSetAddClaimsIatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addClaimsIat",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetAlgorithm(val *string) {
	if err := j.validateSetAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"algorithm",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetAudience(val *string) {
	if err := j.validateSetAudienceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audience",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetExpiresIn(val *float64) {
	if err := j.validateSetExpiresInParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expiresIn",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetHeader(val *string) {
	if err := j.validateSetHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"header",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetInternalValue(val *SyntheticsTestApiStepRequestBasicauth) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetPayload(val *string) {
	if err := j.validateSetPayloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payload",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetResource(val *string) {
	if err := j.validateSetResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetSecret(val *string) {
	if err := j.validateSetSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secret",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetSecretKey(val *string) {
	if err := j.validateSetSecretKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretKey",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetServiceName(val *string) {
	if err := j.validateSetServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceName",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetSessionToken(val *string) {
	if err := j.validateSetSessionTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionToken",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetTokenApiAuthentication(val *string) {
	if err := j.validateSetTokenApiAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenApiAuthentication",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetTokenPrefix(val *string) {
	if err := j.validateSetTokenPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenPrefix",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference)SetWorkstation(val *string) {
	if err := j.validateSetWorkstationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workstation",
		val,
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetAccessKey() {
	_jsii_.InvokeVoid(
		s,
		"resetAccessKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetAccessTokenUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetAccessTokenUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetAddClaimsExp() {
	_jsii_.InvokeVoid(
		s,
		"resetAddClaimsExp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetAddClaimsIat() {
	_jsii_.InvokeVoid(
		s,
		"resetAddClaimsIat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetAlgorithm() {
	_jsii_.InvokeVoid(
		s,
		"resetAlgorithm",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetAudience() {
	_jsii_.InvokeVoid(
		s,
		"resetAudience",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		s,
		"resetClientId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		s,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetDomain() {
	_jsii_.InvokeVoid(
		s,
		"resetDomain",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetExpiresIn() {
	_jsii_.InvokeVoid(
		s,
		"resetExpiresIn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		s,
		"resetHeader",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		s,
		"resetPassword",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetPayload() {
	_jsii_.InvokeVoid(
		s,
		"resetPayload",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		s,
		"resetRegion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetResource() {
	_jsii_.InvokeVoid(
		s,
		"resetResource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		s,
		"resetScope",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		s,
		"resetSecret",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetSecretKey() {
	_jsii_.InvokeVoid(
		s,
		"resetSecretKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetServiceName() {
	_jsii_.InvokeVoid(
		s,
		"resetServiceName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetSessionToken() {
	_jsii_.InvokeVoid(
		s,
		"resetSessionToken",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetTokenApiAuthentication() {
	_jsii_.InvokeVoid(
		s,
		"resetTokenApiAuthentication",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetTokenPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetTokenPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		s,
		"resetType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		s,
		"resetUsername",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ResetWorkstation() {
	_jsii_.InvokeVoid(
		s,
		"resetWorkstation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestBasicauthOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

