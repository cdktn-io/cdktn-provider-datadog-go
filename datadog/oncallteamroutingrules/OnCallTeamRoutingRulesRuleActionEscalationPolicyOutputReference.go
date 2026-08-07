// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oncallteamroutingrules

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/oncallteamroutingrules/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference interface {
	cdktn.ComplexObject
	AckTimeoutMinutes() *float64
	SetAckTimeoutMinutes(val *float64)
	AckTimeoutMinutesInput() *float64
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
	PolicyId() *string
	SetPolicyId(val *string)
	PolicyIdInput() *string
	SupportHours() OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursOutputReference
	SupportHoursInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Urgency() *string
	SetUrgency(val *string)
	UrgencyInput() *string
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
	PutSupportHours(value *OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHours)
	ResetAckTimeoutMinutes()
	ResetPolicyId()
	ResetSupportHours()
	ResetUrgency()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference
type jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) AckTimeoutMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ackTimeoutMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) AckTimeoutMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ackTimeoutMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) SupportHours() OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursOutputReference {
	var returns OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursOutputReference
	_jsii_.Get(
		j,
		"supportHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) SupportHoursInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) Urgency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urgency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) UrgencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urgencyInput",
		&returns,
	)
	return returns
}


func NewOnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewOnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.onCallTeamRoutingRules.OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference_Override(o OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.onCallTeamRoutingRules.OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference)SetAckTimeoutMinutes(val *float64) {
	if err := j.validateSetAckTimeoutMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ackTimeoutMinutes",
		val,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference)SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference)SetUrgency(val *string) {
	if err := j.validateSetUrgencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urgency",
		val,
	)
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) PutSupportHours(value *OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHours) {
	if err := o.validatePutSupportHoursParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSupportHours",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) ResetAckTimeoutMinutes() {
	_jsii_.InvokeVoid(
		o,
		"resetAckTimeoutMinutes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) ResetPolicyId() {
	_jsii_.InvokeVoid(
		o,
		"resetPolicyId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) ResetSupportHours() {
	_jsii_.InvokeVoid(
		o,
		"resetSupportHours",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) ResetUrgency() {
	_jsii_.InvokeVoid(
		o,
		"resetUrgency",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

