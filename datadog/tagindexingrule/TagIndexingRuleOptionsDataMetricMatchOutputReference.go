// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tagindexingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/tagindexingrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TagIndexingRuleOptionsDataMetricMatchOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsQueried() interface{}
	SetIsQueried(val interface{})
	IsQueriedInput() interface{}
	NotQueried() interface{}
	SetNotQueried(val interface{})
	NotQueriedInput() interface{}
	NotUsedInAssets() interface{}
	SetNotUsedInAssets(val interface{})
	NotUsedInAssetsInput() interface{}
	QueriedWindowSeconds() *float64
	SetQueriedWindowSeconds(val *float64)
	QueriedWindowSecondsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UsedInAssets() interface{}
	SetUsedInAssets(val interface{})
	UsedInAssetsInput() interface{}
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
	ResetIsQueried()
	ResetNotQueried()
	ResetNotUsedInAssets()
	ResetQueriedWindowSeconds()
	ResetUsedInAssets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TagIndexingRuleOptionsDataMetricMatchOutputReference
type jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) IsQueried() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isQueried",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) IsQueriedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isQueriedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) NotQueried() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notQueried",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) NotQueriedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notQueriedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) NotUsedInAssets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notUsedInAssets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) NotUsedInAssetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notUsedInAssetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) QueriedWindowSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queriedWindowSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) QueriedWindowSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queriedWindowSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) UsedInAssets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usedInAssets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) UsedInAssetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usedInAssetsInput",
		&returns,
	)
	return returns
}


func NewTagIndexingRuleOptionsDataMetricMatchOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TagIndexingRuleOptionsDataMetricMatchOutputReference {
	_init_.Initialize()

	if err := validateNewTagIndexingRuleOptionsDataMetricMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.tagIndexingRule.TagIndexingRuleOptionsDataMetricMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTagIndexingRuleOptionsDataMetricMatchOutputReference_Override(t TagIndexingRuleOptionsDataMetricMatchOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.tagIndexingRule.TagIndexingRuleOptionsDataMetricMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference)SetIsQueried(val interface{}) {
	if err := j.validateSetIsQueriedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isQueried",
		val,
	)
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference)SetNotQueried(val interface{}) {
	if err := j.validateSetNotQueriedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notQueried",
		val,
	)
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference)SetNotUsedInAssets(val interface{}) {
	if err := j.validateSetNotUsedInAssetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notUsedInAssets",
		val,
	)
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference)SetQueriedWindowSeconds(val *float64) {
	if err := j.validateSetQueriedWindowSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queriedWindowSeconds",
		val,
	)
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference)SetUsedInAssets(val interface{}) {
	if err := j.validateSetUsedInAssetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usedInAssets",
		val,
	)
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) ResetIsQueried() {
	_jsii_.InvokeVoid(
		t,
		"resetIsQueried",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) ResetNotQueried() {
	_jsii_.InvokeVoid(
		t,
		"resetNotQueried",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) ResetNotUsedInAssets() {
	_jsii_.InvokeVoid(
		t,
		"resetNotUsedInAssets",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) ResetQueriedWindowSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetQueriedWindowSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) ResetUsedInAssets() {
	_jsii_.InvokeVoid(
		t,
		"resetUsedInAssets",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TagIndexingRuleOptionsDataMetricMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

