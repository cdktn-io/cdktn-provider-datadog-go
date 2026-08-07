// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tagindexingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/tagindexingrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TagIndexingRuleOptionsDataOutputReference interface {
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
	DynamicTags() TagIndexingRuleOptionsDataDynamicTagsOutputReference
	DynamicTagsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ManagePreexistingMetrics() interface{}
	SetManagePreexistingMetrics(val interface{})
	ManagePreexistingMetricsInput() interface{}
	OverridePreviousRules() interface{}
	SetOverridePreviousRules(val interface{})
	OverridePreviousRulesInput() interface{}
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
	PutDynamicTags(value *TagIndexingRuleOptionsDataDynamicTags)
	ResetDynamicTags()
	ResetManagePreexistingMetrics()
	ResetOverridePreviousRules()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TagIndexingRuleOptionsDataOutputReference
type jsiiProxy_TagIndexingRuleOptionsDataOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) DynamicTags() TagIndexingRuleOptionsDataDynamicTagsOutputReference {
	var returns TagIndexingRuleOptionsDataDynamicTagsOutputReference
	_jsii_.Get(
		j,
		"dynamicTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) DynamicTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) ManagePreexistingMetrics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managePreexistingMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) ManagePreexistingMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managePreexistingMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) OverridePreviousRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overridePreviousRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) OverridePreviousRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overridePreviousRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTagIndexingRuleOptionsDataOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TagIndexingRuleOptionsDataOutputReference {
	_init_.Initialize()

	if err := validateNewTagIndexingRuleOptionsDataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TagIndexingRuleOptionsDataOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.tagIndexingRule.TagIndexingRuleOptionsDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTagIndexingRuleOptionsDataOutputReference_Override(t TagIndexingRuleOptionsDataOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.tagIndexingRule.TagIndexingRuleOptionsDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataOutputReference)SetManagePreexistingMetrics(val interface{}) {
	if err := j.validateSetManagePreexistingMetricsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managePreexistingMetrics",
		val,
	)
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataOutputReference)SetOverridePreviousRules(val interface{}) {
	if err := j.validateSetOverridePreviousRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overridePreviousRules",
		val,
	)
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) PutDynamicTags(value *TagIndexingRuleOptionsDataDynamicTags) {
	if err := t.validatePutDynamicTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDynamicTags",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) ResetDynamicTags() {
	_jsii_.InvokeVoid(
		t,
		"resetDynamicTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) ResetManagePreexistingMetrics() {
	_jsii_.InvokeVoid(
		t,
		"resetManagePreexistingMetrics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) ResetOverridePreviousRules() {
	_jsii_.InvokeVoid(
		t,
		"resetOverridePreviousRules",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

