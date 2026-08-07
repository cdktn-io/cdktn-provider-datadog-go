// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oncallteamroutingrules

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/oncallteamroutingrules/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference interface {
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
	EndDay() *string
	SetEndDay(val *string)
	EndDayInput() *string
	EndTime() *string
	SetEndTime(val *string)
	EndTimeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StartDay() *string
	SetStartDay(val *string)
	StartDayInput() *string
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
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
	ResetEndDay()
	ResetEndTime()
	ResetStartDay()
	ResetStartTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference
type jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) EndDay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) EndDayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) EndTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) StartDay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) StartDayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference {
	_init_.Initialize()

	if err := validateNewOnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.onCallTeamRoutingRules.OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference_Override(o OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.onCallTeamRoutingRules.OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference)SetEndDay(val *string) {
	if err := j.validateSetEndDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endDay",
		val,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference)SetEndTime(val *string) {
	if err := j.validateSetEndTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endTime",
		val,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference)SetStartDay(val *string) {
	if err := j.validateSetStartDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startDay",
		val,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference)SetStartTime(val *string) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) ResetEndDay() {
	_jsii_.InvokeVoid(
		o,
		"resetEndDay",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) ResetEndTime() {
	_jsii_.InvokeVoid(
		o,
		"resetEndTime",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) ResetStartDay() {
	_jsii_.InvokeVoid(
		o,
		"resetStartDay",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		o,
		"resetStartTime",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionEscalationPolicySupportHoursRestrictionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

