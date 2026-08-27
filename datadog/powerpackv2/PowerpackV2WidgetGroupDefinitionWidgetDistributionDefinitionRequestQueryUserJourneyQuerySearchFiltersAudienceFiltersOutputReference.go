// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference interface {
	cdktn.ComplexObject
	Account() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersAccountList
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
	InternalValue() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters
	SetInternalValue(val *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters)
	Segment() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentList
	SegmentInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	User() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersUserList
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

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) Account() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersAccountList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersAccountList
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) AccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) FilterCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) FilterConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) InternalValue() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) Segment() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentList
	_jsii_.Get(
		j,
		"segment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) SegmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"segmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) User() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersUserList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersUserList
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) UserInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference)SetFilterCondition(val *string) {
	if err := j.validateSetFilterConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterCondition",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference)SetInternalValue(val *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) PutAccount(value interface{}) {
	if err := p.validatePutAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAccount",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) PutSegment(value interface{}) {
	if err := p.validatePutSegmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSegment",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) PutUser(value interface{}) {
	if err := p.validatePutUserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putUser",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ResetAccount() {
	_jsii_.InvokeVoid(
		p,
		"resetAccount",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ResetFilterCondition() {
	_jsii_.InvokeVoid(
		p,
		"resetFilterCondition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ResetSegment() {
	_jsii_.InvokeVoid(
		p,
		"resetSegment",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		p,
		"resetUser",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

