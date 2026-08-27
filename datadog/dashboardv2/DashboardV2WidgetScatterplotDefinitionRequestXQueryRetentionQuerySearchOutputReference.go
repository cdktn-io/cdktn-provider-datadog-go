// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference interface {
	cdktn.ComplexObject
	CohortCriteria() DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchCohortCriteriaOutputReference
	CohortCriteriaInput() *DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchCohortCriteria
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
	Filters() DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchFiltersOutputReference
	FiltersInput() *DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchFilters
	// Experimental.
	Fqn() *string
	InternalValue() *DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearch
	SetInternalValue(val *DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearch)
	RetentionEntity() *string
	SetRetentionEntity(val *string)
	RetentionEntityInput() *string
	ReturnCondition() *string
	SetReturnCondition(val *string)
	ReturnConditionInput() *string
	ReturnCriteria() DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchReturnCriteriaOutputReference
	ReturnCriteriaInput() *DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchReturnCriteria
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
	PutCohortCriteria(value *DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchCohortCriteria)
	PutFilters(value *DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchFilters)
	PutReturnCriteria(value *DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchReturnCriteria)
	ResetFilters()
	ResetReturnCriteria()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference
type jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) CohortCriteria() DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchCohortCriteriaOutputReference {
	var returns DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchCohortCriteriaOutputReference
	_jsii_.Get(
		j,
		"cohortCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) CohortCriteriaInput() *DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchCohortCriteria {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchCohortCriteria
	_jsii_.Get(
		j,
		"cohortCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) Filters() DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchFiltersOutputReference {
	var returns DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchFiltersOutputReference
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) FiltersInput() *DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchFilters {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchFilters
	_jsii_.Get(
		j,
		"filtersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) InternalValue() *DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearch {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) RetentionEntity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionEntity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) RetentionEntityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionEntityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) ReturnCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) ReturnConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) ReturnCriteria() DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchReturnCriteriaOutputReference {
	var returns DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchReturnCriteriaOutputReference
	_jsii_.Get(
		j,
		"returnCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) ReturnCriteriaInput() *DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchReturnCriteria {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchReturnCriteria
	_jsii_.Get(
		j,
		"returnCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference_Override(d DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference)SetInternalValue(val *DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference)SetRetentionEntity(val *string) {
	if err := j.validateSetRetentionEntityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionEntity",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference)SetReturnCondition(val *string) {
	if err := j.validateSetReturnConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnCondition",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) PutCohortCriteria(value *DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchCohortCriteria) {
	if err := d.validatePutCohortCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCohortCriteria",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) PutFilters(value *DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchFilters) {
	if err := d.validatePutFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFilters",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) PutReturnCriteria(value *DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchReturnCriteria) {
	if err := d.validatePutReturnCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putReturnCriteria",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) ResetFilters() {
	_jsii_.InvokeVoid(
		d,
		"resetFilters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) ResetReturnCriteria() {
	_jsii_.InvokeVoid(
		d,
		"resetReturnCriteria",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryRetentionQuerySearchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

