// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package costbudget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/costbudget/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CostBudgetBudgetLineOutputReference interface {
	cdktn.ComplexObject
	Amounts() *map[string]*float64
	SetAmounts(val *map[string]*float64)
	AmountsInput() *map[string]*float64
	ChildTagFilters() CostBudgetBudgetLineChildTagFiltersList
	ChildTagFiltersInput() interface{}
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
	ParentTagFilters() CostBudgetBudgetLineParentTagFiltersList
	ParentTagFiltersInput() interface{}
	TagFilters() CostBudgetBudgetLineTagFiltersList
	TagFiltersInput() interface{}
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
	PutChildTagFilters(value interface{})
	PutParentTagFilters(value interface{})
	PutTagFilters(value interface{})
	ResetChildTagFilters()
	ResetParentTagFilters()
	ResetTagFilters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CostBudgetBudgetLineOutputReference
type jsiiProxy_CostBudgetBudgetLineOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CostBudgetBudgetLineOutputReference) Amounts() *map[string]*float64 {
	var returns *map[string]*float64
	_jsii_.Get(
		j,
		"amounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CostBudgetBudgetLineOutputReference) AmountsInput() *map[string]*float64 {
	var returns *map[string]*float64
	_jsii_.Get(
		j,
		"amountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CostBudgetBudgetLineOutputReference) ChildTagFilters() CostBudgetBudgetLineChildTagFiltersList {
	var returns CostBudgetBudgetLineChildTagFiltersList
	_jsii_.Get(
		j,
		"childTagFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CostBudgetBudgetLineOutputReference) ChildTagFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"childTagFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CostBudgetBudgetLineOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CostBudgetBudgetLineOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CostBudgetBudgetLineOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CostBudgetBudgetLineOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CostBudgetBudgetLineOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CostBudgetBudgetLineOutputReference) ParentTagFilters() CostBudgetBudgetLineParentTagFiltersList {
	var returns CostBudgetBudgetLineParentTagFiltersList
	_jsii_.Get(
		j,
		"parentTagFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CostBudgetBudgetLineOutputReference) ParentTagFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parentTagFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CostBudgetBudgetLineOutputReference) TagFilters() CostBudgetBudgetLineTagFiltersList {
	var returns CostBudgetBudgetLineTagFiltersList
	_jsii_.Get(
		j,
		"tagFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CostBudgetBudgetLineOutputReference) TagFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CostBudgetBudgetLineOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CostBudgetBudgetLineOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCostBudgetBudgetLineOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CostBudgetBudgetLineOutputReference {
	_init_.Initialize()

	if err := validateNewCostBudgetBudgetLineOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CostBudgetBudgetLineOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.costBudget.CostBudgetBudgetLineOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCostBudgetBudgetLineOutputReference_Override(c CostBudgetBudgetLineOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.costBudget.CostBudgetBudgetLineOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CostBudgetBudgetLineOutputReference)SetAmounts(val *map[string]*float64) {
	if err := j.validateSetAmountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amounts",
		val,
	)
}

func (j *jsiiProxy_CostBudgetBudgetLineOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CostBudgetBudgetLineOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CostBudgetBudgetLineOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CostBudgetBudgetLineOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CostBudgetBudgetLineOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CostBudgetBudgetLineOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CostBudgetBudgetLineOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CostBudgetBudgetLineOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CostBudgetBudgetLineOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CostBudgetBudgetLineOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CostBudgetBudgetLineOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CostBudgetBudgetLineOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CostBudgetBudgetLineOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CostBudgetBudgetLineOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CostBudgetBudgetLineOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CostBudgetBudgetLineOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CostBudgetBudgetLineOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CostBudgetBudgetLineOutputReference) PutChildTagFilters(value interface{}) {
	if err := c.validatePutChildTagFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putChildTagFilters",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CostBudgetBudgetLineOutputReference) PutParentTagFilters(value interface{}) {
	if err := c.validatePutParentTagFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putParentTagFilters",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CostBudgetBudgetLineOutputReference) PutTagFilters(value interface{}) {
	if err := c.validatePutTagFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTagFilters",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CostBudgetBudgetLineOutputReference) ResetChildTagFilters() {
	_jsii_.InvokeVoid(
		c,
		"resetChildTagFilters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CostBudgetBudgetLineOutputReference) ResetParentTagFilters() {
	_jsii_.InvokeVoid(
		c,
		"resetParentTagFilters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CostBudgetBudgetLineOutputReference) ResetTagFilters() {
	_jsii_.InvokeVoid(
		c,
		"resetTagFilters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CostBudgetBudgetLineOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CostBudgetBudgetLineOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

