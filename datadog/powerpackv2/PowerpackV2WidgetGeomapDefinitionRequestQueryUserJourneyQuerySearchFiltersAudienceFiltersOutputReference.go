// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference interface {
	cdktn.ComplexObject
	Account() PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersAccountList
	AccountInput() interface{}
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
	FilterCondition() *string
	SetFilterCondition(val *string)
	FilterConditionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters
	SetInternalValue(val *PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters)
	Segment() PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentList
	SegmentInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	User() PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersUserList
	UserInput() interface{}
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
	PutAccount(value interface{})
	PutSegment(value interface{})
	PutUser(value interface{})
	ResetAccount()
	ResetFilterCondition()
	ResetSegment()
	ResetUser()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference
type jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) Account() PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersAccountList {
	var returns PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersAccountList
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) AccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) FilterCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) FilterConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) InternalValue() *PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters {
	var returns *PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) Segment() PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentList {
	var returns PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentList
	_jsii_.Get(
		j,
		"segment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) SegmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"segmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) User() PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersUserList {
	var returns PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersUserList
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) UserInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference_Override(p PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference)SetFilterCondition(val *string) {
	if err := j.validateSetFilterConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterCondition",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference)SetInternalValue(val *PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) PutAccount(value interface{}) {
	if err := p.validatePutAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAccount",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) PutSegment(value interface{}) {
	if err := p.validatePutSegmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSegment",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) PutUser(value interface{}) {
	if err := p.validatePutUserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putUser",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ResetAccount() {
	_jsii_.InvokeVoid(
		p,
		"resetAccount",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ResetFilterCondition() {
	_jsii_.InvokeVoid(
		p,
		"resetFilterCondition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ResetSegment() {
	_jsii_.InvokeVoid(
		p,
		"resetSegment",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		p,
		"resetUser",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

