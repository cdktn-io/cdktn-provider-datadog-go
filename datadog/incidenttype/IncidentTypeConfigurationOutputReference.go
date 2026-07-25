// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package incidenttype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/incidenttype/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IncidentTypeConfigurationOutputReference interface {
	cdktn.ComplexObject
	AllowIncidentDeletion() interface{}
	SetAllowIncidentDeletion(val interface{})
	AllowIncidentDeletionInput() interface{}
	AllowWorkflows() interface{}
	SetAllowWorkflows(val interface{})
	AllowWorkflowsInput() interface{}
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
	CreateMessage() *string
	SetCreateMessage(val *string)
	CreateMessageInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EditableTimestamps() interface{}
	SetEditableTimestamps(val interface{})
	EditableTimestampsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PrivateIncidents() interface{}
	SetPrivateIncidents(val interface{})
	PrivateIncidentsByDefault() interface{}
	SetPrivateIncidentsByDefault(val interface{})
	PrivateIncidentsByDefaultInput() interface{}
	PrivateIncidentsInput() interface{}
	SlugSource() *string
	SetSlugSource(val *string)
	SlugSourceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TestIncidents() interface{}
	SetTestIncidents(val interface{})
	TestIncidentsInput() interface{}
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
	ResetAllowIncidentDeletion()
	ResetAllowWorkflows()
	ResetCreateMessage()
	ResetEditableTimestamps()
	ResetPrivateIncidents()
	ResetPrivateIncidentsByDefault()
	ResetSlugSource()
	ResetTestIncidents()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IncidentTypeConfigurationOutputReference
type jsiiProxy_IncidentTypeConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) AllowIncidentDeletion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowIncidentDeletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) AllowIncidentDeletionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowIncidentDeletionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) AllowWorkflows() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowWorkflows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) AllowWorkflowsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowWorkflowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) CreateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) CreateMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) EditableTimestamps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"editableTimestamps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) EditableTimestampsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"editableTimestampsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) PrivateIncidents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateIncidents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) PrivateIncidentsByDefault() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateIncidentsByDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) PrivateIncidentsByDefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateIncidentsByDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) PrivateIncidentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateIncidentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) SlugSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slugSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) SlugSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slugSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) TestIncidents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"testIncidents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference) TestIncidentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"testIncidentsInput",
		&returns,
	)
	return returns
}


func NewIncidentTypeConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IncidentTypeConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewIncidentTypeConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IncidentTypeConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.incidentType.IncidentTypeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIncidentTypeConfigurationOutputReference_Override(i IncidentTypeConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.incidentType.IncidentTypeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference)SetAllowIncidentDeletion(val interface{}) {
	if err := j.validateSetAllowIncidentDeletionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowIncidentDeletion",
		val,
	)
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference)SetAllowWorkflows(val interface{}) {
	if err := j.validateSetAllowWorkflowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowWorkflows",
		val,
	)
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference)SetCreateMessage(val *string) {
	if err := j.validateSetCreateMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createMessage",
		val,
	)
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference)SetEditableTimestamps(val interface{}) {
	if err := j.validateSetEditableTimestampsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"editableTimestamps",
		val,
	)
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference)SetPrivateIncidents(val interface{}) {
	if err := j.validateSetPrivateIncidentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIncidents",
		val,
	)
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference)SetPrivateIncidentsByDefault(val interface{}) {
	if err := j.validateSetPrivateIncidentsByDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIncidentsByDefault",
		val,
	)
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference)SetSlugSource(val *string) {
	if err := j.validateSetSlugSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slugSource",
		val,
	)
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IncidentTypeConfigurationOutputReference)SetTestIncidents(val interface{}) {
	if err := j.validateSetTestIncidentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"testIncidents",
		val,
	)
}

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) ResetAllowIncidentDeletion() {
	_jsii_.InvokeVoid(
		i,
		"resetAllowIncidentDeletion",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) ResetAllowWorkflows() {
	_jsii_.InvokeVoid(
		i,
		"resetAllowWorkflows",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) ResetCreateMessage() {
	_jsii_.InvokeVoid(
		i,
		"resetCreateMessage",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) ResetEditableTimestamps() {
	_jsii_.InvokeVoid(
		i,
		"resetEditableTimestamps",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) ResetPrivateIncidents() {
	_jsii_.InvokeVoid(
		i,
		"resetPrivateIncidents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) ResetPrivateIncidentsByDefault() {
	_jsii_.InvokeVoid(
		i,
		"resetPrivateIncidentsByDefault",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) ResetSlugSource() {
	_jsii_.InvokeVoid(
		i,
		"resetSlugSource",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) ResetTestIncidents() {
	_jsii_.InvokeVoid(
		i,
		"resetTestIncidents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IncidentTypeConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

