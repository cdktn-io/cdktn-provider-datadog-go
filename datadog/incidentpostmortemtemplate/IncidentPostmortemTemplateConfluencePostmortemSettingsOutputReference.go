// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package incidentpostmortemtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/incidentpostmortemtemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference interface {
	cdktn.ComplexObject
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
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
	ParentId() *string
	SetParentId(val *string)
	ParentIdInput() *string
	SpaceId() *string
	SetSpaceId(val *string)
	SpaceIdInput() *string
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
	ResetAccountId()
	ResetParentId()
	ResetSpaceId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference
type jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) ParentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) ParentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) SpaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) SpaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewIncidentPostmortemTemplateConfluencePostmortemSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.incidentPostmortemTemplate.IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference_Override(i IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.incidentPostmortemTemplate.IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference)SetParentId(val *string) {
	if err := j.validateSetParentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentId",
		val,
	)
}

func (j *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference)SetSpaceId(val *string) {
	if err := j.validateSetSpaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spaceId",
		val,
	)
}

func (j *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) ResetAccountId() {
	_jsii_.InvokeVoid(
		i,
		"resetAccountId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) ResetParentId() {
	_jsii_.InvokeVoid(
		i,
		"resetParentId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) ResetSpaceId() {
	_jsii_.InvokeVoid(
		i,
		"resetSpaceId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IncidentPostmortemTemplateConfluencePostmortemSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

