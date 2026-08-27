// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference interface {
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
	Expression() *string
	SetExpression(val *string)
	ExpressionInput() *string
	Filters() DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference
	FiltersInput() *DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchFilters
	// Experimental.
	Fqn() *string
	InternalValue() *DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearch
	SetInternalValue(val *DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearch)
	JoinKeys() DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchJoinKeysOutputReference
	JoinKeysInput() *DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchJoinKeys
	NodeObjects() *string
	SetNodeObjects(val *string)
	NodeObjectsInput() *string
	StepAliases() *string
	SetStepAliases(val *string)
	StepAliasesInput() *string
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
	PutFilters(value *DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchFilters)
	PutJoinKeys(value *DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchJoinKeys)
	ResetFilters()
	ResetJoinKeys()
	ResetStepAliases()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference
type jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) Expression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) Filters() DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchFiltersOutputReference
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) FiltersInput() *DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchFilters {
	var returns *DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchFilters
	_jsii_.Get(
		j,
		"filtersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) InternalValue() *DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearch {
	var returns *DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) JoinKeys() DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchJoinKeysOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchJoinKeysOutputReference
	_jsii_.Get(
		j,
		"joinKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) JoinKeysInput() *DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchJoinKeys {
	var returns *DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchJoinKeys
	_jsii_.Get(
		j,
		"joinKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) NodeObjects() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) NodeObjectsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) StepAliases() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stepAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) StepAliasesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stepAliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference_Override(d DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetExpression(val *string) {
	if err := j.validateSetExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expression",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetInternalValue(val *DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetNodeObjects(val *string) {
	if err := j.validateSetNodeObjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeObjects",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetStepAliases(val *string) {
	if err := j.validateSetStepAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stepAliases",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) PutFilters(value *DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchFilters) {
	if err := d.validatePutFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFilters",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) PutJoinKeys(value *DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchJoinKeys) {
	if err := d.validatePutJoinKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putJoinKeys",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ResetFilters() {
	_jsii_.InvokeVoid(
		d,
		"resetFilters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ResetJoinKeys() {
	_jsii_.InvokeVoid(
		d,
		"resetJoinKeys",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ResetStepAliases() {
	_jsii_.InvokeVoid(
		d,
		"resetStepAliases",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetToplistDefinitionRequestQueryUserJourneyQuerySearchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

