// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatadogcostbudget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/datadatadogcostbudget/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatadogCostBudgetBudgetLineOutputReference interface {
	cdktn.ComplexObject
	Amounts() cdktn.NumberMap
	ChildTagFilters() DataDatadogCostBudgetBudgetLineChildTagFiltersList
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
	ParentTagFilters() DataDatadogCostBudgetBudgetLineParentTagFiltersList
	ParentTagFiltersInput() interface{}
	TagFilters() DataDatadogCostBudgetBudgetLineTagFiltersList
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

// The jsii proxy struct for DataDatadogCostBudgetBudgetLineOutputReference
type jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) Amounts() cdktn.NumberMap {
	var returns cdktn.NumberMap
	_jsii_.Get(
		j,
		"amounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) ChildTagFilters() DataDatadogCostBudgetBudgetLineChildTagFiltersList {
	var returns DataDatadogCostBudgetBudgetLineChildTagFiltersList
	_jsii_.Get(
		j,
		"childTagFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) ChildTagFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"childTagFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) ParentTagFilters() DataDatadogCostBudgetBudgetLineParentTagFiltersList {
	var returns DataDatadogCostBudgetBudgetLineParentTagFiltersList
	_jsii_.Get(
		j,
		"parentTagFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) ParentTagFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parentTagFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) TagFilters() DataDatadogCostBudgetBudgetLineTagFiltersList {
	var returns DataDatadogCostBudgetBudgetLineTagFiltersList
	_jsii_.Get(
		j,
		"tagFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) TagFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatadogCostBudgetBudgetLineOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatadogCostBudgetBudgetLineOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatadogCostBudgetBudgetLineOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dataDatadogCostBudget.DataDatadogCostBudgetBudgetLineOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatadogCostBudgetBudgetLineOutputReference_Override(d DataDatadogCostBudgetBudgetLineOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dataDatadogCostBudget.DataDatadogCostBudgetBudgetLineOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) PutChildTagFilters(value interface{}) {
	if err := d.validatePutChildTagFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putChildTagFilters",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) PutParentTagFilters(value interface{}) {
	if err := d.validatePutParentTagFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putParentTagFilters",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) PutTagFilters(value interface{}) {
	if err := d.validatePutTagFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTagFilters",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) ResetChildTagFilters() {
	_jsii_.InvokeVoid(
		d,
		"resetChildTagFilters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) ResetParentTagFilters() {
	_jsii_.InvokeVoid(
		d,
		"resetParentTagFilters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) ResetTagFilters() {
	_jsii_.InvokeVoid(
		d,
		"resetTagFilters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogCostBudgetBudgetLineOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

