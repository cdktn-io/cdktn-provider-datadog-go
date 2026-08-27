// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference interface {
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
	Filters() DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchFiltersOutputReference
	FiltersInput() *DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchFilters
	// Experimental.
	Fqn() *string
	InternalValue() *DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearch
	SetInternalValue(val *DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearch)
	JoinKeys() DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchJoinKeysOutputReference
	JoinKeysInput() *DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchJoinKeys
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
	PutFilters(value *DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchFilters)
	PutJoinKeys(value *DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchJoinKeys)
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

// The jsii proxy struct for DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference
type jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) Expression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) ExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) Filters() DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchFiltersOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchFiltersOutputReference
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) FiltersInput() *DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchFilters {
	var returns *DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchFilters
	_jsii_.Get(
		j,
		"filtersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) InternalValue() *DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearch {
	var returns *DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) JoinKeys() DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchJoinKeysOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchJoinKeysOutputReference
	_jsii_.Get(
		j,
		"joinKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) JoinKeysInput() *DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchJoinKeys {
	var returns *DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchJoinKeys
	_jsii_.Get(
		j,
		"joinKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) NodeObjects() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) NodeObjectsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) StepAliases() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stepAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) StepAliasesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stepAliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference_Override(d DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference)SetExpression(val *string) {
	if err := j.validateSetExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expression",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference)SetInternalValue(val *DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference)SetNodeObjects(val *string) {
	if err := j.validateSetNodeObjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeObjects",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference)SetStepAliases(val *string) {
	if err := j.validateSetStepAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stepAliases",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) PutFilters(value *DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchFilters) {
	if err := d.validatePutFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFilters",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) PutJoinKeys(value *DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchJoinKeys) {
	if err := d.validatePutJoinKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putJoinKeys",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) ResetFilters() {
	_jsii_.InvokeVoid(
		d,
		"resetFilters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) ResetJoinKeys() {
	_jsii_.InvokeVoid(
		d,
		"resetJoinKeys",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) ResetStepAliases() {
	_jsii_.InvokeVoid(
		d,
		"resetStepAliases",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestFillQueryUserJourneyQuerySearchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

