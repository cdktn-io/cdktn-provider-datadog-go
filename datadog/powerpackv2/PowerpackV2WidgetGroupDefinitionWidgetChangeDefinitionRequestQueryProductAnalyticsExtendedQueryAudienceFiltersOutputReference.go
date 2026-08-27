// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference interface {
	cdktn.ComplexObject
	Account() PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersAccountList
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
	InternalValue() *PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFilters
	SetInternalValue(val *PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFilters)
	Segment() PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersSegmentList
	SegmentInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	User() PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersUserList
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

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) Account() PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersAccountList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersAccountList
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) AccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) FilterCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) FilterConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) InternalValue() *PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFilters {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFilters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) Segment() PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersSegmentList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersSegmentList
	_jsii_.Get(
		j,
		"segment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) SegmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"segmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) User() PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersUserList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersUserList
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) UserInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference)SetFilterCondition(val *string) {
	if err := j.validateSetFilterConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterCondition",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference)SetInternalValue(val *PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFilters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) PutAccount(value interface{}) {
	if err := p.validatePutAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAccount",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) PutSegment(value interface{}) {
	if err := p.validatePutSegmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSegment",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) PutUser(value interface{}) {
	if err := p.validatePutUserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putUser",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) ResetAccount() {
	_jsii_.InvokeVoid(
		p,
		"resetAccount",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) ResetFilterCondition() {
	_jsii_.InvokeVoid(
		p,
		"resetFilterCondition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) ResetSegment() {
	_jsii_.InvokeVoid(
		p,
		"resetSegment",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		p,
		"resetUser",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

