// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityfindingsticketcreationrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/securityfindingsticketcreationrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityFindingsTicketCreationRuleActionOutputReference interface {
	cdktn.ComplexObject
	AssigneeId() *string
	SetAssigneeId(val *string)
	AssigneeIdInput() *string
	AutoDisabledReason() *string
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
	Fields() *string
	SetFields(val *string)
	FieldsInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxTicketsPerDay() *float64
	SetMaxTicketsPerDay(val *float64)
	MaxTicketsPerDayInput() *float64
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	Target() *string
	SetTarget(val *string)
	TargetInput() *string
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
	ResetAssigneeId()
	ResetFields()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecurityFindingsTicketCreationRuleActionOutputReference
type jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) AssigneeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assigneeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) AssigneeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assigneeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) AutoDisabledReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoDisabledReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) Fields() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) FieldsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) MaxTicketsPerDay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTicketsPerDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) MaxTicketsPerDayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTicketsPerDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSecurityFindingsTicketCreationRuleActionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SecurityFindingsTicketCreationRuleActionOutputReference {
	_init_.Initialize()

	if err := validateNewSecurityFindingsTicketCreationRuleActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.securityFindingsTicketCreationRule.SecurityFindingsTicketCreationRuleActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSecurityFindingsTicketCreationRuleActionOutputReference_Override(s SecurityFindingsTicketCreationRuleActionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.securityFindingsTicketCreationRule.SecurityFindingsTicketCreationRuleActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference)SetAssigneeId(val *string) {
	if err := j.validateSetAssigneeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assigneeId",
		val,
	)
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference)SetFields(val *string) {
	if err := j.validateSetFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fields",
		val,
	)
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference)SetMaxTicketsPerDay(val *float64) {
	if err := j.validateSetMaxTicketsPerDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTicketsPerDay",
		val,
	)
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) ResetAssigneeId() {
	_jsii_.InvokeVoid(
		s,
		"resetAssigneeId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) ResetFields() {
	_jsii_.InvokeVoid(
		s,
		"resetFields",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SecurityFindingsTicketCreationRuleActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

